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
	scores              map[string]float64
	evaluation          map[string]float64
	priorities          []string
	admitted            bool
}

func main() {
	var n int
	fmt.Scan(&n)

	candidates := []candidate{}

	{
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
				firstName: data[0],
				lastName:  data[1],
				scores: map[string]float64{
					"Physics":          scores[0],
					"Chemistry":        scores[1],
					"Mathematics":      scores[2],
					"Computer Science": scores[3],
				},
				evaluation: map[string]float64{
					"Physics":     (scores[0] + scores[2]) / 2,
					"Chemistry":   scores[1],
					"Mathematics": scores[2],
					"Engineering": (scores[2] + scores[3]) / 2,
					"Biotech":     (scores[0] + scores[1]) / 2,
				},
				priorities: data[6:],
			})
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	sortCandidates := func(c []candidate, dep string) {
		sort.Slice(c, func(i, j int) bool {
			if c[i].evaluation[dep] != c[j].evaluation[dep] {
				return c[i].evaluation[dep] > c[j].evaluation[dep]
			}
			return (c[i].firstName + c[i].lastName) < (c[j].firstName + c[j].lastName)
		})
	}

	// fmt.Println(candidates)
	applicants := map[string][]candidate{
		"Biotech":     {},
		"Chemistry":   {},
		"Engineering": {},
		"Mathematics": {},
		"Physics":     {},
	}

	departments := []string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"}
	admit := func(itr int) {
		for _, dep := range departments {
			sortCandidates(candidates, dep)

			for j := range candidates {
				if (!candidates[j].admitted && len(applicants[dep]) < n) &&
					(candidates[j].priorities[itr] == dep ||
						itr > 0 && candidates[j].priorities[itr-1] == dep ||
						itr > 1 && candidates[j].priorities[itr-2] == dep) {

					candidates[j].admitted = true
					applicants[dep] = append(applicants[dep], candidates[j])
				}
			}
		}
	}

	admit(0)
	admit(1)
	admit(2)

	for _, dep := range departments {
		fmt.Println(strings.ToLower(dep))
		file, err := os.Create(fmt.Sprintf("%s.txt", strings.ToLower(dep)))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		sortCandidates(applicants[dep], dep)
		
		for i := range applicants[dep] {
			line := fmt.Sprintf("%s %s %.1f\n",
				applicants[dep][i].firstName,
				applicants[dep][i].lastName,
				applicants[dep][i].evaluation[dep],
			)
			fmt.Println(line)
			fmt.Fprintln(file, line)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
