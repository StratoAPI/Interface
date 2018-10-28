package middleware

type Processor interface {
	// Validate request
	// If returns nil, it will pass the request onwards
	// If returns a pointer, request will get replied to with provided response
	Request(resource string, headers map[string][]string) *RequestResponse

	// Validate and filter response
	// If returns nil, it will pass the request onwards with the returned data
	// If returns a pointer, request will get replied to with provided response without data
	Response(resource string, headers map[string][]string, data []map[string]interface{}) ([]map[string]interface{}, *RequestResponse)
}

type RequestResponse struct {
	// Return status code (http or otherwise)
	Code int

	// Return message
	Message string
}
