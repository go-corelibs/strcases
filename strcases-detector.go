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
	"sync"
)

// CaseDetector is a caching system in order to optimize usage of the
// expensive DetectCase function
type CaseDetector interface {
	// Detect wraps DetectCase and caches the results
	Detect(input string) (c Case)
	// Reset will clear the existing cache of cases detected
	Reset()
}

type cCaseDetector struct {
	cache map[string]Case
	sync.Mutex
}

// NewCaseDetector constructs a new CaseDetector instance
func NewCaseDetector() (d CaseDetector) {
	d = &cCaseDetector{
		cache: make(map[string]Case),
	}
	return
}

func (d *cCaseDetector) Detect(input string) (c Case) {
	d.Lock()
	defer d.Unlock()
	var ok bool
	if c, ok = d.cache[input]; ok {
		return
	}
	c = DetectCase(input)
	d.cache[input] = c
	return
}

func (d *cCaseDetector) Reset() {
	d.Lock()
	defer d.Unlock()
	d.cache = make(map[string]Case)
}
