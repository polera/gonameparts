/*
Package gonameparts splits a human name into individual parts.  This is useful
when dealing with external data sources that provide names as a single value, but
you need to store the discrete parts in a database for example.
*/
package gonameparts

import (
	"sort"
	"strings"
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

	p.FullName = strings.Join(fullNameParts, " ")

}

/*
Parse takes a string name as a parameter and returns a populated NameParts object
*/
func Parse(name string) NameParts {
	n := nameString{FullName: name}
	n.normalize()

	p := NameParts{ProvidedName: name, Nickname: n.Nickname}

	// If we're dealing with a business name, just return it back
	if n.looksCorporate() {
		return p
	}

	parts := []string{"salutation", "generation", "suffix", "lnprefix", "nonname", "supplemental"}
	partMap := make(map[string]int)
	var slotted []int

	// Slot and index parts
	for _, part := range parts {
		partIndex := n.find(part)
		partMap[part] = partIndex
		if partIndex > -1 {
			p.slot(part, n.SplitName[partIndex])
			slotted = append(slotted, partIndex)
		}
	}

	// Slot FirstName
	partMap["first"] = partMap["salutation"] + 1
	p.slot("first", n.SplitName[partMap["first"]])
	slotted = append(slotted, partMap["salutation"]+1)

	// Slot prefixed LastName
	if partMap["lnprefix"] > -1 {
		lnEnd := len(n.SplitName)
		if partMap["generation"] > -1 {
			lnEnd = partMap["generation"]
		}
		if partMap["suffix"] > -1 && partMap["generation"] < lnEnd {
			lnEnd = partMap["suffix"]
		}
		// Need to validate the slice parameters make sense
		// Example Name: "I am the Popsicle"
		// This example causes a hit on the generation at position 0,
		// which in turn causes lnEnd to be set to 0, but the lnprefix
		// is greater than 0, causing a slice out of bounds panic
		if lnEnd > partMap["lnprefix"] {
			p.slot("last", strings.Join(n.SplitName[partMap["lnprefix"]:lnEnd], " "))
		}

		// Keep track of what we've slotted
		for i := partMap["lnprefix"]; i <= lnEnd; i++ {
			slotted = append(slotted, i)
		}
	}

	// Slot the rest
	notSlotted := n.findNotSlotted(slotted)

	if len(notSlotted) > 1 {
		lnPrefix := partMap["lnprefix"]
		var multiMiddle []string
		if lnPrefix > -1 {
			for p := range notSlotted {
				multiMiddle = append(multiMiddle, n.SplitName[p])
			}
			p.slot("middle", strings.Join(multiMiddle, " "))

		} else {
			sort.Sort(sort.IntSlice(notSlotted))
			maxNotSlottedIndex := notSlotted[len(notSlotted)-1]
			p.slot("last", n.SplitName[maxNotSlottedIndex])

			for _, p := range notSlotted {
				if p != maxNotSlottedIndex {
					multiMiddle = append(multiMiddle, n.SplitName[p])
				}
			}
			p.slot("middle", strings.Join(multiMiddle, " "))
		}
	}

	if len(notSlotted) == 1 {
		if partMap["lnprefix"] > -1 {
			p.slot("middle", n.SplitName[notSlotted[0]])
		} else {
			p.slot("last", n.SplitName[notSlotted[0]])
		}
	}

	// Process aliases
	for _, alias := range n.Aliases {
		p.Aliases = append(p.Aliases, Parse(alias))
	}

	// Prepare FullName
	p.buildFullName()

	return p
}
