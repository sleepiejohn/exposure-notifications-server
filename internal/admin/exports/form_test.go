// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, softwar
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exports

import (
	"reflect"
	"testing"
)

func TestSplitRegions(t *testing.T) {
	tests := []struct {
		r string
		e []string
	}{
		{"", []string{}},
		{"  test  ", []string{"test"}},
		{"test\n\rfoo", []string{"foo", "test"}},
		{"test\n\rfoo bar\n\r", []string{"foo bar", "test"}},
		{"test\n\rfoo bar\n\r  ", []string{"foo bar", "test"}},
		{"test\nfoo\n", []string{"foo", "test"}},
	}

	for i, test := range tests {
		if res := splitRegions(test.r); !reflect.DeepEqual(res, test.e) {
			t.Errorf("[%d] splitRegions(%v) = %v, expected %v", i, test.r, res, test.e)
		}
	}
}
