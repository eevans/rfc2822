RFC2822 Parser For Go
---------------------
This a simple [rfc2822](http://www.ietf.org/rfc/rfc2822.txt) parser for 
[Go](http://golang.org).  It does not (currently )support the full standard.


### Installation
`goinstall github.com/eevans/rfc2822`

### Usage
```go
package main

import (
    "github.com/eevans/rfc2822"
    "os"
    "fmt"
)

func main() {
    var (
        msg   *rfc2822.Message
        err   os.Error
        value string
    )

    if msg, err = rfc2822.ReadFile("message.txt"); err != nil {
        fmt.Printf("error reading file: %s\n", err)
        os.Exit(2)
    }

    if value, err = msg.GetHeader("subject"); err != nil {
        fmt.Println(err)
        os.Exit(3)
    }

    fmt.Printf("Subject: %s\n", value)
}
```
