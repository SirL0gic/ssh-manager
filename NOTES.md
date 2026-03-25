# Go Basics — Learning Notes

---

## Project Setup

```bash
go mod init ssh-manager
```

This creates a `go.mod` file — Go's way of defining your project and tracking its dependencies. Think of it like `package.json` in Node.js or `requirements.txt` in Python, except it also handles versioning and module identity in one file.

---

## Project Structure

```
manager/
├── main.go         # All application code lives here
├── go.mod          # Module definition and dependencies
├── go.sum          # Checksums for dependency verification (don't edit this)
└── profile.json    # Your SSH profiles (gitignored — contains credentials)
```

---

## Core Language Concepts

### Packages

Every Go file starts with a package declaration:

```go
package main
```

A **package** is a namespace that groups related code — similar to modules in Python. `package main` is special: it tells Go this file is an executable entry point, not a reusable library. If you renamed it `package utils`, Go wouldn't know how to run it.

### The `main()` Function

```go
func main() {
    // program starts here
}
```

Go always starts execution from a function literally named `main` inside `package main`. You cannot rename it. It's equivalent to Python's `if __name__ == "__main__":` block, except it's enforced by the language.

### Imports

```go
import (
    "fmt"
    "os"
    "golang.org/x/crypto/ssh"
)
```

- Standard library packages (`fmt`, `os`, `log`) ship with Go — no installation needed.
- Third-party packages (`golang.org/x/crypto/ssh`) are fetched via `go get`.
- You use a package by prefixing the function with its name: `fmt.Println()`, `os.ReadFile()`.
- Go will **refuse to compile** if you import a package and don't use it. This is intentional — no dead imports.

### Exported vs Unexported Names

Go uses **capitalization** to control visibility — there are no `public`/`private` keywords:

| Capitalized | Lowercase |
|---|---|
| `Println`, `SSHProfile`, `Name` | `connectSSH`, `choice`, `address` |
| Exported — visible outside the package | Unexported — private to the package |

This is why the fields in `SSHProfile` are capitalized (`Name`, `Host`, `Port`) — they need to be exported so the `encoding/json` package can read and write them during JSON parsing.

---

## Types and Structs

Go doesn't have classes. Instead it uses **structs** — named collections of fields:

```go
type SSHProfile struct {
    Name     string `json:"name"`
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Username string `json:"username"`
    Password string `json:"password"`
    Keypath  string `json:"keypath"`
    IsActive bool   `json:"isActive"`
}
```

- `type SSHProfile struct` — defines a new type named `SSHProfile`
- Each line is a **field**: name on the left, type on the right
- The backtick strings (`` `json:"name"` ``) are **struct tags** — metadata that tells the `json` package what key to use when reading/writing JSON. Without the tag `json:"name"`, Go would look for `"Name"` (capital N) in the JSON instead.
- The trailing comma after the last field is required when fields are on separate lines

### Creating a struct value

```go
profile := SSHProfile{
    Name:     "Production",
    Host:     "192.168.1.100",
    Port:     22,
    IsActive: true,
}
```

Access fields with a dot: `profile.Name`, `profile.Port`.

---

## Functions

```go
func connectionString(profile SSHProfile) string {
    return fmt.Sprintf("ssh %s@%s -p %d", profile.Username, profile.Host, profile.Port)
}
```

- `func` keyword starts the definition
- Parameters go inside `()` with the name first, type second — opposite of most languages
- The return type comes after the parameter list (`string` here)
- Go functions can return multiple values: `func loadProfiles() ([]SSHProfile, error)`

---

## Error Handling

Go doesn't have exceptions. Instead, functions return an error as a second value:

```go
data, err := os.ReadFile("profile.json")
if err != nil {
    log.Fatal("Error reading file:", err)
}
```

- `:=` is short variable declaration — Go infers the type
- `err != nil` checks if something went wrong (`nil` means no error)
- `log.Fatal` prints the error and immediately exits the program

This pattern (`value, err := ...` then `if err != nil`) appears constantly in Go code. You'll write it hundreds of times.

---

## Slices

A slice is Go's dynamic array — like a Python list:

```go
var profiles []SSHProfile   // declares an empty slice of SSHProfile
```

Iterating with `range`:

```go
for i, profile := range profiles {
    fmt.Printf("%d. %s\n", i+1, profile.Name)
}
```

- `range` returns the index and value on each iteration
- Use `_` to discard a value you don't need: `for _, profile := range profiles`

---

## Pointers and References

The `&` operator takes the address of a value (creates a pointer):

```go
config := &ssh.ClientConfig{
    User: profile.Username,
    ...
}
```

The `ssh.Dial` function expects a `*ssh.ClientConfig` (pointer to a ClientConfig). Using `&` here creates the struct and immediately gives you a pointer to it. You'll see this pattern often when working with external libraries.

---

## Dependency Management

```bash
go get golang.org/x/crypto/ssh   # add a dependency
go mod tidy                       # remove unused dependencies
```

No virtual environments needed. Go stores packages globally in `GOPATH` (typically `C:\Users\abisa\go\`) and each project's `go.mod` records exactly which versions it uses. Projects are isolated by their `go.mod`, not by a separate environment.

---

## Building & Running

```bash
# Run without building (good for development)
go run main.go

# Compile to an executable
go build -o ssh-manager.exe

# Run the compiled executable
.\ssh-manager.exe
```

`go run` compiles and runs in one step — useful during development. `go build` produces a standalone binary you can distribute.

---

## `fmt.Sprintf` Format Verbs

Used throughout the code for string formatting:

| Verb | Meaning | Example |
|---|---|---|
| `%s` | string | `"root"` |
| `%d` | integer | `22` |
| `%v` | any value (default format) | `true` |
| `%+v` | struct with field names | `{Name:Prod Host:...}` |

```go
address := fmt.Sprintf("%s:%d", profile.Host, profile.Port)
// → "192.168.1.100:22"
```
