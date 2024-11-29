package word_censor_lib

import "strings"

func StringContainCensoredWord(str string) bool {
	if wordCensorLib.CensorRegexV1 == nil {
		return false
	}

	return wordCensorLib.CensorRegexV1.MatchString(str)
}

func StringReplaceCensoredWord(str, rplc string) string {
	if wordCensorLib.CensorRegexV1 != nil {
		censored := false
		replaced := wordCensorLib.CensorRegexV1.ReplaceAllStringFunc(strings.ToLower(str), func(match string) string {
			censored = true
			return strings.Repeat("*", len(match))
		})
		if censored {
			return replaced
		} else {
			return str
		}
	}

	//Old censor in case compiled regex err
	var newSlice []string

	// check for empty slice
	if len(wordCensorLib.Words) <= 0 {
		return str
	}

	// convert str into a slice
	strSlice := strings.Fields(str)

	//check each words in strSlice against censored slice
	for position, word := range strSlice {
		for _, forbiddenWord := range wordCensorLib.Words {
			// NOTE : change between Index and EqualFold to see the different result
			//if test := strings.Index(strings.ToLower(word), forbiddenWord); test > -1 {
			if test := strings.EqualFold(strings.ToLower(word), forbiddenWord); test == true {
				// calculate how many # for replacement
				replacement := strings.Repeat("#", len(word))

				strSlice[position] = replacement
				newSlice = append(strSlice[:position], strSlice[position:]...)
			}
		}
	}

	// convert []string slice back to string
	return strings.Join(newSlice, " ")
}
