package utils

import (
	"distributed-system/map-reduce/models"
	"strings"
	"unicode"
)

type MapFunc func(string, string) []models.KeyValue

func Map(filename string, contents string) []models.KeyValue {
	// r是字母字符返回false，否则返回true
	// 目的是将文件的内容分割成多个字母字符的分片
	ff := func(r rune) bool {
		return !unicode.IsLetter(r)
	}

	words := strings.FieldsFunc(contents, ff)
	kva := []models.KeyValue{}
	for _, w := range words {
		kv := models.KeyValue{
			Key:   w,
			Value: "1", // question: why is signed with "1"
		}
		kva = append(kva, kv)
	}
	return kva
}
