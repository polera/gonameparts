/*
Package gonameparts splits a human name into individual parts.  This is useful
when dealing with external data sources that provide names as a single value, but
you need to store the discrete parts in a database for example.
*/
package gonameparts

import (
	"slices"
	"strings"
)

const (
	EMPTY = ""
)

/*
Identifiable name parts
*/
var (
	salutations      = []string{"MR", "MS", "MRS", "DR", "MISS", "DOCTOR", "CORP", "SGT", "PVT", "JUDGE", "CAPT", "COL", "MAJ", "LT", "LIEUTENANT", "PRM", "PATROLMAN", "HON", "OFFICER", "REV", "PRES", "PRESIDENT", "GOV", "GOVERNOR", "VICE PRESIDENT", "VP", "MAYOR", "SIR", "MADAM", "HONORABLE"}
	generations      = []string{"JR", "SR", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "1ST", "2ND", "3RD", "4TH", "5TH", "6TH", "7TH", "8TH", "9TH", "10TH", "FIRST", "SECOND", "THIRD", "FOURTH", "FIFTH", "SIXTH", "SEVENTH", "EIGHTH", "NINTH", "TENTH"}
	suffixes         = []string{"ESQ", "PHD", "MD"}
	lnPrefixes       = []string{"DE", "DA", "DI", "LA", "DU", "DEL", "DEI", "VDA", "DELLO", "DELLA", "DEGLI", "DELLE", "VAN", "VON", "DER", "DEN", "HEER", "TEN", "TER", "VANDE", "VANDEN", "VANDER", "VOOR", "VER", "AAN", "MC", "BEN", "SAN", "SAINZ", "BIN", "LI", "LE", "DES", "AM", "AUS'M", "VOM", "ZUM", "ZUR", "TEN", "IBN"}
	nonName          = []string{"A.K.A", "AKA", "A/K/A", "F.K.A", "FKA", "F/K/A", "N/K/A"}
	corpEntity       = []string{"NA", "CORP", "CO", "INC", "ASSOCIATES", "SERVICE", "LLC", "LLP", "PARTNERS", "R/A", "C/O", "COUNTY", "STATE", "BANK", "GROUP", "MUTUAL", "FARGO"}
	supplementalInfo = []string{"WIFE OF", "HUSBAND OF", "SON OF", "DAUGHTER OF", "DECEASED", "FICTITIOUS"}
)

/*
NameParts represents the slotted components of a given name
*/
type NameParts struct {
	ProvidedName string      `json:"provided_name"`
	FullName     string      `json:"full_name"`
	Salutation   string      `json:"salutation"`
	FirstName    string      `json:"first_name"`
	MiddleName   string      `json:"middle_name"`
	LastName     string      `json:"last_name"`
	Generation   string      `json:"generation"`
	Suffix       string      `json:"suffix"`
	Nickname     string      `json:"nickname"`
	Aliases      []NameParts `json:"aliases"`
	Size         int         `json:"size"`
}

func (p *NameParts) slot(part string, value string) {
	switch part {
	case "salutation":
		p.Salutation = value
	case "generation":
		p.Generation = value
	case "suffix":
		p.Suffix = value
	case "middle":
		p.MiddleName = value
	case "last":
		p.LastName = value
	case "first":
		p.FirstName = value
	default:

	}

}

func (p *NameParts) buildFullName() {
	var fullNameParts []string

	if len(p.Salutation) > 0 {
		fullNameParts = append(fullNameParts, p.Salutation)
	}

	if len(p.FirstName) > 0 {
		fullNameParts = append(fullNameParts, p.FirstName)
	}

	if len(p.MiddleName) > 0 {
		fullNameParts = append(fullNameParts, p.MiddleName)
	}

	if len(p.LastName) > 0 {
		fullNameParts = append(fullNameParts, p.LastName)
	}

	if len(p.Generation) > 0 {
		fullNameParts = append(fullNameParts, p.Generation)
	}

	if len(p.Suffix) > 0 {
		fullNameParts = append(fullNameParts, p.Suffix)
	}

	p.Size = len(p.FullName)
	p.FullName = strings.Join(fullNameParts, " ")

}

