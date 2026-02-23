package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitGroupBasic(t *testing.T) {
	tests := []struct {
		name       string
		addCount   int
		doneCount  int
		expectPanics bool
	}{
		{"normal", 3, 3, false},
		{"zero count", 0, 0, false},
		{"panic", 1, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := NewWaitGroup()

			if tt.expectPanics {
				defer func() {
					if r := recover(); r == nil {
						t.Fatalf("expected panic but no")
					}
				}()
			}

			wg.Add(tt.addCount)
			for i := 0; i < tt.doneCount; i++ {
				wg.Done()
			}

			if !tt.expectPanics {
				if atomic.LoadInt64(&wg.counter) != 0 {
					t.Fatalf("expected counter=0, got %d", wg.counter)
				}
			}
		})
	}
}

func TestWaitBlocksUntilDone(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(2)

	finished := make(chan struct{})

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
		t.Fatal("Wait() returned before Done()")
	case <-time.After(50 * time.Millisecond):
	}

	wg.Done()
	select {
	case <-finished:
		t.Fatal("Wait() returned after 1 Done()")
	case <-time.After(50 * time.Millisecond):
	}

	wg.Done()
	select {
	case <-finished:
	case <-time.After(50 * time.Millisecond):
		t.Fatal("Wait() not unblock after all Done()")
	}
}
