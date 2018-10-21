package filter

import (
	"encoding/json"
)

type Processor interface {
	// Check if a filter exists
	FilterExists(filter string) bool

	// Create a new instance of the filter
	CreateFilter(filter string) interface{}

	// Validate structure for filter validness
	ValidateFilter(filter ProcessedFilter) (bool, error)
}

type EncodedFilter struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type ProcessedFilter struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
