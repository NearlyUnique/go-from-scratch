# 1: Hello world.

## package main
- packages
- console apps (main)
not java/C# namespaces

## import
Std lib e.g. “fmt”, “encoding/json” “os”
Imported name last part is the prefix to func or type in package e.g.
fmt.Println

## func main
Not like C/C#/Java
Int Main(args string) // not this

## fmt.Println
Pascal case for exported entities

## go build
go build -o output_name
`GOOS=linux, darwin, windows`
`GOARCH=amd64, arm`

```
const goosList = "android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows zos "

const goarchList = "386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64 "
```

# 2: Closures
## Short assignment operator
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

Sync Waitgroups
Waitgroup add
Waitgroup wait
Ahhh deadlock!
Waitgroup done
closure bug
Everything happens out-of order, and that’s a good thing - nondeterministic execution of goroutines, a symptom of their concurrency

# 4: Channels

First bug read from empty chan blocks
Make(chan string)
Read and write to chan
Chan type safety
Deadlocks
Sycronised communication across go routines
Can’t collect return value from goroutine

# 5: type range

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