package main

import (
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
