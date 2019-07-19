package geopoint_test

import (
	"testing"

	"go.zenithar.org/geopoint"
)

// testKey is the key used in the original Crypto-PAn source distribution
// sample.
var testKey = []byte{21, 34, 23, 141, 51, 164, 207, 128, 19, 10, 91, 22, 73, 144, 125, 16, 216, 152, 143, 131, 121, 121, 101, 39, 98, 87, 76, 45, 42, 132, 34, 2}

func TestCryptoPan(t *testing.T) {
	cpan, err := geopoint.NewCryptoPan(testKey)
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

// -----------------------------------------------------------------------------

// BenchmarkCryptopanIPv4 benchmarks annonymizing IPv4 addresses.
func BenchmarkCryptopanPoint(b *testing.B) {
	cpan, err := geopoint.NewCryptoPan(testKey)
	if err != nil {
		b.Fatal("New(testKey) failed:", err)
	}
	b.ResetTimer()

	point := geopoint.Value(75071988303315493)
	for i := 0; i < b.N; i++ {
		_ = cpan.Anonymize(point)
	}
}
