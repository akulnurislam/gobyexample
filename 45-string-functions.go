package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains  :", s.Contains("test", "es"))
	p("Count     :", s.Count("test", "t"))
	p("HasPrefix :", s.HasPrefix("test", "te"))
	p("HasSuffix :", s.HasSuffix("test", "st"))
	p("Index     :", s.Index("test", "e"))
	p("Join      :", s.Join([]string{"a", "b"}, "-"))
	p("Repeat    :", s.Repeat("a", 5))
	p("Replace   :", s.Replace("foo", "o", "0", -1))
	p("Replace   :", s.Replace("foo", "o", "0", 1))
	p("Split     :", s.Split("a-b-c-d-e", "-"))
	p("ToLower   :", s.ToLower("TEST"))
	p("ToUpper   :", s.ToUpper("test"))
	p()
	p("Len       :", len("hello"))
	p("Char      :", "hello"[1])

	arr := s.Split(",,,,A,,,,,,B,,,,,C,D,,,,,", ",")
	filter(&arr, func(a string) bool {
		return a == ""
	})
	p("Filter    :", arr)
}

func filter(arr *[]string, f func(string) bool) {
	var lastIndex int = 0
	var travIndex int = 1

	if len(*arr) == 1 && (*arr)[lastIndex] == "" {
		*arr = (*arr)[:len(*arr)-1]
		return
	} else if len(*arr) == 1 {
		return
	}

	for {
		if (*arr)[lastIndex] == "" && (*arr)[travIndex] != "" {
			(*arr)[lastIndex], (*arr)[travIndex] = (*arr)[travIndex], (*arr)[lastIndex]
			lastIndex++
		} else if (*arr)[lastIndex] != "" && (*arr)[travIndex] == "" {
			lastIndex = travIndex
		}

		if travIndex == len(*arr)-1 && (*arr)[travIndex] == "" {
			*arr = (*arr)[:lastIndex]
			return
		} else if travIndex == len(*arr)-1 {
			return
		}

		travIndex++
	}
}
