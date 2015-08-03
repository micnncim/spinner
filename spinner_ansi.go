// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !windows

package spinner

import "fmt"
import "unicode/utf8"

// erase deletes written characters
func (s *Spinner) erase(a string) {
	n := utf8.RuneCountInString(fmt.Sprintf("%s%s%s ", s.Prefix, s.color(a), s.Suffix))
	for i := 0; i < n; i++ {
		fmt.Fprintf(s.w, "\b")
	}
}