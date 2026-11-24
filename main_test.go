package main

import (
	"reflect"
	"testing"
)

func TestLeftView(t *testing.T) {
	// Level 2 exists only on the right side: its node still shows.
	root := tn(1, tn(2, nil, tn(4, nil, nil)), tn(3, nil, nil))
	if got := LeftView(root); !reflect.DeepEqual(got, []int{1, 2, 4}) {
		t.Errorf("LeftView = %v, want [1 2 4]", got)
	}

	// Deep level reachable only via the right subtree.
	r2 := tn(1, tn(2, nil, nil), tn(3, nil, tn(7, tn(8, nil, nil), nil)))
	if got := LeftView(r2); !reflect.DeepEqual(got, []int{1, 2, 7, 8}) {
		t.Errorf("LeftView = %v, want [1 2 7 8]", got)
	}

	if got := LeftView(nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("LeftView(nil) = %v", got)
	}
	if got := LeftView(tn(9, nil, nil)); !reflect.DeepEqual(got, []int{9}) {
		t.Errorf("single node: %v", got)
	}
}
