package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(2*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(6*time.Second),
	)

	end := time.Since(start)

	if !(end < 2*time.Second) {
		t.Errorf("Reading not from the fastest channel. Expected time : ~1s; Got: %v", end)
	}
}
