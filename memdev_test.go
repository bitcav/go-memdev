package memdev

import "testing"

func TestInfo(t *testing.T) {
	memInfo, err := Info()
	if err != nil {
		t.Fatalf("Expected nil error but got %v", err)
	}
	if len(memInfo) < 1 {
		t.Fatalf("Expected >0 Memory , but got %d", len(memInfo))
	}
}
