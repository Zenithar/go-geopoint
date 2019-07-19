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

package geopoint_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"go.zenithar.org/geopoint"
)

func TestValue_Code(t *testing.T) {
	tcl := []struct {
		name         string
		point        geopoint.Value
		expectedCode string
	}{
		{
			name:         "Place du capitole, Toulouse, France",
			point:        geopoint.Value(75071988303315493),
			expectedCode: "10AB5:935B6:6C225",
		},
	}

	for _, tc := range tcl {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.point.Code()
			if got != tc.expectedCode {
				t.Fatalf("invalid code, expected %s, got %s", tc.expectedCode, got)
			}
		})
	}
}

func TestValue_Base58(t *testing.T) {
	tcl := []struct {
		name         string
		point        geopoint.Value
		expectedCode string
	}{
		{
			name:         "Place du capitole, Toulouse, France",
			point:        geopoint.Value(75071988303315493),
			expectedCode: "B7DEqKRCse",
		},
	}

	for _, tc := range tcl {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.point.Base58()
			if got != tc.expectedCode {
				t.Fatalf("invalid code, expected %s, got %s", tc.expectedCode, got)
			}
		})
	}
}
func TestValue_MarshallJSON(t *testing.T) {
	tcl := []struct {
		name         string
		object       interface{}
		expectedJSON []byte
	}{
		{
			name:         "Value only",
			object:       geopoint.Value(75071988303315493),
			expectedJSON: []byte(`"10AB5:935B6:6C225"`),
		},
		{
			name: "Value in a struct",
			object: struct {
				Point geopoint.Value `json:"point"`
			}{
				Point: geopoint.Value(75071988303315493),
			},
			expectedJSON: []byte(`{"point":"10AB5:935B6:6C225"}`),
		},
	}

	for _, tc := range tcl {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.object)
			if err != nil {
				t.Fatalf("unable to marshal json, got error %v", err)
			}
			if bytes.Compare(body, tc.expectedJSON) != 0 {
				t.Fatalf("invalid json serialization, expected %s, got %s", string(tc.expectedJSON), string(body))
			}
		})
	}
}
