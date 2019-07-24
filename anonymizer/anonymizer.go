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

package anonymizer

import "go.zenithar.org/geopoint"

// Anonymizer defines anonymization strategy contract
type Anonymizer interface {
	Anonymize(point geopoint.Value) geopoint.Value
}

// DeAnonymizer defines de-anonymization strategy contract
type DeAnonymizer interface {
	DeAnonymize(point geopoint.Value) geopoint.Value
}

// Strategy defines complete poin anonymization / de-anonymization contract
type Strategy interface {
	Anonymizer
	DeAnonymizer
}
