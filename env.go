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
	"os"
	"testing"
)

func MakeEnv(t *testing.T, env, value string) {
	t.Helper()
	os.Setenv(env, value)
}

func UnmakeEnv(t *testing.T, env string) {
	t.Helper()
	os.Unsetenv(env)
}

func MakeEnvTempDir(t *testing.T, env, dir string) string {
	t.Helper()
	d := MakeTempDir(t, dir)
	MakeEnv(t, env, d)
	return d
}

func UnmakeEnvTempDir(t *testing.T, env, dir string) {
	t.Helper()
	UnmakeEnv(t, env)
	os.RemoveAll(dir)
}
