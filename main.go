package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	// input is a lot of words
	// count how much each word appears
	// output that in JSON
	// [{ "the": 500 }, ...]

	// url := "loripsum.net/api"
	//
	// resp, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }

	lorem := `<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Qui potest igitur habitare in beata vita summi mali metus? Praeclare Laelius, et recte sofñw, illudque vere: O Publi, o gurges, Galloni! es homo miser, inquit. Nam si propter voluptatem, quae est ista laus, quae possit e macello peti? Quae diligentissime contra Aristonem dicuntur a Chryippo. In quibus doctissimi illi veteres inesse quiddam caeleste et divinum putaverunt. Bork Quaeque de virtutibus dicta sunt, quem ad modum eae semper voluptatibus inhaererent, eadem de amicitia dicenda sunt. Summus dolor plures dies manere non potest? </p>

<p>Duo Reges: constructio interrete. Quid ergo aliud intellegetur nisi uti ne quae pars naturae neglegatur? At Zeno eum non beatum modo, sed etiam divitem dicere ausus est. Si longus, levis; Polemoni et iam ante Aristoteli ea prima visa sunt, quae paulo ante dixi. Hoc est non modo cor non habere, sed ne palatum quidem. Quid, si etiam iucunda memoria est praeteritorum malorum? </p>

<p>Id enim volumus, id contendimus, ut officii fructus sit ipsum officium. Quo modo? Expressa vero in iis aetatibus, quae iam confirmatae sunt. Quodsi vultum tibi, si incessum fingeres, quo gravior viderere, non esses tui similis; Esse enim, nisi eris, non potes. </p>`

	lines := strings.Split(lorem, "\n")

	// fmt.Println(lines)
	// fmt.Println(len(lines))

	counts := map[string]int{}

	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			counts[word]++
		}
	}

	j, err := json.Marshal(counts)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(j))
}
