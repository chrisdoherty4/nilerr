# `nilerr` :facepalm:

Nil pointers are an unfortunate reality in Go. Defining a nil error in projects repeatedly is cumbersome. This package offers a standard nil error type with some utility functions. - nothing more, nothing less.

```go
func Foo(b *Bar) error {
    if b == nil {
        return nilerr.New("b")
    }
}

func main() {
    err := Foo(nil)

    // True
    if nilerr.Is(err) { ... }

    // True
    if errors.Is(err, nilerr.E) { ... }

    // Prints "b is nil"
    fmt.Println(err)
}
```