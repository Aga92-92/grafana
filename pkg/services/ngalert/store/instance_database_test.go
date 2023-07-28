package store_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/grafana/grafana/pkg/services/featuremgmt"
	"github.com/grafana/grafana/pkg/services/ngalert/models"
	"github.com/grafana/grafana/pkg/services/ngalert/tests"
	"github.com/grafana/grafana/pkg/util"
)

const baseIntervalSeconds = 10

func BenchmarkAlertInstanceOperations(b *testing.B) {
	b.StopTimer()
	ctx := context.Background()
	_, dbstore := tests.SetupTestEnv(b, baseIntervalSeconds)

	const mainOrgID int64 = 1

	alertRule := tests.CreateTestAlertRule(b, ctx, dbstore, 60, mainOrgID)

	// Create some instances to write down and then delete.
	count := 10_003
	instances := make([]models.AlertInstance, 0, count)
	keys := make([]models.AlertInstanceKey, 0, count)
	for i := 0; i < count; i++ {
		labels := models.InstanceLabels{"test": fmt.Sprint(i)}
		_, labelsHash, _ := labels.StringAndHash()
		instance := models.AlertInstance{
			AlertInstanceKey: models.AlertInstanceKey{
				RuleOrgID:  alertRule.OrgID,
				RuleUID:    alertRule.UID,
				LabelsHash: labelsHash,
			},
			CurrentState:  models.InstanceStateFiring,
			CurrentReason: string(models.InstanceStateError),
			Labels:        labels,
		}
		instances = append(instances, instance)
		keys = append(keys, instance.AlertInstanceKey)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, instance := range instances {
			_ = dbstore.SaveAlertInstance(ctx, instance)
		}
		_ = dbstore.DeleteAlertInstances(ctx, keys...)
	}
}

