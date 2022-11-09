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
	scores              []float64 // physics, chemistry, math, cs
	priorities          []string
	admitted            bool
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
			if len(data) != 9 {
				log.Fatal(errors.New("fields not match"))
			}
			scores := []float64{}
			for _, raw := range data[2:6] {
				score, _ := strconv.ParseFloat(raw, 64)
				scores = append(scores, score)
			}

			candidates = append(candidates, candidate{
				firstName:  data[0],
				lastName:   data[1],
				scores:     scores,
				priorities: data[6:],
			})
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	sortCandidates := func(c []candidate, course int) {
		sort.Slice(c, func(i, j int) bool {
			if c[i].scores[course] != c[j].scores[course] {
				return c[i].scores[course] > c[j].scores[course]
			}
			return (c[i].firstName + c[i].lastName) < (c[j].firstName + c[j].lastName)
		})
	}

	// fmt.Println(candidates)
	applicants := map[string][]candidate{
		"Mathematics": {}, // 2
		"Physics":     {}, // 0
		"Biotech":     {}, // 1
		"Chemistry":   {}, // 1
		"Engineering": {}, // 3
	}

	mapToCourse := []int{1, 1, 3, 2, 0}
	departments := []string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"}

	admit := func(itr int) {
		for i := 0; i < 5; i++ {
			courseName := departments[i]
			sortCandidates(candidates, mapToCourse[i])

			for j := range candidates {
				if (!candidates[j].admitted && len(applicants[courseName]) < n) &&
					(candidates[j].priorities[itr] == courseName ||
						itr > 0 && candidates[j].priorities[itr-1] == courseName ||
						itr > 1 && candidates[j].priorities[itr-2] == courseName) {

					candidates[j].admitted = true
					applicants[courseName] = append(applicants[courseName], candidates[j])
				}
			}
		}
	}

	admit(0)
	admit(1)
	admit(2)

	printAdmittedStudents := func(i int, s string) {
		fmt.Println(s)
		sortCandidates(applicants[s], mapToCourse[i])

		for j := range applicants[s] {
			fmt.Printf("%s %s %.2f\n",
				applicants[s][j].firstName,
				applicants[s][j].lastName,
				applicants[s][j].scores[mapToCourse[i]],
			)
		}
		fmt.Println()
	}

	for i, s := range departments {
		printAdmittedStudents(i, s)
	}
}
