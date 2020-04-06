package main

import (
	"fmt"
	"os"
	s1 "strings"
)

var p = fmt.Println

type point struct {
	x, y int
}

func stringsDemo() {

	p("Contains:", s1.Contains("test", "es"))
	p("Count:     ", s1.Count("test", "t"))
	p("HasPrefix: ", s1.HasPrefix("test", "te"))
	p("HasSuffix: ", s1.HasSuffix("test", "st"))
	p("Index:     ", s1.Index("test", "e"))
	p("Join:      ", s1.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s1.Repeat("a", 5))
	p("Replace:   ", s1.Replace("foo", "o", "0", -1))
	p("Replace:   ", s1.Replace("foo", "o", "0", 1))
	p("Split:     ", s1.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s1.ToLower("TEST"))
	p("ToUpper:   ", s1.ToUpper("test"))
	p()

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 33)
	fmt.Printf("%x\n", 456)
	fmt.Printf("%f\n", 78.9)
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")
	fmt.Printf("%p\n", &p)
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
