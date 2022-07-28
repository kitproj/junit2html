# junit2html

Convert Junit XML reports (`junit.xml`) into HTML reports using Golang.

* Standalone binary.
* Failed tests are top, that's what's important.
* No JavaScript.
* Look gorgeous.

## Screenshot

![screenshot](screenshot.png)

## Usage

Here is an example that uses trap to always created the test report:

```bash
go install github.com/jstemmer/go-junit-report@latest
go install github.com/alexec/junit2html@latest

trap 'go-junit-report < test.out > junit.xml && junit2html < junit.xml > test-report.html' EXIT

go test -v -cover ./... 2>&1 > test.out
```

💡 Don't use pipes (i.e. `|`) in shell, pipes swallow exit codes. Use `<` and `>` which is POSIX compliant.

## Test

How to test this locally:

```bash
go test -v -cover ./... 2>&1 > test.out
go-junit-report < test.out > junit.xml 
go run . < junit.xml > test-report.html 
# alternatively using flags 
go run . -xml-result junit.xml -html-report test-report.html

# optional for Windows
env GOOS=windows GOARCH=amd64 go build -o ~/junit2html.exe
junit2html.exe -xml-result junit.xml -html-report test-report.html
```
