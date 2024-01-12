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

	"github.com/iancoleman/strcase"
)

// Case is a simple type for indicating a detected string case
type Case uint8

const (
	UnknownCase Case = iota
	LowerCase
	UpperCase
	CamelCase
	LowerCamelCase
	KebabCase
	ScreamingKebabCase
	SnakeCase
	ScreamingSnakeCase
)

// String returns the name of the string case, in that string cases' case
func (c Case) String() (name string) {
	switch c {
	case LowerCase:
		return "lower"
	case UpperCase:
		return "UPPER"
	case CamelCase:
		return "CamelCase"
	case LowerCamelCase:
		return "lowerCamelCase"
	case KebabCase:
		return "kebab-case"
	case ScreamingKebabCase:
		return "SCREAMING-KEBAB-CASE"
	case SnakeCase:
		return "snake_case"
	case ScreamingSnakeCase:
		return "SCREAMING_SNAKE_CASE"
	default:
		name = ""
	}
	return
}

// Apply returns the input text with the detected Case applied. For example:
//
//	CamelCase.Apply("kebab-thing") == "KebabThing"
//	KebabCase.Apply("CamelCase") == "camel-case"
func (c Case) Apply(input string) (modified string) {
	switch c {
	case LowerCase:
		modified = strings.ToLower(input)
	case UpperCase:
		modified = strings.ToUpper(input)
	case CamelCase:
		modified = strcase.ToCamel(input)
	case LowerCamelCase:
		modified = strcase.ToLowerCamel(input)
	case KebabCase:
		modified = strcase.ToKebab(input)
	case ScreamingKebabCase:
		modified = strcase.ToScreamingKebab(input)
	case SnakeCase:
		modified = strcase.ToSnake(input)
	case ScreamingSnakeCase:
		modified = strcase.ToScreamingSnake(input)
	case UnknownCase:
		fallthrough
	default:
		modified = input
	}
	return
}
