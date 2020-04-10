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

go install github.com/methsaan/hello

#cd $GOPATH/src/github.com/methsaan/hello
#go install

#which hello
$GOPATH/bin/hello
hello
