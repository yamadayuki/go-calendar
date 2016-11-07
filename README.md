# Calendar

Provides calendar utilities for golang.

**Usage**

```go
// importing
package main

import (
    "fmt"

    "github.com/yamadayuki/go-calendar"
)

func main() {
    calendar := calendar.CurrentMonth()
    fmt.Println(calendar) // Print current month calendar.
}
```
