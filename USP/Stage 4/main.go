package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type candidate struct {
	firstName, lastName string
	gpa                 float64
	priorities          []string
}

func main() {
	var n int
	fmt.Scan(&n)

	candidates := []candidate{}

	func() {
		file, err := os.Open("applicants.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			data := strings.Split(line, " ")
			if len(data) != 6 {
				log.Fatal(errors.New("fields not match"))
			}
			gpa, _ := strconv.ParseFloat(data[2], 64)
			candidates = append(candidates, candidate{
				firstName:  data[0],
				lastName:   data[1],
				gpa:        gpa,
				priorities: data[3:],
			})
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	sortCandidates := func(c []candidate) {
		sort.Slice(c, func(i, j int) bool {
			if c[i].gpa != c[j].gpa {
				return c[i].gpa > c[j].gpa
			}
			return (c[i].firstName + c[i].lastName) < (c[j].firstName + c[j].lastName)
		})
	}
	sortCandidates(candidates)

	// fmt.Println(candidates)
	status := make([]bool, len(candidates))
	applicants := map[string][]candidate{
		"Mathematics": {},
		"Physics":     {},
		"Biotech":     {},
		"Chemistry":   {},
		"Engineering": {},
	}

	admit := func(itr int) {
		for i, cand := range candidates {
			if !status[i] {
				if len(applicants[cand.priorities[itr]]) < n {
					status[i] = true
					applicants[cand.priorities[itr]] = append(applicants[cand.priorities[itr]], cand)
				} else if itr > 0 && len(applicants[cand.priorities[itr-1]]) < n {
					status[i] = true
					applicants[cand.priorities[itr-1]] = append(applicants[cand.priorities[itr-1]], cand)
				} else if itr > 1 && len(applicants[cand.priorities[itr-2]]) < n {
					status[i] = true
					applicants[cand.priorities[itr-2]] = append(applicants[cand.priorities[itr-2]], cand)
				}
			}
		}
	}

	admit(0)
	admit(1)
	admit(2)

	departments := []string{
		"Biotech",
		"Chemistry",
		"Engineering",
		"Mathematics",
		"Physics",
	}

	printAdmittedStudents := func(s string) {
		fmt.Println(s)
		sortCandidates(applicants[s])
		for _, cand := range applicants[s] {
			fmt.Printf("%s %s %.2f\n",
				cand.firstName,
				cand.lastName,
				cand.gpa,
			)
		}
		fmt.Println()
	}

	for _, s := range departments {
		printAdmittedStudents(s)
	}
}