/*
Parse takes a string name as a parameter and returns a populated NameParts object
*/
func Parse(name string) NameParts {
	s := new(Scanner).init(name)

	// Break string in half along acronymn A.K.A, then find the longer of two names.
	// When the acronmyn is absent, proceed to examine plain name.
	str1, str2 := s.cut()
	if str1 != EMPTY && str2 != EMPTY {
		a := Parse(str1)
		b := Parse(str2)
		if a.Size > b.Size {
			return a
		} else {
			return b
		}
	}

	p := NameParts{ProvidedName: name}

	size := s.Size
	for range size {

		// Read current word.
		token, err := s.current()
		if err != nil {
			return p
		}

		// Is this token in the first half or second half of the whole string?
		latter := s.latterHalf()
		terminus := s.Final == s.Position

		// Create two short-lived stacks for this token.
		stackP := new(PuncStack).init()
		stackL := new(LetterStack).init()

		// Parse each character in word.
		feedStacks(token, stackP, stackL)

		// Are all letters capitalized?
		allCaps := stackL.allCaps()

		// What is the length of the word without punctuation?
		tokenSize := stackL.size()

		// Is this an Honorific? President
		honorific := slices.Contains(salutations, strings.ToUpper(token))
		if s.Position == 0 && honorific {
			p.Salutation = token

			s.next()
			continue
		}

		// Is this a Courtesy Title? Mr. Ms. Mrs. Dr.
		if !latter && tokenSize >= 2 && stackP.period == 1 {
			p.Salutation = token

			s.next()
			continue
		}

		// Is this a first name?
		if !latter && tokenSize >= 2 && stackP.period == 0 && p.FirstName == EMPTY {
			cleanToken := stackL.assemble()
			p.FirstName = cleanToken

			s.next()
			continue
		}

		// Is this a first initial? J.
		if !latter && tokenSize == 1 && stackP.period == 1 && p.FirstName == EMPTY {
			p.FirstName = token

			s.next()
			continue
		}

		// Is this a second initial for a middle name? The D in J. D. Rockefeller
		if tokenSize == 1 && stackP.period == 1 && p.FirstName != EMPTY && p.MiddleName == EMPTY {
			p.MiddleName = token

			s.next()
			continue
		}

		// Is this a middle name?
		suffixes := false
		if s.isNextTokenSuffix() == true || s.isNextTokenGenerational() == true {
			suffixes = true
		}
		if latter && !terminus && p.FirstName != EMPTY && stackP.comma == 0 && stackP.apo == 0 && stackP.quomark == 0 && suffixes == false {

			if p.MiddleName != EMPTY {
				p.MiddleName += " " + token
				s.next()
				continue
			}
			p.MiddleName = token

			s.next()
			continue
		}

		// Is this a simple nickname?
		if latter && stackP.apo == 2 || latter && stackP.quomark == 2 {
			cleanToken := stackL.assemble()
			p.Nickname = cleanToken

			s.next()
			continue
		}

		// Is this a mistyped Irish name? O' Hurley needs to be O'Hurley
		if latter && tokenSize == 1 && stackP.apo == 1 {
			nextToken, err := s.peek()
			if err != nil {
				return p
			}
			tokens := []string{token, nextToken}
			irishSurname := strings.Join(tokens, "")
			p.LastName = irishSurname

			s.next()
			continue
		}

		// Is this a complex nickname?
		if !terminus && stackP.apo == 1 || !terminus && stackP.quomark == 1 {

			first := string(token[0])
			last := string(token[tokenSize])
			if first == QUO || first == APOSTROPHE {
				p.Nickname = token

				s.next()
				continue
			}

			if last == QUO || last == APOSTROPHE {
				p.Nickname += " " + token

				s.next()
				continue
			}
		}

		// Is this a family name preceding a professional suffix? Polera, Esq.
		if latter && !allCaps && stackP.period == 0 && suffixes {
			cleanToken := stackL.assemble()
			p.LastName = cleanToken

			s.next()
			continue
		}

		// Is this a family name? Rockefeller
		if terminus && p.MiddleName != EMPTY && !allCaps && stackP.period == 0 && p.LastName == EMPTY || terminus && p.Nickname != EMPTY && !allCaps && p.LastName == EMPTY {

			// Preserve hyphen in surname.
			if stackP.hyphenated() {
				p.LastName = token
				s.next()
				continue
			}

			cleanToken := stackL.assemble()

			if len(p.MiddleName) > 0 && len(p.MiddleName) <= 3 {
				cleanToken = p.MiddleName + " " + cleanToken
				p.MiddleName = ""
			}

			p.LastName = cleanToken

			s.next()
			continue
		}

		// Is this a family name with a comma? Rockefeller,
		if !terminus && latter && stackP.comma == 1 && p.LastName == EMPTY {
			cleanToken := stackL.assemble()
			p.LastName = cleanToken

			s.next()
			continue
		}

		// Is this an Irish surname? Preserve puncutation in token.
		if terminus && stackP.apo == 1 && tokenSize > 2 && string(token[1]) == APOSTROPHE {
			p.LastName = token

			s.next()
			continue
		}

		// Is this a family name in a two token string?
		if terminus && size == 2 {
			cleanToken := stackL.assemble()
			p.LastName = cleanToken

			s.next()
			continue
		}

		// Is this a generational suffix? II III IV VI
		if latter && allCaps && tokenSize > 1 || latter && stackL.size() == 2 && stackP.period == 1 {
			p.Generation = token

			s.next()
			continue
		}

		// Is this a professional suffix? Esq. M.D.
		if terminus && stackP.period == 1 {
			p.Suffix = token

			s.next()
			continue
		}

		// None of the rules applied. Advance to next token.
		s.next()
	}

	//// Prepare FullName
	p.buildFullName()

	return p
}
