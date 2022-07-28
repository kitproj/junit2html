package main

import (
	"bytes"
	_ "embed"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jstemmer/go-junit-report/formatter"
)

//go:embed style.css
var styles string

func printTest(s formatter.JUnitTestSuite, c formatter.JUnitTestCase) {
	id := fmt.Sprintf("%s.%s.%s", s.Name, c.Classname, c.Name)
	class, text := "passed", "Pass"
	f := c.Failure
	if f != nil {
		class, text = "failed", "Fail"
	}
	k := c.SkipMessage
	if k != nil {
		class, text = "skipped", "Skip"
	}
	fmt.Fprintf(testReport, "<div class='%s' id='%s'>\n", class, id)
	fmt.Fprintf(testReport, "<a href='#%s'>%s <span class='badge'>%s</span></a>\n", id, c.Name, text)
	fmt.Fprintf(testReport, "<div class='expando'>\n")
	if f != nil {
		fmt.Fprintf(testReport, "<div class='content'>%s</div>\n", f.Contents)
	} else if k != nil {
		fmt.Fprintf(testReport, "<div class='content'>%s</div>\n", k.Message)
	}
	d, _ := time.ParseDuration(c.Time)
	fmt.Fprintf(testReport, "<p class='duration' title='Test duration'>%v</p>\n", d)
	fmt.Fprintf(testReport, "</div>\n")
	fmt.Fprintf(testReport, "</div>\n")
}

var (
	flagResult = flag.String("xml-result", "", "xml formatted junit test result")
	flagReport = flag.String("html-report", "", "html report")
)

var (
	testReport io.Writer
)

func main() {
	// Setup
	flag.Parse()
	var testResult io.Reader
	// Read Input (test result)
	//
	testResult = os.Stdin
	if *flagResult != "" {
		res, err := ioutil.ReadFile(*flagResult)
		if err != nil {
			panic(err)
		}
		testResult = bytes.NewReader(res)
	}
	// Output
	//
	testReport = os.Stdout
	if *flagReport != "" {
		f, err := os.OpenFile(*flagReport, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			panic(err)
		}
		testReport = f
	}

	suites := &formatter.JUnitTestSuites{}

	err := xml.NewDecoder(testResult).Decode(suites)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(testReport, "<html>")
	fmt.Fprintln(testReport, "<head>")
	fmt.Fprintln(testReport, "<meta charset=\"UTF-8\">")
	fmt.Fprintln(testReport, "<style>")
	fmt.Fprintln(testReport, styles)
	fmt.Fprintln(testReport, "</style>")
	fmt.Fprintln(testReport, "</head>")
	fmt.Fprintln(testReport, "<body>")
	failures, total := 0, 0
	for _, s := range suites.Suites {
		failures += s.Failures
		total += len(s.TestCases)
	}
	fmt.Fprintf(testReport, "<p>%d of %d tests failed</p>\n", failures, total)
	for _, s := range suites.Suites {
		if s.Failures > 0 {
			printSuiteHeader(s)
			for _, c := range s.TestCases {
				if f := c.Failure; f != nil {
					printTest(s, c)
				}
			}
		}
	}
	for _, s := range suites.Suites {
		printSuiteHeader(s)
		for _, c := range s.TestCases {
			if c.Failure == nil {
				printTest(s, c)
			}
		}
	}
	fmt.Fprintln(testReport, "</body>")
	fmt.Fprintln(testReport, "</html>")
}

func printSuiteHeader(s formatter.JUnitTestSuite) {
	fmt.Fprintln(testReport, "<h4>")
	fmt.Fprintln(testReport, s.Name)
	for _, p := range s.Properties {
		if strings.HasPrefix(p.Name, "coverage.") {
			v, _ := strconv.ParseFloat(p.Value, 10)
			fmt.Fprintf(testReport, "<span class='coverage' title='%s'>%.0f%%</span>\n", p.Name, v)
		}
	}
	fmt.Fprintln(testReport, "</h4>")
}
