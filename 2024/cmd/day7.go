package cmd

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"

	"github.com/justinbrumley/advent-of-code/2024/utils"
	"github.com/spf13/cobra"
)

type Equation struct {
	Solution uint64
	Parts    []uint64
}

// Build list of operators using number, converted to binary, padded to number
// of parts needed.
func (e *Equation) GetOperators(num int) []byte {
	str := strconv.FormatInt(int64(num), 3)

	if len(str) < len(e.Parts)-1 {
		// Pad with zeroes
		for i := len(str); i < len(e.Parts)-1; i++ {
			str = "0" + str
		}
	}

	out := make([]byte, len(e.Parts)-1)

	for i, b := range []byte(str) {
		if b == '0' {
			out[i] = '+'
		} else if b == '1' {
			out[i] = '*'
		} else if b == '2' {
			out[i] = '|'
		}
	}

	return out
}

// IsPossible determines if the parts of the solution can be combined
// using addition and multiplication to find the solution.
func (e *Equation) IsPossible() bool {
	permutations := int(math.Pow(3, float64(len(e.Parts)-1)))

	for i := 0; i < permutations; i++ {
		ops := e.GetOperators(i)

		val := e.Parts[0]

		for j := 1; j < len(e.Parts); j++ {
			op := ops[j-1]

			if op == '+' || op == 0 {
				val += e.Parts[j]
			} else if op == '*' {
				val *= e.Parts[j]
			} else if op == '|' {
				// Concat current value with next, then return to being a number
				res := fmt.Sprintf("%v%v", val, e.Parts[j])
				val, _ = strconv.ParseUint(res, 10, 64)
			}
		}

		if val == e.Solution {
			return true
		}
	}

	return false
}

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use:   "day7",
	Short: "Advent of Code 2024 - Day 7",
	Run: func(cmd *cobra.Command, args []string) {
		defer timer("day7")()

		if len(args) < 1 {
			log.Fatal("Missing required argument: <input_file>")
		}

		lines, err := utils.GetInput(args[0])
		if err != nil {
			log.Fatal(err)
		}

		re := regexp.MustCompile(`\d+`)

		var sum uint64

		for _, line := range lines {
			numbers := re.FindAllString(line, -1)

			var solution uint64
			parts := make([]uint64, 0)

			for i, n := range numbers {
				num, _ := strconv.ParseUint(n, 10, 64)

				if i == 0 {
					solution = num
				} else {
					parts = append(parts, num)
				}
			}

			eq := &Equation{
				Solution: solution,
				Parts:    parts,
			}

			if eq.IsPossible() {
				sum += eq.Solution
			}
		}

		fmt.Printf("Sum of possible equations: %v\n", sum)
	},
}

func init() {
	rootCmd.AddCommand(day7Cmd)
}
