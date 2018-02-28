// mtrand.go -- Mersenne Twister with a known seed
//
// (c) 2015, 2016 Sudhi Herle <sudhi@herle.net>
//
// Licensing Terms: GPLv2
//
// If you need a commercial license for this work, please contact
// the author.
//
// This software does not come with any express or implied
// warranty; it is provided "as is". No claim  is made to its
// suitability for any purpose.

// Package mtrand is a Mersenne-Twister implementation (MT19937) for 32-bit
// output.
//
// Notes:
// * 32-bit Mersenne-Twister MT19937
// * Not safe for calling from multiple goroutines
// * Ref: https://en.wikipedia.org/wiki/Mersenne_Twister
package mtrand // github.com/opencoff/go-lib/mtrand

import (
	"time"
)

type MT struct {
	mt [624]uint32
	i  int
}

func New(seed uint32) *MT {
	m := &MT{}
	mt := m.mt[:]

	if seed == 0 {
		seed = uint32(time.Now().UnixNano())
	}

	m.mt[0] = seed

	for i := 1; i < 624; i++ {
		y := mt[i-1]
		y ^= (y >> 30)
		mt[i] = (1812433253 * y) + uint32(i)
	}

	return m
}

func (m *MT) twist() {
	mt := m.mt[:]

	for i := 0; i < 624; i++ {
		y := (mt[i] & 0x80000000) + (mt[(i+1)%624] & 0x7fffffff)
		mt[i] = mt[(i+397)%624] ^ (y >> 1)

		if (y & 1) != 0 {
			mt[i] ^= 0x9908b0df
		}
	}
	m.i = 0
}

func (m *MT) Next() uint32 {
	mt := m.mt[:]
	i := m.i

	if i >= 624 {
		i = 0
		m.twist()
	}

	y := mt[i]
	y ^= (y >> 11)
	y ^= ((y << 7) & 2636928640)
	y ^= ((y << 15) & 4022730752)
	y ^= (y >> 18)

	m.i = i + 1
	return y
}

// vim: ft=go:sw=4:ts=4:tw=78:expandtab:
