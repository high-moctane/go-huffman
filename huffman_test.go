package huffman

import "testing"

func TestAlwaysTrue(t *testing.T) {
	if !AlwaysTrue() {
		t.Errorf("expect true, but AlwaysTrue() == %v", AlwaysTrue())
	}
}
