package day15

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Lens struct {
	label       string
	focalLength int
}

func LensLibrary(inputFilePath string) (int, int) {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sol1 := 0

	var steps []string

	for scanner.Scan() {
		currentLine := scanner.Text()
		steps = strings.Split(currentLine, ",")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, s := range steps {
		sol1 += stepHash(s)
	}

	out := boxToHashMap(steps)
	sol2 := totalMapValue(out)

	return sol1, sol2

}

func stepHash(input string) int {

	out := 0

	for _, c := range input {
		out += int(c)
		out *= 17
		out %= 256
	}

	return out

}

func boxToHashMap(steps []string) [][]Lens {

	var hashMap [256][]Lens
	fLengthRegex := regexp.MustCompile(`([a-z]+)(=([0-9]+)|-)`)

	for _, s := range steps {

		matches := fLengthRegex.FindStringSubmatch(s)
		label := matches[1]
		idx := stepHash(label)
		currentList := hashMap[idx]

		// remove
		if matches[2] == "-" {
			for i, l := range currentList {
				// remove the value
				if l.label == label {
					currentList = append(currentList[0:i], currentList[i+1:]...)
					hashMap[idx] = currentList
				}
			}
			continue
		}

		focalLength, err := strconv.Atoi(matches[3])
		if err != nil {
			log.Fatal(err)
		}

		// create a new lens
		lens := Lens{label: label, focalLength: focalLength}
		// add
		found := false
		for i, l := range currentList {
			// replace the value
			if l.label == label {
				currentList[i] = lens
				found = true
				break
			}
		}
		if !found {
			currentList = append(currentList, lens)
		}
		// update the list
		hashMap[idx] = currentList
	}

	return hashMap[0:]

}

func totalMapValue(hashMap [][]Lens) int {

	tot := 0
	for box, slots := range hashMap {

		for slotIdx, lens := range slots {
			tot += (box + 1) * (slotIdx + 1) * lens.focalLength
		}

	}

	return tot

}
