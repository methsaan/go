# https://golang.org/doc/code.html
env |grep GO
# GOPATH=/home/test123/programming/go

# where methsaan is your github user name
mkdir -p $GOPATH/src/github.com/methsaan

# To compile and run a simple program, first choose a package path
# (we'll use github.com/user/hello) and create a corresponding
# package directory inside your workspace:
mkdir -p $GOPATH/src/github.com/methsaan/hello

```
cat >$GOPATH/src/github.com/methsaan/hello/hello.go <<EOF
package main

import "fmt"

func main() {
    fmt.Println("Today is the 1st of January, 2019, first program")
}
EOF
```

# Note that you can run this command from anywhere on your system.
# The go tool finds the source code by looking for the github.com/user/hello
# package inside the workspace specified by GOPATH.
go install github.com/methsaan/hello

# This command builds the hello command, producing an executable binary.
# It then installs that binary to the workspace's bin directory as hello
# (or, under Windows, hello.exe). In our example, that will
# be $GOPATH/bin/hello, which is $HOME/go/bin/hello.
cd $GOPATH/src/github.com/methsaan/hello
go install

#which hello
$GOPATH/bin/hello
hello

# If you're using a source control system, now would be a good time to initialize 
# a repository, add the files, and commit your first change. Again,
# this step is optional: you do not need to use source
# control to write Go code.

$ cd $GOPATH/src/github.com/user/hello
$ git init
Initialized empty Git repository in /home/user/work/src/github.com/user/hello/.git/
$ git add hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 1 insertion(+)
  create mode 100644 hello.go
