// Digital Filter Library for Signal Processing
//
// This library consists of easy-to-use Butterworth first & second order digital filters. You can calculate cutoff parameters with:
// 	wc = 2 * pi * (desired cutoff in hz) / (sample rate in hz) = (desired cutoff in rad/s) * (sample period in sec)
// Internally cutoff parameters are prewarped for correct operation. All New*() functions return nil if parameters are invalid.
package butter

// Generic filter type
type Filter interface {
	Next(float64) float64 // Runs filter one step with input u and returns output
	Disable()             // output=input, filter becomes not resettable
	Reset(u, y float64)   // Reset filter such that y will be the next output
}

func valid(wc float64) bool {
	return .0001 < wc && wc < 3.1415 // reasonable bound
}

func valid2(wc, wd float64) bool {
	return .0001 < wc && wc < wd && wd < 3.1415
}

// returns prewarped analog cut-off * dt, given desired cut-off * dt
func prewarp(wc float64) float64 {
	x := wc / 2
	x *= x
	return wc * (15 - x) / (15 - 6*x) // wa * dt
}
