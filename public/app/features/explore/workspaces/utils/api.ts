/*

    const responseObservable = getBackendSrv().fetch<DataQuerySpecResponse>({
        url: BASE_URL
    })
    const response = await lastValueFrom(responseObservable);
    const data = response.data;

 */

import { lastValueFrom } from 'rxjs';

import { FetchResponse, getBackendSrv } from '@grafana/runtime';

import {
  CreateExploreWorkspaceCommand,
  CreateExploreWorkspaceResponse,
  CreateExploreWorkspaceSnapshotCommand,
  CreateExploreWorkspaceSnapshotResponse,
  GetExploreWorkspaceCommand,
  GetExploreWorkspaceResponse,
  GetExploreWorkspacesCommand,
  GetExploreWorkspaceSnapshotCommand,
  GetExploreWorkspaceSnapshotResponse,
  GetExploreWorkspaceSnapshotsCommand,
  GetExploreWorkspaceSnapshotsResponse,
  GetExploreWorkspacesResponse,
  UpdateExploreWorkspaceLatestSnapshotCommand,
  UpdateExploreWorkspaceLatestSnapshotResponse,
} from '../types';

type RequestOptions<Command, Response> = {
  url(command: Command): string;
  body?(command: Command): Object;
  method: string;
  parseResponse?: (response: FetchResponse<Response>) => Response;
};

type ApiCall<Command, Response> = (command: Command) => Promise<Response>;

function apiCall<Command, Response>(options: RequestOptions<Command, Response>): ApiCall<Command, Response> {
  return async function (command: Command) {
    const responseObservable = getBackendSrv().fetch<Response>({
      url: options.url(command),
      method: options.method,
      data: options.method === 'PUT' || options.method === 'POST' ? command : undefined,
    });
    const response = await lastValueFrom(responseObservable);
    return options.parseResponse ? options.parseResponse(response) : response.data;
  };
}

export const getExploreWorkspaces = apiCall<GetExploreWorkspacesCommand, GetExploreWorkspacesResponse>({
  method: 'GET',
  url: () => '/api/exploreworkspaces/',
});

export const getExploreWorkspace = apiCall<GetExploreWorkspaceCommand, GetExploreWorkspaceResponse>({
  method: 'GET',
  url: (command) => `/api/exploreworkspaces/${command.exploreWorkspaceUID}`,
});

export const createExploreWorkspace = apiCall<CreateExploreWorkspaceCommand, CreateExploreWorkspaceResponse>({
  method: 'POST',
  url: () => '/api/exploreworkspaces/',
});

export const updateExploreWorkspaceLatestSnapshot = apiCall<
  UpdateExploreWorkspaceLatestSnapshotCommand,
  UpdateExploreWorkspaceLatestSnapshotResponse
>({
  method: 'POST',
  url: (command) => '/api/exploreworkspaces/' + command.exploreWorkspaceUID,
});

export const createExploreWorkspaceSnapshot = apiCall<
  CreateExploreWorkspaceSnapshotCommand,
  CreateExploreWorkspaceSnapshotResponse
>({
  method: 'POST',
  url: (command) => `/api/exploreworkspaces/${command.exploreWorkspaceUID}/snapshot`,
});

export const getExploreWorkspaceSnapshot = apiCall<
  GetExploreWorkspaceSnapshotCommand,
  GetExploreWorkspaceSnapshotResponse
>({
  method: 'GET',
  url: (command) => `/api/exploreworkspaces/snapshot/${command.uid}`,
});

export const getExploreWorkspaceSnapshots = apiCall<
  GetExploreWorkspaceSnapshotsCommand,
  GetExploreWorkspaceSnapshotsResponse
>({
  method: 'GET',
  url: (command) => `/api/exploreworkspaces/${command.exploreWorkspaceUid}/snapshots`,
});