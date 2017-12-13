package adventofcode2017

import (
	"strings"
	"regexp"
	"strconv"
)

func Day7Part1(input string) string {
	towers := createTowers(input)

	result := make(map[string]struct{}, len(towers))
	for _, v := range towers {
		result[v.name] = struct{}{}
	}

	// delete everything that is above something else
	for _, t := range towers {
		for _, v := range t.above {
			delete(result, v.name)
		}
	}

	if len(result) != 1 {
		panic("towers length is not 1: " + strconv.Itoa(len(towers)))
	}

	var t string
	for v := range result {
		t = v
		break
	}

	return t
}

func Day7Part2(input string) int {
	towers := createTowers(input)

	result := make(map[string]*tower, len(towers))
	for _, v := range towers {
		result[v.name] = v
	}

	// delete everything that is above something else
	for _, t := range towers {
		for _, v := range t.above {
			delete(result, v.name)
		}
	}

	if len(result) != 1 {
		panic("towers length is not 1: " + strconv.Itoa(len(towers)))
	}

	var t *tower
	for _, v := range result {
		t = v
		break
	}

	// we are using a recursive algorithm to find the highest
	// (as in highest up in the tree) unbalanced child,
	// because an imbalance might be caused by an unbalanced
	// child
	var highestUnbalanced map[int][]*tower

	var find func(t *tower)

	find = func(t *tower) {
		if len(t.above) < 1 {
			return
		}

		childWeights := make(map[int][]*tower, len(t.above))

		for _, v := range t.above {
			weight := getWeight(v)

			childWeights[weight] = append(childWeights[weight], v)
		}

		if len(childWeights) != 1 {
			highestUnbalanced = childWeights
		}

		for _, v := range t.above {
			find(v)
		}
	}

	find(t)

	var correct int
	var incorrect int
	var incorrectTower *tower

	for weight, towers := range highestUnbalanced {
		if len(towers) == 1 {
			incorrect = weight
			incorrectTower = towers[0]
		} else {
			correct = weight
		}
	}

	diff := correct - incorrect
	return incorrectTower.weight + diff
}

type towerSpec struct {
	name   string
	weight int
	above  []string
}

type tower struct {
	name   string
	weight int
	above  []*tower
}

func getWeight(t *tower) int {
	sum := t.weight

	for _, v := range t.above {
		sum += getWeight(v)
	}

	return sum
}

func createTowers(input string) map[string]*tower {
	r := regexp.MustCompile("^(\\w+)\\s*\\((\\d+)\\)(\\s*->\\s*(.*))?$")

	lines := strings.Split(input, "\n")
	towerSpecs := make(map[string]towerSpec, len(lines))
	// first get the specs from the lines
	for _, line := range lines {
		result := r.FindStringSubmatch(line)

		name := result[1]
		weight, err := strconv.Atoi(result[2])
		if err != nil {
			panic(err)
		}

		var above []string
		if result[4] != "" {
			above = strings.Split(result[4], ",")
			for i, v := range above {
				above[i] = strings.TrimSpace(v)
			}
		}

		towerSpecs[name] = towerSpec{
			name:   name,
			weight: weight,
			above:  above,
		}
	}

	// now create the towers in a tree structure (although inverted)
	towers := make(map[string]*tower)
	for _, spec := range towerSpecs {
		createTower(towerSpecs, towers, spec.name)
	}

	return towers
}

func createTower(towerSpecs map[string]towerSpec, towers map[string]*tower, name string) *tower {
	if v, ok := towers[name]; ok {
		return v
	}

	v, ok := towerSpecs[name]
	if !ok {
		panic("no tower spec found for " + name)
	}

	above := make([]*tower, len(v.above))
	for i, v := range v.above {
		above[i] = createTower(towerSpecs, towers, v)
	}

	t := &tower{
		name:   v.name,
		weight: v.weight,
		above:  above,
	}

	towers[v.name] = t

	return t
}
