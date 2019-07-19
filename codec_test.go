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
		name           string
		lat            float64
		lon            float64
		expectedPoint  geopoint.Value
		expectedCode   string
		expectedBase58 string
	}{
		{
			name:           "Place du capitole, Toulouse, France",
			lat:            43.603574,
			lon:            1.442917,
			expectedPoint:  geopoint.Value(75071809151126838),
			expectedCode:   "10AB5:69A51:94D36",
			expectedBase58: "B7DA8NMQkm",
		},
		{
			name:           "Mairie de Toulouse, Toulouse, France",
			lat:            43.604297,
			lon:            1.443677,
			expectedPoint:  geopoint.Value(75071809155908323),
			expectedCode:   "10AB5:69A56:242E3",
			expectedBase58: "B7DA8Nmv8A",
		},
		{
			name:           "Tour Eiffel, Paris, France",
			lat:            48.858373,
			lon:            2.292292,
			expectedPoint:  geopoint.Value(77887690747650097),
			expectedCode:   "114B6:712B6:3A031",
			expectedBase58: "BVCUZaWBXe",
		},
		{
			name:           "Montréal, Quebec, Canada",
			lat:            45.558196,
			lon:            -73.870384,
			expectedPoint:  geopoint.Value(76116863733120784),
			expectedCode:   "10E6B:E2603:ABF10",
			expectedBase58: "BFNTwT9CgF",
		},
		{
			name:           "Buenos Aires, Argentina",
			lat:            -34.615662,
			lon:            -58.503337,
			expectedPoint:  geopoint.Value(31659800001010902),
			expectedCode:   "0707A:6B9CB:85CD6",
			expectedBase58: "15GDnG7CTVX",
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
			if out.Base58() != tc.expectedBase58 {
				t.Fatalf("Invalid result: expected %s but got %s", tc.expectedBase58, out.Base58())
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
			point:       geopoint.Value(75071809151126838),
			expectedLat: 43.603574,
			expectedLon: 1.442917,
		},
		{
			name:        "Mairie de Toulouse, Toulouse, France",
			point:       geopoint.Value(75071809155908323),
			expectedLat: 43.604297,
			expectedLon: 1.443677,
		},
		{
			name:        "Tour Eiffel, Paris, France",
			point:       geopoint.Value(77887690747650097),
			expectedLat: 48.858373,
			expectedLon: 2.292292,
		},
		{
			name:        "Montréal, Quebec, Canada",
			point:       geopoint.Value(76116863733120784),
			expectedLat: 45.558196,
			expectedLon: -73.870384,
		},
		{
			name:        "Buenos Aires, Argentina",
			point:       geopoint.Value(31659800001010902),
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
			input:       "10AB5:69A51:94D36",
			expectedLat: 43.603574,
			expectedLon: 1.442917,
		},
		{
			name:        "Mairie de Toulouse, Toulouse, France",
			input:       "10AB5:69A56:242E3",
			expectedLat: 43.604297,
			expectedLon: 1.443677,
		},
		{
			name:        "Tour Eiffel, Paris, France",
			input:       "114B6:712B6:3A031",
			expectedLat: 48.858373,
			expectedLon: 2.292292,
		},
		{
			name:        "Montréal, Quebec, Canada",
			input:       "10E6B:E2603:ABF10",
			expectedLat: 45.558196,
			expectedLon: -73.870384,
		},
		{
			name:        "Buenos Aires, Argentina",
			input:       "0707A:6B9CB:85CD6",
			expectedLat: -34.615662,
			expectedLon: -58.503337,
		},
		{
			name:        "CryptoPan - Place du capitole, Toulouse, France",
			input:       "10A4D:53A5D:54CC6",
			expectedLat: 43.868266,
			expectedLon: -103.116777,
		},
		{
			name:        "CryptoPan - Mairie de Toulouse, Toulouse, France",
			input:       "10A4D:53A58:3A163",
			expectedLat: 43.864537,
			expectedLon: -103.117189,
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
