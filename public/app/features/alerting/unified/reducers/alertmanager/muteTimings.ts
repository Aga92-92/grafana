import { createAction, createReducer, isAnyOf } from '@reduxjs/toolkit';
import { remove } from 'lodash';

import { AlertManagerCortexConfig, MuteTimeInterval } from 'app/plugins/datasource/alertmanager/types';

import { mergeTimeIntervals } from '../../components/mute-timings/util';
import { renameMuteTimings } from '../../utils/alertmanager';

export const addMuteTimingAction = createAction<{ interval: MuteTimeInterval }>('muteTiming/add');
export const updateMuteTimingAction = createAction<{
  interval: MuteTimeInterval;
  originalName: string;
}>('muteTiming/update');
export const deleteMuteTimingAction = createAction<{ name: string }>('muteTiming/delete');

const initialState: AlertManagerCortexConfig = {
  alertmanager_config: {},
  template_files: {},
};

/**
 * This reducer will manage action related to mute timings and make sure all operations on the alertmanager
 * configuration happen immutably and only mutate what they need.
 */
export const muteTimingsReducer = createReducer(initialState, (builder) => {
  builder
    // add a mute timing to the alertmanager configuration
    .addCase(addMuteTimingAction, (draft, { payload }) => {
      const { interval } = payload;
      draft.alertmanager_config.time_intervals = (draft.alertmanager_config.time_intervals ?? []).concat(interval);
    })
    // add a mute timing to the alertmanager configuration
    // make sure we update the mute timing in either the deprecated or the new time intervals property
    .addCase(updateMuteTimingAction, (draft, { payload }) => {
      const { interval, originalName } = payload;
      const nameHasChanged = interval.name !== originalName;

      const timeIntervals = draft.alertmanager_config.time_intervals ?? [];
      const muteTimeIntervals = draft.alertmanager_config.mute_time_intervals ?? [];

      const existingIntervalIndex = timeIntervals.findIndex(({ name }) => name === originalName);
      if (existingIntervalIndex !== -1) {
        timeIntervals[existingIntervalIndex] = interval;
      }

      const existingMuteIntervalIndex = muteTimeIntervals.findIndex(({ name }) => name === originalName);
      if (existingMuteIntervalIndex !== -1) {
        muteTimeIntervals[existingMuteIntervalIndex] = interval;
      }

      if (nameHasChanged && draft.alertmanager_config.route) {
        draft.alertmanager_config.route = renameMuteTimings(
          interval.name,
          originalName,
          draft.alertmanager_config.route
        );
      }
    })
    // delete a mute timing from the alertmanager configuration, since the configuration might be using the "deprecated" mute_time_intervals
    // let's also check there
    .addCase(deleteMuteTimingAction, (draft, { payload }) => {
      const { name } = payload;

      remove(draft.alertmanager_config.time_intervals ?? [], (interval) => interval.name === name);
      remove(draft.alertmanager_config.mute_time_intervals ?? [], (interval) => interval.name === name);
    })
    // when either updating or creating a new time interval, check if we don't already have one with that name
    .addMatcher(isAnyOf(addMuteTimingAction, updateMuteTimingAction), (draft, { payload }) => {
      const newName = payload.interval.name;
      const muteTimings = mergeTimeIntervals(draft.alertmanager_config);

      const intervalAlreadyExists = Boolean(muteTimings.find((interval) => interval.name === newName));
      if (intervalAlreadyExists) {
        throw new Error(`Mute timing already exists with name "${newName}"`);
      }
    })
    .addDefaultCase((_state, action) => {
      throw new Error(`Unknown action for mute timing reducer: ${action.type}`);
    });
});