func TestIntegrationAlertInstanceOperations(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	ctx := context.Background()
	_, dbstore := tests.SetupTestEnv(t, baseIntervalSeconds)

	const mainOrgID int64 = 1

	containsHash := func(t *testing.T, instances []*models.AlertInstance, hash string) {
		t.Helper()
		for _, i := range instances {
			if i.LabelsHash == hash {
				return
			}
		}
		require.Fail(t, "%v does not contain an instance with hash %s", instances, hash)
	}

	alertRule1 := tests.CreateTestAlertRule(t, ctx, dbstore, 60, mainOrgID)
	orgID := alertRule1.OrgID

	alertRule2 := tests.CreateTestAlertRule(t, ctx, dbstore, 60, mainOrgID)
	require.Equal(t, orgID, alertRule2.OrgID)

	alertRule3 := tests.CreateTestAlertRule(t, ctx, dbstore, 60, mainOrgID)
	require.Equal(t, orgID, alertRule3.OrgID)

	alertRule4 := tests.CreateTestAlertRule(t, ctx, dbstore, 60, mainOrgID)
	require.Equal(t, orgID, alertRule4.OrgID)

	t.Run("can save and read new alert instance", func(t *testing.T) {
		labels := models.InstanceLabels{"test": "testValue"}
		_, hash, _ := labels.StringAndHash()
		instance := models.AlertInstance{
			AlertInstanceKey: models.AlertInstanceKey{
				RuleOrgID:  alertRule1.OrgID,
				RuleUID:    alertRule1.UID,
				LabelsHash: hash,
			},
			CurrentState:  models.InstanceStateFiring,
			CurrentReason: string(models.InstanceStateError),
			Labels:        labels,
		}
		err := dbstore.SaveAlertInstance(ctx, instance)
		require.NoError(t, err)

		listCmd := &models.ListAlertInstancesQuery{
			RuleOrgID: instance.RuleOrgID,
			RuleUID:   instance.RuleUID,
		}
		alerts, err := dbstore.ListAlertInstances(ctx, listCmd)
		require.NoError(t, err)

		require.Len(t, alerts, 1)
		require.Equal(t, instance.Labels, alerts[0].Labels)
		require.Equal(t, alertRule1.OrgID, alerts[0].RuleOrgID)
		require.Equal(t, alertRule1.UID, alerts[0].RuleUID)
		require.Equal(t, instance.CurrentReason, alerts[0].CurrentReason)
	})

	t.Run("can save and read new alert instance with no labels", func(t *testing.T) {
		labels := models.InstanceLabels{}
		_, hash, _ := labels.StringAndHash()
		instance := models.AlertInstance{
			AlertInstanceKey: models.AlertInstanceKey{
				RuleOrgID:  alertRule2.OrgID,
				RuleUID:    alertRule2.UID,
				LabelsHash: hash,
			},
			CurrentState: models.InstanceStateNormal,
			Labels:       labels,
		}
		err := dbstore.SaveAlertInstance(ctx, instance)
		require.NoError(t, err)

		listCmd := &models.ListAlertInstancesQuery{
			RuleOrgID: instance.RuleOrgID,
			RuleUID:   instance.RuleUID,
		}

		alerts, err := dbstore.ListAlertInstances(ctx, listCmd)
		require.NoError(t, err)

		require.Len(t, alerts, 1)
		require.Equal(t, alertRule2.OrgID, alerts[0].RuleOrgID)
		require.Equal(t, alertRule2.UID, alerts[0].RuleUID)
		require.Equal(t, instance.Labels, alerts[0].Labels)
	})

	t.Run("can save two instances with same org_id, uid and different labels", func(t *testing.T) {
		labels := models.InstanceLabels{"test": "testValue"}
		_, hash, _ := labels.StringAndHash()
		instance1 := models.AlertInstance{
			AlertInstanceKey: models.AlertInstanceKey{
				RuleOrgID:  alertRule3.OrgID,
				RuleUID:    alertRule3.UID,
				LabelsHash: hash,
			},
			CurrentState: models.InstanceStateFiring,
			Labels:       labels,
		}

		err := dbstore.SaveAlertInstance(ctx, instance1)
		require.NoError(t, err)

		labels = models.InstanceLabels{"test": "testValue2"}
		_, hash, _ = labels.StringAndHash()
		instance2 := models.AlertInstance{
			AlertInstanceKey: models.AlertInstanceKey{
				RuleOrgID:  instance1.RuleOrgID,
				RuleUID:    instance1.RuleUID,
				LabelsHash: hash,
			},
			CurrentState: models.InstanceStateFiring,
			Labels:       labels,
		}
		err = dbstore.SaveAlertInstance(ctx, instance2)
		require.NoError(t, err)

		listQuery := &models.ListAlertInstancesQuery{
			RuleOrgID: instance1.RuleOrgID,
			RuleUID:   instance1.RuleUID,
		}

		alerts, err := dbstore.ListAlertInstances(ctx, listQuery)
		require.NoError(t, err)

		require.Len(t, alerts, 2)
	})

	t.Run("can list all added instances in org", func(t *testing.T) {
		listQuery := &models.ListAlertInstancesQuery{
			RuleOrgID: orgID,
		}

		alerts, err := dbstore.ListAlertInstances(ctx, listQuery)
		require.NoError(t, err)

		require.Len(t, alerts, 4)
	})

	t.Run("should ignore Normal state with no reason if feature flag is enabled", func(t *testing.T) {
		labels := models.InstanceLabels{"test": util.GenerateShortUID()}
		instance1 := models.AlertInstance{
			AlertInstanceKey: models.AlertInstanceKey{
				RuleOrgID:  orgID,
				RuleUID:    util.GenerateShortUID(),
				LabelsHash: util.GenerateShortUID(),
			},
			CurrentState:  models.InstanceStateNormal,
			CurrentReason: "",
			Labels:        labels,
		}
		instance2 := models.AlertInstance{
			AlertInstanceKey: models.AlertInstanceKey{
				RuleOrgID:  orgID,
				RuleUID:    util.GenerateShortUID(),
				LabelsHash: util.GenerateShortUID(),
			},
			CurrentState:  models.InstanceStateNormal,
			CurrentReason: models.StateReasonError,
			Labels:        labels,
		}
		err := dbstore.SaveAlertInstance(ctx, instance1)
		require.NoError(t, err)
		err = dbstore.SaveAlertInstance(ctx, instance2)
		require.NoError(t, err)

		listQuery := &models.ListAlertInstancesQuery{
			RuleOrgID: orgID,
		}

		alerts, err := dbstore.ListAlertInstances(ctx, listQuery)
		require.NoError(t, err)

		containsHash(t, alerts, instance1.LabelsHash)

		f := dbstore.FeatureToggles
		dbstore.FeatureToggles = featuremgmt.WithFeatures(featuremgmt.FlagAlertingNoNormalState)
		t.Cleanup(func() {
			dbstore.FeatureToggles = f
		})

		alerts, err = dbstore.ListAlertInstances(ctx, listQuery)
		require.NoError(t, err)

		containsHash(t, alerts, instance2.LabelsHash)

		for _, instance := range alerts {
			if instance.CurrentState == models.InstanceStateNormal && instance.CurrentReason == "" {
				require.Fail(t, "List operation expected to return all states except Normal but the result contains Normal states")
			}
		}
	})

	t.Run("update instance with same org_id, uid and different state", func(t *testing.T) {
		labels := models.InstanceLabels{"test": "testValue"}
		_, hash, _ := labels.StringAndHash()
		instance1 := models.AlertInstance{
			AlertInstanceKey: models.AlertInstanceKey{
				RuleOrgID:  alertRule4.OrgID,
				RuleUID:    alertRule4.UID,
				LabelsHash: hash,
			},
			CurrentState: models.InstanceStateFiring,
			Labels:       labels,
		}

		err := dbstore.SaveAlertInstance(ctx, instance1)
		require.NoError(t, err)

		instance2 := models.AlertInstance{
			AlertInstanceKey: models.AlertInstanceKey{
				RuleOrgID:  alertRule4.OrgID,
				RuleUID:    instance1.RuleUID,
				LabelsHash: instance1.LabelsHash,
			},
			CurrentState: models.InstanceStateNormal,
			Labels:       instance1.Labels,
		}
		err = dbstore.SaveAlertInstance(ctx, instance2)
		require.NoError(t, err)

		listQuery := &models.ListAlertInstancesQuery{
			RuleOrgID: alertRule4.OrgID,
			RuleUID:   alertRule4.UID,
		}

		alerts, err := dbstore.ListAlertInstances(ctx, listQuery)
		require.NoError(t, err)

		require.Len(t, alerts, 1)

		require.Equal(t, instance2.RuleOrgID, alerts[0].RuleOrgID)
		require.Equal(t, instance2.RuleUID, alerts[0].RuleUID)
		require.Equal(t, instance2.Labels, alerts[0].Labels)
		require.Equal(t, instance2.CurrentState, alerts[0].CurrentState)
	})
}

