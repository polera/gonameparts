package gonameparts

import (
	"errors"
	"strings"
)

type Scanner struct {
	Tokens   []string
	Position uint
	Size     uint
}

func (s *Scanner) init(longString string) *Scanner {
	s.Tokens = strings.Fields(longString)
	s.Position = uint(0)
	s.Size = uint(len(s.Tokens)) - uint(1)
	return s
}

func (s *Scanner) next() (string, error) {
	if s.Position < s.Size {
		s.Position += 1
		return s.Tokens[s.Position], nil
	}
	return "", errors.New("Size of tokens")
}

func (s *Scanner) prior() (string, error) {
	if s.Position > 0 {
		return s.Tokens[s.Position-1], nil
	}
	return "", errors.New("Nothing before the zero-index")
}

func (s *Scanner) peek() (string, error) {
	if s.Position < s.Size {
		return s.Tokens[s.Position+1], nil
	}
	return "", errors.New("No more tokens ahead")
}

func (s *Scanner) latterHalf() bool {
	return s.Position+1 > s.Size/2
}
