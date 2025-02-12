// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type UpdatePostgresConfigBodySessionReplicationRole string

const (
	UpdatePostgresConfigBodySessionReplicationRoleOrigin  UpdatePostgresConfigBodySessionReplicationRole = "origin"
	UpdatePostgresConfigBodySessionReplicationRoleReplica UpdatePostgresConfigBodySessionReplicationRole = "replica"
	UpdatePostgresConfigBodySessionReplicationRoleLocal   UpdatePostgresConfigBodySessionReplicationRole = "local"
)

func (e UpdatePostgresConfigBodySessionReplicationRole) ToPointer() *UpdatePostgresConfigBodySessionReplicationRole {
	return &e
}
func (e *UpdatePostgresConfigBodySessionReplicationRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "origin":
		fallthrough
	case "replica":
		fallthrough
	case "local":
		*e = UpdatePostgresConfigBodySessionReplicationRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdatePostgresConfigBodySessionReplicationRole: %v", v)
	}
}

type UpdatePostgresConfigBody struct {
	EffectiveCacheSize            *string                                         `json:"effective_cache_size,omitempty"`
	LogicalDecodingWorkMem        *string                                         `json:"logical_decoding_work_mem,omitempty"`
	MaintenanceWorkMem            *string                                         `json:"maintenance_work_mem,omitempty"`
	MaxConnections                *int64                                          `json:"max_connections,omitempty"`
	MaxLocksPerTransaction        *int64                                          `json:"max_locks_per_transaction,omitempty"`
	MaxParallelMaintenanceWorkers *int64                                          `json:"max_parallel_maintenance_workers,omitempty"`
	MaxParallelWorkers            *int64                                          `json:"max_parallel_workers,omitempty"`
	MaxParallelWorkersPerGather   *int64                                          `json:"max_parallel_workers_per_gather,omitempty"`
	MaxReplicationSlots           *int64                                          `json:"max_replication_slots,omitempty"`
	MaxSlotWalKeepSize            *string                                         `json:"max_slot_wal_keep_size,omitempty"`
	MaxStandbyArchiveDelay        *string                                         `json:"max_standby_archive_delay,omitempty"`
	MaxStandbyStreamingDelay      *string                                         `json:"max_standby_streaming_delay,omitempty"`
	MaxWalSize                    *string                                         `json:"max_wal_size,omitempty"`
	MaxWalSenders                 *int64                                          `json:"max_wal_senders,omitempty"`
	MaxWorkerProcesses            *int64                                          `json:"max_worker_processes,omitempty"`
	SharedBuffers                 *string                                         `json:"shared_buffers,omitempty"`
	StatementTimeout              *string                                         `json:"statement_timeout,omitempty"`
	TrackCommitTimestamp          *bool                                           `json:"track_commit_timestamp,omitempty"`
	WalKeepSize                   *string                                         `json:"wal_keep_size,omitempty"`
	WalSenderTimeout              *string                                         `json:"wal_sender_timeout,omitempty"`
	WorkMem                       *string                                         `json:"work_mem,omitempty"`
	RestartDatabase               *bool                                           `json:"restart_database,omitempty"`
	SessionReplicationRole        *UpdatePostgresConfigBodySessionReplicationRole `json:"session_replication_role,omitempty"`
}

func (o *UpdatePostgresConfigBody) GetEffectiveCacheSize() *string {
	if o == nil {
		return nil
	}
	return o.EffectiveCacheSize
}

func (o *UpdatePostgresConfigBody) GetLogicalDecodingWorkMem() *string {
	if o == nil {
		return nil
	}
	return o.LogicalDecodingWorkMem
}

func (o *UpdatePostgresConfigBody) GetMaintenanceWorkMem() *string {
	if o == nil {
		return nil
	}
	return o.MaintenanceWorkMem
}

func (o *UpdatePostgresConfigBody) GetMaxConnections() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxConnections
}

func (o *UpdatePostgresConfigBody) GetMaxLocksPerTransaction() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxLocksPerTransaction
}

func (o *UpdatePostgresConfigBody) GetMaxParallelMaintenanceWorkers() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxParallelMaintenanceWorkers
}

func (o *UpdatePostgresConfigBody) GetMaxParallelWorkers() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxParallelWorkers
}

func (o *UpdatePostgresConfigBody) GetMaxParallelWorkersPerGather() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxParallelWorkersPerGather
}

func (o *UpdatePostgresConfigBody) GetMaxReplicationSlots() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxReplicationSlots
}

func (o *UpdatePostgresConfigBody) GetMaxSlotWalKeepSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSlotWalKeepSize
}

func (o *UpdatePostgresConfigBody) GetMaxStandbyArchiveDelay() *string {
	if o == nil {
		return nil
	}
	return o.MaxStandbyArchiveDelay
}

func (o *UpdatePostgresConfigBody) GetMaxStandbyStreamingDelay() *string {
	if o == nil {
		return nil
	}
	return o.MaxStandbyStreamingDelay
}

func (o *UpdatePostgresConfigBody) GetMaxWalSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxWalSize
}

func (o *UpdatePostgresConfigBody) GetMaxWalSenders() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxWalSenders
}

func (o *UpdatePostgresConfigBody) GetMaxWorkerProcesses() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxWorkerProcesses
}

func (o *UpdatePostgresConfigBody) GetSharedBuffers() *string {
	if o == nil {
		return nil
	}
	return o.SharedBuffers
}

func (o *UpdatePostgresConfigBody) GetStatementTimeout() *string {
	if o == nil {
		return nil
	}
	return o.StatementTimeout
}

func (o *UpdatePostgresConfigBody) GetTrackCommitTimestamp() *bool {
	if o == nil {
		return nil
	}
	return o.TrackCommitTimestamp
}

func (o *UpdatePostgresConfigBody) GetWalKeepSize() *string {
	if o == nil {
		return nil
	}
	return o.WalKeepSize
}

func (o *UpdatePostgresConfigBody) GetWalSenderTimeout() *string {
	if o == nil {
		return nil
	}
	return o.WalSenderTimeout
}

func (o *UpdatePostgresConfigBody) GetWorkMem() *string {
	if o == nil {
		return nil
	}
	return o.WorkMem
}

func (o *UpdatePostgresConfigBody) GetRestartDatabase() *bool {
	if o == nil {
		return nil
	}
	return o.RestartDatabase
}

func (o *UpdatePostgresConfigBody) GetSessionReplicationRole() *UpdatePostgresConfigBodySessionReplicationRole {
	if o == nil {
		return nil
	}
	return o.SessionReplicationRole
}
