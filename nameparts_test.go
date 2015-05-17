package gonameparts

import (
	"testing"
)

func TestLooksCorporate(t *testing.T) {
	n := nameString{FullName: "Sprockets Inc"}

	res := n.looksCorporate()

	if res != true {
		t.Errorf("Expected true.  Actual %v", res)
	}

}

func TestSearchParts(t *testing.T) {
	n := nameString{FullName: "Mr. James Polera"}

	res := n.searchParts(&salutations)

	if res != 0 {
		t.Errorf("Expected true.  Actual %v", res)
	}

}

func TestClean(t *testing.T) {
	n := nameString{FullName: "Mr. James Polera"}

	res := n.cleaned()

	if res[0] != "Mr" {
		t.Errorf("Expected 'Mr'.  Actual %v", res[0])
	}

}

func TestLocateSalutation(t *testing.T) {
	n := nameString{FullName: "Mr. James Polera"}

	res := n.find("salutation")

	if res != 0 {
		t.Errorf("Expected 0.  Actual %v", res)
	}
}
