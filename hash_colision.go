package main

import (
	"fmt"
	"math"
)

// hashCollisionProbability calculates the hash collision probability for a given hash algorithm and character length.
func hashCollisionProbability(hashSize, numEntries float64) float64 {
	return (1 - math.Pow(math.E, -(math.Pow(numEntries, 2)/(2*hashSize)))) * 100
}

func main() {
	// Assuming a hash algorithm that produces a 32-bit hash
	hashSize := float64(32) // size of the hash in bits

	// Number of entries or strings being hashed
	numEntries := float64(1000) // number of entries being hashed

	probability := hashCollisionProbability(hashSize, numEntries)

	fmt.Printf("Hash Collision Probability: %.2f%%\n", probability)
}
