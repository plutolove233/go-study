package main

func GetParamsUntilLow(s []rune) func( func(rune) bool) {
	return func(yield func(rune) bool) {
		for _, item := range s {
			if item >= 'A' && item <= 'B' {
				yield(item)
			} else {
				break
			}
		}
	}
}

func main() {
	s := []rune{'A', 'B', 'c', 'D'}
	for item := range GetParamsUntilLow(s) {
		println(item)
	}
}