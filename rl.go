package butter

type rateLimit struct {
	pu, mc float64 // previous input, max change per step
}

// Runs rate limiter one step with input u and returns output
func (rl *rateLimit) Next(u float64) float64 {
	d := u - rl.pu // input - prev_input
	if d > rl.mc {
		u = rl.pu + rl.mc
	} else if d < -rl.mc {
		u = rl.pu - rl.mc
	}
	rl.pu = u
	return u
}

func (rl *rateLimit) NextS(u, y []float64) {
	n := len(u)
	if n > len(y) {
		n = len(y)
	}
	for i := 0; i < n; i++ {
		y[i] = rl.Next(u[i])
	}
}

func (rl *rateLimit) Reset(u, y float64) {
	rl.pu = y
}

// Creates a rate limiter with initial input and max change per step. You can calculate mc with:
// 	mc = (desired rate limit in 1/sec) / (sample rate in hz)
// Note: Reset(u,y) on a rate limiter sets internal state to y (does not use u). The next output will be in [y-mc, y+mc].
func NewRateLimit(iu, mc float64) Filter {
	if mc > 0 {
		return &rateLimit{iu, mc}
	}
	return nil
}