func TestListAlertInstanceData(t *testing.T) {
	ctx := context.Background()
	_, db := tests.SetupTestEnv(t, baseIntervalSeconds)

	t.Run("all instance data is returned", func(t *testing.T) {
		r1 := tests.CreateTestAlertRule(t, ctx, db, 60, 1)
		r2 := tests.CreateTestAlertRule(t, ctx, db, 60, 1)
		require.NoError(t, db.SaveAlertInstanceData(ctx, models.AlertInstanceData{
			OrgID:     1,
			RuleUID:   r1.UID,
			Data:      []byte("data1"),
			ExpiresAt: time.Now().Add(time.Minute),
		}))
		require.NoError(t, db.SaveAlertInstanceData(ctx, models.AlertInstanceData{
			OrgID:     1,
			RuleUID:   r2.UID,
			Data:      []byte("data2"),
			ExpiresAt: time.Now().Add(time.Minute),
		}))
		cmd := models.ListAlertInstancesQuery{}
		data, err := db.ListAlertInstanceData(ctx, &cmd)
		require.NoError(t, err)
		require.Len(t, data, 2)
	})

	t.Run("instance data is returned for just the org", func(t *testing.T) {
		r3 := tests.CreateTestAlertRule(t, ctx, db, 60, 3)
		r4 := tests.CreateTestAlertRule(t, ctx, db, 60, 4)
		require.NoError(t, db.SaveAlertInstanceData(ctx, models.AlertInstanceData{
			OrgID:     3,
			RuleUID:   r3.UID,
			Data:      []byte("data3"),
			ExpiresAt: time.Now().Add(time.Minute),
		}))
		require.NoError(t, db.SaveAlertInstanceData(ctx, models.AlertInstanceData{
			OrgID:     4,
			RuleUID:   r4.UID,
			Data:      []byte("data4"),
			ExpiresAt: time.Now().Add(time.Minute),
		}))
		cmd := models.ListAlertInstancesQuery{
			RuleOrgID: 3,
		}
		data, err := db.ListAlertInstanceData(ctx, &cmd)
		require.NoError(t, err)
		require.Len(t, data, 1)
		assert.Equal(t, data[0].Data, []byte("data3"))
	})

	t.Run("instance data is returned for just the rule", func(t *testing.T) {
		r5 := tests.CreateTestAlertRule(t, ctx, db, 60, 5)
		r6 := tests.CreateTestAlertRule(t, ctx, db, 60, 6)
		require.NoError(t, db.SaveAlertInstanceData(ctx, models.AlertInstanceData{
			OrgID:     5,
			RuleUID:   r5.UID,
			Data:      []byte("data5"),
			ExpiresAt: time.Now().Add(time.Minute),
		}))
		require.NoError(t, db.SaveAlertInstanceData(ctx, models.AlertInstanceData{
			OrgID:     6,
			RuleUID:   r6.UID,
			Data:      []byte("data6"),
			ExpiresAt: time.Now().Add(time.Minute),
		}))
		cmd := models.ListAlertInstancesQuery{
			RuleOrgID: 5,
			RuleUID:   r5.UID,
		}
		data, err := db.ListAlertInstanceData(ctx, &cmd)
		require.NoError(t, err)
		require.Len(t, data, 1)
		assert.Equal(t, data[0].Data, []byte("data5"))
	})

	t.Run("expired instance data is not returned", func(t *testing.T) {
		r7 := tests.CreateTestAlertRule(t, ctx, db, 60, 7)
		require.NoError(t, db.SaveAlertInstanceData(ctx, models.AlertInstanceData{
			OrgID:     7,
			RuleUID:   r7.UID,
			Data:      []byte("data7"),
			ExpiresAt: time.Now(),
		}))
		cmd := models.ListAlertInstancesQuery{}
		data, err := db.ListAlertInstanceData(ctx, &cmd)
		require.NoError(t, err)
		require.Len(t, data, 6)
	})
}

