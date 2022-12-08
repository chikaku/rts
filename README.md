# rts - Runtime Shadow

⚠️ **Only tested in Go1.19**

Access runtime global variable through `go:linkname` you can do such things like

### Get the current timer count in the process

```go
package main

import "github.com/chikaku/rts"

func main() {
    n := rts.NumTimers()
}
```
