# junit2html

Convert Junit XML reports (`junit.xml`) into a simple HTML report. No JavaScript.

```
go install github.com/jstemmer/go-junit-report@latest
go install github.com/alexec/junit2html@latest

trap 'cat test.out | go-junit-report | junit2html > test-report.html' EXIT

go test -v ./... 2>&1 > test.out
```