package main

import (
	"testing"
)

func TestNewQueuering(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want queuering[int]
	}{
		// TODO: Add test cases.
		{},
	}
	testpassed := false
	_ = testpassed
	// defer func() {
	// 	if !testpassed {
	// 		t.Error("expects panic.")
	// 	}
	// }()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if err := recover(); err != nil {
					if err == "oversize" {
						testpassed = true
					} else {
						t.Error("test expects panic")
					}
				}
			}()

			q := NewQueuering[int](4)
			if !q.Empty() {
				t.Error("error in Empty()")
			}
			q.Put(1)
			if q.Get() != 1 {
				t.Error("Get() expects 1")
			}
			if !q.Empty() {
				t.Error("error in Empty()")
			}
			q.Put(1)
			q.Put(2)
			q.Put(3)
			q.Put(4)
			//expects panic
			q.Put(5)
		})
	}
}
