# go-µweb

[![Build Status](https://travis-ci.org/ericmdantas/go-uweb.svg?branch=master)](https://travis-ci.org/ericmdantas/go-uweb)

Micro wrapper on go's HTTP handling. Created for testing purposes.

### Usage

```go
package main

import (
    "net/http"
    uweb "github.com/ericmdantas/go-uweb"
)

func main() {
    u := uweb.New()

    u.GET("/hello_world", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello world!"))
    })

    http.ListenAndServe(":1037", u)
}
```

### Should I use it?

No.