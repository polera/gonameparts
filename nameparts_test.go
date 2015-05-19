package gonameparts

import (
	"testing"
)

func TestLooksCorporate(t *testing.T) {
	t.Parallel()
	n := nameString{FullName: "Sprockets Inc"}

	res := n.looksCorporate()

	if res != true {
		t.Errorf("Expected true.  Actual %v", res)
	}

}

func TestSearchParts(t *testing.T) {
	t.Parallel()
	n := nameString{FullName: "Mr. James Polera"}

	res := n.searchParts(&salutations)

	if res != 0 {
		t.Errorf("Expected true.  Actual %v", res)
	}

}

func TestClean(t *testing.T) {
	t.Parallel()
	n := nameString{FullName: "Mr. James Polera"}

	res := n.cleaned()

	if res[0] != "Mr" {
		t.Errorf("Expected 'Mr'.  Actual %v", res[0])
	}

}

func TestLocateSalutation(t *testing.T) {
	t.Parallel()
	n := nameString{FullName: "Mr. James Polera"}

	res := n.find("salutation")

	if res != 0 {
		t.Errorf("Expected 0.  Actual %v", res)
	}
}

func TestHasComma(t *testing.T) {
	t.Parallel()
	n := nameString{FullName: "Polera, James"}
	res := n.hasComma()

	if res != true {
		t.Errorf("Expected true.  Actual %v", res)
	}

}

func TestNormalize(t *testing.T) {
	t.Parallel()
	n := nameString{FullName: "Polera, James"}
	res := n.normalize()

	if res[0] != "James" {
		t.Errorf("Expected James.  Actual %v", res[0])
	}

	if res[1] != "Polera" {
		t.Errorf("Expected Polera.  Actual %v", res[1])
	}

}

func TestParse(t *testing.T) {
	t.Parallel()
	res := Parse("Mr. James J. Polera Jr. Esq.")

	if res.Salutation != "Mr." {
		t.Errorf("Expected 'Mr.'.  Actual %v", res.Salutation)
	}

	if res.FirstName != "James" {
		t.Errorf("Expected 'James'.  Actual %v", res.FirstName)
	}

	if res.MiddleName != "J." {
		t.Errorf("Expected 'J.'.  Actual %v", res.MiddleName)
	}

	if res.LastName != "Polera" {
		t.Errorf("Expected 'Polera'.  Actual %v", res.LastName)
	}

	if res.Generation != "Jr." {
		t.Errorf("Expected 'Jr.'.  Actual %v", res.Generation)
	}

	if res.Suffix != "Esq." {
		t.Errorf("Expected 'Esq.'.  Actual %v", res.Suffix)
	}
}
