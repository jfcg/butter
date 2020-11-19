## butter [![Go Report Card](https://goreportcard.com/badge/github.com/jfcg/butter)](https://goreportcard.com/report/github.com/jfcg/butter) [![go.dev ref](https://raw.githubusercontent.com/jfcg/.github/main/godev.svg)](https://pkg.go.dev/github.com/jfcg/butter)
Digital Filter Library for Signal Processing.

This library consists of easy-to-use Butterworth first & second order digital filters.
You can calculate cutoff parameters with:
```
wc = 2 * pi * (desired cutoff in Hz) / (sample rate in Hz) =
	(desired cutoff in rad/sec) * (sample period in sec)
```
Internally cutoff parameters are prewarped for correct operation. All `New*()` functions
return nil if parameters are invalid.
Also, butter uses [semantic](https://semver.org) versioning.

### Support
If you use butter and like it, please support via ETH:`0x464B840ee70bBe7962b90bD727Aac172Fa8B9C15`
