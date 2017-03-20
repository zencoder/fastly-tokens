# fastly-tokens

Generate tokens for securing content distributed with the Fastly CDN, in go!

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
	token := ft.GenerateToken("RmFzdGx5IFRva2VuIFRlc3Q=", 60*time.Second, base64.StdEncoding)
	fmt.Printf("Token: %s\n", token)
}
```

## Benchmarks
```shell
BenchmarkGenerateToken-8   	 1000000	      1771 ns/op	     848 B/op	      12 allocs/op
PASS
ok  	github.com/zencoder/fastly-tokens/ft	2.101s
```

## Development

### Dependencies

Tested on go 1.8.

### Build and run unit tests

    make test

### CI

[This library builds on Circle CI, here.](https://circleci.com/gh/zencoder/fastly-tokens/)

## License

[Apache License Version 2.0](LICENSE)
