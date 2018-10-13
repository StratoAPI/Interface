package resource

var processorInstance Processor

func SetProcessor(processor Processor) {
	processorInstance = processor
}

func GetProcessor() Processor {
	if processorInstance == nil {
		panic("resource processor is nil")
	}

	return processorInstance
}
