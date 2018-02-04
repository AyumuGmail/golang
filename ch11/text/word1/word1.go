package word

import "unicode"

func IsPalindrome(s string) bool {
	//var letters []rune
	//すごいこの記述で1/3程度になった。
	letters := make([]rune, 0, len(s))
	/*
		for i := range s {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	*/
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	/*
		for i := range letters {
			if letters[i] != letters[len(letters)-1-i] {
				return false
			}
		}
	*/
	//改善ポイント１　ただし、4%程度の改善
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
