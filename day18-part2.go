//go:build ignore

package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	x, y int
	dir  int
	cost int
}

type nodeHeap []node

var _ heap.Interface = (*nodeHeap)(nil)

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x any) {
	*h = append(*h, x.(node))
}
func (h *nodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type visitedNode struct {
	x, y int
	dir  int
}

type pos struct{ x, y int }

var dirs = [4]pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func makeGrid(mx, my int, bytes []pos, nbytes int) [][]rune {
	grid := [][]rune{}
	for y := 0; y < my+1; y++ {
		row := []rune{}
		for x := 0; x < mx+1; x++ {
			row = append(row, '.')
		}
		grid = append(grid, row)
	}

	for i := 0; i < nbytes; i++ {
		grid[bytes[i].y][bytes[i].x] = '#'
	}

	return grid
}

func main() {
	var bytes []pos
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		bytes = append(bytes, pos{x, y})
	}

	for i := range bytes {
		var found bool
		seen := map[visitedNode]bool{}

		pq := make(nodeHeap, 0)
		heap.Init(&pq)

		heap.Push(&pq, node{0, 0, 1, 0})
		heap.Push(&pq, node{0, 0, 2, 0})

		grid := makeGrid(70, 70, bytes, i)
		if grid[0][0] == '#' || grid[len(grid)-1][len(grid[0])-1] == '#' {
			fmt.Println(i, bytes[i])
			break
		}

		for pq.Len() > 0 {
			n := heap.Pop(&pq).(node)
			if n.y == len(grid)-1 && n.x == len(grid[n.y])-1 {
				found = true
			}
			if seen[visitedNode{n.x, n.y, n.dir}] {
				continue
			}
			grid[n.y][n.x] = 'O'
			seen[visitedNode{n.x, n.y, n.dir}] = true

			next := dirs[n.dir]
			nx := n.x + next.x
			ny := n.y + next.y

			if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) && grid[ny][nx] != '#' {
				heap.Push(&pq, node{nx, ny, n.dir, n.cost + 1})
			}

			for i := 0; i < 4; i++ {
				heap.Push(&pq, node{n.x, n.y, i, n.cost})
			}
		}

		if !found {
			fmt.Println(i-1, bytes[i-1])
			printGrid(grid)
			break
		}
	}
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

var ex = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

