package calculator_test

import (
	"calculator"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	r := rand.New(rand.NewSource(99))

	for i := 0; i < 100; i++ {
		a := r.ExpFloat64()
		b := r.ExpFloat64()
		want := a + b
		name := fmt.Sprintf("TEST adding 2 numbers: %s", strconv.Itoa(i))
		t.Run(name, func(t *testing.T) {
			got := calculator.Add(a, b)
			if want != got {
				t.Errorf("want %f, got %f", want, got)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{
			name: "TEST multiplying 2 numbers",
			a:    5,
			b:    2,
			want: 10,
		},
		{
			name: "TEST multiplying 0 by a negative number",
			a:    0,
			b:    -1,
			want: 0,
		},
		{
			name: "TEST multiplying 2 negative numbers",
			a:    -1,
			b:    -1,
			want: 1,
		},
		{
			name: "TEST multiplying 1 by a fractional number",
			a:    1,
			b:    0.5,
			want: 0.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.Multiply(tt.a, tt.b)
			if tt.want != got {
				t.Errorf("want %f, got %f", tt.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		a           float64
		b           float64
		want        float64
		errExpected bool
	}{
		{
			name:        "TEST dividing 2 numbers",
			a:           6,
			b:           2,
			want:        3,
			errExpected: false,
		},
		{
			name:        "TEST dividing 6 by 0",
			a:           6,
			b:           0,
			want:        0,
			errExpected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Divide(tt.a, tt.b)

			errReceived := err != nil
			if tt.errExpected != errReceived {
				t.Fatalf("Divide(%f, %f): unexpected error status: %v", tt.a, tt.b, errReceived)
			}

			if tt.want != got {
				t.Errorf("want %f, got %f", tt.want, got)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		a           float64
		want        float64
		errExpected bool
	}{
		{
			name:        "Test negative input",
			a:           -1,
			want:        0,
			errExpected: true,
		},
		{
			name:        "Test normal input",
			a:           9,
			want:        3,
			errExpected: false,
		},
		{
			name:        "Test decimal input",
			a:           0.0003,
			want:        0.01732050807,
			errExpected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Sqrt(tt.a)
			errReceived := err != nil
			if tt.errExpected != errReceived {
				t.Fatalf("Sqrt(%f): unexpected error status: %v", tt.a, errReceived)
			}

			if !closeEnough(got, tt.want, 0.000001) {
				t.Errorf("want %f, got %f", tt.want, got)
			}
		})
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
