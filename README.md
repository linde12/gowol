# Introduction

[![GoDoc](https://godoc.org/github.com/linde12/gowol?status.svg)](https://godoc.org/github.com/linde12/gowol)

gowol is a library which provides an easy way to do Wake-on-LAN from your Go application.

The primary goal of `gowol` is to provide an abstraction which enables you to easily send Wake-on-LAN packets. It supports creation and dispatching of Wake-on-LAN packets of the provided MAC to the provided IP(and port).

# Example usage
```go
package main

import "github.com/linde12/gowol"

func main() {
	if packet, err := gowol.NewMagicPacket("03:AA:FF:67:64:05"); err == nil {
		packet.Send("255.255.255.255")          // send to broadcast
		packet.SendPort("255.255.255.255", "7") // specify receiving port
	}
}

```

# Installation

```sh
go get github.com/linde12/gowol
```

# License
MIT
