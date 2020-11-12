package calculator_test

import (
	"calculator"
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

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T){
	t.Parallel()
	tests:= []struct{
		name string
		a float64
		b float64
		want float64
	}{
		{
			name:"Test multiplying 2 positive numbers",
			a:2,
			b:3,
			want:6,
		},
		{
			name:"Test multiplying 0 and a negative number",
			a:0,
			b:-1,
			want:0,
		},
		{
			name:"Test multiplying 2 negative numbers",
			a:-1,
			b:-1,
			want:1,
		},		
	}

	// 	got4 := calculator.Multiply(2)
	// if want2 != got2 && want2 != got4{
	// 	t.Errorf("want %f, got%f",want2,got2)
	// }
	for _,tt:= range tests{
		t.Run(tt.name,func(t *testing.T){
			got:=calculator.Multiply(tt.a,tt.b)
			if got != tt.want{
				t.Errorf("want %g; got %g",tt.want,got)
			} 
		})
	}


}