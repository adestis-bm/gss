# gss

Golang Secure Strings

## Abstract

A very simple package to wrap strings in a secure manner to prevent cleartext buffers.

## Example

```go
var buffer string
// fill buffer with a string, maybe from stdin
ss,key,err := gss.New(buffer)
if err != nil {
    panic(err)
}
defer ss.Destroy()

// access the seacure string
txt,err := ss.String(key)
if err != nil {
    panic(err)
}

// use txt...
```
