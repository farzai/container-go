## Container Service (Go)
[![Go Reference](https://pkg.go.dev/badge/github.com/farzai/container-go.svg)](https://pkg.go.dev/github.com/farzai/container-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/farzai/container-go)](https://goreportcard.com/report/github.com/farzai/container-go)
![Github Actions](https://github.com/farzai/container-go/actions/workflows/main.yml/badge.svg?branch=main)

A simple dependency injection library for Go.

### Installation

```bash
go get -u github.com/farzai/container-go
```

### Usage

```go
import (
    "errors"

    "github.com/farzai/container-go"
)

func main() {
    // Create a new container service instance
    service := container.New()

    // Bind a new dependency
    service.Bind("foo", func(c *container.ContainerService) (interface{}, error) {
        return "bar", nil
    })

    // Resolve a dependency
    foo, err := service.Resolve("foo")
    if err != nil {
        // Handle error
    }

    // Use the resolved dependency
    fmt.Println(foo) // Output: "bar"

    // Use singleton instead of binding to make sure the dependency is only instantiated once
    service.Singleton("baz", func(c *container.ContainerService) (interface{}, error) {
        return "qux", nil
    })

    // Resolving a singleton dependency will always return the same instance
    baz1, err := service.Resolve("baz")
    if err != nil {
        // Handle error
    }

    baz2, err := service.Resolve("baz")
    if err != nil {
        // Handle error
    }

    fmt.Println(baz1 == baz2) // Output: true

    // Unbind a dependency
    service.Unbind("foo")

    _, err = service.Resolve("foo")
    if !errors.Is(err, container.ErrNoBinding) {
        // Handle error
    }
}
```


### License

Container Service (Go) is open-sourced software licensed under the [MIT license](LICENSE.md).