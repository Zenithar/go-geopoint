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

import (
	"fmt"
)

// Value is a type wrapper to define a GPS point
type Value uint64

var (
	// Zero value for comparison
	Zero = Value(0)
)

// Code returns the point encoded as hexadecimal string
func (p Value) Code() string {
	value := uint64(p)
	return fmt.Sprintf("%05X:%05X:%05X", (value >> 40), (value>>20)&0xFFFFF, (value)&0xFFFFF)
}

// -----------------------------------------------------------------------------

// MarshalJSON is used to override JSON marshalling strategy of uint64
func (p Value) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", p.Code())), nil
}
