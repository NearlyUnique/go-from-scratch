# 1: hello_world

## package main
- packages
- console apps (main)
not java/C# namespaces. In go packages are generally smaller

## import
Std lib e.g. “fmt”, “encoding/json” “os”
Imported name last part is the prefix to func or type in package e.g.
`import "fmt"` gives `fmt.Println`
`import "github.com/pkg/errors"` allows `errors.Wrap(err)`

## func main
Not like C/C#/Java
`int main(args string)` // not this
`import "flags"` and or `import "os"` 

## fmt.Println
Uppercase first char for exported entities (types, interfaces, consts)

## go build
cross compile is easy build defaults to directory name, pay attention to `GOPATH`

go build -o output_name
`GOOS=linux, darwin, windows`
`GOARCH=amd64, arm`

```
//
const goosList = "android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows zos "

const goarchList = "386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64 "
```

# 2: loops

_tool_: `go fmt`
_tool_: `goimports`

## Short assignment operator
- declared and not used
`X := rhs`
Or
`var x type`

## time package

Now
Since
Sleep

## loop, one type covers

While, for do

##Printf

C style formatting
%s string %d int %v work it out %T type , plus more

# 3: goroutine

iffy

- go keyword

need to wait for goroutine to complete

- sync.WaitGroup
- struct creation {}
- WaitGroup add
- WaitGroup wait

Deadlock! All goroutines are asleep (one 1 the main one)

- WaitGroup done
- closure bug

Everything happens out-of order, and that’s a good thing - nondeterministic execution of goroutines, a symptom of their concurrency

# 4: Channels

The yin to the yang of goroutines is channels
- create func
- delete WaitGroup
cannot return value, execution has moved on, nothing to collect the value
- make chan

channels are typed

- read and write to chan
comment out write to chan to show deadlock

Synchronised communication across go routines
- show with logging
- const

# 5: types

Create type blog
Create slice blogList
Receivers
Range loop over
!!! Oh wait, incorrect arg (actually it’s type is int as it is index)
So we can use two args
Blank identifier: If we want to ignore index
Function in struct
Add blog to bloglist

# 6: ish

Interfaces
Change the bloglist to a blog collector interface slice
Create a new collector
But implement wrong arguments
Fix args

Read just one
Bring in the random to affect the timeout
Select timeout (pizza collector takes less time)