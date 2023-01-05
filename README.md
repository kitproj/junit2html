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

trap 'go-junit-report < test.out > junit.xml && junit2html --xmlReports junit.xml > test-report.html' EXIT

go test -v -cover ./... 2>&1 > test.out
```

💡 Don't use pipes (i.e. `|`) in shell, pipes swallow exit codes. Use `>` which is POSIX compliant.

## Test

How to test this locally:

```bash
go test -v -cover ./... 2>&1 > test.out
go-junit-report < test.out > junit.xml 
go run .  --xmlReports junit.xml > test-report.html 
```

## Using glob patterns:

Sometimes there is a need to parse multiple xml files and generate single html report.
`junit2html` supports that by using standard [`glob` expression](https://pkg.go.dev/path/filepath#Glob).

```bash
# Explicit single file
junit2html --xmlReports "reports/junit.xml" > report.html

# Multiple files
junit2html --xmlReports "reports/junit.xml,reports/coverage.xml" > report.html

# Single glob pattern
junit2html --xmlReports "reports/*.xml" > report.html

# Multiple glob patterns
junit2html --xmlReports "reports/junit*.xml,reports/coverage*.xml" > report.html
``` 
