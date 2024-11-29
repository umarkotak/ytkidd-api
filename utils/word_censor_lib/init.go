package word_censor_lib

import (
	"regexp"

	"github.com/sirupsen/logrus"
)

type (
	WordCensorLib struct {
		Words         []string
		CensorRegexV1 *regexp.Regexp
		CensorRegexV2 *regexp.Regexp
	}
)

var (
	wordCensorLib WordCensorLib
)

func Initialize(w WordCensorLib) error {
	censorRegexV1, err := buildCensorRegexV1(w.Words)
	if err != nil {
		logrus.Error(err)
		return err
	}

	censorRegexV2, err := buildCensorRegexV1(w.Words)
	if err != nil {
		logrus.Error(err)
		return err
	}

	wordCensorLib = WordCensorLib{
		Words:         w.Words,
		CensorRegexV1: censorRegexV1,
		CensorRegexV2: censorRegexV2,
	}

	return nil
}

func buildCensorRegexV1(words []string) (*regexp.Regexp, error) {
	regexAll := ""
	for _, pattern := range words {
		regexPattern := ""
		/*
				- Every character in the censored word will be appended with: +[\W|_]*,
				  so any 'non word' character will be ignored.
			      e.g: pattern for word 'shit' will be 's+[\W|_]*h+[\W|_]*i+[\W|_]*t+' and 's h!i-t' will match
				- If pattern length <= 6, use whole word matching by start the pattern
				  with : (^|\s+)[\W|_]* and end with +[\W|_]*($|\s+).  If not the pattern is ended with +
				- certain character will have an alias e.g: i will have aliases : i 1 l, s will have alias s,5,c.
			      So `5h1t` still a match
		*/
		if len(pattern) <= 6 {
			regexPattern += `(^|\s+)[\W|_]*`
		}
		for idx, ch := range pattern {
			word := string(ch)
			switch word {
			case "a":
				word = "[a|4]"
			case "i":
				word = "[i|1]"
			case "e":
				word = "[e|3]"
			case "o":
				word = "[o|0]"
			case "l":
				word = "[l|1]"
			case "k":
				word = "[k|q]"
			case "p":
				word = "[p|v]"
			case "s":
				word = "[s|5]"
			}
			if idx < len(pattern)-1 {
				regexPattern = regexPattern + word + "+[\\W|_]*"
			} else {
				if word == "t" || word == "d" {
					word = "[t|d]"
				}
				if len(pattern) <= 6 {
					regexPattern += word + `+[\W|_]*($|\s+)`
				} else {
					regexPattern += word + "+"
				}
			}
		}
		regexAll = regexAll + "|" + regexPattern
	}
	if regexAll == "" {
		logrus.Errorf("Word filter regex is empty")
	}
	regexAll = regexAll[1:]

	censorRegex, err := regexp.Compile(regexAll)
	if err != nil {
		logrus.Errorf("Failed to build censor regex: %s", err.Error())
	}

	return censorRegex, nil
}

func buildCensorRegexV2(words []string) (*regexp.Regexp, error) {
	regexAll := ""
	for _, pattern := range words {
		regexPattern := ""
		/*
				- Every character in the censored word will be appended with: +[\W|_]*,
				  so any 'non word' character will be ignored.
			      e.g: pattern for word 'shit' will be 's+[\W|_]*h+[\W|_]*i+[\W|_]*t+' and 's h!i-t' will match
				- If pattern length <= 6, use whole word matching by start the pattern
				  with : (^|\s+)[\W|_]* and end with +[\W|_]*($|\s+).  If not the pattern is ended with +
				- certain character will have an alias e.g: i will have aliases : i 1 l, s will have alias s,5,c.
			      So `5h1t` still a match
			      	- Improvement
				  - Any non-alphabet character or number will be included, e.g. ?5h1t111 will be detected
				  - Only concern with one space OR one non-alphabet char OR any number, so multiple spaces will
				  	still remains, e.g. take a      5h1t will be take a      ****
		*/
		if len(pattern) <= 6 {
			regexPattern += `(^|\s{1}|[^\w]{1}|\d+)`
		}
		for idx, ch := range pattern {
			word := string(ch)
			switch word {
			case "a":
				word = "[a4]"
			case "i":
				word = "[i1]"
			case "e":
				word = "[e3]"
			case "o":
				word = "[o0]"
			case "l":
				word = "[l1]"
			case "k":
				word = "[kq]"
			case "p":
				word = "[pv]"
			case "s":
				word = "[s5]"
			case "b":
				word = "[b3]"
			}
			if idx < len(pattern)-1 {
				regexPattern = regexPattern + word + "+[\\W|_]*"
			} else {
				if word == "t" || word == "d" {
					word = "[td]"
				}
				if len(pattern) <= 6 {
					regexPattern += word + `+($|\s{1}|[^\w]{1}|\d+)`
				} else {
					regexPattern += word + "+"
				}
			}
		}

		regexAll = regexAll + "|" + regexPattern
	}
	if regexAll == "" {
		logrus.Errorf("Word filter regex is empty")
	}
	regexAll = regexAll[1:]

	censorRegex, err := regexp.Compile(regexAll)
	if err != nil {
		logrus.Errorf("Failed to build censor regex: %s", err.Error())
	}

	return censorRegex, nil
}
