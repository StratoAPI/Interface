package schema

var processorInstance Processor

func SetProcessor(processor Processor) {
	processorInstance = processor
}

func GetProcessor() Processor {
	if processorInstance == nil {
		panic("schema processor is nil")
	}

	return processorInstance
}
