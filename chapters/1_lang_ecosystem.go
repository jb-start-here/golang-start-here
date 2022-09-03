// Basic Hello world in Go

package main // All go code must be defined in a package.
// main function in the Main package is the point of entry
// of all go code. Typically all go code in a folder will probably be under the
// same package. Its basically used for namespacing code.

// All imports are done by import keyword. "fmt" is the formatting package
// from stdlib that houses print to stdout and formatting functions

// Go is not an OOP language - although object like behavior can be modeled via some language features like
// structs and also interfaces
import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
// go build (compiles and places executable in bin directory, which can then be called) 
// or go run (compiles, runs and deletes executable) can be used to run this code.

// How go enviroment and ecosystem works?
// Go - compiled language. Can be compiled against many target archs and os
// When go is installed it comes with few basic binaries - go, gofmt and etc...
// go is the language executable. gofmt - is the opinionated std formatting tool.

// Go installation also depends on 2 env variables by default
// 1. GOROOT - this env var should always point to the go installation that contains actual go source code along with
// stdlib and compiled binaries that this spits out i.e `go` and `gofmt` and other stuff. go binaries are added to the
// $GOROOT/bin directory. So you might want to add $GOROOT/bin to your $PATH if you want to invoke go in your shell.
// Fancy version managers like asdf-vm with golang plugin automatically does this so no need to add stuff to $PATH manually
// 2. GOPATH - this env points to the location where downloaded go modules and binaries live. by default GOPATH/pkg is where
// downloaded code lives and GOPATH/bin is where go binaries live.

// There is also GOBIN with we can override where go binaries should live.

// go also has its own package manager. https://pkg.go.dev/ This is the official
// package manager. go binary already comes with official support to fetch packages from this packages
// from this repository. Go packages are called go modules like how ruby packages are called ruby gems.
// If you compile a go module you may or may not get an executable depending on where there exists a main function (point of entry)

// Since go is a compiled language there are two types of package fetching
// 1. Just fetch the binaries (download the already compiled binaries and begin using it as executables)
// 2. Fetch the actual code (uncompiled) library to include in your code and use the constructs provided by the library in your code.
// Eventually when your module or app that imports the library is compiled - it also compiles the dependencies and uses a linker to link
// everything together. This will remind you of c or c++ thats okay because one of the main go authors is ken thompson :)

// Therefore, we have two ways of getting packages from the pkg.go.dev offical repo

// 1. 'go install path_to_pkg_url@version' this downloads the code from the pkg.go.dev and 
// follows build instructions that comes with the package and places downloaded code in $GOPATH/pkg and any
// resulting binaries in $GOPATH/bin. If you happen to write your own go module that has an accompanying executable then typing
// 'go install' in the root of your project will also compile your current project and install it in the $GOPATH/bin dir so you can use it directly without
// qualifying it with the full path. In fact i believe this is the main purpose of go install. It basically has an added feature where if you pass a vcs url it will also
// fetch the package via the git by default of whatever vcs or git binary you point to in $GOVCS env var.

// 2. `go get -d path_to_pkg_url@version` this will download the code only. It will not build and install it in our $GOPATH. If you skip the -d command then 
// it will also build & install. In the future released at the time of writing go get will have -d option enabled by default.
// So i recommend always using -d for now. go get will actually download the uncompiled go code of a module to the $GOPATH/src directory.
// This is different from $GOPATH/pkg directory. This is called module aware mode. As in go get is used to download modules as uncompiled go code to import in our own code
// whereas, go install is mainly used to download modules that are meant to be built immediately and used as executables and CLI tools. So next time we go install a module
// we have already go installed go quickly checks the $GOPATH/pkg directory and simply rebuilds it instead of downloading it (given that the code didnt change ofcourse, it uses a checksum for this)
// Another important thing is that go get -d can only be run inside another go project. Basically it needs a 'go.mod' in the root of the porject dir for go to become module aware.
// This is like go's version of package.json if you will. 


// The way go stores the module's uncompiled code whether you go install or go get in both $GOPATH/pkg and $GOPATH/src is via an "import path structure"
// Suppose you do `go install github.com/cbergoon/speedtest-go@latest` then you if you go to $GOPATH/pkg/mod directory then you can see folder structure as
// pkg/mod/github.com/cbergoon/speedtest-go@latest/... This makes it easy to traverse to the right directory.
