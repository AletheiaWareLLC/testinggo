/*
 * Copyright 2019 Aletheia Ware LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package testinggo

import (
	"io/ioutil"
	"os"
	"testing"
)

func MakeTempDir(t *testing.T, name string) string {
	t.Helper()
	dir, err := ioutil.TempDir("", name)
	if err != nil {
		t.Fatalf("Could not create temp dir: '%s'", err)
	}
	return dir
}

func UnmakeTempDir(t *testing.T, dir string) {
	t.Helper()
	os.RemoveAll(dir)
}
