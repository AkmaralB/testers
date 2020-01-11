package student

func Split(str, charset string) []string {
	var lenStr int
	var wordIndex int
	for range str {
		lenStr++
	}
	var lenChar int
	for range charset {
		lenChar++
	}
	wordsArray := make([]string, lenStr)
	charsetCount := 0
	word := ""
	if charset != "" {
		for i := 0; i < lenStr; i++ {
			word += string(str[i])
			lenword := 0
			for range word {
				lenword++
			}
			if lenword >= lenChar {
				if word[lenword-lenChar:lenword] == charset {
					charsetCount++
					wordsArray[wordIndex] = word[0 : lenword-lenChar]
					wordIndex++
					word = ""
				}
			}
			if i == lenStr-1 && word != "" {
				wordsArray[wordIndex] = word
				wordIndex++
			}
		}
	}
	if charsetCount == wordIndex {
		return wordsArray[0 : wordIndex+1]
	}
	return wordsArray[0:wordIndex]
}
