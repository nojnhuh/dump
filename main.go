package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
)

func print_frequency_map(frequency_map map[int][]string) {
	for key, element := range frequency_map {
		fmt.Println(key)
		for _, word := range element {
			fmt.Println(word)
		}
	}
}

type WordCount struct {
	Word  string `json:"Word"`
	Count int    `json:"Count"`
}

func word_json(lorem io.Reader) string {
	// url := "loripsum.net/api"

	newScanner := bufio.NewScanner(lorem)

	newScanner.Split(bufio.ScanWords)

	counts := map[string]int{}

	for newScanner.Scan() {
		counts[newScanner.Text()]++
	}

	var wordCountArray []WordCount
	for key, value := range counts {
		wordCountArray = append(wordCountArray, WordCount{key, value})
	}

	sort.Slice(wordCountArray, func(p, q int) bool {
		if wordCountArray[p].Count != wordCountArray[q].Count {
			return wordCountArray[p].Count > wordCountArray[q].Count
		}
		return wordCountArray[p].Word > wordCountArray[q].Word
	})

	j, err := json.Marshal(wordCountArray)
	if err != nil {
		panic(err)
	}
	return string(j)

	/*var s []string

	frequency_map := map[int][]string{}

	frequencies := map[int]bool{}

	for key, element := range counts {
		frequency_map[element] = append(frequency_map[element], key)
		frequencies[element] = true
	}

	frequency_array := make([]int, 0, len(frequencies))

	for key := range frequencies {
		frequency_array = append(frequency_array, key)
	}

	sort.Slice(frequency_array, func(p, q int) bool { // sorts by value in array, ascending to descending
		return frequency_array[p] > frequency_array[q]
	})

	little_map := map[string]int{}

	for _, key := range frequency_array {
		fmt.Println("zoom")
		fmt.Println(frequency_array[key])
		for _, element := range frequency_map[key] {
			little_map[element] = key
			j, err := json.Marshal(little_map)
			if err != nil {
				panic(err)
			}
			delete(little_map, element)
			s = append(s, string(j)+",")
		}
	}
	s[len(s)-1] = s[len(s)-1][:len(s[len(s)-1])-1]
	return s*/
}

func main() {
	lorem := `<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Qui potest igitur habitare in beata vita summi mali metus? Praeclare Laelius, et recte sofñw, illudque vere: O Publi, o gurges, Galloni! es homo miser, inquit. Nam si propter voluptatem, quae est ista laus, quae possit e macello peti? Quae diligentissime contra Aristonem dicuntur a Chryippo. In quibus doctissimi illi veteres inesse quiddam caeleste et divinum putaverunt. Bork Quaeque de virtutibus dicta sunt, quem ad modum eae semper voluptatibus inhaererent, eadem de amicitia dicenda sunt. Summus dolor plures dies manere non potest? </p>

<p>Duo Reges: constructio interrete. Quid ergo aliud intellegetur nisi uti ne quae pars naturae neglegatur? At Zeno eum non beatum modo, sed etiam divitem dicere ausus est. Si longus, levis; Polemoni et iam ante Aristoteli ea prima visa sunt, quae paulo ante dixi. Hoc est non modo cor non habere, sed ne palatum quidem. Quid, si etiam iucunda memoria est praeteritorum malorum? </p>

<p>Id enim volumus, id contendimus, ut officii fructus sit ipsum officium. Quo modo? Expressa vero in iis aetatibus, quae iam confirmatae sunt. Quodsi vultum tibi, si incessum fingeres, quo gravior viderere, non esses tui similis; Esse enim, nisi eris, non potes. </p>`
	myReader := strings.NewReader(lorem)
	s := word_json(myReader)
	fmt.Println(s)
}
