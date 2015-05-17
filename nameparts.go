package gonameparts

import (
	"strings"
)

type NameParts struct {
	ProvidedName string   `json:"provided_name"`
	FullName     string   `json:"full_name"`
	Salutation   string   `json:"salutation"`
	FirstName    string   `json:"first_name"`
	MiddleName   string   `json:"middle_name"`
	LastName     string   `json:"last_name"`
	Suffix       string   `json:"suffix"`
	Aliases      []string `json:"aliases"`
}

type nameString struct {
	FullName string
}

func (n *nameString) cleaned() []string {
	unwanted := []string{",", "."}
	cleaned := []string{}
	for _, x := range n.split() {
		for _, y := range unwanted {
			x = strings.Replace(x, y, "", -1)
		}
		cleaned = append(cleaned, x)
	}
	return cleaned
}

func (n *nameString) searchParts(parts *[]string) int {
	for i, x := range n.cleaned() {
		for _, y := range *parts {
			if strings.ToUpper(x) == strings.ToUpper(y) {
				return i
			}
		}
	}
	return -1
}

func (n *nameString) looksCorporate() bool {
	return n.searchParts(&corpEntity) > -1
}

func (n *nameString) find(part string) int {
	switch part {
	case "salutation":
		return n.searchParts(&salutations)
	case "generation":
		return n.searchParts(&generations)
	case "suffix":
		return n.searchParts(&suffixes)
	case "lnprefix":
		return n.searchParts(&lnPrefixes)
	case "nonname":
		return n.searchParts(&nonName)
	case "supplemental":
		return n.searchParts(&supplementalInfo)
	default:

	}
	return -1
}

func (n *nameString) split() []string {
	return strings.Fields(n.FullName)
}

func Parse(name string) []string {
	n := nameString{FullName: name}
	return n.split()
}
