package main

func PalindromeCheck(text string) bool {

	var length int = len(text)
	var i int
	var j int

	for i = 0; i < length/2; i++ {
		j = length - i - 1
		if text[i] != text[j] {
			return false
		}
	}
	return true

}


