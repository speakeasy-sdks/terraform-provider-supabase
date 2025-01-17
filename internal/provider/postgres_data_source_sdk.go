// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/speakeasy/terraform-provider-supabase/internal/sdk/models/shared"
)

func (r *PostgresDataSourceModel) RefreshFromSharedPostgresConfigResponse(resp *shared.PostgresConfigResponse) {
	if resp != nil {
		r.EffectiveCacheSize = types.StringPointerValue(resp.EffectiveCacheSize)
		r.LogicalDecodingWorkMem = types.StringPointerValue(resp.LogicalDecodingWorkMem)
		r.MaintenanceWorkMem = types.StringPointerValue(resp.MaintenanceWorkMem)
		r.MaxConnections = types.Int64PointerValue(resp.MaxConnections)
		r.MaxLocksPerTransaction = types.Int64PointerValue(resp.MaxLocksPerTransaction)
		r.MaxParallelMaintenanceWorkers = types.Int64PointerValue(resp.MaxParallelMaintenanceWorkers)
		r.MaxParallelWorkers = types.Int64PointerValue(resp.MaxParallelWorkers)
		r.MaxParallelWorkersPerGather = types.Int64PointerValue(resp.MaxParallelWorkersPerGather)
		r.MaxReplicationSlots = types.Int64PointerValue(resp.MaxReplicationSlots)
		r.MaxSlotWalKeepSize = types.StringPointerValue(resp.MaxSlotWalKeepSize)
		r.MaxStandbyArchiveDelay = types.StringPointerValue(resp.MaxStandbyArchiveDelay)
		r.MaxStandbyStreamingDelay = types.StringPointerValue(resp.MaxStandbyStreamingDelay)
		r.MaxWalSenders = types.Int64PointerValue(resp.MaxWalSenders)
		r.MaxWalSize = types.StringPointerValue(resp.MaxWalSize)
		r.MaxWorkerProcesses = types.Int64PointerValue(resp.MaxWorkerProcesses)
		if resp.SessionReplicationRole != nil {
			r.SessionReplicationRole = types.StringValue(string(*resp.SessionReplicationRole))
		} else {
			r.SessionReplicationRole = types.StringNull()
		}
		r.SharedBuffers = types.StringPointerValue(resp.SharedBuffers)
		r.StatementTimeout = types.StringPointerValue(resp.StatementTimeout)
		r.TrackCommitTimestamp = types.BoolPointerValue(resp.TrackCommitTimestamp)
		r.WalKeepSize = types.StringPointerValue(resp.WalKeepSize)
		r.WalSenderTimeout = types.StringPointerValue(resp.WalSenderTimeout)
		r.WorkMem = types.StringPointerValue(resp.WorkMem)
	}
}
