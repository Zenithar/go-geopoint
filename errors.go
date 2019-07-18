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

import "errors"

var (
	// ErrInvalidGeoPointHash is raised when the given hash is syntaxically invalid
	ErrInvalidGeoPointHash = errors.New("geopoint: invalid geopoint hash value")
	// ErrInvalidGeoPointValue is raised when the given hash does not contain a valid value
	ErrInvalidGeoPointValue = errors.New("geopoint: invalid geopoint value")
)
