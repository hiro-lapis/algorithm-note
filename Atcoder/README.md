

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


