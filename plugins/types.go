package plugins

import (
	"github.com/StratoAPI/Interface/filter"
	"github.com/StratoAPI/Interface/middleware"
)

type Plugin interface {
	// The name of the plugin
	Name() string

	// The entrypoint function
	Entrypoint()
}

type Facade interface {
	// Initialize the facade.
	Initialize() error

	// Start the facade. Must be a blocking call.
	Start() error

	// Graceful stopping of the facade with a 30s timeout.
	Stop() error
}

type Storage interface {
	// Initialize the storage.
	Initialize() error

	// Start the storage. Must be a blocking call.
	Start() error

	// Graceful stopping of the storage with a 30s timeout.
	Stop() error

	// Retrieve resources.
	GetResources(resource string, filters []filter.ProcessedFilter) ([]map[string]interface{}, error)

	// Create resources.
	CreateResources(resource string, data []map[string]interface{}) error

	// Update resources.
	UpdateResources(resource string, data map[string]interface{}, filters []filter.ProcessedFilter) error

	// Delete resources.
	DeleteResources(resource string, filters []filter.ProcessedFilter) error
}

type Filter interface {
	// Initialize the filter.
	Initialize() error

	// Start the filter. Does not have to be blocking.
	Start() error

	// Graceful stopping of the filter with a 30s timeout.
	Stop() error

	// Validate structure for filter validness
	ValidateFilter(filter filter.ProcessedFilter) (bool, error)

	// Create a new instance of the filter
	CreateFilter(filter string) (interface{}, error)
}

type Middleware interface {
	// Initialize the middleware.
	Initialize() error

	// Start the middleware. Does not have to be blocking.
	Start() error

	// Graceful stopping of the middleware with a 30s timeout.
	Stop() error

	// Validate request
	// If returns nil, it will pass the request onwards
	// If returns a pointer, request will get replied to with provided response
	Request(resource string, headers map[string][]string, config interface{}) *middleware.RequestResponse

	// Validate and filter response
	// If returns nil, it will pass the request onwards with the returned data
	// If returns a pointer, request will get replied to with provided response without data
	Response(resource string, headers map[string][]string, data []map[string]interface{}, config interface{}) ([]map[string]interface{}, *middleware.RequestResponse)
}

type Registry interface {
	// Register a facade
	RegisterFacade(name string, facade Facade) error

	// Register a store
	RegisterStorage(name string, storage Storage) error

	// Register a filter
	RegisterFilter(name string, filter Filter) error

	// Register a middleware
	RegisterMiddleware(name string, filter Middleware) error

	// Associate a filter with a store
	AssociateFilter(filter string, storage string) error
}
