Digital Filter Library for Signal Processing
===
This library consists of easy-to-use Butterworth first & second order digital filters. You can calculate cutoff parameters with:<br>
wc = 2 * pi * (desired cutoff in hz) / (sample rate in hz) = (desired cutoff in rad/s) * (sample period in sec)<br>
Internally cutoff parameters are prewarped for correct operation. All New*() functions return nil if parameters are invalid. 
