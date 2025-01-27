```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var a []int
        var mu sync.Mutex
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mu.Lock()
                        a = append(a, i)
                        mu.Unlock()
                }(i)
        }
        wg.Wait()
        fmt.Println(len(a))
}
```