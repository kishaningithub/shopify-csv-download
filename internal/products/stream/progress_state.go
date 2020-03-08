package stream

import "sync"

type ProgressState struct {
	NoOfProductsDownloaded     <-chan int
	NoOfProductsConvertedAsCSV <-chan int
}

func (s ProgressState) Ignore() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func() {
		for range s.NoOfProductsDownloaded {
		}
		waitGroup.Done()
	}()
	go func() {
		for range s.NoOfProductsConvertedAsCSV {
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
}
