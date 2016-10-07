Calendar
---
Provides calendar utilities for golang.

**Usage**
```go
// importing
import (
  "fmt"

  "github.com/yamadayuki/calendar"
)

calendar, err := calendar.Weeks()
if err == nil {
  fmt.Println(calendar)
}

// Feb, 2016
// [
//   [0, 1, 2, 3, 4, 5, 6],
//   [7, 8, 9, 10, 11, 12, 13],
//   [14, 15, 16, 17, 18, 19, 20],
//   [21, 22, 23, 24, 25, 26, 27],
//   [28, 29, 0, 0, 0, 0, 0]
// ]
```
