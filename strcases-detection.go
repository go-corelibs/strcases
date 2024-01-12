// Copyright (c) 2024  The Go-Curses Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package strcases

import (
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
)

func canPreserve(r rune) bool {
	switch r {
	case '-', '_':
		return true
	}
	return unicode.IsLower(r) || unicode.IsUpper(r) || unicode.IsDigit(r)
}

// CanPreserve checks if input is only alphanumeric, dash or underscore
// characters
func CanPreserve(input string) (can bool) {
	for _, r := range input {
		if can = canPreserve(r); !can {
			return
		}
	}
	return
}

// ProfileCase scans through each rune in the given input and checks for a
// number of interesting hints about the string case of the input
//
//	breaker     has things aren't alphanumeric, dash or underscore
//	lower       has at least one lower case character
//	upper       has at least one upper case character
//	kebab       has at least one dash character
//	underscore  has at least one underscore character
func ProfileCase(input string) (breaker, lower, upper, kebab, underscore bool) {
	for _, c := range input {
		if !breaker && !canPreserve(c) {
			breaker = true
		} else if !lower && unicode.IsLower(c) {
			lower = true
		} else if !upper && unicode.IsUpper(c) {
			upper = true
		} else if !kebab && c == '-' {
			kebab = true
		} else if !underscore && c == '_' {
			underscore = true
		}
	}
	return
}

// DetectCase uses ProfileCase and some extra efforts to reliably discern the
// obvious string case of the given input
func DetectCase(input string) (detected Case) {
	breaker, lower, upper, kebab, underscore := ProfileCase(input)

	if breaker {
		detected = UnknownCase
		return // early out
	} else if kebab {
		// check for kebab things
		if input == strcase.ToKebab(input) {
			detected = KebabCase
			return
		} else if input == strcase.ToScreamingKebab(input) {
			detected = ScreamingKebabCase
			return
		}
	} else if underscore {
		// check for snake things
		if input == strcase.ToSnake(input) {
			detected = SnakeCase
			return
		} else if input == strcase.ToScreamingSnake(input) {
			detected = ScreamingSnakeCase
			return
		}
	} else if lower && upper {
		// check for camel things
		if input == strcase.ToCamel(input) {
			detected = CamelCase
			return
		} else if input == strcase.ToLowerCamel(input) {
			detected = LowerCamelCase
			return
		}
	}

	// check for default things
	if input == strings.ToLower(input) {
		detected = LowerCase
	} else if input == strings.ToUpper(input) {
		detected = UpperCase
	}
	return
}
