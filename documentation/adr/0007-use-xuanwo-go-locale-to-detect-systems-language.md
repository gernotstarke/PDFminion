# 7. use Xuanwo/go-locale to detect systems language

Date: 2024-12-15

## Status

Accepted

## Context

It turned out that detecting the system language setting is difficult if one relies only on environment variables.

## Decision

The [Xuanwo/go-locale](https://github.com/Xuanwo/go-locale) is a cross-platform library for reliably detecting the system language setting.
It works on MacOS, Linux, and Windows (and some others).

It is easy to use and provides a simple API to get the language setting.

## Consequences

Here you find a sample program showing how to use the library:

```go
package main

import (
"fmt"
"log"

"github.com/Xuanwo/go-locale"
)

func getSystemLanguage() (string, error) {
tag, err := locale.Detect()
if err != nil {
return "", err
}
return tag.String(), nil
}

func main() {
lang, err := getSystemLanguage()
if err != nil {
log.Fatalf("Failed to detect system language: %v", err)
}

fmt.Printf("Detected system language: %s\n", lang)
}
```