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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDetection(t *testing.T) {
	Convey("CanPreserve", t, func() {
		So(CanPreserve("yes"), ShouldEqual, true)
		So(CanPreserve("yesYes"), ShouldEqual, true)
		So(CanPreserve("nope!"), ShouldEqual, false)
	})

	Convey("ProfileCase", t, func() {
		breaker, lower, upper, kebab, underscore := ProfileCase(`lU -_`)
		So(lower, ShouldEqual, true)
		So(upper, ShouldEqual, true)
		So(breaker, ShouldEqual, true)
		So(kebab, ShouldEqual, true)
		So(underscore, ShouldEqual, true)
		breaker, lower, upper, kebab, underscore = ProfileCase(`!.;'"/`)
		So(lower, ShouldEqual, false)
		So(upper, ShouldEqual, false)
		So(breaker, ShouldEqual, true)
		So(kebab, ShouldEqual, false)
		So(underscore, ShouldEqual, false)
	})

	Convey("DetectCase", t, func() {
		So(DetectCase("has space"), ShouldEqual, UnknownCase)
		So(DetectCase("lower"), ShouldEqual, LowerCase)
		So(DetectCase("UPPER"), ShouldEqual, UpperCase)
		So(DetectCase("CamelCase"), ShouldEqual, CamelCase)
		So(DetectCase("lowerCamelCase"), ShouldEqual, LowerCamelCase)
		So(DetectCase("kebab-case"), ShouldEqual, KebabCase)
		So(DetectCase("SCREAMING-SNAKE-CASE"), ShouldEqual, ScreamingKebabCase)
		So(DetectCase("snake_case"), ShouldEqual, SnakeCase)
		So(DetectCase("SCREAMING_SNAKE_CASE"), ShouldEqual, ScreamingSnakeCase)
	})
}
