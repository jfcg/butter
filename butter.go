/*	Copyright (c) 2020, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

// Package butter provides a Digital Filter Library for Signal Processing.
//
// This library consists of easy-to-use Butterworth first & second order digital filters.
// You can calculate cutoff parameters with:
//
//	wc = 2 * pi * (desired cutoff in Hz) / (sample rate in Hz) =
//		(desired cutoff in rad/sec) * (sample period in sec)
//
// Internally cutoff parameters are prewarped for correct operation.
// All New*() functions return nil if parameters are invalid.
package butter

// Filter represents a generic filter
type Filter interface {
	// Next runs filter one step with input u and returns output
	Next(float64) float64

	// NextS runs filter on a slice of input u and writes result in output y.
	// Input and output can be the same buffer.
	NextS(u, y []float64)

	// Reset sets internals of filter such that y will be the next output assuming
	// u is the next input approximately.
	Reset(u, y float64)
}

func valid(wc float64) bool {
	return .0001 < wc && wc < 3.1415 // reasonable bound
}

func valid2(wc, wd float64) bool {
	return .0001 < wc && wc+.0001 < wd && wd < 3.1415
}

// returns prewarped analog cut-off * dt, given desired cut-off * dt
func prewarp(wc float64) float64 {
	x := wc / 2
	x *= x
	return wc * (15 - x) / (15 - 6*x) // wa * dt
}
