

### command 

- Run go-file
```shell
$ go run main.go
# hello world
```

- Run with standard input

in the code  
```go
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Printf("sum is %d", a+b)
}
```

then  

```shell
$ go run sum.go

# listen stdin, input a b
# both space separated and line feed-separated are acceptable  
$ 10 20 
# sum is 30
```

- install package

```
go install <package>@<version>
// ex 
go install golang.org/x/pkgsite/cmd/pkgsite@latest

```
if you look for packages, check [here](https://pkg.go.dev/)


- show docs

run `go doc`
```
go doc fmt.Printf
```

### public/private

Capitalized func,struct are exported.  
Non-capital these can't be referred external file.  
When you make util function, should define with capitalize.  
