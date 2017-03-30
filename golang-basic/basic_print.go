// basic_print.go
package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	fmt.Println("----------")

	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))

	fmt.Println("----------")

	type T struct {
		a int
		b float64
		c string
	}

	var timeZone = map[string]int{"UTC": 0 * 60 * 60, "CST": -6 * 60 * 60}

	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", timeZone)

	fmt.Println("----------")

	fmt.Printf("%T\n", timeZone)
}
