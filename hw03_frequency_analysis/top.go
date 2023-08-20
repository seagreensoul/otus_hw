package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

func Top10(str string) []string {
	str = strings.ToLower(str)
	stringArr := strings.Fields(str)
	frecMap := make(map[string]int)

	for _, s := range stringArr {
		word := strings.TrimRightFunc(s, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsMark(r)
		})
		frecMap[word]++
	}

	delete(frecMap, "")

	keys := make([]string, 0, len(frecMap))

	for key := range frecMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return frecMap[keys[i]] > frecMap[keys[j]]
	})

	var a []string

	for i := 0; i < len(keys); i++ {
		if i < 10 {
			a = append(a, keys[i])
		}
	}

	return a
}
