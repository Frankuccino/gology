# labs/

This directory is the "Research" workspace. It contains isolated experiments, proof-of-concept scripts, and solutions to chapter exercises from *Mastering Go*. 

Each subdirectory corresponds to a chapter or a specific area of Go internals (e.g., `ch02/`, `ch02/`, `ch10_concurrency/`).

## Installation of godoc utility
Run these in the terminal to install the godoc utility.
` go install golang.org/x/tools/cmd/godoc@latest`
This will allow us to run a Go Documentation Server.

### Running the godoc after installation
`~/go/bin/godoc -http :8001`
This will enable the documentation and you can go to your browser's search bar and type `http://localhost:8001/pkg/` to see the documentation server.

This documentation would serve to learn the gory details of the functions and variables you want to use includingthe available packages in the Go standard library.

## GLOBAL RULE IN GO

Rule that applies to function and variable names andis valid for all packages except main: ***everything that begins with a lower-case letter is considered private and is accessible in the current package only.*** Otherwise, you get it, capitalized first letter is considered public and is accessible outside the package.

The only exception to this rule is package names, which can begin with either lowercase or uppercase letters. Having said that, there's not even a Go package that begins with an uppercase letter!

## Introducing packages

Go programs are organized in packages, even the smallest Go program should be delivered as a package.

A very important exception to remember: if you are creating an executable application and not just a package thatwould be shared by other applications or packages, you should name your package `main`.

The `package` keyword helps you define the name of the new package, which can be anything you want.

e.g. `package database` or `package config`, you get it. The word after *package* is the name of the package.

### The `import` keyword

The import keyword is used for importing other Go packages into your Go programs to use some or all of their functionality.

The Go package can either be a part of the rich Standard Library of Go or come from an external source.

Importing from an external package requires to be defined using their full URL.

e.g. 
```go
import "os" // Standard library package import
import "fmt"

import "github.com/spf13/cobra" // External package import

```

## Running Go code
We can run the Go code in 2 ways: through a compiled go that we can execute or through running it like a scripting language.

```zsh
$ go build <myapp.go>
$ ./<myapp>


$ go run <myapp>

```
What's inside the "< >" is your Go file name you compiled or run 


## Important coding rules

- You either use a variable or you do not declare it at all. This rule helps you avoid errors such as misspelling an existing variable or function name.

- Go is delivered in packages. A Go rule that says that if you import a packacge, you should use it in some way(call a function or use a datatype), or the compiler is going to complain. There's an exception to this, and it's with the initialized connections: the use of init() function from a package.

- There is only one way to format curly braces in Go.

- Coding blocks in Go are embedded in curly braces, evenif they contain just a single statement or no statements at all.

- Go functions can return multiple values. 

- You cannot automatically convert between different data types, even if they are of the same kind. As an example, you cannot implicityl convert an integer to a floating point. This also applies to Concrete types and Defined Types.

Go has more rules, but the preceeding ones are the most important. So you can always put them in mind.

```
func main()
{
    fmt.Println("Go has strict rules for curly braces!")
}
// This will throw an error because Go compiler will insert a semicolon at the end of the previous line
```

## Whay you sould know about Go

This discusses important and essential Go features including **variables**, **controlling program flow**, **iterations**, **getting user input**, and Go **concurrency**.

### Defining and using variables
One very important rule: ***if no initial value is given to a variable, the Go compiler will automaticall initialize that variabl to the zero value of its data type.***

There is also the `:=` notation, which can be used instead of the `var` declaration.


