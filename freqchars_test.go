package freqchars

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	items := Count("aab")
	if len(items) != 2 {
		t.Fatalf("len=%d want 2", len(items))
	}
	// 找 a 的计数
	var aCount int
	for _, it := range items {
		if it.Char == "a" {
			aCount = it.Count
		}
	}
	if aCount != 2 {
		t.Errorf("a count=%d want 2", aCount)
	}
}

func TestTop(t *testing.T) {
	top := Top("aabbc", 2)
	if len(top) != 2 {
		t.Fatalf("top len=%d want 2", len(top))
	}
	if top[0].Count < top[1].Count {
		t.Errorf("top not sorted desc: %v", top)
	}
}

func TestBar(t *testing.T) {
	items := []Item{{Char: "a", Count: 3}, {Char: "b", Count: 1}}
	out := Bar(items, "#", 1)
	if !strings.Contains(out, "###") || !strings.Contains(out, "#") {
		t.Errorf("bar output wrong: %q", out)
	}
}
