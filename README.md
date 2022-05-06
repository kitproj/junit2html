# junit2html

Convert Junit XML reports (`junit.xml`) into HTML reports using Golang.

* Standalone binary.
* Failed tests are top, that's what's important.
* No JavaScript.
* Look gorgeous.

![screenshot](screenshot.png)

## Usage

```bash
go install github.com/jstemmer/go-junit-report@latest
go install github.com/alexec/junit2html@latest

trap 'go-junit-report < test.out > junit.xml && junit2html < junit.xml > test-report.html' EXIT

go test -v -cover ./... 2>&1 > test.out
```

## Test

```bash
go test -v -cover ./... 2>&1 > test.out
go-junit-report < test.out > junit.xml 
go run . < junit.xml > test-report.html 
```
