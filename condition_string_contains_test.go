/*
 * Copyright © 2016-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package ladon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringContains(t *testing.T) {
	for _, c := range []struct {
		matches   string
		value     string
		delimiter string
		pass      bool
	}{
		{matches: "foobar", value: "foo", pass: true},
		{matches: "foobar", value: "bar", pass: true},
		{matches: "foo:bar", value: "bar", delimiter: ":", pass: true},
		{matches: "foo:bar:baz", value: "foo:bar", delimiter: ":", pass: true},
		{matches: "foobar", value: "food", pass: false},
	} {
		condition := &StringContainsCondition{
			Contains: c.matches,
		}

		assert.Equal(t, c.pass, condition.Fulfills(c.value, new(Request)), "%s", c.matches)
	}
}
