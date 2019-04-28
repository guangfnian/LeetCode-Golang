package p0684

type UnionFind struct {
	root []int
	N    int
}

func (u *UnionFind) Init(n int) {
	u.N = n
	u.root = make([]int, n)
	for i := 0; i < n; i++ {
		u.root[i] = i
	}
}

func (u *UnionFind) Find(x int) int {
	if x != u.root[x] {
		u.root[x] = u.Find(u.root[x])
	}
	return u.root[x]
}

func (u *UnionFind) Combine(x, y int) bool {
	x = u.Find(x)
	y = u.Find(y)
	if x == y {
		return false
	}
	u.root[x] = y
	return true
}

func findRedundantConnection(edges [][]int) []int {
	max := 0
	for i := range edges {
		if edges[i][1] > max {
			max = edges[i][1]
		}
	}
	u := &UnionFind{}
	u.Init(max + 1)
	for i := range edges {
		if !u.Combine(edges[i][0], edges[i][1]) {
			return []int{edges[i][0], edges[i][1]}
		}
	}
	return nil
}
