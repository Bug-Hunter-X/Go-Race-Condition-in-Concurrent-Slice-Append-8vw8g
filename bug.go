```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var a []int
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        a = append(a, i)
                }(i)
        }
        wg.Wait()
        fmt.Println(len(a))
}
```