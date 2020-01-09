## butter [![Go Report Card](https://goreportcard.com/badge/github.com/jfcg/butter)](https://goreportcard.com/report/github.com/jfcg/butter) [![GoDoc](https://godoc.org/github.com/jfcg/butter?status.svg)](https://godoc.org/github.com/jfcg/butter)
Digital Filter Library for Signal Processing.

This library consists of easy-to-use Butterworth first & second order digital filters.
You can calculate cutoff parameters with:
```
wc = 2 * pi * (desired cutoff in hz) / (sample rate in hz) =
	(desired cutoff in rad/s) * (sample period in sec)
```
Internally cutoff parameters are prewarped for correct operation. All New\*() functions
return nil if parameters are invalid.
