func myPow(x float64, n int) float64 {
    if n == 0 {
		return 1
	}
	if n < 0 {
    // Compute positive power first, then take reciprocal
    x = 1 / x
    n = -n
}
	return solve(x, n)
}
func solve(x float64, n int) float64 {
    if n == 1 {
        return x
    }
	if n % 2 == 0 {
		temp := solve(x, n/2)
		return temp * temp
		//return solve(x, n/2) * solve(x, n/2)
	}
	temp := solve(x, n/2)
	return x * temp * temp
	//return x * solve(x, n/2) * solve(x, n/2)
}