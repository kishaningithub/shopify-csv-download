package progress

// State comprises of progress state that can be used for displaying the consumer the current state of the operation
type State struct {
	NoOfProductsDownloaded     int
	NoOfProductsConvertedAsCSV int
}

// Handler used as callback to process the progress state as the process happens
type Handler func(state State)
