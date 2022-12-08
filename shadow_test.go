package rts_test

import (
	"sync"
	"testing"
	"time"

	"github.com/chikaku/rts"
	"github.com/stretchr/testify/assert"
)

func Test_NumTimers(t *testing.T) {
	ast := assert.New(t)

	const n1 = 32
	count0 := rts.NumTimers()
	for i := 0; i < n1; i++ {
		time.AfterFunc(time.Minute, nil)
	}
	count1 := rts.NumTimers()

	ast.Equal(n1, int(count1-count0))

	const n2 = 128
	wg := new(sync.WaitGroup)
	for i := 0; i < n2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.AfterFunc(time.Minute, nil)
		}()
	}
	wg.Wait()
	count2 := rts.NumTimers()

	ast.Equal(n2, int(count2-count1))
}

func Benchmark_NumTimers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rts.NumTimers()
	}
}
