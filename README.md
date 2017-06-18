[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/b3ntly/insertjson/master/LICENSE.txt) 
[![Build Status](https://travis-ci.org/b3ntly/insertjson.svg?branch=master)](https://travis-ci.org/b3ntly/insertjson)
[![Go Report Card](https://goreportcard.com/badge/github.com/b3ntly/insertjson)](https://goreportcard.com/report/github.com/b3ntly/insertjson)
[![GoDoc](https://godoc.org/github.com/b3ntly/insertjson?status.svg)](https://godoc.org/github.com/b3ntly/insertjson)
[![Coverage Status](https://coveralls.io/repos/github/b3ntly/insertjson/badge.svg?branch=master)](https://coveralls.io/github/b3ntly/insertjson?branch=master?q=1) 

### InsertJSON

A small package for performantly inserting properties into byte-encoded JSON objects using Golang.

### Usage 

```go 
package main 

import (
    "github.com/b3ntly/insertjson"
)

func main(){
    key := "id"
    value := "1"
    example := []byte(`{"name":"fred", "thing":"foo"}`),
    println(string(insertjson.Property(key, value, example))
}
```

```bash
{"id":"1","name":"fred", "thing":"foo"}
```