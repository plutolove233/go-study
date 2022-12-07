package main

type Calculation struct {
	A int
	B int
}

func (c *Calculation) CallBack(fn func(tx *Calculation) int) int {
	return fn(c)
}

func (c *Calculation) Trans(fn func(tx *Calculation) int) int {
	return c.CallBack(fn)
}

func main() {
	c := &Calculation{
		A: 2,
		B: 3,
	}
	println(c.Trans(func(tx *Calculation) int {
		return tx.A * tx.B
	}), c.Trans(func(tx *Calculation) int {
		return tx.A + tx.B
	}))
}
