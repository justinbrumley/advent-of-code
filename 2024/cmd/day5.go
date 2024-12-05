package cmd

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"

	"github.com/justinbrumley/advent-of-code/2024/utils"
	"github.com/spf13/cobra"
)

type Rule struct {
	BaseNumber          int
	PrerequisiteNumbers []int
}

type Update struct {
	Pages   []int
	RuleMap map[int]*Rule
}

func (r *Rule) Contains(num int) bool {
	for _, pre := range r.PrerequisiteNumbers {
		if pre == num {
			return true
		}
	}

	return false
}

// IsOrdered loops through each page in update and makes sure all prequisite pages are met.
// Prereq pages don't HAVE to appear, so we just check that none appear after the base page.
func (u *Update) IsOrdered() bool {
	for i, page := range u.Pages {
		// Nothing to check, or on the last page, so continue
		if u.RuleMap[page] == nil || i == len(u.Pages)-1 {
			continue
		}

		rule := u.RuleMap[page]

		// Check everything AFTER the current page
		for j := i + 1; j < len(u.Pages); j++ {
			if rule.Contains(u.Pages[j]) {
				// Prereq found AFTER number, so this isn't a valid update
				return false
			}
		}
	}

	return true
}

// Order pages (if unordered) using rules attached to update.
func (u *Update) Order() {
	sort.Slice(u.Pages, func(i, j int) bool {
		x := u.Pages[i]
		y := u.Pages[j]

		// Get rule for x
		rule := u.RuleMap[x]
		if rule == nil {
			// No rules, so they are currently fine where they are
			return true
		}

		// If y is a prereq of x, then y is NOT less than x
		if rule.Contains(y) {
			return false
		}

		return true
	})
}

func (u *Update) GetCenterPage() int {
	return u.Pages[len(u.Pages)/2]
}

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "Advent of Code 2024 - Day 5",
	Run: func(cmd *cobra.Command, args []string) {
		lines, err := utils.GetInput("inputs/day5")
		if err != nil {
			log.Fatal(err)
		}

		// Build a map to make rules easier to lookup and edit
		re := regexp.MustCompile(`\d+`)

		readingUpdates := false

		ruleMap := make(map[int]*Rule)

		total := 0
		unorderedTotal := 0

		for _, line := range lines {
			if line == "" {
				// Done parsing rules, switching to reading updates
				readingUpdates = true
				continue
			}

			if readingUpdates {
				// Use parsed rules map to check each line
				parts := re.FindAllString(line, -1)

				update := &Update{
					Pages:   make([]int, 0),
					RuleMap: ruleMap,
				}

				for _, val := range parts {
					num, _ := strconv.Atoi(val)
					update.Pages = append(update.Pages, num)
				}

				if update.IsOrdered() {
					total += update.GetCenterPage()
				} else {
					// Order things, then add to separate total
					update.Order()
					unorderedTotal += update.GetCenterPage()
				}

				continue
			}

			// Still reading in rules
			parts := re.FindAllString(line, -1)

			pre, _ := strconv.Atoi(parts[0])
			base, _ := strconv.Atoi(parts[1])

			if ruleMap[base] == nil {
				ruleMap[base] = &Rule{
					BaseNumber:          base,
					PrerequisiteNumbers: make([]int, 0),
				}
			}

			ruleMap[base].PrerequisiteNumbers = append(ruleMap[base].PrerequisiteNumbers, pre)
		}

		fmt.Printf("Total of center # from ordered updates: %v\n", total)
		fmt.Printf("Total of center # from (fixed) unordered updates: %v\n", unorderedTotal)
	},
}

func init() {
	rootCmd.AddCommand(day5Cmd)
}
