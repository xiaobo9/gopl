// Echo2 prints its command-line arguments.
package echo

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo4() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func echoT1() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func echoT2() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d\t%s\n", i, arg)
	}

}

func Echo() {
	echo4()
}
