# Go Pro

Here I learn the Go language.

### Two main questions I want to answer

- What it is good for?
- When it is better than Python?

## Compiler commands

```bash
# Compile and create "hello" executable
go build ./hello.go

# Compile and run executable in temporary directory
go run ./hello.go
```

# Debugger

Install debugger from GitHub:

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

Start debug session:

```bash
~/go/bin/dlv debug hello.go
```
