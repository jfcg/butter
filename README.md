## butter [![go report card](https://goreportcard.com/badge/github.com/jfcg/butter)](https://goreportcard.com/report/github.com/jfcg/butter) [![go.dev ref](https://pkg.go.dev/static/frontend/badge/badge.svg)](https://pkg.go.dev/github.com/jfcg/butter#pkg-overview)
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
See [Contributing](./.github/CONTRIBUTING.md), [Security](./.github/SECURITY.md) and [Support](./.github/SUPPORT.md) guides. Also if you use butter and like it, please support via [Github Sponsors](https://github.com/sponsors/jfcg) or:
- BTC:`bc1qr8m7n0w3xes6ckmau02s47a23e84umujej822e`
- ETH:`0x3a844321042D8f7c5BB2f7AB17e20273CA6277f6`
