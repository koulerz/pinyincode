package pinyincode

import (
	py "github.com/mozillazg/go-pinyin"
)

const (
	FlagWithEnglish    = 0b0000_0001 // 保留英文字母
	FlagWithNumber     = 0b0000_0010 // 保留数字
	FlagWithHyphen     = 0b0000_0100 // 保留连字符
	FlagWithUnderscore = 0b0000_1000 // 保留下划线

	FlagOnlyPinyin = 0b1111_1111 // 仅保留中文拼音
)

var (
	charactersOnlyPinyin []string
	charactersEnglish    = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}
	charactersNumber     = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	charactersHyphen     = []string{"-"}
	charactersUnderscore = []string{"_"}
)

// Code 获取拼音码，通过 flag 自定义保留的字符
func Code(hans string, flag int) string {
	rules := func(r rune, a py.Args) []string {
		var validCharacter = getValidCharactersByFlag(flag)
		s := string(r)
		if !inStringSlice(s, validCharacter) {
			return []string{}
		}
		return []string{s}
	}

	handler := py.NewArgs()
	handler.Style = py.FirstLetter
	handler.Separator = ""
	handler.Fallback = rules
	return py.Slug(hans, handler)
}

// getValidCharactersByFlag 通过 flag 获取有效的字符
func getValidCharactersByFlag(flag int) []string {
	var characters = charactersOnlyPinyin
	if (FlagOnlyPinyin & flag) == FlagOnlyPinyin {
		return charactersOnlyPinyin
	}
	if (FlagWithEnglish & flag) == FlagWithEnglish {
		characters = append(characters, charactersEnglish...)
	}
	if (FlagWithNumber & flag) == FlagWithNumber {
		characters = append(characters, charactersNumber...)
	}
	if (FlagWithHyphen & flag) == FlagWithHyphen {
		characters = append(characters, charactersHyphen...)
	}
	if (FlagWithUnderscore & flag) == FlagWithUnderscore {
		characters = append(characters, charactersUnderscore...)
	}
	return characters
}

// inStringSlice 验证字符串是否在字符串列表值中
func inStringSlice(s string, ss []string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}
