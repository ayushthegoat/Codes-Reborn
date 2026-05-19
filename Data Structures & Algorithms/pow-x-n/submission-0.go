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
    return x * solve(x, n - 1)
}