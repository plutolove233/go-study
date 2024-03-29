package main

// NoDiff checks whether or not a collection
// of values are all identical.
func NoDiff[V comparable](vs ...V) bool {
	if len(vs) == 0 {
		return true
	}

	v := vs[0]
	for _, x := range vs[1:] {
		if v != x {
			return false
		}
	}
	return true
}

func main() {
	var NoDiffString = NoDiff[string]
	println(NoDiff("Go", "Go", "Go")) // true
	println(NoDiffString("Go", "go")) // false

	println(NoDiff(123, 123, 123, 123)) // true
	println(NoDiff[int](123, 123, 789)) // false

	type A = [2]int
	println(NoDiff(A{}, A{}, A{}))     // true
	println(NoDiff(A{}, A{}, A{1, 2})) // false

	println(NoDiff(new(int)))           // true
	println(NoDiff(new(int), new(int))) // false

	println(NoDiff[bool]())   // true

	// _ = NoDiff() // error: cannot infer V

	// error: *** does not implement comparable
	// _ = NoDiff([]int{}, []int{})
	// _ = NoDiff(map[string]int{})
	// _ = NoDiff(any(1), any(1))
}