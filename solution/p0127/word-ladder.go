package p0127

func ok(x, y string) bool {
	count := 0
	for i := range x {
		if x[i] != y[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return count == 1
}

type Node struct {
	Cur   int
	Value int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	tar := -1
	length := len(wordList)
	edge := make([][]int, length+1)
	for i := range edge {
		edge[i] = make([]int, 0)
	}
	vis := make([]bool, length+1)
	for i := range wordList {
		if wordList[i] == endWord {
			tar = i + 1
		}
		if ok(beginWord, wordList[i]) {
			edge[0] = append(edge[0], i+1)
		}
		for j := i + 1; j < length; j++ {
			if ok(wordList[i], wordList[j]) {
				edge[i+1] = append(edge[i+1], j+1)
				edge[j+1] = append(edge[j+1], i+1)
			}
		}
	}
	if tar == -1 {
		return 0
	}
	que := make([]*Node, 1)
	que[0] = &Node{0, 1}
	vis[0] = true
	for len(que) > 0 {
		cur := que[0].Cur
		value := que[0].Value
		if cur == tar {
			return value
		}
		que = que[1:]
		for i := range edge[cur] {
			if vis[edge[cur][i]] {
				continue
			}
			vis[edge[cur][i]] = true
			que = append(que, &Node{edge[cur][i], value + 1})
		}
	}
	return 0
}
