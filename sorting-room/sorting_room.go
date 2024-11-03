package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if fancyNumber, ok := fnb.(FancyNumber); ok {
		if num, err := strconv.Atoi(fancyNumber.n); err == nil {
			return num
		}
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	if fancyNumber, ok := fnb.(FancyNumber); ok {
		num, err := strconv.Atoi(fancyNumber.n)
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(num))
	}
	return "This is a fancy box containing the number 0.0"
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	if a, ok := i.(int); ok {
		return DescribeNumber(float64(a))
	}
	if b, ok := i.(float64); ok {
		return DescribeNumber(b)
	}
	if c, ok := i.(NumberBox); ok {
		return DescribeNumberBox(c)
	}
	if d, ok := i.(FancyNumberBox); ok {
		return DescribeFancyNumberBox(d)
	}
	return "Return to sender"
}
