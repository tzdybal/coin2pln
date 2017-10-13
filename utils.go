package main

// removes dumplicates from sorted slice, in place (slice is not usable anymore)
func uniqueStrings(slice []string) []string {
	if len(slice) < 1 {
		return nil
	}
	insert := 1

	for read := 1; read < len(slice); read++ {
		if slice[read] != slice[insert-1] {
			slice[insert] = slice[read]
			insert++
		}
	}
	return slice[:insert]
}
