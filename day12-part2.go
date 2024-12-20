//go:build ignore

package main

import (
	"fmt"
	"slices"
	"strings"
)

type direction int

const (
	top direction = iota
	right
	bottom
	left
)

type perimeterEdge struct {
	x1, y1 int
	x2, y2 int
	dir    direction
}

type plot struct {
	plantType  rune
	x, y       int
	perimeters []perimeterEdge
}

type plotMarker struct {
	plantType rune
	x, y      int
}

var marker = map[plotMarker]bool{}

func calcPerimeter(garden [][]rune, mx, my int, p plot) []perimeterEdge {
	edges := []perimeterEdge{}
	if p.x-1 < 0 || garden[p.y][p.x-1] != p.plantType {
		edges = append(edges, perimeterEdge{x1: p.x, y1: p.y, x2: p.x, y2: p.y + 1, dir: left})
	}

	if p.y-1 < 0 || garden[p.y-1][p.x] != p.plantType {
		edges = append(edges, perimeterEdge{x1: p.x, y1: p.y, x2: p.x + 1, y2: p.y, dir: top})
	}

	if p.x+1 >= mx || garden[p.y][p.x+1] != p.plantType {
		edges = append(edges, perimeterEdge{x1: p.x + 1, y1: p.y, x2: p.x + 1, y2: p.y + 1, dir: right})
	}

	if p.y+1 >= my || garden[p.y+1][p.x] != p.plantType {
		edges = append(edges, perimeterEdge{x1: p.x, y1: p.y + 1, x2: p.x + 1, y2: p.y + 1, dir: bottom})
	}

	return edges
}

func findPlotInSameRegion(garden [][]rune, mx, my int, plantType rune, x, y int) []plot {
	var plots []plot
	pt := plot{plantType: plantType, x: x, y: y}
	plotMark := plotMarker{plantType: plantType, x: x, y: y}
	marker[plotMark] = true
	// calculate perimeter
	pt.perimeters = calcPerimeter(garden, mx, my, pt)

	plots = append(plots, pt)

	if y-1 >= 0 && garden[y-1][x] == plantType && !marker[plotMarker{plantType: plantType, x: x, y: y - 1}] {
		plots = append(plots, findPlotInSameRegion(garden, mx, my, plantType, x, y-1)...)
	}

	if x+1 < mx && garden[y][x+1] == plantType && !marker[plotMarker{plantType: plantType, x: x + 1, y: y}] {
		plots = append(plots, findPlotInSameRegion(garden, mx, my, plantType, x+1, y)...)
	}

	if y+1 < my && garden[y+1][x] == plantType && !marker[plotMarker{plantType: plantType, x: x, y: y + 1}] {
		plots = append(plots, findPlotInSameRegion(garden, mx, my, plantType, x, y+1)...)
	}

	if x-1 >= 0 && garden[y][x-1] == plantType && !marker[plotMarker{plantType: plantType, x: x - 1, y: y}] {
		plots = append(plots, findPlotInSameRegion(garden, mx, my, plantType, x-1, y)...)
	}

	return plots
}

func main() {
	lines := strings.Split(input, "\n")
	var garden [][]rune
	var regions [][]plot
	for _, line := range lines {
		garden = append(garden, []rune(line))
	}

	mx, my := len(garden[0]), len(garden)

	for y := range garden {
		for x := range garden[y] {
			if !marker[plotMarker{plantType: garden[y][x], x: x, y: y}] {
				regions = append(regions, findPlotInSameRegion(garden, mx, my, garden[y][x], x, y))
			}
		}
	}

	var price int
	for _, r := range regions {
		edges := make(map[direction]map[int][]int)
		for _, p := range r {
			for _, e := range p.perimeters {
				if _, ok := edges[e.dir]; !ok {
					edges[e.dir] = make(map[int][]int)
				}
				switch e.dir {
				case top:
					edges[top][e.y1] = append(edges[top][e.y1], e.x1)
				case right:
					edges[right][e.x1] = append(edges[right][e.x1], e.y1)
				case bottom:
					edges[bottom][e.y1] = append(edges[bottom][e.y1], e.x1)
				case left:
					edges[left][e.x1] = append(edges[left][e.x1], e.y1)
				}
			}
		}

		for _, es := range edges {
			for _, e := range es {
				slices.Sort(e)
			}
		}

		var perimeterCount int
		for _, es := range edges {
			for _, e := range es {
				j := 1
				perimeterCount++
				for j < len(e) {
					if e[j]-e[j-1] > 1 {
						perimeterCount++
					}
					j++
				}
			}
		}

		price += len(r) * perimeterCount
	}
	fmt.Println(price)
}

var ex1 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

var ex2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

var ex3 = `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

var ex4 = `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

