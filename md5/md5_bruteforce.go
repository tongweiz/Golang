package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	// Target hash to match
	targetHash := "xxxxxxxxxxxxxxxxx"

	// Prefix string
	prefixes := []string{"transient", "fleeting", "momentary", "impermanent", "temporary"}

	start := time.Now()
	// Iterate over all possible 6-digit numbers (000000 to 999999)
	for _, prefix := range prefixes {
		for i := 0; i <= 999999; i++ {
			// Format the number as a 6-digit string (e.g., "000001", "000002", etc.)
			numStr := fmt.Sprintf("%06d", i)

			// Combine prefix with the number
			fullStr := prefix + numStr

			// Compute the MD5 hash
			hash := md5.Sum([]byte(fullStr))

			// Convert the hash to a hex string
			hashStr := hex.EncodeToString(hash[:])

			// Compare the generated hash with the target hash
			if hashStr == targetHash {
				elapsed := time.Since(start)
				fmt.Printf("Match found! The 6-digit number is: %s\n", numStr)
				fmt.Printf("Time taken: %s\n", elapsed)
				return
			}
		}
	}

	fmt.Println("No match found.")
}
