package resource

import "github.com/StratoAPI/Interface/plugins"

type Processor interface {
	// Get the store of a resource
	GetStore(resource string) *plugins.Storage

	// Get a list of all resources
	GetResources() []string
}
