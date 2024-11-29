package utils

import "strings"

func SplitString(v, sep string) []string {
	vs := []string{}
	if v != "" {
		vs = strings.Split(v, sep)
	}
	return vs
}

func SliceStringContain(strs []string, str string) bool {
	for _, tmp := range strs {
		if tmp == str {
			return true
		}
	}
	return false
}
