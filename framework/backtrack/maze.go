package main

import (
	"fmt"
)

type Maze struct {
	Map            [][]int
	StartX, StartY int
	EndX, EndY     int
	Width, Heigh   int
}

func NewMaze(mp [][]int, StartX, StartY, EndX, EndY int) *Maze {
	m := &Maze{
		Map:    mp,
		StartX: StartX,
		StartY: StartY,
		EndX:   EndX,
		EndY:   EndY,
		Heigh:  len(mp),
	}
	if m.Heigh != 0 {
		m.Width = len(mp[0])
	}

	return m
}

func (m *Maze) ShowMap() {

	for i := 0; i < len(m.Map); i++ {
		for j := 0; j < len(m.Map[i]); j++ {
			switch m.Map[i][j] {
			case 0:
				fmt.Printf("▢ ")
			case 1:
				fmt.Printf("○ ")
			case 2:
				fmt.Printf("■ ")
			}
		}
		fmt.Println()
	}
}

func (m *Maze) FindPath() {
	m.findPath(m.StartX, m.StartY)
}

func (m *Maze) findPath(x, y int) {
	if x >= m.Heigh || x < 0 { // 越界
		return
	}
	if y >= m.Width || y < 0 { // 越界
		return
	}

	m.Map[x][y] = 1 // 选择

	if x == m.EndX && y == m.EndY {
		fmt.Println("solution:")
		m.ShowMap()
	}

	// 进入下轮
	if y >= 1 && m.Map[x][y-1] == 0 { // 左边能走, 往左走
		m.findPath(x, y-1)
	}
	if y+1 < m.Width && m.Map[x][y+1] == 0 { // 右边能走, 往右走
		m.findPath(x, y+1)
	}
	if x >= 1 && m.Map[x-1][y] == 0 { // 上边能走, 往上走
		m.findPath(x-1, y)
	}
	if x+1 < m.Heigh && m.Map[x+1][y] == 0 { // 下边能走, 往下走
		m.findPath(x+1, y)
	}

	// 撤销选择
	m.Map[x][y] = 0
}

// 非递归的方式来回溯
// 想一想选择列表, 抽象出来问题后, 对应的是 上、下、左、右 这四个选择, 每个选择可以重复选取, 升级的排列问题
// 剪枝函数就是根据坐标来判断如何选择, 需要记录坐标
func (m *Maze) FindPath2() {
	x, y := m.StartX, m.StartY

	n := 4             // 上、下、左、右
	track := []int{-1} // 解空间索引

	k := 0

	// 移动位置, 先移动位置, 成功则选中新坐标
	move := func(direction int) bool {
		var canMove bool
		// 尝试移动坐标
		switch direction {
		case 0:
			if x >= 1 && m.Map[x-1][y] == 0 { // 向上
				x--
				canMove = true
			}
		case 1:
			if x+1 < m.Heigh && m.Map[x+1][y] == 0 { // 向下
				x++
				canMove = true
			}
		case 2:
			if y >= 1 && m.Map[x][y-1] == 0 { // 向左
				y--
				canMove = true
			}
		case 3:
			if y+1 < m.Width && m.Map[x][y+1] == 0 { // 向右
				y++
				canMove = true
			}

		}

		if canMove { // 已经移动, 选中坐标
			m.Map[x][y] = 1 // 新坐标选中
		}

		return canMove
	}

	// 回退, 先取消选中, 再回退坐标, 回退失败则恢复选中 (上下左右的反向操作, 只有值为 1 (被选中)才能回退过去)
	comeBack := func(direction int) {
		m.Map[x][y] = 0 // 取消选中

		// 回退坐标
		switch direction {
		case 0:
			if x+1 < m.Heigh && m.Map[x+1][y] == 1 { // 向下
				x++
				return
			}

		case 1:
			if x >= 1 && m.Map[x-1][y] == 1 { // 向上
				x--
				return
			}

		case 2:
			if y+1 < m.Width && m.Map[x][y+1] == 1 { // 向右
				y++
				return
			}
		case 3:
			if y >= 1 && m.Map[x][y-1] == 1 { // 向左
				y--
				return
			}
		}

		m.Map[x][y] = 1 // 回退失败, 恢复选中状态
	}

	for k >= 0 {
		track[k]++                            // 选择当前操作
		comeBack(track[k] - 1)                // 如果能回退, 则回退上一步 x,y 坐标
		for track[k] < n && !move(track[k]) { // 进行剪枝操作
			track[k]++
		}

		if track[k] > n-1 { // 回溯
			k--
		} else if x == m.EndX && y == m.EndY { // 找到结果输出一组解
			fmt.Println("solution:")
			m.ShowMap()
		} else {

			k++
			if k >= len(track) { // 有需要则扩容
				track = append(track, -1)
			} else {
				track[k] = -1
			}

		}

	}
}

func main() {
	m := NewMaze([][]int{
		[]int{2, 0, 2, 2, 2, 2, 2, 2},
		[]int{2, 0, 2, 0, 0, 0, 0, 2},
		[]int{2, 0, 0, 0, 2, 2, 0, 2},
		[]int{2, 2, 2, 0, 2, 0, 0, 2},
		[]int{2, 0, 0, 0, 0, 0, 2, 2},
		[]int{2, 0, 2, 0, 2, 0, 0, 2},
		[]int{2, 0, 2, 0, 2, 2, 0, 2},
		[]int{2, 2, 2, 2, 2, 2, 0, 2},
	}, 0, 1, 7, 6)
	fmt.Println("map:")
	m.ShowMap()
	m.FindPath2()
}
