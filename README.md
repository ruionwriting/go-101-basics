# go-101-basics
Introduction to Go, the tooling and basic language features

## Authors
Rob Reid [@codingconcepts](https://github.com/codingconcepts)

Nick Lanng [@nicklanng](https://github.com/nicklanng)

Rui Marques [@ruionwriting](https://github.com/ruionwriting)

## Topics
1. [Install Go](#install-go)
2. What is the Go workspace and why is it different to project workspaces
3. Data types
4. Slices and arrays
5. Pointers vs values
6. Functions
7. [FizzBuzz Exercise](#fizzbuzz-exercise)
8. String Calculator Exercise


<a name="install-go"/>
## 1. Install Go

### Windows

### Linux
Depending on the distro, you use you might already have Go available with your package manager. For example, Ubuntu has a [PPA](https://github.com/golang/go/wiki/Ubuntu) and, in most variants, is already added on your system. The install goes as follows:
```bash
sudo apt-get install golang
```

In the case that you want a bit more control, or if you want to use a different version than the one available from the PPA, just install it directly; visit the [official Go downloads page](https://golang.org/dl/) and find the URL for the current binary release's tarball, along with its SHA256 hash.
```bash
cd ~
curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
```

Verify the tarball using `sha256sum`:
```bash
sha256sum go1.8.linux-amd64.tar.gz
53ab94104ee3923e228a2cb2116e5e462ad3ebaeea06ff04463479d7f12d27ca  go1.8.linux-amd64.tar.gz
```
Extract the files:
```bash
tar xvf go1.8.linux-amd64.tar.gz
```
You should now have a directory called `go` in your home directory. Recursively change `go`'s owner and group to `root`, and move it to `/usr/local`:
```bash
sudo chown -R root:root ./go
sudo mv go /usr/local
```

Create a folder for your Go workspace, we recommend to make a folder at `~/work/go`.

Open your ~/.bashrc or ~/.zshrc (or whatever the config file is for your shell of choice).
You need to set up your environmental variables used for Go tooling and binaries pulled down with ```go get```.

Add the following lines to your shell config, be sure to set the right path if different:
```bash
export GOPATH=$HOME/work/go
export PATH=$PATH:$GOPATH/bin
```

### Mac OS X
The easiest way to install go is to use [Homebrew](http://brew.sh/).
```bash
brew install go
```

Create a folder for your Go workspace, we recommend to make a folder at `~/work/go`.

Open your ~/.bashrc or ~/.zshrc (or whatever the config file is for your shell of choice).
You need to set up your environmental variables used for Go tooling and binaries pulled down with `go get`.

Add the following lines to your shell config, be sure to set the right path if different:
```bash
export GOPATH=$HOME/work/go
export PATH=$PATH:$GOPATH/bin
```

<a name="fizzbuzz-exercise"/>
## 7. FizzBuzz Exercise

In this classic coding kata, you are tasked with printing the numbers from 1 to 100.

However, there are some extra rules.

If the number is a multiple of 3 then, instead of the number, print 'Fizz'.

If the number is a multiple of 5 then, instead of the number, print 'Buzz'.

If the number is a multiple of BOTH then instead of the number, print 'FizzBuzz'.

<details><summary>Click here to see the solution!</summary><p>
```go
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
```
</p></details>

