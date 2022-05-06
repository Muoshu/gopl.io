package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

//自身和上下左右五个位置
var positions = [][]int{{0, 0}, {0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var m, n int = rowAndColumn()

type Node struct {
	//(i,j)当前位置
	i, j int
	//机器人数 被监控的房间数
	robotNum, roomNum int
	//机器人位置
	robotPos [30][30]int
	//被监视的房间位置
	roomWatched [30][30]int
	//堆中的索引
	index int
}

// PriorityQueue 一个实现了 heap.Interface 接口的优先队列，队列中包含任意多个 Item 结构。
type PriorityQueue []*Node

var pq PriorityQueue

func main() {
	//机器人最多的数量
	ans := m*n/3 + 2
	var ansArr [30][30]int
	//初始化
	var node Node
	node = initNode(node)

	//放入快照队列
	heap.Push(&pq, node)
	heap.Init(&pq)
	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*Node)

		//如果房间没有全部被监控，则分别在当前点的下方、本身和右方放置机器人
		//注意这三种情况是互不干扰的，它们会生成三种快照，判断出用机器人最少的一个
		if p.roomNum < m*n {
			//1.放在下方，判断条件为下方有位置，不能放在最后一行
			if p.i < m {
				setRobot(*p, p.i+1, p.j)
			}
			//2.放在本身
			//第一个判断条件是在它已经没有下方和右方的点的情况下，只能选择自身
			//第二个判断条件是它的右边没有被监控

			if (p.i == m && p.j == n) || p.roomWatched[p.i][p.j+1] == 0 {
				setRobot(*p, p.i, p.j)
			}

			//3. 在右边放置
			//第一个判断条件是遍历点右边是没被监控的点
			//第二个判断条件是遍历点右边的右边是没有监控的点
			if p.j < n && (p.roomWatched[p.i][p.j] == 0 || p.roomWatched[p.i][p.j+2] == 0) {
				setRobot(*p, p.i, p.j)
			}
		} else {
			if p.roomNum < ans {
				ans = p.roomNum
				ansArr = p.robotPos
			}
		}
	}

	fmt.Printf("机器人数量：%d\n", ans)
	fmt.Println(ansArr)
}

func rowAndColumn() (int, int) {
	file, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Printf("打开文件失败，err:%v\n\n", err)
		return 0, 0
	}
	data := string(file)
	re := regexp.MustCompile("[0-9]+")
	info := re.FindAllString(data, -1)
	m, _ := strconv.Atoi(info[0])
	n, _ := strconv.Atoi(info[1])
	return m, n
}

func setRobot(p Node, x, y int) {
	//复制一份快照p
	var node Node
	node = initNode(node)
	node.i = p.i
	node.j = p.j
	node.index = p.index
	node.robotNum = p.robotNum
	node.robotPos = p.robotPos
	node.roomWatched = p.roomWatched

	//在(x,y)点新增机器人，机器人数量+1
	node.robotPos[x][y] = 1
	node.robotNum = p.roomNum + 1

	//对这个新增的机器人上下左右和自身标记监控
	for d := 0; d < 5; d++ {
		//posX,posY表示机器人上下左右位置，标志这些位置的房间被监控
		posX := x + positions[d][0]
		posY := y + positions[d][1]

		//标记一个房间，roomNum就加1
		//一要等于1，因为有些房间会被重复监控
		if node.roomWatched[posX][posY] == 1 {
			node.roomNum++
		}
	}

	//如果行数不越界且当前点被监控了
	for node.i <= m && node.roomWatched[node.i][node.j] == 1 {
		//当前点的列右移一个单位
		node.j++
		if node.j > n {
			node.i++
			node.j = 1
		}
	}
	heap.Push(&pq, node)
	pq.update(&node, node.roomNum)
}

func initNode(node Node) Node {
	node.i = 1
	node.j = 1

	//在博物馆上下扩充两行
	for i := 0; i < m+1; i++ {
		node.roomWatched[i][0] = 1
		node.roomWatched[i][m+1] = 1
	}
	//在博物馆左右扩充两列
	for i := 0; i < n+1; i++ {
		node.roomWatched[0][i] = 1
		node.roomWatched[n+1][i] = 1
	}
	return node
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 我们希望 Pop 返回的是最大值而不是最小值，
	// 因此这里使用大于号进行对比。
	if pq[i].roomNum > pq[j].roomNum {
		return true
	}
	return false
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // 为了安全性考虑而做的设置
	*pq = old[0 : n-1]
	return item
}

// 更新函数会修改队列中指定元素的优先级以及值。
func (pq *PriorityQueue) update(item *Node, priority int) {
	item.roomNum = priority
	heap.Fix(pq, item.index)
}
