# sliceflag

slice flag for ```flag``` package of go

## USAGE

```go
var (
       strs = sliceflag.String(flag.CommandLine, "str", []string{}, "str")
)

func main() {
    flag.Parse()
    log.Println(*strs)
}
```

## SEEALSO

* [flag - The Go Programming Language](https://golang.org/pkg/flag/)

## LICENSE

* MIT
