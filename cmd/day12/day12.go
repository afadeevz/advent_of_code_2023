package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	foldRatePart1 = 1
	foldRatePart2 = 5
)

func getFoldRate(isPart2 bool) int {
	if !isPart2 {
		return foldRatePart1
	} else {
		return foldRatePart2
	}
}

type Elem = int
type Spring = []Elem

const (
	Operational Elem = iota
	Damaged
	Unknown
)

func main() {
	const isPart2 = true

	foldRate := getFoldRate(isPart2)
	slog.Info("got answer", "answer", run(os.Stdin, foldRate))
}

func run(input io.Reader, foldRate int) (answer int) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		answer += processLine(line, foldRate)
	}

	return
}

func processLine(s string, foldRate int) (answer int) {
	spring, nums := parseLine(s, foldRate)
	solver := &Solver{spring, nums}
	return solver.Solve()
}

type Solver struct {
	spring Spring
	nums   []int
}

func (s *Solver) Solve() int {
	return s.solve(0, 0, 0)
}

func (s *Solver) solve(idx int, numIdx int, curLen int) int {
	if idx >= len(s.spring) {
		if curLen > 0 {
			numIdx++
		}

		if numIdx != len(s.nums) {
			// fmt.Println("BD too low\t", s.spring, s.nums)
			return 0
		}

		// fmt.Println("OK", s.spring, s.nums)
		return 1
	}

	switch s.spring[idx] {
	case Operational:
		if curLen != 0 {
			if curLen != s.nums[numIdx] {
				// fmt.Println("BD wrong len\t", s.spring, s.nums)
				return 0
			}

			if numIdx == len(s.nums) {
				// fmt.Println("BD too many\t", s.spring, s.nums)
				return 0
			}

			return s.solve(idx+1, numIdx+1, 0)
		}
		return s.solve(idx+1, numIdx, 0)

	case Damaged:
		if numIdx >= len(s.nums) {
			// fmt.Println("BD too many\t", s.spring, s.nums)
			return 0
		}
		if curLen >= s.nums[numIdx] {
			// fmt.Println("BD too long\t", s.spring, s.nums)
			return 0
		}
		return s.solve(idx+1, numIdx, curLen+1)

	case Unknown:
		defer func() {
			s.spring[idx] = Unknown
		}()

		s.spring[idx] = Operational
		answer := s.solve(idx, numIdx, curLen)

		s.spring[idx] = Damaged
		answer += s.solve(idx, numIdx, curLen)

		return answer

	default:
		panic("unknown elem value")
	}
}

func solveSlow(spring Spring, nums []int) (answer int) {
	indices := make([]int, 0, len(spring))
	combinations := 1
	for i, elem := range spring {
		if elem != Unknown {
			continue
		}

		combinations *= 2
		indices = append(indices, i)
	}

	for combination := 0; combination < combinations; combination++ {
		combination := combination
		for _, i := range indices {
			spring[i] = combination % 2
			combination /= 2
		}

		answer += validateSpring(spring, nums)
	}

	return
}

func validateSpring(spring Spring, nums []int) int {
	lens := make([]int, 0, len(spring)/2)
	curLen := 0

	for _, elem := range spring {
		if elem == Damaged {
			curLen++
		} else if curLen > 0 {
			lens = append(lens, curLen)
			curLen = 0
		}
	}
	if curLen > 0 {
		lens = append(lens, curLen)
	}

	if !slices.Equal(nums, lens) {
		return 0
	}

	return 1
}

func parseLine(s string, foldRate int) (spring Spring, nums []int) {
	parts := strings.Split(s, " ")

	springStr := parts[0]
	spring = parseSpring(springStr, foldRate)

	numsStr := parts[1]
	nums = parseNums(numsStr, foldRate)

	return
}

func parseNums(numsStr string, foldRate int) (nums []int) {
	numsStrs := strings.Split(numsStr, ",")
	numsNotFolded := make([]int, 0, len(numsStrs))
	for _, numStr := range numsStrs {
		num, _ := strconv.ParseInt(numStr, 10, 64)
		numsNotFolded = append(numsNotFolded, int(num))
	}

	nums = make([]int, 0, len(numsNotFolded)*foldRate)
	for i := 0; i < foldRate; i++ {
		nums = append(nums, numsNotFolded...)
	}

	return
}

func parseSpring(s string, foldRate int) (spring Spring) {
	springNotFolded := make([]int, 0, len(s))

	for _, r := range s {
		springNotFolded = append(springNotFolded, parseElem(r))
	}

	spring = make(Spring, 0, (len(springNotFolded)+1)*foldRate-1)
	for i, first := 0, true; i < foldRate; i++ {
		if !first {
			// spring = append(spring, Unknown)
		} else {
			first = false
		}

		spring = append(spring, springNotFolded...)

	}

	return
}

func parseElem(r rune) Elem {
	switch r {
	case '.':
		return Operational
	case '#':
		return Damaged
	case '?':
		return Unknown
	default:
		panic("unknown elem: " + string(r))
	}
}
