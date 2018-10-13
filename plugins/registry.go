package plugins

var registryInstance Registry

func SetRegistry(registry Registry) {
	registryInstance = registry
}

func GetRegistry() Registry {
	if registryInstance == nil {
		panic("registry is nil")
	}

	return registryInstance
}
