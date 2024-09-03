<h1 align="center">Cache</h1>

<p align="center">
   <img alt="Static Badge" src="https://img.shields.io/badge/Version-v1.0-blue?style=flat&color=blueviolet">
   <img alt="Static Badge" src="https://img.shields.io/badge/License-MIT-green?style=flat">
</p>

## About

Library for working with cache

## Documentation
``` go
...
  Set(key string, value any, ttl time.Duration)
  Get(key string) (any, error)
  Delete(key string) error
...
```

## Example usage:

```golang
package main

import (
	"fmt"
	memorycache "github.com/pelicanch1k/in-memory-cache/pkg/cache"
)

func main(){
	cache := memorycache.New()
	cache.Set("user", 19)
	
	name, _ := cache.Get("user")
	fmt.Println(name)

	cache.Delete("user")

	name, _ = cache.Get("user")
	fmt.Println(name)
}
```
