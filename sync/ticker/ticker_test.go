package ticker

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	wg := sync.WaitGroup{}
	ticker := time.NewTicker(time.Second)
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			t := <-ticker.C
			fmt.Println(t)
		}
		wg.Done()
	}()
	wg.Wait()
	ticker.Stop()
}
