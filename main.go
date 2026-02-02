package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

// Part-1
// Factorial calculates n!
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// IsPrime returns true if n is a prime number
func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}
	if n == 2 {
		return true, nil
	}
	if n%2 == 0 {
		return false, nil
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

// Power calculates base^exponent
func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

// Part-2
// MakeCounter
func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

// MakeMultiplier
func MakeMultiplier(factor int) func(int) int {
	return func(input int) int {
		return input * factor
	}
}

// MakeAccumulator
func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	state := initial

	add = func(n int) {
		state += n
	}
	subtract = func(n int) {
		state -= n
	}
	get = func() int {
		return state
	}

	return add, subtract, get
}

// Part-3
// Apply returns a new slice
func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = operation(v)
	}
	return result
}

// Filter returns a new slice
func Filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce collapses a slice into a single value
func Reduce(nums []int, initial int, operation func(int, int) int) int {
	accumulator := initial
	for _, v := range nums {
		accumulator = operation(accumulator, v)
	}
	return accumulator
}

// Compose returns a function f(g(x))
func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// Part-4
// Demonstrates process metadata and memory isolation.
func ExploreProcess() {
	//Get Process IDs
	pid := os.Getpid()
	ppid := os.Getppid()

	fmt.Printf("Current Process ID: %d\n", pid)
	fmt.Printf("Parent Process ID: %d\n", ppid)

	//Create sample data
	data := []int{1, 2, 3, 4, 5}

	// Print memory addresses
	fmt.Printf("Slice Header Address: %p\n", &data)
	fmt.Printf("First Element Address: %p\n", &data[0])

	// Explanation of Isolation
	fmt.Println("\n--- Security Note ---")
	fmt.Println("Other processes cannot access these memory addresses.")
	fmt.Println("Each process operates in its own virtual address space,")
	fmt.Println("preventing one program from accidentally (or maliciously)")
	fmt.Println("reading or overwriting another program's data.")
}

// part-5
func DoubleValue(x int) {
	x = x * 2
}

func DoublePointer(x *int) {
	*x = *x * 2
}

func CreateOnStack() int {
	x := 42
	return x
}

func CreateOnHeap() *int {
	x := 42
	return &x
}

// SwapValues swaps two values and returns them.
func SwapValues(a, b int) (int, int) {
	return b, a
}

// SwapPointers swaps the values that two pointers point to.
func SwapPointers(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// AnalyzeEscape calls the creation functions to trigger compiler analysis.
func AnalyzeEscape() {
	s := CreateOnStack()
	h := CreateOnHeap()

	fmt.Printf("Stack value: %d, Heap value pointer: %p\n", s, h)
}

/*
ESCAPE ANALYSIS REPORT:
1. Which variables escaped to the heap?
   The variable 'x' inside CreateOnHeap() escaped to the heap. Depending on
   compiler optimizations, the arguments passed to fmt.Printf (like 's' and 'h')
   may also escape because they are passed to an interface{} parameter.

2. Why did they escape?
   In CreateOnHeap(), the function returns the memory address (&x) of a local
   variable. If this variable remained on the stack, it would be invalidated
   as soon as the function returned. To keep the data alive for the caller,
   the Go compiler moves ('escapes') it to the heap.

3. What does "escapes to heap" mean?
   It means the compiler has decided to allocate the memory for a variable in
   the dynamic heap storage rather than the function's local stack frame.
   Stack memory is automatically reclaimed when a function ends, while heap
   memory is managed by the Garbage Collector (GC) to ensure pointers
   remain valid as long as they are being referenced.
*/

func main() {
	fmt.Println("=== Part 4: Process Explorer ===")
	ExploreProcess()

}
