package main

import "fmt"

type Option struct {
	A string
	B string
	C int
}

func NewOption(opts ...OptionFunc) *Option {
	opt := defaultOption
	for _, o := range opts {
		o(opt)
	}
	return opt
}

func WithA(a string) OptionFunc {
	return func(option *Option) {
		option.A = a
	}
}

func WithB(b string) OptionFunc {
	return func(option *Option) {
		option.B = b
	}
}

func WithC(c int) OptionFunc {
	return func(option *Option) {
		option.C = c
	}
}

type OptionFunc func(*Option)

var (
	defaultOption = &Option{
		A: "A",
		B: "B",
		C: 100,
	}
)

func main() {
	x := NewOption(
		WithA("沙河娜扎"),
		WithC(250),
	)
	fmt.Println(x)
}
