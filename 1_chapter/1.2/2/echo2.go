// Echo2 выводит аргументы командной строки
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	//var s, sep string
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	s := strings.Join(os.Args[1:], " ")

	fmt.Println(s)
	fmt.Print(time.Since(start))
}