var input = `GGGGGDDDDSSSEEEEEEEEEEEEEEEUFFOOOOOOOYYKKKKXKKKKKKKKKKKKKNVVVVVVVVVVVVVVVVVVVYYEEZNNNNNNNNNNBNSSNNNWUUUUUUKKKKKKKKKKKKKKKKKKKKHHHHHHHHHHHHHH
GGDDDDDDDSSSSEEEEEEEEEEEEEEFFFOOOOOOOYYKKKKKKKKKKKKKKKKKNNNVVVVVVVVVVVVVVVVEYYEEENNNNNNNNNNNNNSNNNWWWWWUUIITTKKKKKKKKKKKKKKKKHHHHHHHHHHHHHHH
GGGDDDDDDSSSEEEEEEEEEEEEEFFFFFOKOOOOOOOOKOKKKKKKKKKKKKKKKNNVVVVVVFVVVVVVVVVEEEEEENENNNNNNNNNNNNNNNWWWWWUUTTTTTTKKKKKKKKKKKKKKKHHHHHHHHHHHHHH
GGGGDDDDSSSSESEEEEEEEEEEEEEFFFOOOOOOOOOOKKKKKKKKKKKKKKKKNNNVVVVVVFFFVVVVEEEEEEEEEEEENNNNNNNNNNNNNWWWWWUUUUTTTTTKKKKKKKKKKKKKKKKHHHHHHHHHHHHH
GGGDDDDDTSSSSSSESSEEEEEEEEEFFFFOOOOOOOOOKKKKKKKKKKKKKKKKNNNVVVVVVFFFVVVVEEEEEEEEEEBBBNNNNNNNNNNNWWWWUUUUUTTTTTTTTKKKKKKKKKKKKKKHNHHHHHHHHHHH
GGGBDDDDTSSSSSSSSSSEEEEECFFFFFONOOOOOOOOOOKKKKKKKKKKKKNNNNVVVVVVVFFFVVVVEEEEEEEEEEBBBBBBNNNNNNNNWWWXUUUUUTTTGTTTTTTTKKKKKKKKKKKHNHHHHHHHHHHH
GBBBDDDDSSSSSSSSSSSJEEELFFFFFFOOOOOOOOOOOOKKKSKKKKKKKKKNNNNVVVVVVVVVVVVEEEEEEEEEEBBBBBBNNNNNNNWWWWWFFFFFFFFFFFFFFFTKKKKKKKKKKKHHNHHHHHHHHHHH
GGBBBBBUUUUUUUSSSSSSSLLLFFLFFFOOOOOOOOOOOOKKKKKKKKKKKNNNNNNMMMMVVVVVVVVEEEEEEEEEEBBBBBBBNNNNNNWWWXXFFFFFFFFFFFFFFFFFFFFFFFFKKKHHHHHHHHHHHHHN
BBBBBUUUUUUUUUUSSSSSLLLLLLLFFLOOOOOOOOOOOOKRKKRRKKKKKKKNMMMMMVVVVVVVVVVVEEVVEEEEEBBBBBBBNNNNNXXXXXXFFFFFFFFFFFFFFFFFFFFFFFFKKKHHHZZVHHHHXHHN
BBBBBUUUUUUUUUUAAAAALLLAAAAAAALOOOOOOOOOOOORRRRRRKKKKKQMMMMMMMVMMVGGVVVVVVVVEEEEEBBBBBBNNNUNNXXXXJXFFFFFFFFFFFFFFFFFFFFFFFFKKKHHHZZVFFHHHHNN
BBBBBUUUUUUUUUUAAAAAAAAAAAAAAAZOOOOOOOAOOOOAARRRRRKKRRRMMMMMMMMMMGGGGGVVVVVVEEUEEBBBUUNNNNUUXXXXXXXFFFFFFFFFFFFFFFFFFFFFFFFKKKKKZZVVVVVEHVNN
BBBBBUUUUUUUUUUUAAAAAAAAAAAAAAJJJJOOAAAAAORAARRRRRRRRRMMMMMMMMMMMGGGGGVGGGGEREEWESSUUUUUUUUXXXXXXXXFFFFFFFFFFFFFFFFFFFFFFFFKKKKKKZZVVVVVVVNN
BBBBBUUUUUUUUUUUAAAAAAAAAAAAAAJJJOOOAAAAAAAAARRRRRRRRRMMMMMMMMMMMGGGGGGGGGEEEEWWWUUUUUUUUUUXXXXXXXXXXXXXVFFFFFFFFFFFFFFFFFFQQQKKKVVVVVVVVVVV
BBBBBUUUUUUUUUUUUUUAAAAAAAAAAAZZIDAAAAAAAAAAAAAARRRRRRRFMMMMMMMMMGGGGGGGGGGGFEWWWUUUUUUUUUUXXXXXXXXXXXXZZFFFFFFFFFFFFFFFFFFSQQKQQVVVVVVVVVVV
KKKBBUUUUUUUUUUUUUUAAAAAAAAAAAIIIDAAAAAAAAAAAAAAARRRRRFFMFMMMMMMMGGGGGGGGGEGGWWWWUUUUUUUUXXXXXXXXXXXXAZZZFFFFFFFFFFFFFFFFFFSQQKQQQQVVVVVVVVV
KKBBBUUUUUUUUUUUUUUAAAAAAAAAAAIIAAAAAAAAAAAAAAAAAARRRRFFFFFMFFMMNGGGGGGHEEEEWWWWUUUUUUUXXXXXXXXXXXXXXAZZZFFFFFFFFFKKKKSSCSSSQQQQVVVVVVVVVVVV
KKBBKKKKSUUUUUUUUUUAAAAAAAAZZZIIIIAAAAAAAAAAAAAAAARRRRFFFFFFFFFFGGGGGGGHEEEEEWWWUUUUTUUXXXXXXXXXXXXXXZZZZZZZZZOKKKKKSSSSSSSSSSVVVVVVVVVVVVVV
KKKKKKKKSUUUUUUUAAAAAAAAAAAZZZIIIIAAAAAAAAAAAAAAAARRRFFFFFFFFFFFFGGGGGGEEEEEEWTTYYTTTDUXXXXXXXXXXXXXXEZAAAAAAAAAKSSKSSSSSSSSSSXVVVVVVVVVVVVV
KKKKKKKKKUUUUUUUIEIAAAAAAAZZZZIIIIKIAAAAAAAAAAAAAARRRFFFFFFFFFFFRGGGGRREEEEEEWTTTTTTTTUXXXXXXXXXXXXXIIZAAAAAAAAAGSSSSSSSSSSSSSXXVVVVVVDVVVVV
KKKKKKKKKUUUUUUUIEIEZZZZZZZZZZZIIIIIAAAAAAAAAAAARRRRRFFEFFFFFFFFRGGGGRRRRRRRWWTTTTTTTGGGGGXXGXXXXGXAAAAAAAAAAAAASSSSSSSSSSSSSXXVVVVVVVDVVVVV
KKKKKKKKKUUUUUUUEEIYZZZZZZZZZZGWOIUAAAAAAAAABBAARRRRZTTFFFFFZFRRRRRRRRRRRRRRRTTTTTTTTTTGGGGGGXXXXGGAAAAAAAAAAAAAGSSSSSSSSSSSXXXVVVVVDVDDVVVV
KKKKKKKKKEEEEEEEEEIYYYYTTGWWWZWWIIUUAXAAAAAAYYYYRRRRZZZZZZZZZFRRTTTTTRRRRRRRRRTTTTTTTTGGGGGGXXXGGGGAAAAAAAAAAAAAGSSSSSSSSSSSXXXXXVVVDDDDDDVV
KKKKKKKKKEEEEEEEEEEYYYYTGGWHWZWWWIUHHXAAAAAAYYYYYRRZZZZZZZZZZZRRRTRRRRVRRRVVTTTTTTTTTTTGGGBBBBBGGGGAAAAAAAAAAAAAGGSSSSSSSSSXXXXXVVVDDDDDDDVV
KKKKKKKKKEEEEEEEEEEYGVGTGGHHHWWHIIUHHHYYYYAYYYYBBZZZZZZZZZZZZZZRRRRRRRVVVRVVAATTTTTTTTTTTBBBBBBBBGGAAAAAAAAYGGGGGGWGSPPPPPSXXXXXXVDXDDDDDDVV
KKKKKKKKKEEEEEEGGGGGGGGGGGHHHHWHHHHHHHHYYYAYYYYBZZZZZZZZZZZZZZRRRRRRRRVVVVVVAATTTTTTTTTTHBBBBBBYGGGAAAAAAAAGGGGGGGGGGPPPPPPPXXXXXXDDDDDDDDDV
KKKKKYYKKEEEEEGGGGGGGGGGGGHHHHHHHHHHHHHYYYYYYYZZZZZZZZZZZZZZZZZRRRRRRVVVVVVVVAATTTTTTTTHHBBBBBYYYYYAAAAAAAAGYGGGGGGGGPPPPPPPPXXXXXXDDDDDDDDD
KKKYYYYKKKEEEEGGGGGGGGGGGGHHHHHHHHHHHHPYYJYYYYYYZZZZZZZZZZZZZZZRRRRRRVVVVVVVVVAATTTTTTTTBBBBBUUYYYYYYYAAAAAYYGGGGGGPPPPPPPPPXXXXDDDDDDDDDDDM
KKFFYYYEEEETEGGGMMMMMMMGGGHHNNHHHHHHHHPPPJYJJJJJJGGZZZZZZZZSSSRRRRRVVVVVVVVVCCVTTTTTTTTTTUBUUUUYGYAYYYAAAAAPYGGGGGGPPPPPPOPPXXXXDDDDDDDDDDDD
YFFFYQYYYYEEEEGGMMMMMMMGGGGNNNNNNNNHHHPJJJJJJJJJJJGZZZZZZZZSSRRRRRRRRHVVVVVVVCVTTTJTTTTTUUUUUAAAAAAAAAAAAAAGGGGGPGPPPPPOOOPXXXXXDDDDDDDDDDDD
YYYYYYYYYYEEEYLLMMMMMMMGGGGNNNNNNNNHPPPJJJJJJJJJJJJZZZIZZZZRSRRRHHRHHHVVVVVVVVVTTVVTUUUTUUUUUAAAAAAAAAAYYYBGGGGGPPPPPPPOOOOOOXOOODDDDDDDDDDZ
YYYYYYYYYYYYYYYLMMMMMMMGGGGNNNNNNNNNPPPJJJJJJJJJJIIIIIIIZZRRSRRHHHHHHHHVHVVVIIVVVVVVUUUUUUUUUAAAAAAAAAAAYBBGGGGPPPPPPPOOOOOOOOOOOODDDDDDZZDZ
YYYYYYYYYYYYYYYGMMMMMMMGGGNNNNNNNNRSRPPPPJJJJJJJJIIIIIIIZRRRRRHBBHHHHHHHHHHVVVVVVVNUUUUUUUUUUAAAAAAAAAAAAAAGGGPPPPPPPPPPOOOOOOOOOOOODZZDZZZZ
YYYYYYYYYYYYYYYGMMMMMMMGGGNNNNNNNNRRRRRRJJJJJJJJJIIIIIIZPRPHHHHHHHHHHHHHHJJVVVVVVVNUUUUUUUUUOAAAAAAAAAAAAAAAGGPPPPPPPPOOOOOOOOOOOOOOOOZZZZZZ
YYYYYYYYYYYYYYYXMMMMMMMGGGGGNNNNNRRRRRRRRJJJJJJJJIIIIIIZPPPHHHHHHHHHHHHHHHHVVVVVVVNUUUUUUUUUOOJOOOAAAAAAAAAAGPPPPPPPPPOOOOOOOOOOOOOOOOZZZZZZ
YYYYYYYYYYYYYYYYMMMMMMMGGGNNNNNNNNRRRRRRRRRJJJJJJIIIIZZZZZPHHHHHHHHHHHHHHHHHVVVVVVNUUUUUUUUUOOOOOOAAAAAAAAAAAPPPPPPPPPPPPOOOOOOOOOOOOOOZSSZZ
YYYYYFYYYYYYYYYYMMMMMMMNGNNNNNNNNNRRRRRRRRRHJJJRJIIIIIZZZZZHHHHHHHHHHHHHHHHVVVVVVNNUUUUUOUOOOOOOOMAAAAAAAAAAPPPPPPPPPPPPPOOOOOOOOOSSOSSSSZZZ
YYYYYKPYTYYYYYYYMMMMMMMNNNNNNNNNNRRRRRRRRRRHRTJJJIIZZZZZZZZZHHHHHHHHHHHHHPPPVVVVVGNUUUUOOOOOOOOOOMMMOAAAAAAPPPPPPPPPPPPPOOOOOOOOOSSSSSSSSSSZ
YYYYKKKTTTYYYYYYMMMMMMMNNNNNNANAARRRRRRRRRRRRJJIIIIIZZZZZZZZHHHHHHHHHHHHZZZVVVVVGGNUUUUOOOOOOOOOMMOOOAPAHHHPPPPPPPPPPPPPPOOOOOOOOOSSSSSSSSSZ
YYYYYKKKKTKKKKKMMMMMMMMNNNNNNAAAARLRRRRRRRRRRRIIIIIIZZZZZZZZHMHHHHHHWHZZZZZVVVVVGUUUUUUTOOOOOOOOOOOOLOAAHHHHPPPPPPPPPPPPPOOOOOOOOOSSSSSSSSSZ
YYYYYYYKKKKKKKZZZMMMMNNNNNNNAAAAARLLRRRRRRRRRIIIIIIIZZZZIZZZZHHHHTHHZZZZVZZHHVVUUUUUUUZOOOOOOOOOOOOOOOCAAHHHHHPHPPPPPPPPPPPOSSOOOOOSBSSSSSJJ
YYYYYYYKKKKKKKKKKMMMMMMNNNNAAAAAAAARRRRRRRRRRRPPIIIIIZIIIZZZZZHLHTHHZZZZZZZHHUUUUUUHUUUHOOOOOOKKKSSSSSCAAHHHHHHHHPPPPPPPPPSSSSSSOOOOBSSSSSJJ
YYYYYYYKKKKKKKKMMMMMMMMNNNNAAAAAAAAAARRRRRRRRPPPPPIIIIIIZZZZZHHLZTZIZZZZZZZZZUUUUHHHUUUHHOOHOOKKKKKSSSSHHHHHHHHHPPPPPPPPPPPSNNNXNNNNNSSSSJJJ
YYYYYYYKKKKKKKKKKMMMMMNNNNNAAAAAAAAAARRRRRRRRPPPPIIIIIIZZZZZZZZZZTZZZZZZZZZUUUUUUUHHHHHHHHOHHKKKKKKSSSSSSHHHHHHHHPPPPPPPCCCNNNNNNNNNNWSSJJJJ
YRYYYYKKKKYYYKKKMMMMMMNNKKKKKKAAMMAAARRRRRRRPPPPPIIIIIIZZZZZZZZZZZZZZZZZZZTUUUUUUUUHHHHHHHHHHKKKKKDSSSSSSSHHHGGGHHHPPPPCCCCNNNNNNNNNNWWWJJJJ
RRRYYYYYYYYYYYYYMMMMMNNNKKKKKKKAAMMAARRRRRRRPPPPPPPIIOIZZZZZZZZZZZVVZZZZZZTTTUUUTTUPPHHHHHHHHKKKKKKFSSHHHSHHGGGGGHNPPPCCRRCNNNNNNNNNNWWWJJJC
RRRRYYYYYYYYYYYMMMMMHNNNVVKKKKKAHHMARRRRRRRRRPPPPFFOIOZZZZZZZZZVVVVVVZZZZZTTTUUUTTTPPPPHHHHHHKKKKKKKKKKHHHHHHGGGGGNNNPCCRRRRNNNNNNNNNWWWWJJC
RRRRRRRYYYYYYYOOEOOMHHNNNNHHNNNHHIMMIRRRRRRRPPPPYOOOOOOZOOOOOVVVVVVVVZLZZZTTTTTTTTPPPPPPHHHHHKKKKKKKKKKHHHHHHGGGGGGNGRRRRRRRRNNNNMNNWWWWWVVC
RRRRRRYYYYYYYOOOOOOMMHHNNTNNNNNHHIIIIRIIIPRRRRLOOFOOOOOOOOOOOVPVVVVVVVLZZTTTTTTTTTPPPPHHHKKKKKKKKKKKHHHHHHHHHGGGGGGGGGRRRRRRRNNNMMMMMWWWVVVV
GRRRRYYYYYYYYOOOOOOOOOHOITNNNNNNHHIIRRIIIPPRRROOOOOOOOOOOOOOOOPVLLLLLLLZZTTTTTTTTTTPPPPHHKKKKKKKKKKHHHHHHMHHHGGGGGGGGGRRRRRRRNNNMMMMMWMWVVVV
GRRRRRYYYYYVYOOOOOOOOOOOIINNNNNCCHIIIIIIIIPPPPOOOOOOOOOOOOOOOOPVVYLLLLLZTTTTTTTTTTTTPPGRHKKKKKKKKKKHHHHHMMMMMGGGGGGGGGRRRRRRRRNNMMMMMMMWWVVV
GRRRRRYVYYYVYBOOOOOOOOIIIIOONNNCCIIIIIIIPPPPPPPPOOOOOOOOOOOOOOOVYYLLLLLTTTTTTTTTTNNNRRGRHHRKKFKKBKKBHHHMMMMMMGGGGGGGGGGRRRRRRNNMMMMMMMMWVVVV
GGGGGGYVVVVVVVVOOOOOOOIIIIIINDNICCIIIIIIPPPPPPPPOOOOOOOOOOOOOOOOYLLLLLTTTTTTTTTTTNNRRRRRRRRRFFFBBBBBMMMMMMMMGGGGGGGGGGGRRRRRRRRRMMVVVMMMVVVV
GGGGGGYVVVVVVVVVOOOOOOIIIIIIZIIIICIIIIPPPPPPPPPPOOOOOOOOOOOOMMMLLLLLLLLLTTTTTTNNNNNNRRRRRRRRBFFFBBBBMMMMMMMGGGGGGGGGGGGGSRRRRRRRRCVVVVVMVVVV
GGGGGGGGVVVVVVVVOOOOOOIIIIIIIIIIIIIIIPPPPPPPPPPPOOOOOOOOOOOOOMMLLLLLLLLLLTTTENNNNNNNRRRRRRRRRFYFBBBBMMMMMMMMGGGGSSSAGGGKSSRCCCRRCCVVVVVVVVVV
GGGGGGGGVVVVVVVVOOOOOOIIIIIIIIIIIPPIIPPPPPPPPPOPPOOOROOOOOOOOOMMLLLLLLLELLTTENENNNNRRRRRRRFFBFFFFFFFMMMMMMMMMGGMISSSSSGSSSSJCCCCCVVVVVFFFVVV
GVGGGGGRRVVVVLVOOOOOOOIIIIIIIIIIIAPPPPOOOOOPOOOPPPOOOOOOOMMMMMMLLLLLLLLEEEEEEEENNNNNIIIRRRFFFFFFFFFMMMMMMMMMMGMMMSSSSSSSSSSSVCCCCVVVVVFFFFFF
VVGGGGRRRVVVVLVOOOOOOOOIIIIIIIIIIAPAAAAOOOOOOPOOOOOOOOOOOSSMMMMMMLLLLLLEEEEEENNNNNNNIIIRRRFFFFFFFFFMMMMMPPMMMMMMSSSSSSSSSSSSCCCCVVVVVVFFFFFF
VVRRRRRRRVVVLLKLOOOLLLOIIMMIIIAAIAPAAAAOOOOOOOOOOOOOOOHHOSSSSMSMLLLLLLLLEEEEENNNNNNNIIIRRRFFFFFFFFFMFMMMPMMMMMMMMMSSSSSSSSKKCCCCCCCVVVFFFFFF
VVVVVRRRRVVVLLLLLLLLLLOOOIIIAAAAAAAAAAAAOOOOOOOOOOOOOOOSSSSSSSSMLLLLLLEEEEDNNNNNNNNNIIIRRGFFFFFFFFFMFMMHPMMPPPPMMMSSSSSSSSKCCCCVCCVVVVFFFFFF
VVVVVRRRRVVLLLLLLLULLLOOOOOIAAAAAAAAAAABOOOOOOOOOOOOOOOSSSSSSSSMMLSSSLEEDEDLNDNNNNNNIIIGGGFFFFFFFFFMFMMPPPMMPPPPSSSSSSSSZKKKKVVVVVVVFFFFFFFF
VVVVVVVRRVVLLLLLLLLLLLLOOOOOAAAAAAAAAAAAAAOOOOOOOOOOOOOSSSSSSSSSMLLSSEEYDDDDDDDNNNNNIIIGUFFFFFFFFFFFFMKPPPPPPPPSSSSZSZZZZKKKVVVVVVVFFFFFFFFF
VVVVVVVVRVLLLLLLLLLLLLLOOOOOOAAAAAAAAAAAAAAOOOOOOOOOOOSSSSSSSSSLLLSSEEEEDDDDDDIIIIIIIIIUUFFFFFFFFFFFFMKPPPPPPPPPSSSZSZZZZZKVVVVVVVVVFFFFFFFF
VVVVVVVVVVVLLLLLLLLLLLLLLOOOAAAAAAAAAAAAAOOOOOOOOOOOOOSSSSSSSSSSSSSSSESSSDDDDDIIIIIIIIIUUFFFFFFFFFFFFPPPPPPPPPPSSSSZZZZZZZZVVVVVVVVFFFFFFFFF
IIVVVVVVVVVVLLLJLLLLLLLOLOOUAAAAAAAAAAAAAOOOOOOOOOOOSSSSSSSSSSSSSSSSSSSSSDDDDYIIIIIIIUUUFFFFFFLFFFSFFVVPPPPPPPPPPSSZZZZZZZZIVIVVIVVIFFFFFFFF
IIIVHHHHHHVMMMLMSLLLLOOOOOUUAAAAAAAAAAAAAOOOOOOOEEEEEEESSSSSSSSSSMSSSSSSSDDDYYIIIIIIIGUUIFFPPFFFFFFFVVLLLLPPPPPPPPZZZZZZZZIIIIIIIIIIIIFFFFFF
IVVVVHHHHHMMMMMMMMMLLOOUUOUAAAUAAAAAAAAAALNOOOOOEEEEEEEEEEEEESSSSSSSSSSSSYDYYYYYYYYGGHIIIIFYYYFPFFFVVVLLLLPQPPPPPZZZZZZZZZIIQIIIIIIIIIFFFFFF
PVVVHHHHHHBMMMMMMMMMOOUUUUUUUUUAUAAAAAFANLNNNOOOEEEEEEEEEEEEESSSSSSSSSYYYYYYYYYYYYGGGHHIIIIYYYFVVSVVVVLLLLPPPPPPPPZZZZZZZZQQQQIIILIIIIIFFFFK
PTVHHHHHHHBMMMMMMMMMOOOUUUUUUUUUUAUAAAANNNNNOOOOEEEEEEEEEEEEESSSSSSSSSSYYYYYYYYYYGGGGHHHHIIYYYYYVVVVVVLLLLVPEPPPZZZZZZZZZQQQQQQLLLLIXFFFFFFF
TTHHHHHHHHMMMMMMMMMMMOUUUUUUUUUUUUUAAANNNNNNNOOAEEEEEEEEEEEEESSSSSSSSYYYYYYYYYYYYYGYYHHHHHIYYYYYYVVVVVLLLLLVVPPPZZZZZZZZZVVVVVVVVVVLXXFFFFFF
DTTTHHHHHHMZMMMMMMMMMMUUUUUUUUUUUUAAAAANNNNNNNAAEEEEEEEEEEEEESSSSSSSSSYYYYYYYYYYYYYYYYHHHHISYYYYVVVVVVLLLLLVPPAAZZZZZZZZZVVVVVVVVVVLXXXFFFFF
DTTTTNHHHZZZMMMMMMMMMMUUUUUUUUUUUUUQANANNNNNNNAAAGUUAEEEEEEEESSSSSSYYYYYYYYYYYYYYYYYYYHHHSSSSYYVVVVVVVLLLLLVVVAZZZZZZZZZZVVVVVVVVVVVVVXXXXFF
DTTTTTHHHHZZMMMMMMMIIIUFUUUJWUUUDDXXNNNNNNNNNNNAGGUCAAAACCCCCSASSSSSSYYYYYYYYYYYYYYYYOOHKKSSSYVVVVVVVVLLLLVVVVVZZEEZZZZZZVVVVVVVVVVVVVXXXXFF
DDTTTTHHHHHHWMMMMMIIIIIUUUUJWWWDDXXXXXNNNNNNNNNGGUUCACACCCCAAAAAASSSSYYYYYYYYYYYYYPYOOOHSSSSSSSWVVVVGVVVVVVVVVOOZEZZZZZZZVVVVVVVVVVVVVXXXXXX
DTTTTTHHQHHHMMBMMMIIIIIIIIIWWWWWDXXXXXNNNNNNNNNUUUUCCCCCCCCAAAAAAASSOOYYYYYYYYYYYYYYOOOOOSSSSSSSBVBBVVVVVVOVVYOOOZZZVVVVVVVVVVVVVVVVVVXXXXXX
DTTTDHHHQHHHBBBBMBIIIIIIIWWBWWWFXXXXXXXXNNNNNUKKUUUCCCCCCCCCCCAAAASOOVYGGGYYYYYYYYYYYOOOSSSSSSSSBBBBBBVVVVOOOOEOOFZZVVVVVVVVVVVVVVVVVVXXXXXK
DDTDDDJHJJBBBBBBBBIIIIIIWYWWWWWWEXXXXXXXNQUUUUUUUUUUCCCCCCCAAAAAAOOOOOOGGGYYYYYYYOOOOOOOSSSSSSSSBBBBBVVVVVOOOOOOOOOZVVVVVVVVVVVVVVVVVVXXXXXK
DDDDJDJJJBBBBBBBBBBBIIIWWYWWWWWWEXXXXXXDNNUUUUUUUUUUUUUCCCCAAAAAAAAOOBOGGGYYJYYYYYNOOOOOEESSSSSBBBOOOOOVOOOOOOOOOOORVVVVVVVVVVVVVVVVVVXXXXXX
DDDDJDJJJBBZZBBBBBBBBBIWWWWWWWWWXXXXXXXXNNUUUUUUUUUUUUCCCCCAAAAAAAOOOOOGGGGJJYYYYYNONNOOOESEBBBBBBOOOOOOOOOOOOOOOOORVVVVVVVVVVVVVVVVVVXXXXVV
DDDDJJJJJBBBBBBBBBBBBYWWWWWWWWWWKXXXXXXXNNXXUUUUUUUUUUCCCCCAAAAAAAOOOOGGGGGCJJYYYNNOONOOEEEEBBBBBBBBOOOOOOOOOOOOORRRVVVVVVVVVVVVVVVVVVTXXXVV
DDDJJJJJJHBBBBBBBBBBYYQQQWWWWWWKKKXXXXXXXXXXUUUUUUUUUUUCECCAAAAAAAAGOOGGGGCCCJJNNNNNNNNNNNNNNBBBBBBBXXXOOOOOOOOOORRRVVVVVVVVVVVVVXXXXXTLXXVV
DDDDDJJJJHBBBBBBBBBBBQQHHWWWWWWWKXXXKKXXXUUUUUUUUUUUUUEEECAAAAAAAAGGGGGGGGCCCCNNNNNNNNNNNNNNBBBBBBBBXXXXXXOOOOJOORRRVVVVVVVVVVVVVXTXXXTLXXVV
DDDDDJJJJBBBBBBBBBBBQQQNWWWNWWWKKKXKKKKKXUUUUUUUUUUUUUUEEAAAAAAAAAGGGGGGGGCCCCCNNNNNNNNNNNNNBBBBBBBBBXXXXXXOXCXRRRRRVVVVVVVVVVVVVTTTTTTTVVVV
DDDDVVJJXXBBBBBBMBBBQQQNNWWNWWWWKKKKKKKKXLUUUUUUUUUUUUUEEAAAAAAAAAAGTGNCCGCCCCCCCNNNCCVXNXBBBBBXBBBXXXXXXXXXXXXXRRRRVVVVVVVVVVTTTTTTTTTMEEVV
VDVDVVVXXXBBBBBBBBBQQQQNNNNNNWWWWKKKKKKKXXVVVVUUUUUUUUUEEEEEEEEAAAATTGNNCCCCCCCCCCCCCCCXXXXBXXXXXXXXXXXXXXXXXXXXRRRRRRBBBBTTTTTTTTTTTTTMEEVV
VVVVVVVXXBBBNBBBQQBBQQQNNNNNNNNNVKVKKKKVXXVVVVUUOUUUEUEEEEEEDEEEWCCCTGNNNCCCCCCCCCCCCCCXXXXXXXXXXXXXXXXXXXXXXXRXRRRBBBBBBBBTTTTTTTTTTTTEEEEE
VVVVVVVVXXQBNBQQQQQQQQQQNNNNNNNNVVVKKKKVVVVJIIIIIIIIIIIIEEEEEEEECCCCCNNNCCCCCCCCCCCCCXXXXXXXXXXXXXXXXXXXXXXXRRRRRRBBBBBBBBBTTTTTTTTTTTTEEEEE
VVVVVVVXXXQQQQQQQQQQQQQQNNNNNNNNVVVVKVVVVJJJIIIIIIIIIIIIEEEEFFCCCCCCCCOOCCCCCCCCCCCCXXXXXXXXXXXXXXXXXXXXXXXXRRRRRBBBBBBBBBBTHHHHHHHTETTEEEEE
VVVVVVVXXQQQQQQQQQQQQQQQQNNNLNNNNVVVVVVVJJJJIIIIIIIIIIIIEEEFFFCCCCCCCOOOOOCCCCCCCCCCXXXXXXXXXXXFXXXXXXXXXXXXXRRRRRBBBBBBBBBTHHHHHHHTETIEEEEE
VVVVVVVVXQQQQQQQQQQQQQQQTTTTOOOOVVVVVVVVJJJJJJEIIIIIIIIIEEFFFFFFFCCCCOOOOOFCCCCCCCCCCCXXXXXXXXXXUXXXXXXXXXXXTTRRBBBBBBBBBVVTHHHHHHHEEEEEEEEE
VVVVVVHHQQQQQQQQQQQQQQQQQTTTOOOOOOVVVVVVPJJJJJJIIIIIIIIIEEFEEEEFZCCCOOOOOFFCCCCCCCCCCXXXXXXXXXXUUXXXXXXXXXXXTTTRRBBBBBBBVVVTHHHHHHHEEEEEEEEE
VVVVVVHEQQQQQQQQQQQQQQQQTTTOOOOOOOVVVVVPPJJJJIIIIIIIIIIIEEEEEEZZZZCOOOOOOOFCCCCCCCCCCCXXXXXXXXXUQXTTXXXXXXXXTTTTRRBBBVVVVVVTHHHHHHHEEEEEEEEE
VVVVVREETEEEEQQQQQQQTTTTTTTTOOOOOOVOVVVPPJJJJIIIIIIIIIIIEEEEEEEZZZZZOOAAOAFCCCCCCCCCCCXXXXXXXXQQQQQTTTXXXXXTTTTTRRRBBBBVVVVVHHHHHHHEEEEEEEEE
VVVVDEEEEEEYQQQQQQKQKKKKTTOOOOOOOOOOOOOPPPPJJIIIIIIIIIIIEEEEEEZZZZZZAAAAOAFPCCCCCCCDFXXXXXXXXQQQQQQTTTTTTTTTTTTTTTRBCCVVVVVVVHHHHHHTTEEEEEEE
VVVVDDEEEEEEHQQQQQKKKKKKKOOOOOOOOOOOPPPPPPPJJJJIIIIIIIIIEEEEEEEEIZZZZZAAAAADCCCCDCCDXXXGXXQQXQQQQQTTTTTTTTTTTTTTTTTTCCCVVVVVVHHHHHHTEEEEEEEE
VVVDDDEEEEEYHQQQQQQKKKKKKKKOHOOOOOOPPPQPPQQJQQQIIIIIIIIIEEEEEEEEZZZZAAAAAADDDDDDDCCDDXXXXXQQQQQQQQTTTTTTTTTTTTTTTTCCCCCVVVVVVVVVVTTTEEEEEEEE
VVVDDEEEEEYYYYYYQYGKKKKKKKKKOOOOOOOPPQQQQQQQQQQQQQPPEEEEAEEEEEEEZFZZKAAAADDDDDDDDDDDDDDHHHQQQQQQQZATTTTTTTTTTTTTTTCCCCCVVVVVVVVVVVEEEEEEEEEE
VVVDEEEEEEEYYYYYQYGKKKKKKKKKKOOOOMPPPQQQQQQQQQQQQPPPPLELAEEEEEEEEFKKKAADDDDDDDDDDDDDDDDAHHQQQQQQQZANTTTTTTTTTTTTTTCCCVVVVVVVVVVVVVEEEEEEEEEE
LVVDEEEEEEYYYYYYYYYKKKKKKKKKKOOOMMMMPPQQQQQQQQQQQPPPPPELAAAEEEEPEFKKKDDDDDDDDDDDDDDDDDDAHHQQQQQQQZATTTTTTTTTTTTTTTCCCVVVVVVVVVVVIEEEEEEEEEEE
VVVEEEEEEEYYYYYYYYYYYKKKKKKKKKMOMMMMPPPQQQQQQQQQQQPPAPAAAAAMEEEPPKKKKDDDDDNDDDDDDDDDDDAAQQQQQQQQQAATTTTTTTTTTTTTTCCCCVAVAVVVVVVVVUUUUCCPPEEE
RRRRRRREEYYYYYYYYYWYYYYKKKKKKPMMMMPPPQPQQQQQQQQQQQQPAAAAAAAEEEPPPKKKDDDDDTNDPCDDDDDDDDAAAQQQQQQYQQAAATAATTTTCTTTCCCCAAAAAAAVVVVVVVUCCCPPCCCC
RRRRRRRRRRYYYYYYYYYYYYYKKKKKKKBMMMPPQQQQQQQQQQQQQQAAAAAAAATTEEPPKKKKKKKTTTTCCCCDDDDAAAAAAQQQQQQYQYAAAAAATTTTCCCCCCCAAAAAAAAVAAVVPCCCCCCCCCCJ
RRRRRRRRRYYYYYYYYYYYYYYKKKKKKKKMMMCCWWQQQQQQQQQQQAAAAAAAAAEEEEKKKKKKKKKTTTTCCCSDDDDDAAAYYGQQQQYYYYAAAAAAAAATCCCCCCCCCAAAAAAAAAAACCCCCCCCCCCC
RRRRRRRRRYYYYYYYYYYYYYYKMMKKMKMMMMCCCCCQQQQQQQQAAAAAAAAAAAAAAADDKKKKKKTTTTCCVCCCCCDDDAAYGGQQYYYYYYAAAAAAAAATAACCCCCCCAAAAAAAAAAAACCCCCCCCCCC
RRRRRRRRRYYYYYYYYYYYYYYKKMMMMMMMCCCCCCCQQQJQQQAAAAAAAAAAAAAADDDDDKTVKTTTTCCCVCCCCCCDDCAAGGGGGYYYYAAAAAAAAAAAAACDDDDCCAAVAAAAAAFACCCCCCCCCCCC
RRRRRRRRRYYYYYYYYYYYYYYYMMMMMMMCCCCCCCCQAIISQQQAAAAAAAAAAAAAADDDDKTTTTTTTCCCCCCCCCDDDCAAGGGGGYYYYAYAAAAAAAAACCCCDDDCCDDVAABAAAACCCCCCCCCCCCC
TRRRRRRORYYYYYYYYYYKYTTMMMMMMCCCCCCIIICCIXISQQAADAAAAAAAAAAAADDDDTTTTTTTTCTTCCCCCCCCCCAGGGGGGYGYYYYYAAAAAAAAACCCDDDDDDDDDAAAAAAACCCCCCCCCCCC
TTRRRROOOYYYYYYYWYYTYTTTTMMMMMCCCCCCCIIIIIIQQQADDDDAAMAAAAAADDODOOTTTTTTTTTCCECCCCCCCCGGGGGGGGGGGYYYYYAAAAACCCCDDDDDDDDDDAAAAAACCCCCCCCCCCCC
TTTRTROOOOYYYYYYTTTTTTTTTTMMMMCCCCCIIIIIIIIQZZDDDDAAMMAAAAADDOOOOTTTTTTTTTTCCCCCCCCCCCGGGGGGGGGGGGGAAAAAAAAACCDDDDDDDDDDDCAAAAAACCCCCCCCCCCC
TTTTTOOOOYYLLYLYLTTTTTTTTTMMTTCZCICIIIIIIIIIDDDDDMMMMMMAAAOOOOOOOOJTTTTTTTTTTCCCCCCCCGGGGGGGGGGGGGAAAAAAAAACCCDDDDDDDDDCCCCAACCCCCCCCCCCCCCC
TTTTTOOOOOOOLLLLLTTTTTTTTTTTTTCIIIIIIIIIIIIIDDDDDDDMMMMMAAOOOOOOOOOTTTTTTTTCCCCCCCCCGGGGGGGGGGGGGNNNAAALACACCDDDDDDDDDDCCCCAACCCCCCCCFFCCCCQ
TTTTTTOOOOOOOLLLLLLATTTTTTTTTTIIIIIIIIIIIIINNNNNNDMMKKMNOOOOOOOOOOOOOOTTTTTTCTCSCCCSGGGGGGGGGGGOOONNNAACACCCDDDDDDDDDDDDCCCCCCCFFFFFFFFFFCCQ
TTTTTTOOOOOOOLLOLLLAATTTTTTTKTIIIIIIIIIIIIINNNNNNDUUKKKOOOOOOOOOOOOOOOTTTTTTTTCSSCCSDGGGGGGGGGGGOOONNACCACCCDDDDDDDDDDDDCCCCCCCCFFFFFFFFFFFF
TTTTTOOOOOOOOOOOLAAAATTTTTKKKTIIIIIIIIIIIIINNNNNNUUUUUUOOOOOOOOOOOOOOOOTTTTTTTSSSSSSDDSSGGGGFGGGOOOCCCCCCCDDDDDDDDDDDDDDCCCCCCCFFFFFFFFFFFFF
TTTTTOOOOOOOOOOOOAHAATATKKKKTTTKIIIIIISIIEEEEUUUUUUUUUUJOOOOOOOOOOOOOOTTTTTTTSSSSSSSSSSFFFFFFFFGGOOOOACCCCCDDDDDDDDDDDDCCCCCCFFFFFFFFFFFFFFF
TTTTTOOOOOOOOOOOOOAAATAKKKKKTKKKEEIIIEEIIEEEEEUUUUUUUUUJJOOOOOOOOOOOOOTTTTTTTTSSSSSSSSFFFFFFFFGGOOOOOCCCCCCIIDDDDDDDDDDCCCCCCFFFFFFFFFFFFFFF
TTTTTTOOOOOOOOOXOOAAAAAAAKKKKKKKKKEEEEEIIIIEEEEUUUUUUUUPOOOOOOOOOOOOTTTTTTTXTSSSSSSSSSSSFFFOOOOOOOOOOOCCCICIIDDDDIDDDDDDDDCCCCCFFFFFFFFFFFFF
TTTTTTOOOOOOOOOOAAAAXXXAKKKKKKKKKKEEEEEEEEEEEEEUUUUUUUUPOPPOOOOOOOOTTTTTTTFSSSSSSSSSSSSSFFOOOOOOOOOOOOCCCIIIIIDDDIIIDTDDDCCCCCCFFFFFFFFFFFFF
TTTTTTOOOOOOOOOOXAAAAAXXKKKKKKKPPEEEEEEEEEEEEEEEEUUUUUPPPPPPUUUUOUUTTTTTTTFFFBBSSSSSSSKXFFFOOOOOOOOOOOCCIIIIIIIIIIIIITTDDFHHCHCHHFFFFFFFFFFF
TTTTTTTOOOOOOOXXXXXXXXXXXKKKKKPPEEEEEEEEEEEEEEEEEPUUUPPPPPPGPUUUUUUTTTTTTTTFFFFSSSSSKKKXFFFXOOOOOOOOOOOOOIIIIIIIIIIIITTTTHHHHHHHHFFFFFFFFFOF
TTRRTTTOOOOQOOXDDXXXXXXXXXKKPPPPEEEEEEEEEEEEEEEEEPPPPPPPPPPPPPUUUUURRTTTTTBFFFFFSSSSSKKXXXXXOOOOOOOOOOOOOIIIIIIIIIITTTGGTHHHHHHHHHHFFFFFFFFF
RRRRRRTTOOOQQDXXDDXXXXXXXXKKPPAAEEEEEEEEEEEEEEEEEEPPPPPPPPPPPPUUUURRRRRTTBBFFFFVFSSSSSKXXXXXXOOOOOOOOOOOIIIIIIIIIIITTTGGEEHHHHHHHHHFFFFFFFFF
WWRRRRTTTTOOQDDDDDXXXXXXXXKKPPAAAAENNEEAEEEEEEEEEEPPPPPPPXPPHPUUUURRRRRTTBBBFFFFFSSSSSXXXXXXXXOOOOOOOOOIIIIIIIIIIIIIIGGGEHHHHHHHHHFFFFFFFFCC
WWRRRRRTRRROQQQDDDDDDXXXXXXXXPAAACNNNEEAAAEEEEEEEQQQQPPPXXPXXXXXXRRRRREETTFFFFFFFFFFSKXXXXXXYXOOOOAAOOOOOIIIIIIIIIIITTGGHHHHHHHHHHFFFFWWFCCC
WURRRRRRRRRDDQDDDDDDDXXXXXXXPPAAANNNNNAAAAAEEELLPQQQQPPMXXXXXXXXXRRREREEEEFFFFFFFFUSSSXXXXXXXXXXXOAAOFFFFIIIIIIIIIITTTGGGHHHHHHHHHFFZFNNCCCC
UUUUURRRRRRDDDDDDDDDDDDXXXXXPXAAAANNAAAAAAAALELLLQQQQPMMMMMXXXXXXRRREEEEEEEEFFFFFFUSSUUXXXXXXXXXOOAAOFFAFIIIIIIPPPPPPPPGGGHHHHHHHHHHHNNNNNCC
TTUUUURRRRDDDDDDDDDDDDDDXXXXPXAAAANNAAAAAAOOLLLLLQQQQMMMMMMPPXXRRRRRREEEEEEEEEFFFUUUUUUUXXXXXOXOOOAAAAAAAIIIIIIPPPPPPPPPPPGHHHHHHHHHINNNNNCC
TTUUURRRNRRRNDDDDDDDDDXXXXXXXXXXAAAAAAAAAAOLLLQQQQQQQMMMMMPPPXRRRRRRREERRRREEEEKUUUUUUUCXXOOOOOOOAAAAAAAACCIXHHPPPPPPPPPPPGHHHHHHHHHINNNNNNC
TTUNNRRRNRNRNNDDDDDDDDDDXHXXXXXXAAAAAAAAAQQQQQQQQQQQQMMPPPPPPXRRRRRRRRRRRRREEEKKUUUUUUUCOOOOCOOOOAAAAAAAACXXXXSPPPPPPPPPPPGHHHHHHHHHHNNNNNNN
TTTTNNNNNNNNNNDDDDDDDDDDXHXXXAAAAAAAAAAAAQQQQQQQQQQQQOMPPPPPPXRRRRRRRRRRRRKKKEKKKUUCCCUCCCCCCOOOOAAAAAALACXWXXSISZZPPPPPPPGGHHTHHTHNNNNNNNNN
TTTTNNNNNNNNNNNNDDDDDDDDDDAAAAAAAAAAAAAAOQQQQQQQQQQQPPPPPPPPPVRRRRRRRRRRRKKKKKKKUUUCCCCCCCCCOOOOOOHAAAAAACXWWXSSSZZPPPPPPPKKKTTTTTHHNNNNNNNN
TTTTNNNNNNNNNNNNDDDDYYYYDDDAAAAAAAAAAAAAHHOORRROOQQQOPPPPPPPPPRRRRRRRRRRRRKKKKKKUUUCCCCCCCCCCCOOOOHHAAAAAHHWWWHHSSZPPPPPPPGTTTTTTTTNNNNNNQQQ
TTTTNNNNNNNNNNNDDYDYYYYYYAAAAAAAAAAAAAAHHHORRRRROOOOPPPPPPPPPPPPRRRRRRRKKKKKKKKKUUCCCCCCCCCCCCOOOHHHHHAHHHHWWHHSSSSPPPPPPPTTTTTTTTNNNNNNQQQQ
TTTNHNNNNHHHHNNDYYYYYYYYYAAYYAAAAAAAAAHHHRMMRRRROOOPAPPPPPPPPPPPPKRRRRRRKKKKKKCCUUCCCCCCCCCCCCCHHHHHHHHHHHHHHHHHHSGPPPPPPPTTTTTTTTTNNNNNNNNQ
TTTNNNNNNHHHHHHDYYYYYYYYYKKYYYBYAAAAAAHHMRMMMRRRRRPPPPPPPPPPPPPKPKRRRRRIKKKKKKCCCCCCCCCCCCCCCCCHHHHHHHHHHHHHHHHHHHGPPPPPPPTTTTTTTTTTNNNNNNNN
TTNNTNNNNHHHHHHYYYYYYYYYYKKYYYYYAAAAAAHMMMMMMRRRRRRRRRPPPPPPLPPKKKKRRRRIIKKKKKKCCCCCCCCCCCCCCCLLLCCHHHHHHHHHHHHHHHGPPPPPPPTTTTTTTTTTTNNNNNNN
TTTTTTNHNHHHHHYYYYYYYYYYYKJYYYYYAAAAAHHMMMMXRRRRRRRRRRRPPPPPRKKKSSSSRRRRSSSKKKKKCCCCCCCCCCCCCCCCCCCHHHHHHHHHHHXXXXXXXXUUUTTTTTTTTTTNNNNNNNNN
NNNTNNNHHHHHHHHYYYYYYYYYYYYYYYYYYAHAHHHHHMMXXRXRRRRRRRRRPPRRRRRKSSSSSSSSSSSKKKKKKCCCCCCCCCCCCCCCCHHHHHHHHHHHHHHHXXXXXUUUUTTTTTTTTTTNNNNNNNNN
NNNNNHHHHHHHHHHYYYYYYYYYYYYYYYYYYHHHHHHHHXXXXXXRRRRRRRRRRPRRRRRKSSSSSSSSSSSSKKKKQCCCCCCCCCCCCMCCHHHHHHHHHHHHHHHHXXXXXXUUUTTMTTNNNTTNNNNNNNNN
NNNNNNNNHHHHHHHYHYYYYYYYYYYYYYYYYHHHHHHHXXXXXXXRRRRRRRRRRRRRRRRKSSSSSSSSSSSSKKKKQQQQCCQCQQQQQHHHHHHHHHHHHHHHHHHXXXXXXXXXXTTNNNNNNNNNNNNNNNNN
NNNNNNNNHHHHHHHHHYYYYYYYYYYYYYYYYYYHHHHHXXXXXXXRRRRRRRRRRRRRRRSSSSSSSSSSSSSSKKKKQQQQQQQQQQQQQQHHHHHHHHHHHHHHHXXXXXXXXXXXXXTNNNNNNNNNNNNNNNNN`
