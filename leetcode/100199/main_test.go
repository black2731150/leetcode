package main

import "testing"

func TestNumOfOne(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		if numOfOne(34) != 4 {
			t.Errorf("fail")
		}
	})
}
