[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/strcases)
[![codecov](https://codecov.io/gh/go-corelibs/strcases/graph/badge.svg?token=zl7oAXem0u)](https://codecov.io/gh/go-corelibs/strcases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/strcases)](https://goreportcard.com/report/github.com/go-corelibs/strcases)

# strcases - string case utilities

This package provides a rudimentary detector and other string case utilities.

# Installation

``` shell
> go get github.com/go-corelibs/strcases@latest
```

# Examples

## CaseDetector

``` go
func main() {
    // create a new CaseDetector instance
    cd := strcases.NewCaseDetector()
    c := cd.Detect(`ThisIsCamelCase`)
    // c == strcases.CamelCase
    modified := c.Apply(`kebab-case`)
    // modified == "KebabCase"
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
