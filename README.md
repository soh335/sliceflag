[![wercker status](https://app.wercker.com/status/eef7d91dc4e3bfbbd7e74029a5254f8d/s/master "wercker status")](https://app.wercker.com/project/bykey/eef7d91dc4e3bfbbd7e74029a5254f8d)

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
