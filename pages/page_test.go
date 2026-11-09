package pages

import "testing"

func TestZeroPageHasEightKbData(t *testing.T) {
	var p Page
	if len(p.data) != 8192{
		t.Errorf("len(p.data) = %d, want 8192", len(p.data))
	}
}