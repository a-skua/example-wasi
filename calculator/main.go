package main

import (
	"github.com/a-skua/example-wasi/calculator/internal/example/calculator/calc"
)

func init() {
	calc.Exports.Add = add
}

func add(a, b int32) int32 {
	return a + b
}

func main() {}