func TestDeleteAlertInstanceData(t *testing.T) {
	ctx := context.Background()
	_, db := tests.SetupTestEnv(t, baseIntervalSeconds)

	r := tests.CreateTestAlertRule(t, ctx, db, 60, 1)
	require.NoError(t, db.SaveAlertInstanceData(ctx, models.AlertInstanceData{
		OrgID:     1,
		RuleUID:   r.UID,
		Data:      []byte("data1"),
		ExpiresAt: time.Now().Add(time.Minute),
	}))

	cmd := models.ListAlertInstancesQuery{}
	data, err := db.ListAlertInstanceData(ctx, &cmd)
	require.NoError(t, err)
	require.Len(t, data, 1)

	ok, err := db.DeleteAlertInstanceData(ctx, models.AlertRuleKey{OrgID: 1, UID: r.UID})
	require.NoError(t, err)
	require.True(t, ok)

	cmd = models.ListAlertInstancesQuery{}
	data, err = db.ListAlertInstanceData(ctx, &cmd)
	require.NoError(t, err)
	require.Len(t, data, 0)

	ok, err = db.DeleteAlertInstanceData(ctx, models.AlertRuleKey{OrgID: 1, UID: r.UID})
	require.NoError(t, err)
	require.False(t, ok)
}

func TestDeleteExpiredAlertInstanceData(t *testing.T) {
	ctx := context.Background()
	_, db := tests.SetupTestEnv(t, baseIntervalSeconds)

	r1 := tests.CreateTestAlertRule(t, ctx, db, 60, 1)
	require.NoError(t, db.SaveAlertInstanceData(ctx, models.AlertInstanceData{
		OrgID:     1,
		RuleUID:   r1.UID,
		Data:      []byte("data1"),
		ExpiresAt: time.Now().Add(time.Minute),
	}))

	// Has not expired so should not be deleted
	n, err := db.DeleteExpiredAlertInstanceData(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(0), n)

	r2 := tests.CreateTestAlertRule(t, ctx, db, 60, 1)
	require.NoError(t, db.SaveAlertInstanceData(ctx, models.AlertInstanceData{
		OrgID:     1,
		RuleUID:   r2.UID,
		Data:      []byte("data2"),
		ExpiresAt: time.Now(),
	}))

	// Last one has expired and should be deleted
	n, err = db.DeleteExpiredAlertInstanceData(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(1), n)
}
