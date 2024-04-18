package notifier

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"sort"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/prometheus/alertmanager/pkg/labels"

	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
	"github.com/grafana/grafana/pkg/services/ngalert/models"
)

type autogenRuleStore interface {
	ListNotificationSettings(ctx context.Context, q models.ListNotificationSettingsQuery) (map[models.AlertRuleKey][]models.NotificationSettings, error)
}

// AddAutogenConfig creates the autogenerated configuration and adds it to the given apiAlertingConfig.
// If skipInvalid is true, then invalid notification settings are skipped, otherwise an error is returned.
func AddAutogenConfig[R receiver](ctx context.Context, logger log.Logger, store autogenRuleStore, orgId int64, cfg apiAlertingConfig[R], skipInvalid bool) error {
	autogenRoute, err := newAutogeneratedRoute(ctx, logger, store, orgId, cfg, skipInvalid)
	if err != nil {
		return err
	}

	err = autogenRoute.addToRoute(cfg.GetRoute())
	if err != nil {
		return err
	}

	return nil
}

// newAutogeneratedRoute creates a new autogenerated route based on the notification settings for the given org.
// cfg is used to construct the settings validator and to ensure we create a dedicated route for each receiver.
// skipInvalid is used to skip invalid settings instead of returning an error.
func newAutogeneratedRoute[R receiver](ctx context.Context, logger log.Logger, store autogenRuleStore, orgId int64, cfg apiAlertingConfig[R], skipInvalid bool) (autogeneratedRoute, error) {
	settings, err := store.ListNotificationSettings(ctx, models.ListNotificationSettingsQuery{OrgID: orgId})
	if err != nil {
		return autogeneratedRoute{}, fmt.Errorf("failed to list alert rules: %w", err)
	}

	var notificationSettings []models.NotificationSettings
	// Add a default notification setting for each contact point. This is to ensure that we always have a route for each
	// contact point even if no rules are using it. This will prevent race conditions between AM sync and rule sync.
	for _, receiver := range cfg.GetReceivers() {
		notificationSettings = append(notificationSettings, models.NewDefaultNotificationSettings(receiver.GetName()))
	}

	validator := NewNotificationSettingsValidator(cfg)
	for ruleKey, ruleSettings := range settings {
		for _, setting := range ruleSettings {
			// TODO we should register this errors and somehow present to the users or make sure the config is always valid.
			if err = validator.Validate(setting); err != nil {
				if skipInvalid {
					logger.Error("Rule notification settings are invalid. Skipping", append(ruleKey.LogContext(), "error", err)...)
					continue
				}
				return autogeneratedRoute{}, fmt.Errorf("invalid notification settings for rule %s: %w", ruleKey.UID, err)
			}
			notificationSettings = append(notificationSettings, setting)
		}
	}
	if len(notificationSettings) == 0 {
		return autogeneratedRoute{}, nil
	}

	newAutogenRoute, err := generateRouteFromSettings(cfg.GetRoute().Receiver, notificationSettings, skipInvalid, logger)
	if err != nil {
		return autogeneratedRoute{}, fmt.Errorf("failed to create autogenerated route: %w", err)
	}
	return newAutogenRoute, nil
}

type autogeneratedRoute struct {
	Route *definitions.Route
}

// generateRouteFromSettings generates a route and fingerprint for this route. The route is a tree of 3 layers:
//  1. with matcher by label models.AutogeneratedRouteLabel equals 'true'.
//  2. with matcher by receiver name.
//  3. with matcher by unique combination of optional settings. It is created only if there are optional settings.
func generateRouteFromSettings(defaultReceiver string, settings []models.NotificationSettings, skipInvalid bool, logger log.Logger) (autogeneratedRoute, error) {
	rootMatcher, err := labels.NewMatcher(labels.MatchEqual, models.AutogeneratedRouteLabel, "true")
	if err != nil {
		return autogeneratedRoute{}, err
	}

	autoGenRoot := &definitions.Route{
		Receiver:       defaultReceiver,
		ObjectMatchers: definitions.ObjectMatchers{rootMatcher},
		Continue:       false, // We explicitly don't continue toward user-created routes if this matches.
	}

	seenFingerprint := make(map[data.Fingerprint]struct{})
	receiverRoutes := make(map[string]*definitions.Route)
	for _, s := range settings {
		receiverRoute, ok := receiverRoutes[s.Receiver]
		if !ok {
			receiverRoute = models.ReceiverRoute(s.Receiver)
			receiverRoutes[s.Receiver] = receiverRoute
			autoGenRoot.Routes = append(autoGenRoot.Routes, receiverRoute)
		}

		// Do not create hash specific route if all group settings such as mute timings, group_wait, group_interval, etc are default
		if s.IsAllDefault() {
			continue
		}

		fingerprint := s.Fingerprint()
		if _, ok := seenFingerprint[fingerprint]; ok {
			continue
		}
		seenFingerprint[fingerprint] = struct{}{}

		hashedRoute, err := s.ToRoute()
		if err != nil {
			// This should not be possible as we are validating the settings on save, but we should be defensive.
			if skipInvalid {
				logger.Error("Receiver notification settings are invalid. Skipping", "receiver", s.Receiver, "fingerprint", fingerprint.String(), "error", err)
				continue
			}
			return autogeneratedRoute{}, fmt.Errorf("invalid notification settings for receiver %s with fingerprint %s: %w", s.Receiver, fingerprint.String(), err)
		}
		receiverRoute.Routes = append(receiverRoute.Routes, &hashedRoute)
	}

	// Sorted to ensure resulting route is stable.
	sort.SliceStable(autoGenRoot.Routes, func(i, j int) bool {
		return autoGenRoot.Routes[i].Receiver < autoGenRoot.Routes[j].Receiver
	})
	for _, route := range autoGenRoot.Routes {
		sort.SliceStable(route.Routes, func(i, j int) bool {
			return route.Routes[i].ObjectMatchers[0].Value < route.Routes[j].ObjectMatchers[0].Value // Fingerprint.
		})
	}

	return autogeneratedRoute{
		Route: autoGenRoot,
	}, nil
}

// addToRoute adds this autogenerated route to the given route as the first top-level route under the root.
func (ar *autogeneratedRoute) addToRoute(route *definitions.Route) error {
	if route == nil {
		return errors.New("route does not exist")
	}
	if ar == nil || ar.Route == nil {
		return nil
	}
	// Combine autogenerated route with the user-created route.
	ar.Route.Receiver = route.Receiver

	// Remove existing autogenerated route if it exists.
	RemoveAutogenConfigIfExists(route)

	route.Routes = append([]*definitions.Route{ar.Route}, route.Routes...)
	return nil
}

// RemoveAutogenConfigIfExists removes all top-level autogenerated routes from the provided route.
// If no autogenerated routes exist, this function does nothing.
func RemoveAutogenConfigIfExists(route *definitions.Route) {
	route.Routes = slices.DeleteFunc(route.Routes, func(route *definitions.Route) bool {
		return isAutogeneratedRoot(route)
	})
}

// isAutogeneratedRoot returns true if the route is the root of an autogenerated route.
func isAutogeneratedRoot(route *definitions.Route) bool {
	return len(route.ObjectMatchers) == 1 && route.ObjectMatchers[0].Name == models.AutogeneratedRouteLabel
}
