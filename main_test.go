package main

import (
	"reflect"
	"testing"
)

// Part-1
func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{name: "factorial of 0", input: 0, want: 1, wantErr: false},
		{name: "factorial of 1", input: 1, want: 1, wantErr: false},
		{name: "factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "factorial of 10", input: 10, want: 3628800, wantErr: false},
		{name: "negative input error", input: -1, want: 0, wantErr: true},
		{name: "large negative input error", input: -10, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{name: "minimum prime 2", input: 2, want: true, wantErr: false},
		{name: "small prime 3", input: 3, want: true, wantErr: false},
		{name: "even composite 4", input: 4, want: false, wantErr: false},
		{name: "larger prime 17", input: 17, want: true, wantErr: false},
		{name: "large composite 100", input: 100, want: false, wantErr: false},
		{name: "error for input 1", input: 1, want: false, wantErr: true},
		{name: "error for negative", input: -5, want: false, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool
	}{
		{name: "positive power", base: 2, exponent: 3, want: 8, wantErr: false},
		{name: "base to power 0", base: 5, exponent: 0, want: 1, wantErr: false},
		{name: "0 to positive power", base: 0, exponent: 5, want: 0, wantErr: false},
		{name: "1 to any power", base: 1, exponent: 100, want: 1, wantErr: false},
		{name: "negative base even power", base: -2, exponent: 2, want: 4, wantErr: false},
		{name: "negative exponent error", base: 2, exponent: -1, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Power() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Part-2
func TestMakeCounter(t *testing.T) {
	tests := []struct {
		name       string
		startValue int
		increments int
		want       int
	}{
		{name: "start at 0, increment 3 times", startValue: 0, increments: 3, want: 3},
		{name: "start at 10, increment 1 time", startValue: 10, increments: 1, want: 11},
		{name: "start at -5, increment 2 times", startValue: -5, increments: 2, want: -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := MakeCounter(tt.startValue)
			var got int
			for i := 0; i < tt.increments; i++ {
				got = counter()
			}
			if got != tt.want {
				t.Errorf("%s: after %d increments got %d, want %d", tt.name, tt.increments, got, tt.want)
			}
		})
	}

	// Independence check
	t.Run("counters are independent", func(t *testing.T) {
		c1 := MakeCounter(0)
		c2 := MakeCounter(0)
		c1() // becomes 1
		if c1Val, c2Val := c1(), c2(); c1Val == c2Val {
			t.Errorf("expected independent counters, but both returned %d", c1Val)
		}
	})
}

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		input  int
		want   int
	}{
		{name: "double 5", factor: 2, input: 5, want: 10},
		{name: "triple 10", factor: 3, input: 10, want: 30},
		{name: "multiply by zero", factor: 0, input: 100, want: 0},
		{name: "multiply by negative", factor: -2, input: 4, want: -8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mult := MakeMultiplier(tt.factor)
			if got := mult(tt.input); got != tt.want {
				t.Errorf("%s: got %d, want %d", tt.name, got, tt.want)
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	type op struct {
		isAdd bool // true for add, false for subtract
		val   int
	}

	tests := []struct {
		name    string
		initial int
		ops     []op
		want    int
	}{
		{
			name:    "basic add and subtract",
			initial: 100,
			ops:     []op{{true, 50}, {false, 30}}, // 100 + 50 - 30
			want:    120,
		},
		{
			name:    "sequential additions",
			initial: 0,
			ops:     []op{{true, 10}, {true, 20}, {true, 30}},
			want:    60,
		},
		{
			name:    "result goes negative",
			initial: 10,
			ops:     []op{{false, 20}},
			want:    -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add, sub, get := MakeAccumulator(tt.initial)
			for _, operation := range tt.ops {
				if operation.isAdd {
					add(operation.val)
				} else {
					sub(operation.val)
				}
			}
			if got := get(); got != tt.want {
				t.Errorf("%s: got %d, want %d", tt.name, got, tt.want)
			}
		})
	}
}

// Part-3
func TestApply(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		operation func(int) int
		want      []int
	}{
		{"double all", []int{1, 2, 3}, func(x int) int { return x * 2 }, []int{2, 4, 6}},
		{"square all", []int{4, 5, -2}, func(x int) int { return x * x }, []int{16, 25, 4}},
		{"negate all", []int{10, -5, 0}, func(x int) int { return -x }, []int{-10, 5, 0}},
		{"add one", []int{-1, 0, 1}, func(x int) int { return x + 1 }, []int{0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.nums, tt.operation)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		want      []int
	}{
		{"only evens", []int{1, 2, 3, 4, 5, 6}, func(x int) bool { return x%2 == 0 }, []int{2, 4, 6}},
		{"only positives", []int{-2, -1, 0, 1, 2}, func(x int) bool { return x > 0 }, []int{1, 2}},
		{"greater than ten", []int{5, 10, 15, 20}, func(x int) bool { return x > 10 }, []int{15, 20}},
		{"no matches", []int{1, 3, 5}, func(x int) bool { return x%2 == 0 }, nil}, // returns empty or nil
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.nums, tt.predicate)
			// Handle empty slice vs nil slice if necessary
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		initial   int
		operation func(int, int) int
		want      int
	}{
		{"summation", []int{1, 2, 3, 4}, 0, func(acc, curr int) int { return acc + curr }, 10},
		{"product", []int{1, 2, 3, 4}, 1, func(acc, curr int) int { return acc * curr }, 24},
		{"find maximum", []int{5, 12, 3, 9}, 0, func(acc, curr int) int {
			if curr > acc {
				return curr
			}
			return acc
		}, 12},
		{"find minimum", []int{5, 12, 3, 9}, 100, func(acc, curr int) int {
			if curr < acc {
				return curr
			}
			return acc
		}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.nums, tt.initial, tt.operation)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	addOne := func(x int) int { return x + 1 }
	triple := func(x int) int { return x * 3 }
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	tests := []struct {
		name  string
		f     func(int) int
		g     func(int) int
		input int
		want  int
	}{
		{"triple then add one: (5*3)+1", addOne, triple, 5, 16},
		{"add one then triple: (5+1)*3", triple, addOne, 5, 18},
		{"negate then absolute: abs(-10)", abs, func(x int) int { return -x }, 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := Compose(tt.f, tt.g)
			if got := fn(tt.input); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

// Part-5
func TestSwapValues(t *testing.T) {
	tests := []struct {
		name  string
		a, b  int
		wantA int
		wantB int
	}{
		{"swap small positives", 5, 10, 10, 5},
		{"swap negatives", -1, -5, -5, -1},
		{"swap zero", 0, 100, 100, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := SwapValues(tt.a, tt.b)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf("SwapValues(%d, %d) = (%d, %d); want (%d, %d)",
					tt.a, tt.b, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}

func TestSwapPointers(t *testing.T) {
	tests := []struct {
		name  string
		a, b  int
		wantA int
		wantB int
	}{
		{"swap by pointer basic", 1, 2, 2, 1},
		{"swap identical values", 10, 10, 10, 10},
		{"swap large values", 1000, 5000, 5000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valA, valB := tt.a, tt.b
			SwapPointers(&valA, &valB)
			if valA != tt.wantA || valB != tt.wantB {
				t.Errorf("After SwapPointers, a=%d, b=%d; want a=%d, b=%d",
					valA, valB, tt.wantA, tt.wantB)
			}
		})
	}
}
