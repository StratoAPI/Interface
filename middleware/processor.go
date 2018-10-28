package middleware

var processorInstance Processor

func SetProcessor(processor Processor) {
	processorInstance = processor
}

func GetProcessor() Processor {
	if processorInstance == nil {
		panic("middleware processor is nil")
	}

	return processorInstance
}
