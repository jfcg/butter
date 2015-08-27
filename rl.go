package butter

type RateLimit struct {
	pu, mc float64 // previous input, max change per step
}

// Runs rate limiter one step with input u and returns output
func (rl *RateLimit) Next(u float64) float64 {
	d := u - rl.pu // input - prev_input
	if d > rl.mc {
		u = rl.pu + rl.mc
	} else if d < -rl.mc {
		u = rl.pu - rl.mc
	}
	rl.pu = u
	return u
}

func (rl *RateLimit) NextS(u, y []float64) {
	n := len(u)
	if n > len(y) {
		n = len(y)
	}
	for i := 0; i < n; i++ {
		y[i] = rl.Next(u[i])
	}
}

// Creates a rate limiter with initial input and max change per step. You can calculate mc with:
// 	mc = (desired rate limit in 1/sec) / (sample rate in hz)
func NewRateLimit(iu, mc float64) *RateLimit {
	if mc > 0 {
		return &RateLimit{iu, mc}
	}
	return nil
}
