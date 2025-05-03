package main

import "fmt"

// TODO:
// Is it possible to have List as a defined type?
// So we can define Map() as a receiver function.

type List[T any] = []T

type Mapper[T any, U any] = func(T) U

func Map[T any, U any](l *List[T], fn Mapper[T, U]) *List[U] {
	var out List[U]
	for _, v := range *l {
		out = append(out, fn(v))
	}

	return &out
}

func main() {
	nums := List[int]{1, 2, 3}
	squares := Map(&nums, func(n int) int { return n * n })
	fmt.Printf("%v\n", *squares)
}
