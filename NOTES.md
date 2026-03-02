go mod init ssh-manager - This creates a go.mod file in your manager directory.  Go manages your project and its dependencies. Think of it like npm init in Node.js or pip init in Python

package main — in Go, every file belongs to a package. A package is just a way to group related code (similar to modules in Python). package main is special — it tells Go "this is an executable program, not a library." If you named it package utils or anything else, Go wouldn't know where to start running your program.

import "fmt" — yes, very similar to Python's from datetime import time. fmt is a package from Go's standard library that handles formatted input/output (printing, string formatting, reading input). The name stands for "format." You're importing the whole package, then using its functions like fmt.Println(). So fmt is the package and Println is a function inside that package.

Why main.go? — honestly, it's just convention. You could name it app.go or banana.go and it would still work. What matters is that the file contains package main and a main() function. But main.go is standard practice so other developers know where to look for the entry point.

Does the function have to be named main? — yes, absolutely. Go looks for a function called main inside package main as the starting point of your program. It's like how Python uses if __name__ == "__main__". You can't rename it — if you do, Go won't know what to run.

In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package. Like time.Now()

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package. 

Go handles this differently from Python — no virtual environments needed! When you install a package, it goes to your GOPATH directory (usually C:\Users\abisa\go). Each project's go.mod file tracks its own dependencies, so projects are already isolated. Think of go.mod as your requirements.txt and virtual environment combined into one.

Let's learn structs — this is how Go handles objects. Go doesn't have classes like Python, it uses structs instead.

Define a struct - like a Python class but just data
// Create a profile - like instantiating a class

type SSHProfile struct defines the structure — it's your blueprint. The field names are capitalized (Name, Host) — this is important in Go. Capitalized means "exported" (public), lowercase means "unexported" (private). That trailing comma after the last field (IsActive: true,) is required in Go when you write it on multiple lines.