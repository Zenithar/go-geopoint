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
	"testing"

	"go.zenithar.org/geopoint"
)

func TestEncoder_EncodeDecode(t *testing.T) {
	lat, lon := 43.603574, 1.442917

	p := geopoint.Encode(lat, lon)
	pLat, pLon, err := geopoint.Decode(p)
	if err != nil {
		t.Fatalf("error should not be raise, %v", err)
	}
	if lat != pLat {
		t.Fatalf("invalid latitude, expected: %v, got %v", lat, pLat)
	}
	if lon != pLon {
		t.Fatalf("invalid longitude, expected: %v, got %v", lon, pLon)
	}
}

func TestEncoder_Encode(t *testing.T) {

	tcl := []struct {
		name          string
		lat           float64
		lon           float64
		expectedPoint geopoint.Value
		expectedCode  string
	}{
		{
			name:          "Place du capitole, Toulouse, France",
			lat:           43.603574,
			lon:           1.442917,
			expectedPoint: geopoint.Value(75071988303315493),
			expectedCode:  "10AB5:935B6:6C225",
		},
		{
			name:          "Mairie de Toulouse, Toulouse, France",
			lat:           43.604297,
			lon:           1.443677,
			expectedPoint: geopoint.Value(75071989061436701),
			expectedCode:  "10AB5:93889:6C51D",
		},
		{
			name:          "Tour Eiffel, Paris, France",
			lat:           48.858373,
			lon:           2.292292,
			expectedPoint: geopoint.Value(77888104758015428),
			expectedCode:  "114B6:D1905:475C4",
		},
		{
			name:          "Montréal, Quebec, Canada",
			lat:           45.558196,
			lon:           -73.870384,
			expectedPoint: geopoint.Value(76116476767848432),
			expectedCode:  "10E6B:88474:D47F0",
		},
		{
			name:          "Buenos Aires, Argentina",
			lat:           -34.615662,
			lon:           -58.503337,
			expectedPoint: geopoint.Value(31659983379082793),
			expectedCode:  "0707A:964EE:7AE29",
		},
	}

	for _, tc := range tcl {
		t.Run(tc.name, func(t *testing.T) {
			out := geopoint.Encode(tc.lat, tc.lon)
			if out != tc.expectedPoint {
				t.Fatalf("Invalid result: expected %d but got %d", tc.expectedPoint, uint64(out))
			}
			if out.Code() != tc.expectedCode {
				t.Fatalf("Invalid result: expected %s but got %s", tc.expectedCode, out.Code())
			}
		})
	}
}

func TestEncoder_Decode(t *testing.T) {

	tcl := []struct {
		name        string
		point       geopoint.Value
		expectedLat float64
		expectedLon float64
		expectedErr error
	}{
		{
			name:        "Place du capitole, Toulouse, France",
			point:       geopoint.Value(75071988303315493),
			expectedLat: 43.603574,
			expectedLon: 1.442917,
		},
		{
			name:        "Mairie de Toulouse, Toulouse, France",
			point:       geopoint.Value(75071989061436701),
			expectedLat: 43.604297,
			expectedLon: 1.443677,
		},
		{
			name:        "Tour Eiffel, Paris, France",
			point:       geopoint.Value(77888104758015428),
			expectedLat: 48.858373,
			expectedLon: 2.292292,
		},
		{
			name:        "Montréal, Quebec, Canada",
			point:       geopoint.Value(76116476767848432),
			expectedLat: 45.558196,
			expectedLon: -73.870384,
		},
		{
			name:        "Buenos Aires, Argentina",
			point:       geopoint.Value(31659983379082793),
			expectedLat: -34.615662,
			expectedLon: -58.503337,
		},
	}

	for _, tc := range tcl {
		t.Run(tc.name, func(t *testing.T) {
			lat, lon, err := geopoint.Decode(tc.point)
			if err != tc.expectedErr {
				t.Fatalf("Invalid result: Error expected %v, go %v.", tc.expectedErr, err)
			}
			if lat != tc.expectedLat {
				t.Fatalf("Invalid result: expected latitude %0.10f but got %0.10f", tc.expectedLat, lat)
			}
			if lon != tc.expectedLon {
				t.Fatalf("Invalid result: expected latitude %0.10f but got %0.10f", tc.expectedLon, lon)
			}
		})
	}
}

func TestEncoder_DecodeString(t *testing.T) {

	tcl := []struct {
		name        string
		input       string
		expectedLat float64
		expectedLon float64
		expectedErr error
	}{
		{
			name:        "Place du capitole, Toulouse, France",
			input:       "10AB5:935B6:6C225",
			expectedLat: 43.603574,
			expectedLon: 1.442917,
		},
		{
			name:        "Mairie de Toulouse, Toulouse, France",
			input:       "10AB5:93889:6C51D",
			expectedLat: 43.604297,
			expectedLon: 1.443677,
		},
		{
			name:        "Tour Eiffel, Paris, France",
			input:       "114B6:D1905:475C4",
			expectedLat: 48.858373,
			expectedLon: 2.292292,
		},
		{
			name:        "Montréal, Quebec, Canada",
			input:       "10E6B:88474:D47F0",
			expectedLat: 45.558196,
			expectedLon: -73.870384,
		},
		{
			name:        "Buenos Aires, Argentina",
			input:       "0707A:964EE:7AE29",
			expectedLat: -34.615662,
			expectedLon: -58.503337,
		},
		{
			name:        "CryptoPan - Place du capitole, Toulouse, France",
			input:       "10A4D:937C9:8A1DA",
			expectedLat: 43.604105,
			expectedLon: -103.565722,
		},
		{
			name:        "CryptoPan - Mairie de Toulouse, Toulouse, France",
			input:       "10A4D:9390E:73D33",
			expectedLat: 43.604430,
			expectedLon: -103.474419,
		},
		{
			name:        "CryptoPan - Buenos Aires, Argentina",
			input:       "08E35:69AEF:9AE2B",
			expectedLat: -19.432879,
			expectedLon: -127.634411,
		},
	}

	for _, tc := range tcl {
		t.Run(tc.name, func(t *testing.T) {
			lat, lon, err := geopoint.DecodeString(tc.input)
			if err != tc.expectedErr {
				t.Fatalf("Invalid result: Error expected %v, go %v.", tc.expectedErr, err)
			}
			if lat != tc.expectedLat {
				t.Fatalf("Invalid result: expected latitude %0.10f but got %0.10f", tc.expectedLat, lat)
			}
			if lon != tc.expectedLon {
				t.Fatalf("Invalid result: expected longitude %0.10f but got %0.10f", tc.expectedLon, lon)
			}
		})
	}
}

// -----------------------------------------------------------------------------

func BenchmarkDecoder_Decode(b *testing.B) {
	p := geopoint.Value(31659983379082793)
	for i := 0; i < b.N; i++ {
		_, _, _ = geopoint.Decode(p)
	}
}

func BenchmarkCodec_Encode(b *testing.B) {
	lat, lon := 43.603574, 1.442917
	for i := 0; i < b.N; i++ {
		_ = geopoint.Encode(lat, lon)
	}
}
