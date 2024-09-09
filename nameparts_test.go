package gonameparts

import (
	"fmt"
	"reflect"
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

func TestParseOnlySalutation(t *testing.T) {
	t.Parallel()

	res := Parse("Mr.")
	if res.FirstName != "" {
		t.Errorf("Expected ''.  Actual %v", res.FirstName)
	}

	if res.LastName != "" {
		t.Errorf("Expected ''.  Actual %v", res.LastName)
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
	t.Skip("Too many things happening...")

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

func TestObviouslyBadName(t *testing.T) {
	// make sure we don't panic on a clearly bad name
	defer func() {
		if r := recover(); r != nil {
			// panic happened, fail the test
			t.Errorf("Panic happened, where it shouldn't have")
		}
	}()
	Parse("I am a Popsicle")
}

func TestLastNameSalutation(t *testing.T) {
	// make sure we don't panic if the last name looks like a salutation
	defer func() {
		if r := recover(); r != nil {
			// panic happened, fail the test
			t.Errorf("Panic happened, where it shouldn't have")
		}
	}()
	res := Parse("Alan Hon")

	if res.FirstName != "Alan" {
		t.Errorf("Expected 'Alan'.  Actual: %v", res.FirstName)
	}

	if res.LastName != "Hon" {
		t.Errorf("Expected 'Hon'.  Actual: %v", res.LastName)
	}

	if res.FullName != "Alan Hon" {
		t.Errorf("Expected 'Alan Hon'.  Actual: %v", res.FullName)
	}
}

func TestLastNameNonName(t *testing.T) {
	// make sure we don't panic if the last name looks like a nonname
	defer func() {
		if r := recover(); r != nil {
			// panic happened, fail the test
			t.Errorf("Panic happened, where it shouldn't have")
		}
	}()
	res := Parse("Jessica Aka")

	if res.FirstName != "Jessica" {
		t.Errorf("Expected 'Jessica'.  Actual: %v", res.FirstName)
	}

	if res.LastName != "Aka" {
		t.Errorf("Expected 'Aka'.  Actual: %v", res.LastName)
	}

	if res.FullName != "Jessica Aka" {
		t.Errorf("Expected 'Jessica Aka'.  Actual: %v", res.FullName)
	}
}

func TestNameEndsWithApostrophe(t *testing.T) {
	// make sure we don't panic on a clearly bad name
	defer func() {
		if r := recover(); r != nil {
			// panic happened, fail the test
			t.Errorf("Panic happened, where it shouldn't have")
		}
	}()
	res := Parse("James Polera'")
	if res.FirstName != "James" {
		t.Errorf("Expected 'James'. Actual: %v", res.FirstName)
	}

	if res.LastName != "Polera" {
		t.Errorf("Expected 'Polera'. Actual: %v", res.LastName)
	}
}

func TestSuffix(t *testing.T) {
	t.Parallel()
	res := Parse("John A. Smith, Jr.")

	if res.FirstName != "John" {
		t.Errorf("Expected 'John'. Actual: %v", res.FirstName)
	}

	if res.LastName != "Smith" {
		t.Errorf("Expected 'Smith'. Actual: %v", res.LastName)
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

func TestScannerCreation(t *testing.T) {
	t.Parallel()
	name := "John D. Rockefeller, Jr."
	s := new(Scanner).init(name)

	if s.Position != 0 {
		t.Errorf("Expected 0. Actual: %v", s.Position)
	}

	if s.Size != 3 {
		t.Errorf("Expected 4. Actual: %v", s.Size)
	}

	tokens := []string{"John", "D.", "Rockefeller,", "Jr."}
	if reflect.DeepEqual(tokens, s.Tokens) == false {
		t.Errorf("Expected list of strings. Actual: %v", s.Tokens)
	}

}

func TestScanNext(t *testing.T) {
	t.Parallel()
	name := "John D. Rockefeller, Jr."
	s := new(Scanner).init(name)

	token, _ := s.next()

	if token != "D." {
		t.Errorf("Expected 'D.' - Actual: %v", s.Tokens[s.Position])
	}
}

func TestScanNextNothing(t *testing.T) {
	t.Parallel()
	name := "John D. Rockefeller, Jr."
	s := new(Scanner).init(name)
	s.Position = 3

	token, err := s.next()

	if token != "" {
		t.Errorf("Expected '' - Actual: %v", s.Tokens[s.Position])
	}

	if err == nil {
		t.Errorf("Expected an error")
	}
}

func TestScanPrior(t *testing.T) {
	t.Parallel()
	name := "John D. Rockefeller, Jr."
	s := new(Scanner).init(name)
	s.Position = 3

	token, _ := s.prior()

	if token != "Rockefeller," {
		t.Errorf("Expected 'Rockefeller,' - Actual: %v", s.Tokens[s.Position-1])
	}
}

func TestScanPriorNothing(t *testing.T) {
	t.Parallel()
	name := "John D. Rockefeller, Jr."
	s := new(Scanner).init(name)

	token, err := s.prior()

	if token != "" {
		t.Errorf("Expected '' - Actual: %v", s.Tokens[s.Position])
	}

	if err == nil {
		t.Errorf("Expected an error")
	}
}

func TestScanPeek(t *testing.T) {
	t.Parallel()
	name := "John D. Rockefeller, Jr."
	s := new(Scanner).init(name)

	peekedToken, _ := s.peek()
	nextToken, _ := s.next()

	if peekedToken != "D." {
		t.Errorf("Expected 'D.' - Actual: %v", peekedToken)
	}

	if peekedToken != nextToken {
		t.Errorf("Expected 'D.' - Actual: %v", peekedToken)
	}
}

func TestPunctuationStack(t *testing.T) {
	t.Parallel()

	stack := new(PuncStack).init()
	stack.push(PERIOD)
	stack.push("D")

	c, present := stack.pop()
	if c != PERIOD {
		t.Errorf("Expected '.' - Actual: %v", c)
	}

	if present != true {
		t.Errorf("Expected 'true' - Actual: %v", present)
	}
}

func TestStackOrder(t *testing.T) {
	t.Parallel()

	stack := new(PuncStack).init()
	stack.push(PERIOD)
	stack.push(COMMA)
	stack.push(SLASH)
	s, _ := stack.pop()
	c, _ := stack.pop()
	p, _ := stack.pop()

	if p != PERIOD {
		t.Errorf("Expected '.' - Actual: %v", p)
	}

	if c != COMMA {
		t.Errorf("Expected ',' - Actual: %v", c)
	}

	if s != SLASH {
		t.Errorf("Expected '/' - Actual: %v", s)
	}

}

func TestStackQuotationCount(t *testing.T) {
	t.Parallel()

	stack := new(PuncStack).init()
	stack.push(PERIOD)
	stack.push(QUO)
	stack.push(COMMA)
	stack.push(SLASH)
	stack.push(QUO)

	if stack.quomark != 2 {
		t.Errorf("Expected 2 - Actual: %v", stack.quomark)
	}
}

func TestLetterStack(t *testing.T) {
	t.Parallel()

	stack := new(LetterStack).init()
	stack.push("D")
	stack.push("r")
	stack.push(PERIOD)

	if stack.size() != 2 {
		t.Errorf("Expected 2 - Actual: %v", stack.size())
	}

	if stack.allCaps() != false {
		t.Errorf("Expected false - Actual: %v", stack.allCaps())
	}
}

func TestLetterStackAllCaps(t *testing.T) {
	t.Parallel()

	stack := new(LetterStack).init()
	stack.push("I")
	stack.push("I")
	stack.push("I")
	stack.push(PERIOD)

	if stack.size() != 3 {
		t.Errorf("Expected 2 - Actual: %v", stack.size())
	}

	if stack.allCaps() != true {
		t.Errorf("Expected true - Actual: %v", stack.allCaps())
	}
}
