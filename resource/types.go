package resource

import (
	"github.com/StratoAPI/Interface/filter"
	"github.com/StratoAPI/Interface/plugins"
)

type Processor interface {
	// Get the store of a resource
	GetStore(resource string) *plugins.Storage

	// Get a list of all resources
	GetResourceList() []string

	// Retrieve resources.
	GetResources(resource string, filters []filter.ProcessedFilter) ([]map[string]interface{}, error)

	// Create resources.
	CreateResources(resource string, data []map[string]interface{}) error

	// Update resources.
	UpdateResources(resource string, data map[string]interface{}, filters []filter.ProcessedFilter) error

	// Delete resources.
	DeleteResources(resource string, filters []filter.ProcessedFilter) error
}
