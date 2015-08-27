package butter

// Second order filter
type filter2 struct {
	s2, s1, a1, a2, b0, b1, b2 float64
}

func (f *filter2) Next(u float64) float64 {
	t := u - f.a1*f.s1 - f.a2*f.s2 // Direct Form 2
	y := f.b0*t + f.b1*f.s1 + f.b2*f.s2
	f.s2, f.s1 = f.s1, t // shift memory
	return y
}

func (f *filter2) NextS(u, y []float64) {
	n := len(u)
	if n > len(y) {
		n = len(y)
	}
	for i := 0; i < n; i++ {
		y[i] = f.Next(u[i])
	}
}

func (f *filter2) Disable() {
	f.b0 = 1
	f.s1, f.s2, f.b1, f.b2, f.a1, f.a2 = 0, 0, 0, 0, 0, 0
}

func (f *filter2) Reset(u, y float64) {
	if f.b0 == 1 { // b0=1 only when second order filter is disabled
		return
	}
	// Internal states will be in arithmetic progression except HP
	// All divisions are safe

	hp := f.b1 == -2*f.b0 // HP?

	if hp {
		f.s1 = -y / (2 * f.b0)
		f.s2 = u - f.a1*f.s1
	} else {
		f.s1 = y / (f.b1 + 2*f.b0)
		f.s2 = u - (f.a1+2)*f.s1
	}

	if f.b2 < 0 || hp { // BP or HP?
		f.s2 /= 1 + f.a1 + f.a2
		f.s1 += f.s2
	} else {
		f.s2 /= f.a2 - 1
	}
}

// Create second order Low-Pass filter
func NewLowPass2(wc float64) Filter {
	if !valid(wc) {
		return nil
	}
	var f filter2
	wat := prewarp(wc) // wa * dt
	wat2 := wat * wat
	wat *= 2.82842712474619 // sqrt(8)

	a0 := wat2 + 4 + wat
	f.a1 = 2 * (wat2 - 4) / a0
	f.a2 = (wat2 + 4 - wat) / a0
	f.b0 = wat2 / a0
	f.b2 = f.b0
	f.b1 = 2 * f.b0
	f.s2, f.s1 = 0, 0
	return &f
}

// Create second order High-Pass filter
func NewHighPass2(wc float64) Filter {
	if !valid(wc) {
		return nil
	}
	var f filter2
	wat := prewarp(wc) // wa * dt
	wat2 := wat * wat
	wat *= 2.82842712474619 // sqrt(8)

	a0 := wat2 + 4 + wat
	f.a1 = 2 * (wat2 - 4) / a0
	f.a2 = (wat2 + 4 - wat) / a0
	f.b0 = 4 / a0
	f.b2 = f.b0
	f.b1 = -2 * f.b0
	f.s2, f.s1 = 0, 0
	return &f
}

// Create second order Band-Pass filter
func NewBandPass2(wc, wd float64) Filter {
	if !valid2(wc, wd) {
		return nil
	}
	var f filter2
	wat := prewarp(wc) // wa * dt
	wbt := prewarp(wd) // wb * dt
	wbwa := wbt * wat  // wb * wa * dt^2
	wd = 2 * (wbt - wat)

	a0 := wbwa + 4 + wd
	f.a1 = 2 * (wbwa - 4) / a0
	f.a2 = (wbwa + 4 - wd) / a0
	f.b0 = wd / a0
	f.b2 = -f.b0
	f.b1, f.s2, f.s1 = 0, 0, 0
	return &f
}

// Create second order Band-Stop filter
func NewBandStop2(wc, wd float64) Filter {
	if !valid2(wc, wd) {
		return nil
	}
	var f filter2
	wat := prewarp(wc) // wa * dt
	wbt := prewarp(wd) // wb * dt
	wbwa := wbt * wat  // wb * wa * dt^2
	wd = 2 * (wbt - wat)

	a0 := wbwa + 4 + wd
	f.a1 = 2 * (wbwa - 4) / a0
	f.b1 = f.a1
	f.a2 = (wbwa + 4 - wd) / a0
	f.b0 = (wbwa + 4) / a0
	f.b2 = f.b0
	f.s2, f.s1 = 0, 0
	return &f
}
