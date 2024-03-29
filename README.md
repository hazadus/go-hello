# Go

Here I learn the Go language.

## References

- Book: [Learning Go. An Idiomatic Approach to Real-World Go Programming](http://library.hazadus.ru/books/82/details/) by Jon Bodner. O'Reilly, 2021, 374 p.
	- 🔒 [My Notes](https://github.com/hazadus/Hazadus-Vault/blob/main/Dev/Reading/Learning%20Go%20(Bodner).md)
	- [Examples](https://github.com/hazadus/go-hello/tree/main/bodner)
- Book: [Pro Go: The Complete Guide to Programming Reliable and Efficient Software
  Using Golang](https://doi.org/10.1007/978-1-4842-7355-5) by Adam Freeman, ISBN 978-1-4842-7355-5, 2022, 1078 p.
- [The Go Programming Language Specification](https://go.dev/ref/spec) - This is the reference manual for the Go programming language.
- [Effective Go](https://go.dev/doc/effective_go) - This document gives tips for writing clear, idiomatic Go code.
- [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments).

## Basic Algorithms in Go

- Quicksort: [algorithms/quicksort.go](algorithms/quicksort.go)

## Learning Go, by Jon Bodner

- Undertanding capacity: [bodner/ch3/example_3_1.go](bodner/ch3/example_3_1.go)
- Slicing slices: [bodner/ch3/example_3_4.go](bodner/ch3/example_3_4.go)
- Slices with overlapping storage: [bodner/ch3/example_3_5.go](bodner/ch3/example_3_5.go)
- `append` makes overlapping slices more confusing: [bodner/ch3/example_3_6.go](bodner/ch3/example_3_6.go)
- Using `copy`: [bodner/ch3/example_3_6a.go](bodner/ch3/example_3_6a.go)
- Converting strings to slices: [bodner/ch3/example_3_9.go](bodner/ch3/example_3_9.go)
- Using a map: [bodner/ch3/example_3_10.go](bodner/ch3/example_3_10.go)
- Using a map as a set: [bodner/ch3/example_3_11.go](bodner/ch3/example_3_11.go)
- The for-range loop: [bodner/ch4/example_4_13.go](bodner/ch4/example_4_13.go)
- The `switch` statement: [bodner/ch4/example_4_19.go](bodner/ch4/example_4_19.go)
- Variadic Input Parameters and Slices, Multiple Return Values: [bodner/ch5/example_5_1.go](bodner/ch5/example_5_1.go)
- Functions As Values: [bodner/ch5/example_5_2a.go](bodner/ch5/example_5_2a.go)
- Closures: [bodner/ch5/example_5_2b.go](bodner/ch5/example_5_2b.go)
- Implementation of a binary tree that takes advantage of `nil` values for the receiver: [bodner/ch7/example_7_a.go](bodner/ch7/example_7_a.go)
- Using implicit inter‐ faces to compose applications via dependency injection: [bodner/ch7/example_7_b.go](bodner/ch7/example_7_b.go)

----

## Compiler commands

```bash
# Compile and create "hello" executable
go build hello.go

# Compile and run executable in temporary directory
go run hello.go
```

## Debugger

Install debugger from GitHub:

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

Start debug session:

```bash
~/go/bin/dlv debug hello.go
```

Creating a Breakpoint:

```
# Inside debugger (21 is number of line counting from `func main() {` line):
break bp1 main.main:21
# Breakpoint bp1 set at 0x104bc3c38 for main.main() ./hello.go:28

# Execution will be halted only when condition evaluates to true
condition bp1 i == 2
continue
# print <expr>
# set <varible> = <value>
# locals
# whatis <expr>
```

Debugging can be also started from VS code via `Run` menu.

## Code style. Linting Go code

References:

- [Effective Go](https://go.dev/doc/effective_go) - This document gives tips for writing clear, idiomatic Go code.
- [Go Code Style](https://google.github.io/styleguide/go/decisions) by Google.
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) - This page collects common comments made during reviews of Go code, so that a single detailed explanation can be referred to by shorthands. This is a laundry list of common style issues, not a comprehensive style guide.

### Using `revive` linter

Install: `go install github.com/mgechev/revive@latest`

Use: `~/go/bin/revive`

[Available rules](https://github.com/mgechev/revive#available-rules) for `revive` can be used to exclude some warnings.

We can also create [configuration file](https://github.com/mgechev/revive#configuration) to set global project-level rules for `revive`:

```
# revive.toml
ignoreGeneratedHeader = false
severity = "warning"
confidence = 0.8
errorCode = 0
warningCode = 0
[rule.blank-imports]
[rule.context-as-argument]
[rule.context-keys-type]
[rule.dot-imports]
[rule.error-return]
[rule.error-strings]
[rule.error-naming]
# Disable one rule for an example:
#[rule.exported]
[rule.if-return]
[rule.increment-decrement]
[rule.var-naming]
[rule.var-declaration]
[rule.package-comments]
[rule.range]
[rule.receiver-naming]
[rule.time-naming]
[rule.unexported-return]
[rule.indent-error-flow]
[rule.errorf]
```

Running the Linter with a Configuration File: `revive -config revive.toml`

## Configuring linter in VS Code

Open "Code" -> "Preferences" -> "Settings" -> "Extensions" -> "Go", and find "Lint tool" dropdown there.

As for me, it's set to `staticcheck` by default, but can be changed to `gofmt` or `revive`.

If you want to use `revive` with a custom configuration file, use the "Lint Flags" configuration option (above "Lint tool") to add a flag with the value `-config=./revive.toml`, which will select the `revive.toml` file.

### Function descriptions

Comment [must be placed before the function](https://go.dev/doc/comment) declaration, and should be started with function name.

```go
// PrintNumber writes a `number` to the console.
func PrintNumber(number int) {
	fmt.Println(number)
}
```

All documentation in the package can be viewed with `go doc -all` command.

## Multiple modules in the same workspace

[Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)

In the workspace directory, issue following commands:

```bash
go work init ./hello
go work use ./basicFeatures
# ...
```

This will create `go.work` file with following contents:

```
go 1.19

use (
	./basicFeatures
	./hello
)
```

Running modules from workspace directory:

```bash
go run ./hello
go run ./basicFeatures
# ...
```
