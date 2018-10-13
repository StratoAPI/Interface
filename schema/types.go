package schema

type Processor interface {
	// Get the store of a resource
	ResourceExists(resource string) bool

	// Get a list of all resources
	ResourceValid(resource string, data string) (bool, error)

	GetSchema(resource string) *Schema
}

type Schema struct {
	Source map[string]interface{}
	Meta   ResourceMeta
}

type ResourceMeta struct {
	Resource string `json:"resource"`
	Store    string `json:"store"`
}
