# color

Terminal output with ANSI colors in Go.

+ **Support chain call**
+ **Support nested color**
+ **Support enable/disable color** (default follow terminal support)

![](./preview.png)

```sh
go get github.com/mengdu/color
```

```go
func main() {
  // color.Std = color.New(false) // reset global instance
  fmt.Println(color.Red().String("Hello, world!"))
  fmt.Println(color.Bolb().Italic().Underline().Red().String("Hello, world!"))

  c := color.New(color.EnableColor(true)) // new instance
  fmt.Println(c.Red().String("Red"))
}
```

[Examples](./examples/main.go)
