package main

import "fmt"

// Data types in Go:
//
// * int, int8, int16, int32, int64
//   - Signed integers (can be positive or negative)
//   - int size is platform-dependent (32-bit or 64-bit)
//
// * uint, uint8, uint16, uint32, uint64
//   - Unsigned integers (positive values only)
//
// * float32, float64
//   - Floating-point numbers (decimal values)
//
// * byte
//   - Alias for uint8
//   - Commonly used to represent raw binary data or ASCII characters
//
// * rune
//   - Alias for int32
//   - Represents a Unicode code point
//
// * bool
//   - true or false
//
// * string
//   - Immutable sequence of bytes
//   - Uses double quotes: "hello"
//
// * nil
//   - Zero value for pointers, slices, maps, channels, interfaces, and functions
//   - Similar to Python's None, but NOT a type itself

func main() {
	var x uint8 = 100
	fmt.Println(x)
}
