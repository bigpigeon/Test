package main

func main() {
	a := make([]int, 0, 5)
	b := []int{1, 1, 1, 1, 1, 1, 1}
	copy(a, b)
}
