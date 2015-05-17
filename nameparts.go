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

func (n *nameString) searchParts(parts *[]string) (index int) {
	for i, x := range n.split() {
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

func (n *nameString) split() []string {
	return strings.Fields(n.FullName)
}

func Parse(name string) []string {
	n := new(nameString)
	n.FullName = name
	return n.split()

}
