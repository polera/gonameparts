package gonameparts

import (
	"fmt"
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

	res := n.searchParts(salutations)

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

func TestParseAllFields(t *testing.T) {
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

func TestParseFirstLast(t *testing.T) {
	t.Parallel()

	res := Parse("James Polera")
	if res.FirstName != "James" {
		t.Errorf("Expected 'James'.  Actual %v", res.FirstName)
	}

	if res.LastName != "Polera" {
		t.Errorf("Expected 'Polera'.  Actual %v", res.LastName)
	}
}

func TestLastNamePrefix(t *testing.T) {
	t.Parallel()

	res := Parse("Otto von Bismark")

	if res.FirstName != "Otto" {
		t.Errorf("Expected 'Otto'.  Actual %v", res.FirstName)
	}

	if res.LastName != "von Bismark" {
		t.Errorf("Expected 'von Bismark'.  Actual %v", res.LastName)
	}

}

func TestAliases(t *testing.T) {
	t.Parallel()

	res := Parse("James Polera a/k/a Batman")

	if res.Aliases[0].FirstName != "Batman" {
		t.Errorf("Expected 'Batman'.  Actual: %v", res.Aliases[0].FirstName)
	}

}

func TestNickname(t *testing.T) {
	t.Parallel()

	res := Parse("Philip Francis 'The Scooter' Rizzuto")

	if res.Nickname != "'The Scooter'" {
		t.Errorf("Expected 'The Scooter'.  Actual: %v", res.Nickname)
	}
}

func TestStripSupplemental(t *testing.T) {
	t.Parallel()

	res := Parse("Philip Francis 'The Scooter' Rizzuto, deceased")

	if res.FirstName != "Philip" {
		t.Errorf("Expected 'Philip'.  Actual: %v", res.FirstName)
	}

	if res.MiddleName != "Francis" {
		t.Errorf("Expected 'Francis'.  Actual: %v", res.MiddleName)
	}

	if res.Nickname != "'The Scooter'" {
		t.Errorf("Expected 'The Scooter'.  Actual: %v", res.Nickname)
	}

	if res.LastName != "Rizzuto" {
		t.Errorf("Expected 'Rizzuto'.  Actual: %v", res.LastName)
	}
}

func TestLongPrefixedLastName(t *testing.T) {
	t.Parallel()

	res := Parse("Saleh ibn Tariq ibn Khalid al-Fulan")

	if res.FirstName != "Saleh" {
		t.Errorf("Expected 'Saleh'.  Actual: %v", res.FirstName)
	}

	if res.LastName != "ibn Tariq ibn Khalid al-Fulan" {
		t.Errorf("Expected 'ibn Tariq ibn Khalid al-Fulan'.  Actual: %v", res.LastName)

	}
}

func TestMisplacedApostrophe(t *testing.T) {
	t.Parallel()

	res := Parse("John O' Hurley")

	if res.FirstName != "John" {
		t.Errorf("Expected 'John'.  Actual: %v", res.FirstName)
	}

	if res.LastName != "O'Hurley" {
		t.Errorf("Expected 'O'Hurley'.  Actual: %v", res.LastName)
	}

}

func TestMultipleAKA(t *testing.T) {
	t.Parallel()

	res := Parse("Tony Stark a/k/a Ironman a/k/a Stark, Anthony a/k/a Anthony Edward \"Tony\" Stark")

	if len(res.Aliases) != 3 {
		t.Errorf("Expected 3 aliases.  Actual: %v", len(res.Aliases))
	}

	if res.FirstName != "Tony" {
		t.Errorf("Expected 'Tony'.  Actual: %v", res.FirstName)
	}

	if res.LastName != "Stark" {
		t.Errorf("Expected 'Stark'.  Actual: %v", res.LastName)
	}

}

func TestBuildFullName(t *testing.T) {
	res := Parse("President George Herbert Walker Bush")

	if res.FullName != "President George Herbert Walker Bush" {

		t.Errorf("Expected 'President George Herbert Walker Bush'.  Actual: %v", res.FullName)
	}

}

func TestDottedAka(t *testing.T) {
	res := Parse("James Polera a.k.a James K. Polera")
	if len(res.Aliases) != 1 {
		t.Errorf("Expected 1 alias.  Actual: %v", len(res.Aliases))
	}
}

func TestUnicodeCharsInName(t *testing.T) {
	res := Parse("König Ludwig")

	if res.FirstName != "König" {
		t.Errorf("Expected 'König'.  Actual: %v", res.FirstName)
	}
}

func TestTabsInName(t *testing.T) {
	res := Parse("Dr. James\tPolera\tEsq.")

	if res.Salutation != "Dr." {
		t.Errorf("Expected 'Dr.'.  Actual: %v", res.Salutation)
	}

	if res.FirstName != "James" {
		t.Errorf("Expected 'James'.  Actual: %v", res.FirstName)
	}

	if res.LastName != "Polera" {
		t.Errorf("Expected 'Polera'.  Actual: %v", res.LastName)
	}

	if res.Suffix != "Esq." {
		t.Errorf("Expected 'Esq.'.  Actual: %v", res.Suffix)
	}
}

func ExampleParse() {
	res := Parse("Thurston Howell III")
	fmt.Println("FirstName:", res.FirstName)
	fmt.Println("LastName:", res.LastName)
	fmt.Println("Generation:", res.Generation)

	// Output:
	// FirstName: Thurston
	// LastName: Howell
	// Generation: III

}

func ExampleParse_second() {

	res := Parse("President George Herbert Walker Bush")
	fmt.Println("Salutation:", res.Salutation)
	fmt.Println("FirstName:", res.FirstName)
	fmt.Println("MiddleName:", res.MiddleName)
	fmt.Println("LastName:", res.LastName)

	// Output:
	// Salutation: President
	// FirstName: George
	// MiddleName: Herbert Walker
	// LastName: Bush

}
