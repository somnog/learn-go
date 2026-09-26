package main

import "cmp"

type B interface {
}

// Generics
func Add[T cmp.Ordered](a T, b T) T {
	return a + b
}
