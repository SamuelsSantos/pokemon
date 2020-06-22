package rand

var pool = "NSEO"

// GetRandomJorney return a string that contains a number o steps
func GetRandomJorney(steps int) string {
	bytes := make([]byte, steps)
	for i := 0; i < steps; i++ {
		bytes[i] = pool[seededRand.Intn(len(pool))]
	}
	return string(bytes)
}
