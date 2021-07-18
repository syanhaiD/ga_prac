package main

import (
	"testing"
)

func TestTarget(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		// t.Parallel()
		if target() != true {
			t.Fatalf("sonnna wake naiyan")
		}
	})
}
