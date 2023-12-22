package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Solution struct {
	sol  []string
	nums []int
	ops  map[string]struct{}
}

type PartSol struct {
	sum  string
	answ int
}

func absi(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func find_solution(target int, sol Solution) {
	var solutions []Solution
	solutions = append(solutions, sol)
	for len(solutions) > 0 {
		s := solutions[0]
		solutions = solutions[1:]
		if len(s.nums) < 2 {
			continue
		}
		//fmt.Printf("Considering: %#v\n", s)
		for i, n1 := range s.nums {
			for j, n2 := range s.nums {
				if j <= i {
					continue
				}
				sols := make([]PartSol, 0)
				if _, canadd := sol.ops["+"]; canadd {
					sols = append(sols, PartSol{sum: fmt.Sprintf("%d+%d=%d", n1, n2, n1+n2), answ: n1 + n2})
				}
				var lrg, sml int
				if n1 > n2 {
					lrg = n1
					sml = n2
				} else {
					lrg = n2
					sml = n1
				}
				if _, cansub := sol.ops["-"]; cansub {
					sols = append(sols, PartSol{sum: fmt.Sprintf("%d-%d=%d", lrg, sml, lrg-sml), answ: lrg - sml})
				}
				if _, canmult := sol.ops["*"]; canmult {
					sols = append(sols, PartSol{sum: fmt.Sprintf("%d*%d=%d", n1, n2, n1*n2), answ: n1 * n2})
				}
				if _, candiv := sol.ops["/"]; candiv && sml > 0 && lrg%sml == 0 {
					sols = append(sols, PartSol{sum: fmt.Sprintf("%d/%d=%d", lrg, sml, lrg/sml), answ: lrg / sml})
				}
				for _, psol := range sols {
					var soladd Solution
					soladd.sol = make([]string, len(s.sol), len(s.sol)+1)
					copy(soladd.sol, s.sol)
					soladd.sol = append(soladd.sol, psol.sum)
					soladd.nums = []int{psol.answ}
					for c, n := range s.nums {
						if c != i && c != j {
							soladd.nums = append(soladd.nums, n)
						}
					}
					if absi(target-psol.answ) <= 2 {
						fmt.Printf("Possible solution: %d (%d): %s\n", target, psol.answ, strings.Join(soladd.sol, "; "))
					}
					solutions = append(solutions, soladd)
				}
			}
		}
	}
}

func main() {
	validops := map[string]struct{}{
		"+": struct{}{},
		"*": struct{}{},
		"-": struct{}{},
		"/": struct{}{},
	}
	if len(os.Args) < 3 {
		panic("Provide target and input numbers")
	}
	target, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Target should be integer: %s", err))
	}
	sol := Solution{ops: make(map[string]struct{})}
	for _, argstr := range os.Args[2:] {
		if _, isop := validops[argstr]; isop {
			sol.ops[argstr] = struct{}{}
		} else if arg, err := strconv.Atoi(argstr); err != nil {
			panic(fmt.Sprintf("Digits should be integers: %s", err))
		} else {
			if arg == target {
				panic("No fun, the target is present in the input")
			}
			sol.nums = append(sol.nums, arg)
		}
	}
	if len(sol.ops) == 0 {
		sol.ops = validops
	}
	find_solution(target, sol)
}
