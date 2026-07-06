# labs/

This directory is the "Research" workspace. It contains isolated experiments, proof-of-concept scripts, and solutions to chapter exercises from *Mastering Go*. 

Each subdirectory corresponds to a chapter or a specific area of Go internals (e.g., `ch02/`, `ch02/`, `ch10_concurrency/`).

## Installation of godoc utility
Run these in the terminal to install the godoc utility.
` go install golang.org/x/tools/cmd/godoc@latest`
This will allow us to run a Go Documentation Server.

### Running the godoc after installation
`~/go/bin/godoc -http :8001`
This will enable the documentation and you can go to your browser's `http://localhost:8001/pkg/`

This documentation would serve to learn the gory details of the functions and variables you want to use includingthe available packages in the Go standard library.

## GLOBAL RULE IN GO

Rule that applies to function and variable names andis valid for all packages except main: ***everything that begins with a lower-case letter is considered private and is accessible in the current package only.***.

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
```
import "os" // Standard library package import
import "fmt"

import "github.com/spf13/cobra" // External package import

```



