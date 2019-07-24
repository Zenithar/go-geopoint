// +build cryptopan experimental

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

package anonymizer_test

import (
	"testing"

	"go.zenithar.org/geopoint"
	"go.zenithar.org/geopoint/anonymizer"
)

// testKey is the key used in the original Crypto-PAn source distribution
// sample.
var testKey = []byte{21, 34, 23, 141, 51, 164, 207, 128, 19, 10, 91, 22, 73, 144, 125, 16, 216, 152, 143, 131, 121, 121, 101, 39, 98, 87, 76, 45, 42, 132, 34, 2}

func TestCryptoPan(t *testing.T) {
	cpan, err := anonymizer.CryptoPan(testKey)
	if err != nil {
		t.Fatal("New(testKey) failed:", err)
	}

	vectors := []struct {
		point        geopoint.Value
		expected     geopoint.Value
		expectedCode string
	}{
		{
			point:        geopoint.Value(75071809151126838),
			expected:     geopoint.Value(74957365464878278),
			expectedCode: "10A4D:53A5D:54CC6",
		},
		{
			point:        geopoint.Value(75071809155908323),
			expected:     geopoint.Value(74957365459525987),
			expectedCode: "10A4D:53A58:3A163",
		},
		{
			point:        geopoint.Value(77887690747650097),
			expected:     geopoint.Value(78593030310600448),
			expectedCode: "11737:F1D57:C1F00",
		},
		{
			point:        geopoint.Value(31659800001010902),
			expected:     geopoint.Value(40028805103861206),
			expectedCode: "08E35:FC74F:FBDD6",
		},
	}

	for _, v := range vectors {
		result := cpan.Anonymize(v.point)
		if v.expected != result {
			t.Fatalf("invalid result, expected %d got %d", v.expected, result)
		}
		if v.expectedCode != result.Code() {
			t.Fatalf("invalid code, expected '%s' got '%s'", v.expectedCode, result.Code())
		}
	}
}

func TestCryptoPan_World(t *testing.T) {
	cpan, err := anonymizer.CryptoPan(testKey)
	if err != nil {
		t.Fatal("New(testKey) failed:", err)
	}

	vectors := []struct {
		lat                 float64
		lon                 float64
		point               geopoint.Value
		expectedCryptoPoint geopoint.Value
		expectedCryptoLat   float64
		expectedCryptoLon   float64
	}{
		{
			lat:                 0,
			lon:                 0,
			expectedCryptoPoint: geopoint.Value(14860500948746467),
			expectedCryptoLat:   -64.132617,
			expectedCryptoLon:   23.656909,
		},
		{
			lat:                 0,
			lon:                 90,
			expectedCryptoPoint: geopoint.Value(14918225837949661),
			expectedCryptoLat:   -64.161807,
			expectedCryptoLon:   76.145434,
		},
		{
			lat:                 0,
			lon:                 180,
			expectedCryptoPoint: geopoint.Value(14650440576121728),
			expectedCryptoLat:   -64.986960,
			expectedCryptoLon:   -168.496120,
		},
		{
			lat:                 0,
			lon:                 -180,
			expectedCryptoPoint: geopoint.Value(14650440576121728),
			expectedCryptoLat:   -64.986960,
			expectedCryptoLon:   -168.496120,
		},
		{
			lat:                 0,
			lon:                 -90,
			expectedCryptoPoint: geopoint.Value(14736189554943261),
			expectedCryptoLat:   -64.981431,
			expectedCryptoLon:   -90.424386,
		},
		{
			lat:                 45,
			lon:                 45,
			expectedCryptoPoint: geopoint.Value(76384112167864092),
			expectedCryptoLat:   45.788406,
			expectedCryptoLon:   170.853906,
		},
		{
			lat:                 -45,
			lon:                 -135,
			expectedCryptoPoint: geopoint.Value(46376032511439846),
			expectedCryptoLat:   -8.654170,
			expectedCryptoLon:   14.557021,
		},
		{
			lat:                 90,
			lon:                 90,
			expectedCryptoPoint: geopoint.Value(101619606975423775),
			expectedCryptoLat:   90.982583,
			expectedCryptoLon:   82.491075,
		},
		{
			lat:                 90,
			lon:                 0,
			expectedCryptoPoint: geopoint.Value(101519575039615292),
			expectedCryptoLat:   90.130582,
			expectedCryptoLon:   -9.654918,
		},
		{
			lat:                 -90,
			lon:                 0,
			expectedCryptoPoint: geopoint.Value(71833354328081954),
			expectedCryptoLat:   37.161824,
			expectedCryptoLon:   128.210965,
		},
	}

	for _, v := range vectors {
		v.point = geopoint.Encode(v.lat, v.lon)

		cryptoPoint := cpan.Anonymize(v.point)
		if cryptoPoint != v.expectedCryptoPoint {
			t.Errorf("invalid result, expected %d, got %d", uint64(v.expectedCryptoPoint), cryptoPoint)
		}

		lat, lon, err := geopoint.Decode(cryptoPoint)
		if err != nil {
			t.Fatalf("error %d should not be raised, got %v", uint64(v.expectedCryptoPoint), err)
		}
		if v.expectedCryptoLat != lat {
			t.Errorf("invalid latitude, %d expected %f, got %f", uint64(v.expectedCryptoPoint), v.expectedCryptoLat, lat)
		}
		if v.expectedCryptoLon != lon {
			t.Errorf("invalid longitude, %d expected %f, got %f", uint64(v.expectedCryptoPoint), v.expectedCryptoLon, lon)
		}
	}
}

// -----------------------------------------------------------------------------

// BenchmarkCryptopanIPv4 benchmarks annonymizing IPv4 addresses.
func BenchmarkCryptopanPoint(b *testing.B) {
	cpan, err := anonymizer.CryptoPan(testKey)
	if err != nil {
		b.Fatal("New(testKey) failed:", err)
	}
	b.ResetTimer()

	point := geopoint.Value(75071988303315493)
	for i := 0; i < b.N; i++ {
		_ = cpan.Anonymize(point)
	}
}
