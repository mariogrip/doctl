package match

import (
	"fmt"
)

type Row struct {
	Matchers    Matchers
	RunesLength int
}

func (self Row) matchAll(s string) bool {
	var idx int
	for _, m := range self.Matchers {
		length := m.Len()

		var next, i int
		for next = range s[idx:] {
			i++
			if i == length {
				break
			}
		}

		if i < length || !m.Match(s[idx:idx+next+1]) {
			return false
		}

		idx += next + 1
	}

	return true
}

func (self Row) lenOk(s string) bool {
	var i int
	for range s {
		i++
		if i >= self.RunesLength {
			return true
		}
	}

	return false
}

func (self Row) Match(s string) bool {
	return self.lenOk(s) && self.matchAll(s)
}

func (self Row) Len() (l int) {
	return self.RunesLength
}

func (self Row) Index(s string) (int, []int) {
	if !self.lenOk(s) {
		return -1, nil
	}

	for i := range s {
		// this is not strict check but useful
		// when glob will be refactored for usage with []rune
		// it will be better
		if len(s[i:]) < self.RunesLength {
			break
		}

		if self.matchAll(s[i:]) {
			return i, []int{self.RunesLength}
		}
	}

	return -1, nil
}

func (self Row) String() string {
	return fmt.Sprintf("<row_%d:[%s]>", self.RunesLength, self.Matchers)
}