var input = `62,31
53,52
66,65
61,34
7,21
11,4
45,67
67,62
59,43
52,65
5,16
55,41
41,62
64,35
61,53
8,11
58,61
62,57
9,18
69,51
34,29
26,25
60,41
22,35
9,32
1,10
65,41
40,63
51,45
23,10
39,70
20,15
15,17
64,59
5,28
57,44
19,39
54,35
61,58
26,19
69,60
21,3
13,1
59,41
55,54
5,19
67,65
63,35
69,61
59,55
67,47
67,38
13,9
33,17
37,18
57,33
10,11
33,55
57,52
5,6
69,43
29,21
29,25
39,54
35,57
64,63
61,50
59,70
19,20
63,37
16,15
28,29
6,17
43,70
2,3
25,25
55,69
13,23
28,19
63,44
3,23
29,15
61,47
32,11
63,65
17,37
39,67
37,59
4,15
49,45
57,64
9,33
28,21
20,35
18,31
31,28
27,1
5,9
65,59
7,5
60,51
24,17
57,47
17,20
13,24
34,5
30,23
67,69
29,32
22,1
61,37
49,68
7,35
8,29
12,47
69,64
31,2
15,33
25,9
27,27
59,56
41,53
54,57
13,16
27,38
2,7
15,13
61,39
11,5
2,17
29,7
53,67
31,1
5,25
13,7
65,52
24,11
30,3
28,17
25,23
11,27
9,6
54,39
31,17
15,32
53,40
3,2
31,20
21,38
67,50
58,43
67,67
41,58
9,27
3,17
66,31
20,23
19,2
14,5
25,32
11,33
13,13
55,62
29,26
33,60
29,69
8,35
20,9
11,9
33,67
6,33
59,37
31,15
30,17
28,3
17,18
12,21
53,57
26,37
33,61
27,17
17,31
61,43
21,18
66,27
32,25
33,66
25,15
27,16
51,41
17,27
53,69
67,57
10,33
8,19
57,36
8,23
17,23
42,59
14,7
35,65
33,57
39,57
15,8
20,5
67,40
21,37
55,44
60,45
47,65
67,43
23,1
70,47
42,55
15,31
27,19
52,63
34,25
39,60
17,3
15,47
15,23
5,14
41,65
7,31
11,7
29,5
69,52
6,3
13,21
70,59
51,44
36,7
11,19
68,33
14,17
31,5
59,35
28,1
56,37
37,23
68,43
40,57
28,31
15,52
65,33
28,27
17,6
5,27
47,45
35,60
29,17
19,11
3,4
27,21
28,35
5,23
16,31
58,63
57,65
49,51
35,66
9,23
22,15
53,63
16,19
62,63
15,43
56,59
0,5
37,65
52,67
49,39
49,69
58,55
63,57
11,10
9,30
19,26
12,1
13,43
30,9
69,29
57,48
43,61
25,21
23,30
26,1
7,28
16,27
69,33
51,46
39,5
55,42
25,37
47,57
55,59
19,44
17,43
27,9
17,15
41,63
18,15
13,4
9,22
48,63
59,32
34,63
7,17
56,67
69,42
25,26
59,57
63,69
19,19
51,47
64,43
69,56
19,25
19,41
19,21
61,64
25,2
55,39
45,58
62,51
1,3
29,59
65,63
11,31
14,25
27,29
18,43
9,31
30,5
23,27
61,55
59,66
25,18
69,55
17,7
2,11
60,61
51,43
6,25
55,65
53,65
47,47
14,21
15,25
43,67
22,13
17,0
19,33
65,65
27,7
15,44
59,68
64,67
67,52
68,35
16,3
53,53
63,59
7,7
5,3
23,4
65,55
22,37
65,34
53,41
57,46
63,43
17,35
49,36
31,8
11,21
68,63
1,16
49,47
49,56
13,17
1,13
15,9
5,17
62,33
65,47
6,7
13,11
49,43
61,69
35,3
15,19
36,19
63,38
55,60
23,9
42,65
70,39
61,63
1,14
3,11
51,69
27,8
23,22
26,21
9,3
30,29
27,4
7,24
65,58
57,69
21,39
1,21
27,37
59,67
25,33
35,14
64,55
42,67
57,43
39,59
65,50
16,43
23,11
53,39
9,19
22,23
45,68
55,32
59,47
35,25
17,19
19,23
39,69
37,5
66,69
27,3
15,29
69,68
63,45
19,27
53,34
21,44
15,10
13,27
1,5
23,28
7,25
15,15
41,61
67,41
43,68
53,47
22,25
59,62
25,27
35,61
65,45
57,37
18,17
21,31
59,59
9,16
61,45
60,37
25,24
9,8
11,28
56,47
67,35
43,65
63,46
28,23
51,67
33,35
31,0
31,27
69,36
57,68
24,39
66,47
57,61
63,52
57,49
11,17
20,41
11,24
33,56
57,41
35,20
31,63
7,3
23,37
56,65
7,29
47,66
34,7
55,55
9,29
23,17
38,67
7,11
35,67
58,49
9,7
22,9
13,42
15,1
62,69
25,31
17,46
16,23
70,33
11,11
12,33
19,24
37,3
62,53
3,12
69,67
61,59
66,35
65,43
59,61
55,47
5,5
56,55
12,27
11,1
68,57
31,23
61,31
29,3
27,15
61,67
23,5
66,67
59,63
45,63
14,11
67,51
17,29
14,1
59,34
19,12
23,3
9,5
69,63
63,41
37,64
65,61
17,34
9,26
61,42
15,45
28,7
8,1
3,1
13,19
67,49
14,35
5,2
38,55
69,31
65,53
25,7
21,12
58,35
18,29
69,69
3,13
37,57
39,63
55,63
35,17
33,18
23,23
23,20
62,47
67,54
69,57
23,31
13,18
9,20
67,39
55,67
67,59
9,25
48,67
50,67
67,63
18,7
7,13
13,15
15,38
67,32
33,9
3,6
13,35
5,1
21,40
17,17
65,51
19,29
13,33
45,47
5,22
22,41
57,31
62,39
63,66
63,51
17,45
23,25
32,19
59,69
68,67
66,61
68,47
58,41
17,21
45,66
31,16
8,15
18,23
33,1
35,56
25,20
58,59
40,67
32,23
53,43
59,33
36,59
22,17
11,18
32,63
19,13
26,5
59,65
57,66
10,7
65,31
56,33
19,4
49,64
37,67
21,7
63,47
12,13
68,41
41,67
39,61
1,9
3,21
53,62
3,0
27,31
19,42
31,7
27,14
69,65
19,36
60,31
65,46
25,17
61,35
43,5
25,35
21,11
1,22
69,54
18,39
59,54
58,51
69,35
17,33
29,19
37,62
37,58
35,58
29,35
65,67
21,17
15,12
45,65
31,25
13,14
24,5
27,10
5,47
25,11
5,13
11,2
69,45
33,24
67,33
10,15
4,9
21,26
9,9
21,19
7,23
17,40
51,70
12,7
15,7
21,33
21,25
55,38
17,26
63,61
11,12
19,32
63,55
33,63
38,61
53,35
33,7
5,8
6,13
13,25
53,46
1,19
37,66
64,31
11,22
14,45
6,19
5,11
67,53
37,61
1,7
55,35
27,11
7,1
63,49
13,45
61,30
58,31
17,41
13,34
22,31
41,55
24,3
68,37
29,12
64,49
61,36
6,1
24,29
37,69
24,37
23,29
6,31
37,63
17,10
17,9
27,28
7,27
35,4
20,19
24,25
4,19
22,7
21,13
17,25
17,22
23,39
19,17
35,18
60,65
20,29
21,15
9,2
35,21
26,31
54,65
9,4
27,35
7,19
11,37
39,64
9,13
53,51
37,1
69,47
61,29
23,19
61,33
7,10
12,31
27,23
44,59
69,49
0,1
49,60
55,53
3,15
63,33
27,25
52,69
53,44
13,5
65,60
5,15
29,39
39,65
26,13
17,13
25,34
43,69
17,49
1,1
55,57
38,57
13,31
57,63
62,61
27,13
21,23
61,61
15,11
18,1
58,39
46,9
69,41
25,36
32,3
31,21
19,3
59,44
3,33
63,63
19,15
63,68
11,29
67,61
25,10
49,58
65,69
18,13
0,13
25,1
64,33
66,43
9,21
70,67
8,5
53,45
39,2
19,34
39,26
13,48
66,45
61,65
11,3
59,49
66,49
47,69
23,14
63,56
22,21
33,25
19,7
6,11
1,27
67,55
5,35
63,31
61,66
30,21
44,65
40,69
66,55
56,41
13,47
10,25
15,20
61,41
21,6
9,1
60,57
68,49
3,7
10,27
55,61
5,7
47,55
47,48
46,69
11,15
33,29
21,9
69,37
9,17
5,4
35,7
17,8
33,3
65,40
57,35
57,59
69,59
62,41
21,5
15,30
65,70
69,53
54,53
45,9
25,13
43,63
16,35
40,65
61,49
12,19
17,36
65,42
65,57
11,30
53,54
7,9
23,7
19,16
41,57
13,3
54,47
31,68
57,67
51,65
62,29
63,48
29,24
7,26
21,29
29,31
41,59
65,62
19,5
49,57
9,12
33,22
15,36
68,65
9,14
45,60
31,13
12,45
33,62
21,35
45,69
21,21
54,69
55,43
29,1
3,19
25,8
1,15
25,40
68,59
54,67
33,5
41,60
51,63
39,17
11,23
3,18
7,15
15,14
47,53
19,37
55,33
33,58
23,13
57,45
17,1
25,6
47,63
48,69
1,11
3,14
17,38
69,39
66,57
55,31
57,39
20,31
25,12
33,59
2,9
21,27
19,10
35,19
35,16
49,63
17,11
61,60
3,9
3,5
1,17
60,47
12,9
32,17
21,28
26,29
29,6
25,22
23,35
25,29
25,16
47,59
46,61
13,26
15,50
32,65
53,42
23,15
29,33
59,38
65,39
4,23
57,57
41,36
43,19
28,33
37,13
23,64
53,33
21,1
19,43
45,37
13,59
43,27
54,49
49,25
9,39
10,67
51,42
55,45
57,21
42,53
7,39
7,20
43,12
36,29
9,70
51,17
61,3
6,69
3,57
47,41
39,9
37,47
47,31
52,23
0,33
44,3
20,67
18,63
67,4
3,32
13,50
1,33
28,55
9,64
33,54
49,59
59,16
47,27
43,21
33,33
13,63
10,49
29,55
10,61
35,47
47,33
30,61
53,37
63,15
27,44
63,17
48,41
33,68
41,47
31,59
27,45
19,47
37,14
47,7
53,31
17,56
1,49
61,51
51,39
7,50
25,55
16,59
23,57
1,51
36,49
58,1
43,53
51,54
43,56
65,11
41,7
69,27
53,11
51,31
29,27
68,15
31,40
68,3
7,33
59,45
6,45
58,17
57,20
5,49
1,64
53,25
45,2
41,23
63,14
63,3
55,22
45,23
1,30
23,66
67,19
29,29
37,43
0,19
61,23
65,6
57,40
48,61
28,67
38,17
48,45
24,53
15,37
45,18
39,37
65,35
51,59
48,11
7,61
3,31
11,45
28,47
65,37
18,51
28,45
35,23
29,42
53,13
17,67
41,8
42,49
38,43
12,55
33,69
53,29
57,51
19,59
25,69
14,29
19,63
47,58
26,47
43,44
3,42
59,28
33,13
69,10
46,63
1,54
41,28
45,41
41,31
47,43
11,39
12,43
50,47
46,5
53,9
69,9
39,47
49,49
3,63
5,53
44,39
52,59
47,37
46,11
65,3
33,42
37,33
51,19
51,61
59,39
45,40
29,65
29,70
25,46
17,57
15,63
65,19
27,55
3,43
0,57
41,29
41,3
36,9
3,52
50,25
35,13
40,9
46,15
17,68
51,3
30,13
59,25
9,41
53,49
10,47
40,21
43,15
16,41
57,24
5,33
37,51
31,35
55,28
5,29
67,25
70,25
31,32
53,55
47,17
45,57
15,49
37,37
61,2
43,23
7,45
5,61
22,69
37,19
55,25
51,49
45,3
53,18
36,63
37,10
69,13
3,67
38,45
41,22
49,29
9,49
51,14
65,25
45,33
47,61
39,53
44,15
65,29
42,15
8,67
68,1
3,30
11,38
66,21
69,17
69,11
30,55
49,18
61,21
55,21
31,29
22,33
20,57
41,48
70,13
54,23
55,7
16,5
37,25
46,27
14,41
43,41
33,41
59,11
3,29
43,11
35,39
31,11
60,9
44,55
15,54
54,27
5,37
31,30
14,65
61,20
67,23
11,56
53,23
30,35
19,45
33,31
43,17
19,65
50,27
7,62
9,63
22,59
49,48
53,38
13,37
41,54
43,52
2,29
20,49
65,26
33,27
61,9
19,64
21,45
55,49
11,59
67,21
57,23
67,6
69,19
63,25
65,16
23,33
65,12
39,51
50,7
68,23
31,51
53,17
17,39
45,55
42,37
51,37
52,35
59,17
41,35
9,67
13,51
17,59
30,65
67,5
41,18
9,11
34,51
50,19
1,35
15,21
4,47
43,3
39,43
41,27
35,38
17,65
67,11
32,5
62,27
5,38
26,65
5,52
22,53
29,45
41,15
56,29
27,69
18,27
63,0
55,12
51,15
41,4
9,43
57,15
50,5
67,1
11,58
41,11
49,33
4,57
41,43
1,45
11,68
7,64
12,5
43,64
1,37
7,52
53,1
49,26
49,67
67,14
43,9
43,31
59,22
35,45
69,25
47,29
59,31
21,53
45,21
43,35
21,2
49,4
67,37
3,66
57,5
45,30
35,5
27,40
27,41
25,49
57,9
23,43
35,43
55,1
56,9
34,41
31,19
5,57
64,21
49,27
21,70
39,15
36,69
3,51
50,23
25,45
28,11
45,24
63,20
18,67
40,19
37,4
50,63
27,57
41,51
55,27
5,63
47,21
66,1
13,61
31,69
31,46
45,6
9,36
69,1
59,6
45,22
4,65
36,37
5,59
31,9
51,57
38,41
1,53
39,20
69,15
46,43
29,50
30,37
53,20
55,37
9,53
42,5
59,18
41,19
49,13
65,8
33,46
61,54
34,35
45,19
2,37
47,35
12,63
69,28
49,54
31,3
13,41
52,5
4,49
36,43
47,19
33,11
51,12
57,7
39,27
48,27
27,48
25,65
19,56
23,63
45,46
25,19
7,43
45,45
60,23
47,49
55,20
59,19
68,11
49,17
49,5
39,29
13,62
53,56
25,41
47,0
31,37
41,41
4,39
37,48
44,43
51,5
1,62
49,3
1,34
39,41
1,39
11,57
39,23
64,9
3,25
33,45
31,54
54,61
20,47
5,45
30,49
43,42
35,28
4,55
21,67
50,49
55,16
53,5
42,39
48,53
15,3
47,50
25,51
38,13
33,34
53,15
25,57
36,11
60,15
63,19
51,33
64,27
43,55
35,35
52,49
23,69
13,60
39,39
28,53
69,30
41,37
65,49
41,1
33,36
37,31
68,19
45,25
37,24
55,9
53,14
1,67
9,35
9,46
16,47
51,0
39,19
59,23
46,55
55,6
9,51
31,43
36,35
3,39
63,2
4,35
67,22
43,37
36,53
7,63
21,42
25,53
19,67
61,15
53,27
29,47
31,48
14,55
41,21
22,49
5,40
43,49
2,41
15,70
47,51
36,23
51,7
53,10
37,41
51,58
15,55
3,50
50,33
21,49
21,59
3,47
11,43
45,53
24,63
58,13
55,2
61,13
57,17
58,21
59,4
33,65
63,4
42,19
57,4
35,9
28,39
1,23
19,55
68,27
33,44
9,61
17,12
25,50
27,68
41,9
15,53
59,26
6,37
27,47
39,49
23,45
1,58
39,3
35,41
13,57
54,11
59,15
9,69
39,13
51,11
29,37
49,62
19,69
23,34
11,51
15,41
5,31
49,44
49,41
3,49
57,55
32,37
34,39
47,67
3,41
3,53
3,60
65,17
1,40
7,53
18,69
62,17
70,17
11,47
67,29
37,55
56,15
56,51
7,58
3,62
37,45
39,35
45,15
40,5
11,63
10,39
31,47
41,45
17,61
35,69
31,58
45,39
39,1
56,19
59,14
4,27
17,53
15,64
26,53
49,16
8,33
54,1
8,53
47,5
49,31
69,3
19,35
11,64
23,49
9,55
63,18
69,23
29,51
39,31
51,30
35,37
39,50
56,27
7,46
8,59
38,5
61,6
45,61
54,15
26,55
67,31
1,59
52,19
66,23
35,49
52,51
55,3
6,67
47,46
9,45
39,8
7,47
45,54
25,68
41,14
45,4
14,57
34,47
38,1
37,34
2,23
56,25
1,61
17,63
13,67
20,51
59,53
13,53
35,63
13,49
2,45
11,53
15,28
15,2
23,53
7,57
5,43
37,49
45,1
55,30
49,23
47,39
17,62
55,5
33,51
30,51
61,19
61,27
27,62
51,16
43,32
9,15
35,51
42,45
18,47
65,38
44,7
51,55
22,67
1,26
40,31
11,52
41,49
37,50
29,38
48,33
55,51
65,18
37,27
1,46
7,51
48,29
3,59
33,37
34,15
43,1
45,20
54,25
57,11
52,21
9,38
50,21
9,65
59,5
30,67
36,3
51,2
43,47
45,28
45,59
27,43
15,59
21,55
34,21
47,3
19,9
21,61
27,61
43,30
33,21
5,44
57,10
38,31
15,48
31,67
5,41
49,9
51,53
31,41
70,21
27,51
18,65
6,61
69,8
5,48
15,5
47,1
16,63
35,31
66,9
63,27
18,59
44,37
55,29
63,13
60,27
24,43
3,3
29,11
51,27
63,10
31,45
1,43
43,45
31,53
47,15
35,55
51,29
11,54
51,23
37,17
5,21
49,32
24,45
51,51
7,55
6,59
23,51
15,61
27,65
55,50
47,13
3,45
1,48
13,40
41,24
21,47
45,36
19,1
29,41
11,69
25,5
25,67
57,25
40,35
37,11
9,59
50,41
23,54
61,12
43,25
1,50
49,11
29,53
1,29
17,55
50,53
5,69
5,39
69,7
25,61
23,61
47,11
63,29
43,26
15,68
29,60
11,61
3,27
49,21
6,35
33,47
18,53
46,39
19,61
39,40
1,65
13,39
7,41
65,23
47,10
31,56
35,15
37,46
9,54
38,35
52,27
24,69
3,35
61,1
50,11
21,57
31,33
7,49
59,9
1,31
13,65
38,21
3,36
13,38
57,6
49,19
32,51
43,57
52,57
25,60
42,1
25,42
26,43
33,43
27,63
35,52
29,14
29,64
60,11
61,25
15,58
47,24
29,63
27,39
5,30
36,39
47,9
43,18
44,23
33,49
3,68
69,21
45,51
59,21
65,24
3,44
55,19
63,7
12,59
2,51
27,56
14,61
46,21
45,29
48,21
60,49
52,29
31,39
52,37
69,4
40,51
55,15
9,56
39,45
46,51
39,55
31,57
43,13
43,2
32,7
34,9
49,38
50,39
9,57
21,54
7,66
61,17
63,22
31,31
11,35
48,23
0,43
59,51
48,7
1,25
15,35
49,35
69,5
35,12
39,33
37,53
46,1
30,45
5,55
1,60
24,49
35,46
23,60
41,25
65,9
56,13
3,20
48,15
41,13
41,33
36,31
67,17
63,11
68,7
3,58
21,69
28,59
61,57
2,67
41,16
47,42
49,1
51,9
25,3
23,47
47,36
37,39
34,1
53,2
39,52
37,36
45,13
34,69
63,26
35,1
23,67
15,27
37,35
63,53
67,15
23,56
52,31
1,57
12,49
43,29
36,41
17,51
14,67
43,7
37,7
59,7
41,44
35,54
51,25
42,29
23,41
10,43
23,55
4,61
41,12
3,61
11,49
39,38
25,43
39,21
55,4
18,49
35,2
60,19
65,13
31,65
6,49
29,9
62,15
58,29
57,29
31,52
17,50
40,39
45,50
39,46
21,63
43,22
49,65
3,55
1,24
61,5
21,46
33,14
25,62
17,60
65,14
55,18
59,29
21,41
35,11
45,43
51,6
67,9
21,62
17,5
16,53
63,5
33,15
57,13
41,5
1,41
1,55
53,61
20,53
11,13
37,21
51,1
63,40
67,3
66,3
35,44
27,67
40,1
46,19
27,53
7,37
44,35
19,51
59,13
67,13
49,15
19,49
29,43
45,31
63,9
49,2
26,63
17,47
49,37
19,53
57,3
29,58
30,41
53,32
0,69
35,33
26,59
27,66
61,7
5,64
44,9
40,47
63,21
35,53
37,26
63,1
69,12
50,9
41,69
37,15
59,1
6,55
26,51
21,51
61,24
25,63
11,67
44,27
57,53
65,21
61,8
15,67
23,59
11,55
21,65
15,57
25,59
68,45
54,9
45,5
64,7
67,7
32,39
51,60
65,1
11,65
20,65
40,43
67,27
29,57
7,67
21,4
7,42
55,17
25,39
52,13
38,11
65,15
22,51
33,53
31,61
20,61
57,19
33,39
39,24
22,63
51,21
28,51
19,58
1,63
15,39
32,13
13,55
29,67
16,55
56,57
31,55
53,19
11,66
49,7
5,32
49,53
23,21
47,25
5,56
3,37
32,31
11,40
27,33
59,10
7,40
53,21
51,13
57,1
48,51
67,20
35,27
41,39
62,5
2,27
3,64
48,3
65,7
40,13
45,17
68,25
32,61
35,59
43,43
41,34
66,29
64,1
9,58
38,23
24,57
33,48
33,28
55,13
13,69
55,11
21,58
42,33
35,32
46,31
51,36
51,35
23,65
45,11
54,5
43,10
42,47
50,29
53,24
7,69
27,59
27,5
27,49
39,25
65,27
60,3
46,13
35,68
37,29
49,55
5,65
62,11
29,49
7,65
1,47
22,45
55,23
15,65
3,65
45,49
9,50
46,33
39,7
47,38
52,9
9,52
43,59
57,27
64,15
0,37
29,23
9,62
39,11
19,57
47,16
45,7
46,45
25,66
13,66
21,43
58,3
48,31
3,69
47,23
33,23
39,30
4,25
1,69
56,1
67,16
53,59
43,51
47,56
57,8
63,39
31,49
61,11
15,51
65,5
53,3
4,69
9,47
47,8
25,58
59,3
33,19
8,39
9,44
5,67
17,69
49,34
11,36
39,32
63,23
29,13
43,50
67,45
31,10
13,29
3,34
12,69
2,55
63,67
62,23
39,16
28,43
11,41
5,42
34,31
44,47
9,42
49,61
11,25
44,13
38,7
59,27
37,9
53,8
41,6
28,63
7,59
45,35
41,17
15,69
29,61
19,31
41,26
23,48
13,52
50,51
25,47
49,14
32,43
35,26
5,54
38,53
45,34
43,33
43,39
45,27
35,29
5,51
53,7
2,39
9,37
39,28
42,41
31,34
67,18
64,29
9,60
42,9
43,62
66,11
47,32
34,8
20,59
28,36
34,22
11,42
6,41
10,40
20,45
30,33
2,54
0,60
4,26
64,0
18,55
20,6
39,10
14,36
60,26
50,54
44,34
24,48
57,26
12,40
20,64
38,22
26,17
64,42
27,58
44,26
2,43
36,68
18,34
18,40
6,68
36,54
70,4
12,66
5,26
52,47
1,28
32,8
56,64
48,22
69,46
70,46
36,58
40,61
4,20
16,38
23,40
44,63
8,31
4,33
54,22
44,1
50,69
6,40
47,70
52,55
4,48
21,48
3,40
12,39
22,65
36,60
35,22
42,40
39,48
14,18
53,22
40,48
8,34
60,14
14,31
56,61
56,10
14,44
17,52
8,38
46,56
40,66
2,14
13,44
16,26
16,49
54,24
26,2
42,36
46,24
46,68
55,64
43,66
2,61
4,17
0,50
8,70
18,25
15,60
60,10
22,48
0,35
17,32
19,46
46,18
64,62
48,32
38,68
12,52
36,25
68,8
40,41
54,62
42,21
68,4
36,28
25,64
27,32
44,56
42,11
24,52
3,28
6,22
58,50
9,28
22,3
58,45
44,33
40,24
66,22
14,34
12,0
40,25
23,0
30,32
20,42
48,12
12,28
0,28
27,64
34,59
0,70
24,4
22,64
2,68
2,56
10,36
46,53
13,22
11,8
50,52
32,68
4,66
43,36
7,36
18,64
63,12
14,49
18,62
32,69
34,66
0,11
64,58
54,17
6,50
44,41
30,48
37,52
44,69
24,6
46,52
22,0
63,6
67,70
18,0
32,55
60,40
1,12
52,18
64,10
4,30
16,69
65,30
38,33
12,15
55,66
58,5
42,30
32,29
24,30
11,44
38,50
20,36
45,52
6,23
42,22
28,30
12,14
12,48
49,6
40,62
7,56
4,56
58,25
5,70
38,66
29,10
37,68
10,22
36,24
51,52
63,30
46,70
12,23
7,30
62,12
35,34
8,66
42,24
33,52
29,2
47,52
42,62
53,30
38,64
20,32
66,38
30,12
52,54
52,16
13,36
8,54
4,22
26,58
42,18
0,21
30,26
30,28
28,18
3,24
24,68
70,10
22,32
6,5
20,55
25,44
41,52
2,18
64,36
52,52
18,52
48,60
15,66
13,0
9,66
53,66
44,66
17,4
4,31
40,34
64,26
16,39
6,26
5,68
6,57
31,6
62,1
56,68
57,58
64,22
5,24
46,46
14,9
68,10
14,27
60,4
24,51
24,8
6,15
37,20
31,24
6,18
5,60
14,63
20,58
67,46
50,57
6,60
34,44
20,4
14,52
6,6
14,14
66,52
6,43
14,28
58,6
66,42
44,70
59,58
38,28
21,14
24,54
0,40
64,60
54,58
62,37
60,62
6,39
18,6
24,65
34,43
42,66
16,70
34,27
60,66
59,8
44,12
51,62
27,36
2,69
7,38
4,64
46,40
42,51
2,24
48,55
54,59
12,38
4,43
68,42
34,53
36,21
8,26
24,2
66,40
45,62
27,2
30,52
12,53
59,40
18,33
18,35
45,14
68,53
4,28
46,22
10,44
54,26
22,42
18,66
52,53
54,68
6,14
40,32
36,50
66,50
12,36
40,54
26,16
70,9
64,37
7,34
3,10
68,9
10,53
38,6
1,36
33,64
20,68
38,10
52,43
42,35
0,68
62,6
0,36
11,50
50,26
0,39
50,66
45,70
13,58
32,36
38,16
14,10
36,64
56,35
30,47
20,43
23,44
10,65
38,47
61,0
8,55
26,56
10,68
46,20
62,26
60,6
15,16
19,50
41,66
32,32
11,60
32,60
30,1
52,45
64,14
39,62
56,56
34,37
51,24
54,36
49,22
16,51
60,64
57,62
67,10
26,40
13,70
4,62
56,23
46,62
68,46
56,58
27,34
28,60
14,26
52,24
28,61
38,8
4,59
16,66
30,2
30,53
54,19
66,62
4,10
24,59
64,38
30,19
40,27
8,17
14,30
23,36
52,61
44,54
2,52
8,40
12,30
18,68
52,46
0,18
42,63
28,20
34,40
40,14
4,29
22,10
42,4
66,70
8,36
42,17
70,26
30,57
50,68
15,42
68,6
40,11
44,57
67,8
51,8
50,42
27,50
13,10
22,43
10,12
38,0
45,8
46,37
44,61
64,51
7,22
38,44
14,53
70,36
20,17
42,60
28,56
44,6
50,59
30,24
50,30
30,39
64,46
56,2
60,69
60,29
32,1
68,34
29,18
32,20
8,3
4,24
68,69
0,47
60,1
16,57
40,36
33,70
34,3
70,32
50,32
54,20
62,13
29,52
12,22
19,52
16,46
61,4
52,15
17,54
12,54
43,28
26,26
68,58
50,37
31,36
58,11
12,42
58,4
8,52
33,10
8,69
12,35
10,59
28,66
4,4
28,38
66,46
22,56
2,30
26,67
34,17
4,3
23,32
40,60
48,10
26,44
62,62
56,0
25,54
10,34
59,50
38,27
28,37
1,6
3,48
52,3
62,36
45,32
24,32
0,64
12,20
10,4
64,34
36,40
45,38
13,30
64,12
8,46
14,24
58,46
8,62
6,38
1,42
8,68
30,66
21,24
32,48
52,26
19,8
40,58
11,48
70,44
22,28
26,70
46,26
48,50
14,60
44,48
46,58
70,28
58,27
26,46
7,18
30,50
60,60
70,43
36,2
69,26
45,26
10,32
0,31
42,14
32,59
6,56
27,0
60,48
40,4
46,4
34,4
4,58
51,34
46,6
6,10
68,38
28,8
38,49
18,18
2,36
44,68
16,58
14,50
27,18
36,34
5,10
7,44
36,57
64,16
4,46
36,46
11,62
16,34
64,23
4,2
6,42
2,66
0,7
67,48
48,36
56,21
40,8
4,18
6,54
27,42
30,46
25,30
10,8
0,61
1,66
2,60
36,66
68,51
68,32
51,64
7,6
10,26
63,8
38,30
46,14
68,31
60,7
22,39
64,69
4,34
65,66
48,57
28,10
66,66
51,26
38,26
40,55
55,36
6,66
8,32
14,0
40,15
2,70
18,24
30,40
31,26
8,44
49,50
60,30
4,41
70,54
58,69
69,32
44,38
43,4
34,52
68,55
6,51
34,38
22,57
16,64
18,2
8,28
29,56
5,36
25,14
16,62
21,8
26,22
28,4
36,33
10,64
2,15
0,38
68,70
65,20
44,8
46,42
63,58
41,42
44,28
20,70
8,9
44,58
2,12
67,64
54,41
60,34
50,3
54,8
29,28
8,25
54,28
33,38
70,18
57,38
38,70
3,16
70,14
14,64
32,35
29,40
4,21
7,14
14,16
46,38
8,57
14,38
64,25
66,24
56,54
32,38
54,38
70,20
63,70
10,0
47,40
23,18
22,19
66,28
42,61
0,17
38,56
52,22
68,18
7,68
1,68
1,38
13,8
56,4
36,12
32,27
2,47
62,22
48,8
54,16
48,20
4,5
54,37
18,50
55,58
3,26
36,15
36,26
60,33
34,46
0,14
60,50
40,12
67,28
28,69
8,12
53,12
32,26
36,51
42,64
3,8
54,2
64,48
15,4
38,58
70,40
32,24
5,18
68,2
36,67
38,29
60,2
12,57
24,22
50,44
69,2
56,62
66,56
12,46
13,54
66,36
10,56
64,41
47,54
37,12
30,7
50,16
52,17
60,35
15,6
69,20
47,68
6,62
35,64
33,12
26,54
3,46
67,58
12,61
70,49
58,7
66,15
64,5
62,65
34,56
10,28
32,44
58,30
22,36
52,32
47,26
39,18
50,40
34,50
41,32
39,0
2,33
52,41
58,42
6,12
20,30
44,20
32,21
42,50
13,28
59,2
28,40
38,3
9,68
51,22
70,3
66,60
62,34
44,67
62,40
58,0
70,30
38,63
44,50
70,66
3,38
38,48
24,0
15,24
62,45
50,62
51,28
56,30
18,4
34,26
9,10
55,26
43,40
10,62
62,49
28,34
63,36
36,14
20,52
13,68
9,48
56,17
5,12
48,0
20,2
41,2
24,67
40,29
23,42
29,44
18,57
2,40
46,47
41,40
56,46
28,52
46,8
64,50
45,44
57,14
6,16
49,66
24,1
46,7
68,48
19,38
63,50
29,36
20,37
34,42
50,61
37,6
14,8
49,12
66,63
56,44
62,42
69,0
2,49
22,52
4,6
60,16
48,1
22,6
31,12
55,70
34,12
60,68
46,28
70,31
70,69
26,7
60,0
64,44
58,22
14,54
56,34
24,28
60,17
12,2
4,1
62,56
22,24
67,36
56,3
58,57
68,40
67,44
5,46
59,0
18,45
56,45
18,54
50,48
62,3
50,38
45,16
35,40
53,28
53,4
52,66
34,32
60,54
64,47
30,44
17,16
33,50
32,53
41,68
61,40
12,64
22,16
65,4
53,68
21,56
16,25
20,21
8,24
7,16
7,2
54,14
1,18
60,39
32,42
69,44
68,20
24,50
56,70`
