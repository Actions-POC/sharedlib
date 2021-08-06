# shared-library
A library project to import into our main "microservice" project.

## Installation 

```sh
$ go get github.com/Actions-POC/sharedlib
```

## Example

```go
package main

import (
	"fmt"
	"github.com/Actions-POC/sharedlib"
)
	
func main() {
	fmt.Println(sharedlib.Hi("github"))
	fmt.Printf("Nice to meet you, %s\n", sharedlib.Name())
}
```