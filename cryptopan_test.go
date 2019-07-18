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
			point:        geopoint.Value(75071988303315493),
			expected:     geopoint.Value(74957639650943450),
			expectedCode: "10A4D:937C9:8A1DA",
		},
		{
			point:        geopoint.Value(75071989061436701),
			expected:     geopoint.Value(74957639991639347),
			expectedCode: "10A4D:9390E:73D33",
		},
		{
			point:        geopoint.Value(77888104758015428),
			expected:     geopoint.Value(78592826413445701),
			expectedCode: "11737:C25C4:47245",
		},
		{
			point:        geopoint.Value(31659983379082793),
			expected:     geopoint.Value(40028174716349995),
			expectedCode: "08E35:69AEF:9AE2B",
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
