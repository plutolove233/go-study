package utils

import "strconv"

type ReduceFunc func(string, []string) string

func Reduce(key string, value []string) string {
	return strconv.Itoa(len(value))
}
