package gonameparts

import (
	"errors"
	"strings"
)

type Scanner struct {
	Tokens   []string
	Position uint
	Size     uint
	Final    uint
}

func (s *Scanner) init(longString string) *Scanner {
	s.Tokens = strings.Fields(longString)
	s.Position = uint(0)
	s.Size = uint(len(s.Tokens))
	s.Final = uint(len(s.Tokens) - 1)
	return s
}

func (s *Scanner) current() (string, error) {
	if s.Size == 0 {
		return "", errors.New("Empty")
	}

	return s.Tokens[s.Position], nil
}

func (s *Scanner) next() (string, error) {
	if s.Position < s.Final {
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
	if s.Position < s.Final {
		return s.Tokens[s.Position+1], nil
	}
	return "", errors.New("No more tokens ahead")
}

func (s *Scanner) latterHalf() bool {
	return s.Position >= s.Size/2
}
