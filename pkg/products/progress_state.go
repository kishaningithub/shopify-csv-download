package products

// ProgressState comprises of progress state that can be used for displaying the consumer the current state of the operation
type ProgressState struct {
	NoOfProductsDownloaded     int
	NoOfProductsConvertedAsCSV int
}

// ProgressHandler used as callback to process the progress state as the process happens
type ProgressHandler func(state ProgressState)
