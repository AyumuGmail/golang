package word

import (
	"unicode"
)

func IsAnagram(s1 string, s2 string) bool {
	m1 := createMap(s1)
	m2 := createMap(s2)
	for k, v := range m1 {
		if v != m2[k] {
			return false
		} else {
			delete(m2, k)
		}

	}
	if len(m2) > 0 {
		return false
	} else {
		return true
	}
}

func createMap(s string) map[rune]int {
	letters := str2rune(s)
	m := map[rune]int{}
	for _, letter := range letters {
		if _, exist := m[letter]; exist == true {
			m[letter] = m[letter] + 1
		} else {
			m[letter] = 1
		}
	}
	return m
}

func str2rune(s string) []rune {
	var letters []rune

	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	return letters
}
