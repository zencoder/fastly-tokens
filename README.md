# fastly-tokens
Generate tokens for securing content distributed with the Fastly CDN

## Usage
Install the ```fastly-tokens``` library:
```go get github.com/zencoder/fastly-tokens/ft```

```go
package main

import (
  "encoding/base64"
  "fmt"
  "log"
  "time"

  "github.com/zencoder/fastly-tokens/ft"
)

func main() {
	// Generate token that is valid for 60 seconds using the default secret
	var token string
	var err error
	if token, err = ft.GenerateToken("RmFzdGx5IFRva2VuIFRlc3Q=", 60*time.Second, base64.StdEncoding); err != nil {
		log.Fatal("Error while generating token", err)
	}
	fmt.Printf("Token: %s\n", token)
}
```

## Benchmarks
```shell
PASS
BenchmarkGenerateToken    200000        6523 ns/op
ok    github.com/zencoder/fastly-tokens/ft  1.551s
```
