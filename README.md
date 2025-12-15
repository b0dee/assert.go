# assert.go

I'm  sure there are better assertion libraries better fit for your needs... This
one is simple, at just 5 lines of code, no dependencies...

Putting this here because I know I'll want it in other places...

## Installation
go get github.com/b0dee/assert.go

## Usage

To use, I recommend dot-importing as you likely do not have name conflicts for
the 2 functions defined in this library...

```go
import . "github.com/b0dee/assert.go"
```

### Runtime assert

Use this as a mosqueto net for your code base, Go's error handling ideology is
great, though sometimes we want to abort

```go
func foo(a int) { 
    Assert(a == 69 || a == 420, "Pick a better number")
    // code...
}
```
