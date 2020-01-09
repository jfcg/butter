/*	Copyright (c) 2020, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package butter

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

type newFilter func(float64) Filter
type newFilter2 func(x, y float64) Filter

func Test1(tst *testing.T) { // test invalid cutoffs
	nf := []uint64{
		0x7FF8000000000001, // nan
		0xFFF8000000000010, // nan
		0x7FF0000000000000, // inf
		0xFFF0000000000000, // -inf
	}
	badwc := []float64{-1, 0, math.Pi, 4,
		*(*float64)(unsafe.Pointer(&nf[0])),
		*(*float64)(unsafe.Pointer(&nf[1])),
		*(*float64)(unsafe.Pointer(&nf[2])),
		*(*float64)(unsafe.Pointer(&nf[3])),
		1.5, // for nfl2
	}

	nfl := []newFilter{
		NewHighPass1,
		NewLowPass1,
		NewHighPass2,
		NewLowPass2,
	}
	nfl2 := []newFilter2{
		NewBandPass2,
		NewBandStop2,
	}

	for i := 0; i < len(badwc)-1; i++ {
		for k := 0; k < len(nfl); k++ {
			if nfl[k](badwc[i]) != nil {
				tst.Fatal("nfl must return nil", i, k)
			}
		}
	}

	for i := 0; i < len(badwc); i++ {
		for r := 0; r < len(badwc); r++ {
			for k := 0; k < len(nfl2); k++ {
				if nfl2[k](badwc[i], badwc[r]) != nil {
					tst.Fatal("nfl2 must return nil", i, r, k)
				}
			}
		}
	}
}

func Test2(tst *testing.T) {
	dt := 1. / 64              // 64 hz sample rate
	wc := 2 * math.Pi * 7 * dt // 7 hz cutoff
	fl := []Filter{
		NewHighPass1(wc),
		NewLowPass1(wc),
		NewHighPass2(wc),
		NewLowPass2(wc),
		NewBandPass2(wc/2, 2*wc),
		NewBandStop2(wc/2, 2*wc),
		NewRateLimit(0, wc),
	}
	fmt.Println("u  hp1  lp1  hp2  lp2  bp2  bs2  rl")
	n := 3 * 64 // 3 seconds
	t := .0     // time

	for i := 0; i < n; i++ {
		t += dt
		// linear chirp 3..15 hz
		u := math.Sin(2 * math.Pi * t * (3 + 2*t))
		fmt.Print(u, " ")

		for k := 0; k < len(fl); k++ {
			if i == 3*32 { // reset at t=1.5
				fl[k].Reset(u, -.8+float64(k)*.32)
			}
			fmt.Print(fl[k].Next(u), " ")
		}
		fmt.Println()
	}
}
