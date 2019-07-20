// +build gofuzz

/*
 * Copyright 2019 Thibault NORMAND
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
 */

package geopoint

// Fuzz usage:
//   go get github.com/dvyukov/go-fuzz/...
//
//   go-fuzz-build go.zenithar.org/geopoint && go-fuzz -bin=./geopoint-fuzz.zip -workdir=/tmp/geopint-fuzz
func Fuzz(data []byte) int {
	code := string(data)
	if err := Check(code); err != nil {
		return 0
	}
	lat, lon, err := FromString(code)
	if err != nil {
		return 2
	}
	if _, _, err := Decode(Encode(lat, lon)); err != nil {
		return 2
	}
	return 1
}
