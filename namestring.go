package gonameparts

import (
	"sort"
	"strings"
)

type nameString struct {
	FullName  string
	SplitName []string
	Nickname  string
	Aliases   []string
}

func (n *nameString) cleaned() []string {
	unwanted := []string{",", "."}
	cleaned := []string{}
	for _, x := range n.split() {
		for _, y := range unwanted {
			x = strings.Replace(x, y, "", -1)
		}
		cleaned = append(cleaned, strings.Trim(x, " "))
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

func (n *nameString) hasComma() bool {

	for _, x := range n.split() {
		if strings.ContainsAny(x, ",") {
			return true
		}
	}
	return false
}

func (n *nameString) slotNickname() {

	var nickNameBoundaries []int

	for index, x := range n.split() {
		if string(x[0]) == "'" || string(x[0]) == "\"" {
			nickNameBoundaries = append(nickNameBoundaries, index)
		}
		if string(x[len(x)-1]) == "'" || string(x[len(x)-1]) == "\"" {
			nickNameBoundaries = append(nickNameBoundaries, index)
		}
	}

	if len(nickNameBoundaries) > 0 && len(nickNameBoundaries)%2 == 0 {
		nickStart := nickNameBoundaries[0]
		nickEnd := nickNameBoundaries[1]

		nick := n.SplitName[:nickStart]
		postNick := n.SplitName[nickEnd+1:]

		n.Nickname = strings.Join(n.SplitName[nickStart:nickEnd+1], " ")
		nick = append(nick, postNick...)
		n.FullName = strings.Join(nick, " ")
	}
}

func (n *nameString) hasAliases() (bool, string) {
	for _, x := range nonName {
		if strings.Contains(strings.ToUpper(n.FullName), x) {
			return true, x
		}
	}
	return false, ""
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
	n.SplitName = strings.Fields(n.FullName)
	return n.SplitName
}

func (n *nameString) normalize() []string {

	// Handle any aliases in our nameString
	hasAlias, aliasSep := n.hasAliases()

	if hasAlias {
		n.splitAliases(aliasSep)
	}

	// Strip Supplemental info
	supplementalIndex := n.find("supplemental")
	if supplementalIndex > -1 {
		n.FullName = strings.Join(n.SplitName[:supplementalIndex], " ")
	}

	// Handle quoted Nicknames
	n.slotNickname()

	// Swap Lastname, Firstname to Firstname Lastname
	if n.hasComma() {
		commaSplit := strings.SplitN(n.FullName, ",", 2)
		sort.StringSlice(commaSplit).Swap(1, 0)
		n.FullName = strings.Join(commaSplit, " ")
	}

	return n.cleaned()

}

func (n *nameString) splitAliases(aliasSep string) {
	splitNames := n.split()

	for index, part := range splitNames {
		if strings.ToUpper(part) == aliasSep {
			splitNames[index] = "*|*"
		}
	}

	names := strings.Split(strings.Join(splitNames, " "), "*|*")

	n.FullName = names[0]
	n.Aliases = names[1:]
}

func (n *nameString) findNotSlotted(slotted []int) []int {

	var notSlotted []int

	for i, _ := range n.SplitName {
		found := false
		for _, j := range slotted {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			notSlotted = append(notSlotted, i)
		}
	}

	return notSlotted
}
