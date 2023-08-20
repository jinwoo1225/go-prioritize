# go-prioritize
Priority Queue for Go with Generic support

## Installation

```bash
go get github.com/jinwoo1225/go-prioritize
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/jinwoo1225/go-prioritize"
)

func main() {
	// Create a new priority queue
    pq := prioritize.NewPriorityQueue[string](nil)

    // Add some items
    pq.Push(prioritize.NewItem[string]("foo", 1))
    pq.Push(prioritize.NewItem[string]("bar", 2))
    pq.Push(prioritize.NewItem[string]("baz", 3))

    // Pop the highest priority item
    item := pq.Pop()
    fmt.Println(item) // "baz"

    // Pop the next highest priority item
    item = pq.Pop()
    fmt.Println(item) // "bar"

    // Pop the next highest priority item
    item = pq.Pop()
    fmt.Println(item) // "foo"

    // Pop the next highest priority item
    item = pq.Pop()
    fmt.Println(item) // nil (queue is empty)
}

```

