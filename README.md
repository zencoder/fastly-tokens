# fastly-tokens
Generate tokens for securing content distributed with the Fastly CDN

## Usage
Install the ```fastly-tokens``` library:
```go get github.com/zencoder/fastly-tokens```

```go
package main

import (
  "fmt"
  "log"

  "github.com/zencoder/fast-tokens"
)

func main() {
	// Generate token that is valid for 60 seconds using the default secret
	var token string
	var err error
	if token, err = GenerateToken("RmFzdGx5IFRva2VuIFRlc3Q=", 60); err != nil {
		log.Fatal("Error while generating token", err)
	}
	fmt.Printf("Token: %s\n", token)
}
```

## Benchmarks
```shell
PASS
BenchmarkGenerateToken	  200000	      8687 ns/op
ok  	github.com/zencoder/fastly-tokens	2.359s
```
