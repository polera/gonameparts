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

func (s *Scanner) cut() (string, string) {
	size := s.Size
	for range size {
		// Read current word.
		token, err := s.current()
		if err != nil {
			s.Position = 0
			return EMPTY, EMPTY
		}

		// Create two short-lived stacks for this token.
		stackP := new(PuncStack).init()
		stackL := new(LetterStack).init()

		// Parse each character in word.
		feedStacks(token, stackP, stackL)

		if stackL.aka() {
			// Cut string into two pieces around AKA
			divider := s.Position
			first := strings.Join(s.Tokens[:divider], " ")
			second := strings.Join(s.Tokens[divider+1:], " ")

			if second == EMPTY {
				s.Position = 0
				return EMPTY, EMPTY
			}

			return first, second
		}

		s.next()
	}

	s.Position = 0
	return EMPTY, EMPTY
}

func (s *Scanner) isNextTokenPro() bool {
	token, err := s.peek()
	if err != nil {
		return false
	}
	terminus := s.Tokens[s.Final] == token

	// Create two short-lived stacks for this token.
	stackP := new(PuncStack).init()
	stackL := new(LetterStack).init()

	// Parse each character in word.
	feedStacks(token, stackP, stackL)

	if stackP.period == 1 && terminus {
		return true
	}

	return false
}

func (s *Scanner) isNextTokenSuffix() bool {
	token, err := s.peek()
	if err != nil {
		return false
	}

	suffix := false
	if s.Tokens[s.Final] == token {
		suffix = true
	} else if s.Tokens[s.Final-1] == token {
		suffix = true
	}

	// Create two short-lived stacks for this token.
	stackP := new(PuncStack).init()
	stackL := new(LetterStack).init()

	// Parse each character in word.
	feedStacks(token, stackP, stackL)

	if stackP.period == 1 && suffix {
		return true
	}

	return false
}

func (s *Scanner) isNextTokenGenerational() bool {
	token, err := s.peek()
	if err != nil {
		return false
	}

	// Create two short-lived stacks for this token.
	stackP := new(PuncStack).init()
	stackL := new(LetterStack).init()

	// Parse each character in word.
	feedStacks(token, stackP, stackL)

	return stackL.allCaps()
}
