package main

import (
	"encoding/xml"
	"fmt"
	"github.com/jstemmer/go-junit-report/formatter"
	"os"
)

func main() {
	suites := &formatter.JUnitTestSuites{}

	err := xml.NewDecoder(os.Stdin).Decode(suites)
	if err != nil {
		panic(err)
	}

	fmt.Println("<html>")
	fmt.Println("<head>")
	fmt.Println("<meta charset=\"UTF-8\">")
	fmt.Println("<style>")
	fmt.Println("body {font-family: sans-serif}")
	fmt.Println("</style>")
	fmt.Println("</head>")
	fmt.Println("<body>")
	failures, total := 0, 0
	for _, s := range suites.Suites {
		failures += s.Failures
		total += len(s.TestCases)
	}
	fmt.Printf("<p>%d of %d tests failed</p>\n", failures, total)
	for _, s := range suites.Suites {
		if s.Failures > 0 {
			fmt.Printf("<h2>%s</h2>\n", s.Name)
			for _, c := range s.TestCases {
				if f := c.Failure; f != nil {
					fmt.Printf("<p><span style='color:red'>𒒬</span> %s</p>\n", c.Name)
					fmt.Printf("<pre>%s</pre>\n", f.Contents)
				}
			}
		}
	}
	for _, s := range suites.Suites {
		fmt.Printf("<h2>%s</h2>\n", s.Name)
		for _, c := range s.TestCases {
			if c.Failure == nil {
				fmt.Printf("<p><span style='color:green'>✔</span> %s</p>\n", c.Name)
			}
		}
	}
	fmt.Println("</body>")
	fmt.Println("</html>")
}
