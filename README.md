# GO Stream API  based on golang generic

## Usage

### Install

``go get -u github.com/aak1247/go-stream``

### Basic Usage


```go
// define a stream of int
s := Of([]int{0, 0, 1, 2, 3})
// map to float32
s2 := Map(s, func(a int) float32 {
    return float32(a)
})
// foreach
res := s2.
// filter if not zero
Filter(func(i float32) bool {
    return i > 0
}).
// for each
Foreach(func(i float32){
    fmt.Println(i) // print 1 2 3 4
}).
Reduce(0, func(a,b float32) float32 {
    return a+b
})

fmt.Println(res) // print 10
```

## Future Work

+ [ ] Add support to parallel
+ [ ] Lazy evaluation
+ [ ] other useful methods (according to user needs)