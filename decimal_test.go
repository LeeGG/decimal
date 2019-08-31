package decimal

import (
	"fmt"
	"testing"
)



func TestDemo(t *testing.T) {
	a :=  NewDecFromFloat(0.0000001)
	b := new(Decimal)
	b.FromFloat64(0.00001)
	c := new(Decimal)
	d := NewDecFromFloat(0)

	if error := DecimalDiv(a, d, c, 10); error != nil {
		fmt.Println(error)
	}
	DecimalMul(b, a, c)
	fmt.Println(c)
	fmt.Println(c.String())
}


func TestFromInt(t *testing.T) {
	tests := []struct {
		input  int64
		output string
	}{
		{-12345, "-12345"},
		{-1, "-1"},
		{1, "1"},
		{-9223372036854775807, "-9223372036854775807"},
		{-9223372036854775808, "-9223372036854775808"},
	}
	for _, tt := range tests {
		dec := NewDecFromInt(tt.input)
		str := dec.ToString()
		fmt.Println(str)
		fmt.Println(dec.String())

	}
}

