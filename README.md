# junit2html

Convert Junit XML reports (`junit.xml`) into HTML reports using a single standalone binary.

* Standalone binary.
* Failed tests are top, that's what's important.
* No JavaScript.
* Look gorgeous.

## Screenshot

![screenshot](screenshot.png)

## Install

Like `jq`, `junit2html` is a tiny (8Mb) standalone binary. You can download it from the [releases page](https://github.com/kitproj/junit2html/releases/latest).

If you're on MacOS, you can use `brew`:

```bash
brew tap kitproj/junit2html --custom-remote https://github.com/kitproj/junit2html
brew install junit2html
```

Otherwise, you can use `curl`:

```bash
curl -q https://raw.githubusercontent.com/kitproj/junit2html/main/install.sh | sh
```

## Usage

Here is an example that uses trap to always created the test report:

```bash
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
