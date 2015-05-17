package gonameparts

import (
	"testing"
)

func TestLooksCorporate(t *testing.T) {
	n := nameString{FullName: "Sprockets Corp"}

	res := n.looksCorporate()

	if res != true {
		t.Errorf("Expected true.  Actual %v", res)
	}

}

func TestSearchParts(t *testing.T) {
	n := nameString{FullName: "Mr James Polera"}

	res := n.searchParts(&salutations)

	if res != 0 {
		t.Errorf("Expected true.  Actual %v", res)
	}

}
