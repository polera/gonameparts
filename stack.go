package gonameparts

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	COMMA  = ","
	PERIOD = "."
	SLASH  = "/"
	QUO    = "\""
)

type Stack[T any] []T

func (s *Stack[T]) push(v T) {
	*s = append([]T{v}, (*s)...)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(*s) == 0 {
		var v T
		return v, false
	}
	v := (*s)[0]
	*s = (*s)[1:]
	return v, true
}

func (s *Stack[T]) size() int {
	return len(*s)
}

type PuncStack struct {
	s       Stack[string]
	quomark uint
	comma   uint
	period  uint
	slash   uint
}

func (p *PuncStack) init() *PuncStack {
	p.quomark = 0
	p.comma = 0
	p.period = 0
	p.slash = 0
	return p
}

func (p *PuncStack) push(c string) {
	switch c {
	case COMMA:
		p.s.push(c)
		p.comma += 1
	case PERIOD:
		p.s.push(c)
		p.period += 1
	case SLASH:
		p.s.push(c)
		p.slash += 1
	case QUO:
		p.s.push(c)
		p.quomark += 1
	default:
		return
	}
}

func (p *PuncStack) pop() (string, bool) {
	return p.s.pop()
}

type LetterStack struct {
	s        Stack[string]
	capitals uint
}

func (l *LetterStack) init() *LetterStack {
	l.capitals = 0
	return l
}

func (l *LetterStack) pop() (string, bool) {
	return l.s.pop()
}

func (l *LetterStack) push(c string) {
	ch, _ := utf8.DecodeRuneInString(c)
	switch unicode.IsLetter(ch) {
	case true:
		if unicode.IsUpper(ch) {
			l.capitals += 1
		}
		l.s.push(c)
	default:
		return
	}
}

func (l *LetterStack) size() int {
	return l.s.size()
}

func (l *LetterStack) allCaps() bool {
	return uint(l.size()) == l.capitals
}

func (l *LetterStack) assemble() string {
	var token string
	for i := l.size() - 1; i >= 0; i-- {
		token += l.s[i]
	}
	return token
}

func feedStacks(token string, stack1 *PuncStack, stack2 *LetterStack) {
	characters := strings.Split(token, "")
	for _, ch := range characters {
		stack1.push(ch)
		stack2.push(ch)
	}
}
