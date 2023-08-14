package hw02unpackstring

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	runeArr := []rune(str)

	newRuneArr := []rune{}
	if (len(runeArr) > 0) && (unicode.IsDigit(runeArr[0])) {
		return "", ErrInvalidString
	}

	for i := 0; i < len(runeArr); i++ {
		tempRune := runeArr[i]
		if (i+1 < len(runeArr)) && (unicode.IsDigit(runeArr[i+1])) {
			digtIndex := i + 1
			if (i+2 < len(runeArr)) && (unicode.IsDigit(runeArr[i+2])) {
				return "", ErrInvalidString
			}
			digt := int(runeArr[digtIndex] - '0')
			for j := 0; j < digt; j++ {
				newRuneArr = append(newRuneArr, tempRune)
			}
			i++
		} else {
			newRuneArr = append(newRuneArr, tempRune)
		}
	}

	return string(newRuneArr), nil
}
