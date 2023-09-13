// Code generated by grpcvtctldclient-generator. DO NOT EDIT.

/*
Copyright 2021 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package grpcvtctldclient

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	vtctldatapb "vitess.io/vitess/go/vt/proto/vtctldata"
	vtctlservicepb "vitess.io/vitess/go/vt/proto/vtctlservice"
)

// AddCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) AddCellInfo(ctx context.Context, in *vtctldatapb.AddCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.AddCellInfoResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.AddCellInfo(ctx, in, opts...)
}

// AddCellsAlias is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) AddCellsAlias(ctx context.Context, in *vtctldatapb.AddCellsAliasRequest, opts ...grpc.CallOption) (*vtctldatapb.AddCellsAliasResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.AddCellsAlias(ctx, in, opts...)
}

// ApplyRoutingRules is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ApplyRoutingRules(ctx context.Context, in *vtctldatapb.ApplyRoutingRulesRequest, opts ...grpc.CallOption) (*vtctldatapb.ApplyRoutingRulesResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ApplyRoutingRules(ctx, in, opts...)
}

// ApplySchema is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ApplySchema(ctx context.Context, in *vtctldatapb.ApplySchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.ApplySchemaResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ApplySchema(ctx, in, opts...)
}

// ApplyShardRoutingRules is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ApplyShardRoutingRules(ctx context.Context, in *vtctldatapb.ApplyShardRoutingRulesRequest, opts ...grpc.CallOption) (*vtctldatapb.ApplyShardRoutingRulesResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ApplyShardRoutingRules(ctx, in, opts...)
}

// ApplyVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ApplyVSchema(ctx context.Context, in *vtctldatapb.ApplyVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.ApplyVSchemaResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ApplyVSchema(ctx, in, opts...)
}

// Backup is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) Backup(ctx context.Context, in *vtctldatapb.BackupRequest, opts ...grpc.CallOption) (vtctlservicepb.Vtctld_BackupClient, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.Backup(ctx, in, opts...)
}

// BackupShard is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) BackupShard(ctx context.Context, in *vtctldatapb.BackupShardRequest, opts ...grpc.CallOption) (vtctlservicepb.Vtctld_BackupShardClient, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.BackupShard(ctx, in, opts...)
}

// CancelSchemaMigration is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) CancelSchemaMigration(ctx context.Context, in *vtctldatapb.CancelSchemaMigrationRequest, opts ...grpc.CallOption) (*vtctldatapb.CancelSchemaMigrationResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.CancelSchemaMigration(ctx, in, opts...)
}

// ChangeTabletType is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ChangeTabletType(ctx context.Context, in *vtctldatapb.ChangeTabletTypeRequest, opts ...grpc.CallOption) (*vtctldatapb.ChangeTabletTypeResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ChangeTabletType(ctx, in, opts...)
}

// CleanupSchemaMigration is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) CleanupSchemaMigration(ctx context.Context, in *vtctldatapb.CleanupSchemaMigrationRequest, opts ...grpc.CallOption) (*vtctldatapb.CleanupSchemaMigrationResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.CleanupSchemaMigration(ctx, in, opts...)
}

// CompleteSchemaMigration is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) CompleteSchemaMigration(ctx context.Context, in *vtctldatapb.CompleteSchemaMigrationRequest, opts ...grpc.CallOption) (*vtctldatapb.CompleteSchemaMigrationResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.CompleteSchemaMigration(ctx, in, opts...)
}

// CreateKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) CreateKeyspace(ctx context.Context, in *vtctldatapb.CreateKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.CreateKeyspaceResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.CreateKeyspace(ctx, in, opts...)
}

// CreateShard is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) CreateShard(ctx context.Context, in *vtctldatapb.CreateShardRequest, opts ...grpc.CallOption) (*vtctldatapb.CreateShardResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.CreateShard(ctx, in, opts...)
}

// DeleteCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) DeleteCellInfo(ctx context.Context, in *vtctldatapb.DeleteCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteCellInfoResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.DeleteCellInfo(ctx, in, opts...)
}

// DeleteCellsAlias is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) DeleteCellsAlias(ctx context.Context, in *vtctldatapb.DeleteCellsAliasRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteCellsAliasResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.DeleteCellsAlias(ctx, in, opts...)
}

// DeleteKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) DeleteKeyspace(ctx context.Context, in *vtctldatapb.DeleteKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteKeyspaceResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.DeleteKeyspace(ctx, in, opts...)
}

// DeleteShards is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) DeleteShards(ctx context.Context, in *vtctldatapb.DeleteShardsRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteShardsResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.DeleteShards(ctx, in, opts...)
}

// DeleteSrvVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) DeleteSrvVSchema(ctx context.Context, in *vtctldatapb.DeleteSrvVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteSrvVSchemaResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.DeleteSrvVSchema(ctx, in, opts...)
}

// DeleteTablets is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) DeleteTablets(ctx context.Context, in *vtctldatapb.DeleteTabletsRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteTabletsResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.DeleteTablets(ctx, in, opts...)
}

// EmergencyReparentShard is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) EmergencyReparentShard(ctx context.Context, in *vtctldatapb.EmergencyReparentShardRequest, opts ...grpc.CallOption) (*vtctldatapb.EmergencyReparentShardResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.EmergencyReparentShard(ctx, in, opts...)
}

// ExecuteFetchAsApp is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ExecuteFetchAsApp(ctx context.Context, in *vtctldatapb.ExecuteFetchAsAppRequest, opts ...grpc.CallOption) (*vtctldatapb.ExecuteFetchAsAppResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ExecuteFetchAsApp(ctx, in, opts...)
}

// ExecuteFetchAsDBA is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ExecuteFetchAsDBA(ctx context.Context, in *vtctldatapb.ExecuteFetchAsDBARequest, opts ...grpc.CallOption) (*vtctldatapb.ExecuteFetchAsDBAResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ExecuteFetchAsDBA(ctx, in, opts...)
}

// ExecuteHook is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ExecuteHook(ctx context.Context, in *vtctldatapb.ExecuteHookRequest, opts ...grpc.CallOption) (*vtctldatapb.ExecuteHookResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ExecuteHook(ctx, in, opts...)
}

// FindAllShardsInKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) FindAllShardsInKeyspace(ctx context.Context, in *vtctldatapb.FindAllShardsInKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.FindAllShardsInKeyspaceResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.FindAllShardsInKeyspace(ctx, in, opts...)
}

// GetBackups is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetBackups(ctx context.Context, in *vtctldatapb.GetBackupsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetBackupsResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetBackups(ctx, in, opts...)
}

// GetCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetCellInfo(ctx context.Context, in *vtctldatapb.GetCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.GetCellInfoResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetCellInfo(ctx, in, opts...)
}

// GetCellInfoNames is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetCellInfoNames(ctx context.Context, in *vtctldatapb.GetCellInfoNamesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetCellInfoNamesResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetCellInfoNames(ctx, in, opts...)
}

// GetCellsAliases is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetCellsAliases(ctx context.Context, in *vtctldatapb.GetCellsAliasesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetCellsAliasesResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetCellsAliases(ctx, in, opts...)
}

// GetFullStatus is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetFullStatus(ctx context.Context, in *vtctldatapb.GetFullStatusRequest, opts ...grpc.CallOption) (*vtctldatapb.GetFullStatusResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetFullStatus(ctx, in, opts...)
}

// GetKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetKeyspace(ctx context.Context, in *vtctldatapb.GetKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.GetKeyspaceResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetKeyspace(ctx, in, opts...)
}

// GetKeyspaces is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetKeyspaces(ctx context.Context, in *vtctldatapb.GetKeyspacesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetKeyspacesResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetKeyspaces(ctx, in, opts...)
}

// GetPermissions is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetPermissions(ctx context.Context, in *vtctldatapb.GetPermissionsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetPermissionsResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetPermissions(ctx, in, opts...)
}

// GetRoutingRules is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetRoutingRules(ctx context.Context, in *vtctldatapb.GetRoutingRulesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetRoutingRulesResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetRoutingRules(ctx, in, opts...)
}

// GetSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetSchema(ctx context.Context, in *vtctldatapb.GetSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSchemaResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetSchema(ctx, in, opts...)
}

// GetSchemaMigrations is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetSchemaMigrations(ctx context.Context, in *vtctldatapb.GetSchemaMigrationsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSchemaMigrationsResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetSchemaMigrations(ctx, in, opts...)
}

// GetShard is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetShard(ctx context.Context, in *vtctldatapb.GetShardRequest, opts ...grpc.CallOption) (*vtctldatapb.GetShardResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetShard(ctx, in, opts...)
}

// GetShardRoutingRules is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetShardRoutingRules(ctx context.Context, in *vtctldatapb.GetShardRoutingRulesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetShardRoutingRulesResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetShardRoutingRules(ctx, in, opts...)
}

// GetSrvKeyspaceNames is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetSrvKeyspaceNames(ctx context.Context, in *vtctldatapb.GetSrvKeyspaceNamesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvKeyspaceNamesResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetSrvKeyspaceNames(ctx, in, opts...)
}

// GetSrvKeyspaces is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetSrvKeyspaces(ctx context.Context, in *vtctldatapb.GetSrvKeyspacesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvKeyspacesResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetSrvKeyspaces(ctx, in, opts...)
}

// GetSrvVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetSrvVSchema(ctx context.Context, in *vtctldatapb.GetSrvVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvVSchemaResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetSrvVSchema(ctx, in, opts...)
}

// GetSrvVSchemas is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetSrvVSchemas(ctx context.Context, in *vtctldatapb.GetSrvVSchemasRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvVSchemasResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetSrvVSchemas(ctx, in, opts...)
}

// GetTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetTablet(ctx context.Context, in *vtctldatapb.GetTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.GetTabletResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetTablet(ctx, in, opts...)
}

// GetTablets is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetTablets(ctx context.Context, in *vtctldatapb.GetTabletsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetTabletsResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetTablets(ctx, in, opts...)
}

// GetTopologyPath is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetTopologyPath(ctx context.Context, in *vtctldatapb.GetTopologyPathRequest, opts ...grpc.CallOption) (*vtctldatapb.GetTopologyPathResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetTopologyPath(ctx, in, opts...)
}

// GetVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetVSchema(ctx context.Context, in *vtctldatapb.GetVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.GetVSchemaResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetVSchema(ctx, in, opts...)
}

// GetVersion is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetVersion(ctx context.Context, in *vtctldatapb.GetVersionRequest, opts ...grpc.CallOption) (*vtctldatapb.GetVersionResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetVersion(ctx, in, opts...)
}

// GetWorkflows is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) GetWorkflows(ctx context.Context, in *vtctldatapb.GetWorkflowsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetWorkflowsResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.GetWorkflows(ctx, in, opts...)
}

// InitShardPrimary is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) InitShardPrimary(ctx context.Context, in *vtctldatapb.InitShardPrimaryRequest, opts ...grpc.CallOption) (*vtctldatapb.InitShardPrimaryResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.InitShardPrimary(ctx, in, opts...)
}

// LaunchSchemaMigration is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) LaunchSchemaMigration(ctx context.Context, in *vtctldatapb.LaunchSchemaMigrationRequest, opts ...grpc.CallOption) (*vtctldatapb.LaunchSchemaMigrationResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.LaunchSchemaMigration(ctx, in, opts...)
}

// MoveTablesComplete is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) MoveTablesComplete(ctx context.Context, in *vtctldatapb.MoveTablesCompleteRequest, opts ...grpc.CallOption) (*vtctldatapb.MoveTablesCompleteResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.MoveTablesComplete(ctx, in, opts...)
}

// MoveTablesCreate is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) MoveTablesCreate(ctx context.Context, in *vtctldatapb.MoveTablesCreateRequest, opts ...grpc.CallOption) (*vtctldatapb.WorkflowStatusResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.MoveTablesCreate(ctx, in, opts...)
}

// PingTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) PingTablet(ctx context.Context, in *vtctldatapb.PingTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.PingTabletResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.PingTablet(ctx, in, opts...)
}

// PlannedReparentShard is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) PlannedReparentShard(ctx context.Context, in *vtctldatapb.PlannedReparentShardRequest, opts ...grpc.CallOption) (*vtctldatapb.PlannedReparentShardResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.PlannedReparentShard(ctx, in, opts...)
}

// RebuildKeyspaceGraph is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) RebuildKeyspaceGraph(ctx context.Context, in *vtctldatapb.RebuildKeyspaceGraphRequest, opts ...grpc.CallOption) (*vtctldatapb.RebuildKeyspaceGraphResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.RebuildKeyspaceGraph(ctx, in, opts...)
}

// RebuildVSchemaGraph is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) RebuildVSchemaGraph(ctx context.Context, in *vtctldatapb.RebuildVSchemaGraphRequest, opts ...grpc.CallOption) (*vtctldatapb.RebuildVSchemaGraphResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.RebuildVSchemaGraph(ctx, in, opts...)
}

// RefreshState is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) RefreshState(ctx context.Context, in *vtctldatapb.RefreshStateRequest, opts ...grpc.CallOption) (*vtctldatapb.RefreshStateResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.RefreshState(ctx, in, opts...)
}

// RefreshStateByShard is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) RefreshStateByShard(ctx context.Context, in *vtctldatapb.RefreshStateByShardRequest, opts ...grpc.CallOption) (*vtctldatapb.RefreshStateByShardResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.RefreshStateByShard(ctx, in, opts...)
}

// ReloadSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ReloadSchema(ctx context.Context, in *vtctldatapb.ReloadSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.ReloadSchemaResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ReloadSchema(ctx, in, opts...)
}

// ReloadSchemaKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ReloadSchemaKeyspace(ctx context.Context, in *vtctldatapb.ReloadSchemaKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.ReloadSchemaKeyspaceResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ReloadSchemaKeyspace(ctx, in, opts...)
}

// ReloadSchemaShard is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ReloadSchemaShard(ctx context.Context, in *vtctldatapb.ReloadSchemaShardRequest, opts ...grpc.CallOption) (*vtctldatapb.ReloadSchemaShardResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ReloadSchemaShard(ctx, in, opts...)
}

// RemoveBackup is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) RemoveBackup(ctx context.Context, in *vtctldatapb.RemoveBackupRequest, opts ...grpc.CallOption) (*vtctldatapb.RemoveBackupResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.RemoveBackup(ctx, in, opts...)
}

// RemoveKeyspaceCell is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) RemoveKeyspaceCell(ctx context.Context, in *vtctldatapb.RemoveKeyspaceCellRequest, opts ...grpc.CallOption) (*vtctldatapb.RemoveKeyspaceCellResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.RemoveKeyspaceCell(ctx, in, opts...)
}

// RemoveShardCell is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) RemoveShardCell(ctx context.Context, in *vtctldatapb.RemoveShardCellRequest, opts ...grpc.CallOption) (*vtctldatapb.RemoveShardCellResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.RemoveShardCell(ctx, in, opts...)
}

// ReparentTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ReparentTablet(ctx context.Context, in *vtctldatapb.ReparentTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.ReparentTabletResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ReparentTablet(ctx, in, opts...)
}

// ReshardCreate is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ReshardCreate(ctx context.Context, in *vtctldatapb.ReshardCreateRequest, opts ...grpc.CallOption) (*vtctldatapb.WorkflowStatusResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ReshardCreate(ctx, in, opts...)
}

// RestoreFromBackup is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) RestoreFromBackup(ctx context.Context, in *vtctldatapb.RestoreFromBackupRequest, opts ...grpc.CallOption) (vtctlservicepb.Vtctld_RestoreFromBackupClient, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.RestoreFromBackup(ctx, in, opts...)
}

// RetrySchemaMigration is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) RetrySchemaMigration(ctx context.Context, in *vtctldatapb.RetrySchemaMigrationRequest, opts ...grpc.CallOption) (*vtctldatapb.RetrySchemaMigrationResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.RetrySchemaMigration(ctx, in, opts...)
}

// RunHealthCheck is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) RunHealthCheck(ctx context.Context, in *vtctldatapb.RunHealthCheckRequest, opts ...grpc.CallOption) (*vtctldatapb.RunHealthCheckResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.RunHealthCheck(ctx, in, opts...)
}

// SetKeyspaceDurabilityPolicy is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) SetKeyspaceDurabilityPolicy(ctx context.Context, in *vtctldatapb.SetKeyspaceDurabilityPolicyRequest, opts ...grpc.CallOption) (*vtctldatapb.SetKeyspaceDurabilityPolicyResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.SetKeyspaceDurabilityPolicy(ctx, in, opts...)
}

// SetShardIsPrimaryServing is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) SetShardIsPrimaryServing(ctx context.Context, in *vtctldatapb.SetShardIsPrimaryServingRequest, opts ...grpc.CallOption) (*vtctldatapb.SetShardIsPrimaryServingResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.SetShardIsPrimaryServing(ctx, in, opts...)
}

// SetShardTabletControl is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) SetShardTabletControl(ctx context.Context, in *vtctldatapb.SetShardTabletControlRequest, opts ...grpc.CallOption) (*vtctldatapb.SetShardTabletControlResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.SetShardTabletControl(ctx, in, opts...)
}

// SetWritable is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) SetWritable(ctx context.Context, in *vtctldatapb.SetWritableRequest, opts ...grpc.CallOption) (*vtctldatapb.SetWritableResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.SetWritable(ctx, in, opts...)
}

// ShardReplicationAdd is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ShardReplicationAdd(ctx context.Context, in *vtctldatapb.ShardReplicationAddRequest, opts ...grpc.CallOption) (*vtctldatapb.ShardReplicationAddResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ShardReplicationAdd(ctx, in, opts...)
}

// ShardReplicationFix is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ShardReplicationFix(ctx context.Context, in *vtctldatapb.ShardReplicationFixRequest, opts ...grpc.CallOption) (*vtctldatapb.ShardReplicationFixResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ShardReplicationFix(ctx, in, opts...)
}

// ShardReplicationPositions is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ShardReplicationPositions(ctx context.Context, in *vtctldatapb.ShardReplicationPositionsRequest, opts ...grpc.CallOption) (*vtctldatapb.ShardReplicationPositionsResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ShardReplicationPositions(ctx, in, opts...)
}

// ShardReplicationRemove is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ShardReplicationRemove(ctx context.Context, in *vtctldatapb.ShardReplicationRemoveRequest, opts ...grpc.CallOption) (*vtctldatapb.ShardReplicationRemoveResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ShardReplicationRemove(ctx, in, opts...)
}

// SleepTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) SleepTablet(ctx context.Context, in *vtctldatapb.SleepTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.SleepTabletResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.SleepTablet(ctx, in, opts...)
}

// SourceShardAdd is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) SourceShardAdd(ctx context.Context, in *vtctldatapb.SourceShardAddRequest, opts ...grpc.CallOption) (*vtctldatapb.SourceShardAddResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.SourceShardAdd(ctx, in, opts...)
}

// SourceShardDelete is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) SourceShardDelete(ctx context.Context, in *vtctldatapb.SourceShardDeleteRequest, opts ...grpc.CallOption) (*vtctldatapb.SourceShardDeleteResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.SourceShardDelete(ctx, in, opts...)
}

// StartReplication is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) StartReplication(ctx context.Context, in *vtctldatapb.StartReplicationRequest, opts ...grpc.CallOption) (*vtctldatapb.StartReplicationResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.StartReplication(ctx, in, opts...)
}

// StopReplication is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) StopReplication(ctx context.Context, in *vtctldatapb.StopReplicationRequest, opts ...grpc.CallOption) (*vtctldatapb.StopReplicationResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.StopReplication(ctx, in, opts...)
}

// TabletExternallyReparented is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) TabletExternallyReparented(ctx context.Context, in *vtctldatapb.TabletExternallyReparentedRequest, opts ...grpc.CallOption) (*vtctldatapb.TabletExternallyReparentedResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.TabletExternallyReparented(ctx, in, opts...)
}

// UpdateCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) UpdateCellInfo(ctx context.Context, in *vtctldatapb.UpdateCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.UpdateCellInfoResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.UpdateCellInfo(ctx, in, opts...)
}

// UpdateCellsAlias is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) UpdateCellsAlias(ctx context.Context, in *vtctldatapb.UpdateCellsAliasRequest, opts ...grpc.CallOption) (*vtctldatapb.UpdateCellsAliasResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.UpdateCellsAlias(ctx, in, opts...)
}

// UpdateThrottlerConfig is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) UpdateThrottlerConfig(ctx context.Context, in *vtctldatapb.UpdateThrottlerConfigRequest, opts ...grpc.CallOption) (*vtctldatapb.UpdateThrottlerConfigResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.UpdateThrottlerConfig(ctx, in, opts...)
}

// Validate is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) Validate(ctx context.Context, in *vtctldatapb.ValidateRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.Validate(ctx, in, opts...)
}

// ValidateKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ValidateKeyspace(ctx context.Context, in *vtctldatapb.ValidateKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateKeyspaceResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ValidateKeyspace(ctx, in, opts...)
}

// ValidateSchemaKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ValidateSchemaKeyspace(ctx context.Context, in *vtctldatapb.ValidateSchemaKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateSchemaKeyspaceResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ValidateSchemaKeyspace(ctx, in, opts...)
}

// ValidateShard is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ValidateShard(ctx context.Context, in *vtctldatapb.ValidateShardRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateShardResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ValidateShard(ctx, in, opts...)
}

// ValidateVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ValidateVSchema(ctx context.Context, in *vtctldatapb.ValidateVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateVSchemaResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ValidateVSchema(ctx, in, opts...)
}

// ValidateVersionKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ValidateVersionKeyspace(ctx context.Context, in *vtctldatapb.ValidateVersionKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateVersionKeyspaceResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ValidateVersionKeyspace(ctx, in, opts...)
}

// ValidateVersionShard is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) ValidateVersionShard(ctx context.Context, in *vtctldatapb.ValidateVersionShardRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateVersionShardResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.ValidateVersionShard(ctx, in, opts...)
}

// WorkflowDelete is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) WorkflowDelete(ctx context.Context, in *vtctldatapb.WorkflowDeleteRequest, opts ...grpc.CallOption) (*vtctldatapb.WorkflowDeleteResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.WorkflowDelete(ctx, in, opts...)
}

// WorkflowStatus is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) WorkflowStatus(ctx context.Context, in *vtctldatapb.WorkflowStatusRequest, opts ...grpc.CallOption) (*vtctldatapb.WorkflowStatusResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.WorkflowStatus(ctx, in, opts...)
}

// WorkflowSwitchTraffic is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) WorkflowSwitchTraffic(ctx context.Context, in *vtctldatapb.WorkflowSwitchTrafficRequest, opts ...grpc.CallOption) (*vtctldatapb.WorkflowSwitchTrafficResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.WorkflowSwitchTraffic(ctx, in, opts...)
}

// WorkflowUpdate is part of the vtctlservicepb.VtctldClient interface.
func (client *gRPCVtctldClient) WorkflowUpdate(ctx context.Context, in *vtctldatapb.WorkflowUpdateRequest, opts ...grpc.CallOption) (*vtctldatapb.WorkflowUpdateResponse, error) {
	if client.c == nil {
		return nil, status.Error(codes.Unavailable, connClosedMsg)
	}

	return client.c.WorkflowUpdate(ctx, in, opts...)
}
