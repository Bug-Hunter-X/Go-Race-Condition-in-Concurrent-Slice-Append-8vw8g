# Go Race Condition in Concurrent Slice Appending

This repository demonstrates a common race condition in Go that can occur when multiple goroutines concurrently modify a shared slice.  The code appends numbers to a slice within a loop of goroutines. Without proper synchronization, this leads to unpredictable results. The solution demonstrates how to resolve the issue using a mutex.