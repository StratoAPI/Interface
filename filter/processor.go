package filter

var processorInstance Processor

func SetProcessor(processor Processor) {
	processorInstance = processor
}

func GetProcessor() Processor {
	if processorInstance == nil {
		panic("filter processor is nil")
	}

	return processorInstance
}
