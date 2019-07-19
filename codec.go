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
	"regexp"
	"strconv"
	"strings"
)

var (
	geoPointRegex = regexp.MustCompile("^[0-9A-Za-z]+:[0-9A-Za-z]{5}:[0-9A-Za-z]{5}$")
)

// Encode a point using Crypto-PAn algorithm
func Encode(latitude float64, longitude float64) Value {

	// Render as string (precision 10^6)
	coords := fmt.Sprintf("%.6f.%.6f", latitude, longitude)

	// Split in parts
	parts := strings.SplitN(coords, ".", 4)

	// Deocde parts as integer
	highLat, _ := strconv.Atoi(parts[0]) // [-90; 90] => log2(180) => 8bits
	highLat = ((highLat + 90) % 180)
	lowLat, _ := strconv.Atoi(parts[1])  // 10^6 => log2(10^6) => 20bits
	highLon, _ := strconv.Atoi(parts[2]) // [-180; 180] => log2(360) => 9bits
	highLon = ((highLon + 180) % 360)
	lowLon, _ := strconv.Atoi(parts[3]) // 10^6 => log2(10^6) => 20bits

	// Fill an uint64
	encoded := uint64(0)

	// Interleave LSB
	encoded |= interleave((uint32(lowLat)&0xFFFFF), uint32(lowLon&0xFFFFF)) & 0xFFFFFFFFFF
	// Rebase longitude to 0 to remove sign
	encoded = encoded | uint64((highLon&0x1FF)<<40)
	// Rebase latitude to 0 to remove sign
	encoded = encoded | uint64((highLat&0xFF)<<49)

	// Return a point
	return Value(encoded)
}

// Decode a point to retrieve (lat,lon)
func Decode(raw Value) (float64, float64, error) {

	value := uint64(raw)

	// Decode packed value
	highLat := int64((value>>49)&0xFF-90) % 180
	highLon := int64((value>>40)&0x1FF-180) % 360
	lowLat, lowLon := deinterleave(uint64(value & 0xFFFFFFFFFF))

	// Assemble values
	latStr := fmt.Sprintf("%d.%d", highLat, lowLat)
	lonStr := fmt.Sprintf("%d.%d", highLon, lowLon)

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return 0, 0, ErrInvalidGeoPointValue
	}
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return 0, 0, ErrInvalidGeoPointValue
	}

	return lat, lon, nil
}

// Check the given encoded point
func Check(raw string) error {
	// Check
	if !geoPointRegex.MatchString(raw) {
		return ErrInvalidGeoPointHash
	}

	// Syntaxically correct
	return nil
}

// DecodeString a point to retrieve (lat,lon)
func DecodeString(raw string) (float64, float64, error) {

	// Check given raw string
	if err := Check(raw); err != nil {
		return 0, 0, err
	}

	// Remove all ':'
	raw = strings.ReplaceAll(raw, ":", "")

	// Decode hexadecimal
	value, err := strconv.ParseUint(raw, 16, 64)
	if err != nil {
		return 0, 0, ErrInvalidGeoPointValue
	}

	// Delegate to decoder
	return Decode(Value(value))
}

// -----------------------------------------------------------------------------
// copied from https://github.com/mmcloughlin/geohash/blob/master/geohash.go

// Spread out the 32 bits of x into 64 bits, where the bits of x occupy even
// bit positions.
func spread(x uint32) uint64 {
	X := uint64(x)
	X = (X | (X << 16)) & 0x0000ffff0000ffff
	X = (X | (X << 8)) & 0x00ff00ff00ff00ff
	X = (X | (X << 4)) & 0x0f0f0f0f0f0f0f0f
	X = (X | (X << 2)) & 0x3333333333333333
	X = (X | (X << 1)) & 0x5555555555555555
	return X
}

// Interleave the bits of x and y. In the result, x and y occupy even and odd
// bitlevels, respectively.
func interleave(x, y uint32) uint64 {
	return spread(x) | (spread(y) << 1)
}

// Squash the even bitlevels of X into a 32-bit word. Odd bitlevels of X are
// ignored, and may take any value.
func squash(X uint64) uint32 {
	X &= 0x5555555555555555
	X = (X | (X >> 1)) & 0x3333333333333333
	X = (X | (X >> 2)) & 0x0f0f0f0f0f0f0f0f
	X = (X | (X >> 4)) & 0x00ff00ff00ff00ff
	X = (X | (X >> 8)) & 0x0000ffff0000ffff
	X = (X | (X >> 16)) & 0x00000000ffffffff
	return uint32(X)
}

// Deinterleave the bits of X into 32-bit words containing the even and odd
// bitlevels of X, respectively.
func deinterleave(X uint64) (uint32, uint32) {
	return squash(X), squash(X >> 1)
}
