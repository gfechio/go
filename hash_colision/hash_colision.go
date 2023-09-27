package main

import (
        "fmt"
        "os"
        "math"

        "github.com/spf13/cobra"
)


func calculateCollisionProbability(hashLength int, numEntries float64) float64 {
        hashSize := float64(hashLength) * 8 // Size in bits
        return (1 - math.Pow(math.E, -(math.Pow(numEntries, 2)/(2*hashSize)))) * 100
}

func calculateCollisionProbabilityCmd(cmd *cobra.Command, args []string) {
        hashLength, _ := cmd.Flags().GetInt("hash-length")
        numEntries, _ := cmd.Flags().GetFloat64("num-entries")

        probability := calculateCollisionProbability(hashLength, numEntries)

        fmt.Printf("Hash Collision Probability with hash length %d: %.2f%%\n", hashLength, probability)
}

func main() {
        rootCmd := &cobra.Command{
                Use:   "hash-collision-calculator",
                Short: "Calculate hash collision probability",
                Run:   calculateCollisionProbabilityCmd,
        }

        // Flags
        rootCmd.Flags().IntP("hash-length", "l", 128, "Hash length in bits")
        rootCmd.Flags().Float64P("num-entries", "n", 1000, "Number of entries")

        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
