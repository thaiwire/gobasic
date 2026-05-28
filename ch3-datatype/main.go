package main

import (
	"fmt"
	"html"
)

func main() {

	// datatype in GO
	// มี 2 ประเภท คือ
	// 1. Primitive data type (int, float, string, bool)
	// 2. Composite data type (array, slice, map, struct)

	// basic type
	// Numeric type
	var a int8 = 127                    // range -128 to 127
	var b int16 = 32767                 // range -32768 to 32767
	var c int32 = 2147483647            // range -2147483648 to 2147483647
	var d int64 = 9223372036854775807   // range -9223372036854775808 to 9223372036854775807
	var e int = 100                     // int is platform dependent (32 or 64 bit)
	var f uint8 = 255                   // range 0 to 255
	var g uint16 = 65535                // range 0 to 65535
	var h uint32 = 4294967295           // range 0 to 4294967295
	var i uint64 = 18446744073709551615 // range 0 to 18446744073709551615
	// Floating point type
	var k float32 = 3.14                   // range 1.18e-38 to 3.4e+38
	var l float64 = 3.14159265358979323846 // range 2.23e-308 to 1.8e+308

	// string type
	var m string = "Hello, World!"
	// boolean type
	var n bool = true

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)
	fmt.Println("g:", g)
	fmt.Println("h:", h)
	fmt.Println("i:", i)
	fmt.Println("k:", k)
	fmt.Println("l:", l)
	fmt.Println("m:", m)
	fmt.Println("n:", n)

	// Interpret string
	var o string = "123 \n\n"
	fmt.Println("o:", o)

	// raw string
	var p string = `This is a raw string literal.
It can contain multiple lines and does not interpret escape sequences like \n or \t.`
	fmt.Println("p:", p)

	var p1 string = `AAABBB
	CCC
	DDD`
	fmt.Println("p1:", p1)

	// \n is a newline character, it will create a new line when printed
	var q string = "Hello\nWorld"
	fmt.Println("q:", q)

	// \t is a tab character, it will create a horizontal tab when printed
	var r string = "Hello\tWorld"
	fmt.Println("r:", r)
	// \\ is a backslash character, it will print a single backslash when printed
	var s string = "Hello\\World"
	fmt.Println("s:", s)
	// \" is a double quote character, it will print a single double quote when printed
	var t string = "Hello \"World\""
	fmt.Println("t:", t)

	// \' is a single quote character, it will print a single quote when printed
	var u string = "Hello 'World'"
	fmt.Println("u:", u)

	// \r is a carriage return character, it will move the cursor to the beginning of the line when printed
	var v string = "Hello\rWorld"
	fmt.Println("v:", v)

	// \b is a backspace character, it will move the cursor back one position when printed
	var w string = "Hello\bWorld"
	fmt.Println("w:", w)

	// escape HTML
	var x string = "<div>Hello, World!</div>"
	fmt.Println("x:", x)
	fmt.Println("escaped x:", html.EscapeString(x))

	// unescape HTML
	var y string = "&lt;div&gt;Hello, World!&lt;/div&gt;"
	fmt.Println("y:", y)
	fmt.Println("unescaped y:", html.UnescapeString(y))

}
