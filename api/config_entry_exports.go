package api

import "encoding/json"

// ServiceExportsConfigEntry manages the exported services for a single admin partition.
// Admin Partitions are a Consul Enterprise feature.
type ServiceExportsConfigEntry struct {
	// Partition is the partition the ServiceExportsConfigEntry applies to.
	// Partitioning is a Consul Enterprise feature.
	Partition string `json:",omitempty"`

	// Services is a list of services to be exported and the list of partitions
	// to expose them to.
	Services []ExportedService

	Meta map[string]string `json:",omitempty"`

	// CreateIndex is the Raft index this entry was created at. This is a
	// read-only field.
	CreateIndex uint64

	// ModifyIndex is used for the Check-And-Set operations and can also be fed
	// back into the WaitIndex of the QueryOptions in order to perform blocking
	// queries.
	ModifyIndex uint64
}

// ExportedService manages the exporting of a service in the local partition to
// other partitions.
type ExportedService struct {
	// Name is the name of the service to be exported.
	Name string

	// Namespace is the namespace to export the service from.
	Namespace string `json:",omitempty"`

	// Consumers is a list of downstream consumers of the service to be exported.
	Consumers []ServiceConsumer
}

// ServiceConsumer represents a downstream consumer of the service to be exported.
type ServiceConsumer struct {
	// Partition is the admin partition to export the service to.
	Partition string
}

func (e *ServiceExportsConfigEntry) GetKind() string            { return ServiceExports }
func (e *ServiceExportsConfigEntry) GetName() string            { return e.Partition }
func (e *ServiceExportsConfigEntry) GetPartition() string       { return e.Partition }
func (e *ServiceExportsConfigEntry) GetNamespace() string       { return IntentionDefaultNamespace }
func (e *ServiceExportsConfigEntry) GetMeta() map[string]string { return e.Meta }
func (e *ServiceExportsConfigEntry) GetCreateIndex() uint64     { return e.CreateIndex }
func (e *ServiceExportsConfigEntry) GetModifyIndex() uint64     { return e.ModifyIndex }

// MarshalJSON adds the Kind field so that the JSON can be decoded back into the
// correct type.
func (e *ServiceExportsConfigEntry) MarshalJSON() ([]byte, error) {
	type Alias ServiceExportsConfigEntry
	source := &struct {
		Kind string
		*Alias
	}{
		Kind:  ServiceExports,
		Alias: (*Alias)(e),
	}
	return json.Marshal(source)
}
