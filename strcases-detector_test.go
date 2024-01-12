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

func TestDetector(t *testing.T) {
	Convey("CaseDetector", t, func() {
		d := NewCaseDetector()
		cd, _ := d.(*cCaseDetector)
		So(len(cd.cache), ShouldEqual, 0)
		So(d.Detect(`lower`), ShouldEqual, LowerCase)
		So(len(cd.cache), ShouldEqual, 1)
		So(d.Detect(`lower`), ShouldEqual, LowerCase)
		So(len(cd.cache), ShouldEqual, 1)
		So(d.Detect(`CamelCase`), ShouldEqual, CamelCase)
		So(len(cd.cache), ShouldEqual, 2)
		d.Reset()
		So(len(cd.cache), ShouldEqual, 0)
	})
}
