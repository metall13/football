package main

import (
	"fmt"
	"strings"
)

type resultsMach struct {
	team
	r int
}

type inputData struct {
	c1 string
	c2 string
	r1 string
	r2 string
}

type team string
type allTeams map[team]int
type list map[team][]resultsMach

var ALL = allTeams{}

func main() {
	newMach := "Спартак Динамо 3 1\nСпартак Зенит 1 1\nЗенит Динамо 1 3"
	asdasd := AddData(newMach)
	fmt.Println(asdasd)
}

func AddData(resultsNew string) string {

	//bytes, err := ioutil.ReadAll(os.Stdin)
	//if err != nil {
	//	log.Fatal(err)
	//}

	dataSlise := strings.Split(string(resultsNew), "\n")
	masData := make([]inputData, 0, 0)

	for _, val := range dataSlise {
		data := strings.Split(val, " ")
		id := inputData{data[0], data[1], data[2], data[3]}
		masData = append(masData, id)
	}

	listParse := parse(masData)
	outData := output(listParse)
	fmt.Println(outData)
	return outData
	//os.Exit(0)
}

func parse(data []inputData) map[team]map[team]resultsMach {
	var lists = map[team]map[team]resultsMach{}
	for _, j := range data {
		ALL[team(j.c1)]++
		ALL[team(j.c2)]++
		if _, ok := lists[team(j.c1)]; !ok {
			lists[team(j.c1)] = map[team]resultsMach{team(j.c2): resultsMach{team(j.c2), 0}}
		} else {
			lists[team(j.c1)][team(j.c2)] = resultsMach{team(j.c2), 0}
		}

		if _, ok := lists[team(j.c2)]; !ok {
			lists[team(j.c2)] = map[team]resultsMach{team(j.c1): resultsMach{team(j.c1), 0}}
		} else {
			lists[team(j.c2)][team(j.c1)] = resultsMach{team(j.c1), 0}
		}
	}

	return lists
}

func output(l map[team]map[team]resultsMach) string {
	var scoreboard string

	scoreboard += fmt.Sprintf("%v \n", ALL)

	for k, c := range l {
		var res string
		for r, _ := range ALL {
			res += fmt.Sprintf("%5v", c[r].r)
		}
		scoreboard += fmt.Sprintf("%10v %5v\n", k, res)
	}

	return scoreboard
}
