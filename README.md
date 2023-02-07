# round-robin-go

round-robin-go is a round robin balancing algorithm written in golang. 

The project uses go-generics, that allows you to use round-robin for any data type. With help of go generics, it has become easier for you to plugin `roundrobin` to your project.

## Installation
```shell
go get github.com/thegeekyasian/round-robin-go
```

## Version:
> 1.18

Since Generics were introduced in go with version 1.18, the project requires go 1.18 to work.

## Usage

### any type
```go
type resource struct {
    id   int
    name string
}

...

rr, _ := roundrobin.New[resource](
    &resource{1, "resource-1"},
    &resource{2, "resource-2"},
    &resource{3, "resource-3"},
)

rr.Next() // resource-1
rr.Next() // resource-2
rr.Next() // resource-3
rr.Next() // resource-1
```

### string
```go
one := "One"
two := "Two"
three := "Three"

rr, _ := roundrobin.New[string](&one, &two, &three)

rr.Next()	// One
rr.Next()	// Two
rr.Next()	// Three
rr.Next()	// One
```

### URLs
```go
rr, _ := roundrobin.New[url.URL](
    url.URL{Host: "192.168.0.1"},
    url.URL{Host: "192.168.0.2"},
    url.URL{Host: "192.168.0.3"},
    url.URL{Host: "192.168.0.4"},
)

rr.Next() // 192.168.0.1
rr.Next() // 192.168.0.2
rr.Next() // 192.168.0.3
rr.Next() // 192.168.0.4
rr.Next() // 192.168.0.1
rr.Next() // 192.168.0.2
```

### Author
* The Geeky Asian
  * [Github](https://github.com/thegeekyasian/)
  * [Website](https://thegeekyasian.com/)

### License
round-robin-go is released under [MIT license](https://github.com/thegeekyasian/round-robin-go/blob/master/LICENSE) (2023).

