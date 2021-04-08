package dp

import "fmt"

/*
	编辑距离

  输入两个字符串s1和s2，计算出将 s1 转换成 s2 所使用的最小操作数

  你可以对一个字符串进行如下三种操作：
	1. 插入一个字符
	2. 删除一个字符
	3. 替换一个字符

  例如
  	s1="horse" s2="ros" 将 s1 变为 s2
	输出 3
	horse->rorse  h 换为 r
	rorse->rose	删除 r
	rose->ros 删除 e

	s1="intention" s2="execution"
	输出 5
	intention->inention 删除 t
	inention->enention i 替换为 e
	enention->exention n 替换为 x
	exention->exection n 替换为 c
	exection->execution 插入 u
*/

/*
	思路
		只求距离，两个字符串，和 lcs 系列问题相似，都是两个指针，从一边到另一边遍历

		两个指针 i,j 指向 s1 s2 的尾

	三种操作的分析
		删除： i 走一步，操作数 + 1
		替换： i、j 都走一步，操作数 + 1
		插入： j 走一步，操作数 + 1
		相等，什么都不做：i、j 都走一步，操作数不变

	base case
		s2 已经走完，s1 剩下的都删除
		s1 已经走完，s2 剩下的都插入
*/

func editDistanceDp(s1 string, i int, s2 string, j int) int {
	// base case
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}

	if s1[i] == s2[j] {
		// 什么都不做
		return editDistanceDp(s1, i-1, s2, j-1)
	} else {
		resInsert := editDistanceDp(s1, i, s2, j-1) + 1
		resDelete := editDistanceDp(s1, i-1, s2, j) + 1
		resReplace := editDistanceDp(s1, i-1, s2, j-1) + 1

		// 选取三个子问题的最小的
		minRes := resInsert
		if resDelete < minRes {
			minRes = resDelete
		}
		if resReplace < minRes {
			minRes = resReplace
		}

		return minRes
	}
}

func editDistance(s1, s2 string) int {

	return editDistanceDp(s1, len(s1)-1, s2, len(s2)-1)
}

// 使用备忘录
func editDistanceDp2(memo *[][]int, s1 string, i int, s2 string, j int) int {
	// base case
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}

	if (*memo)[i][j] != -1 {
		return (*memo)[i][j]
	}

	if s1[i] == s2[j] {
		// 什么都不做
		(*memo)[i][j] = editDistanceDp(s1, i-1, s2, j-1)
	} else {
		resInsert := editDistanceDp(s1, i, s2, j-1) + 1
		resDelete := editDistanceDp(s1, i-1, s2, j) + 1
		resReplace := editDistanceDp(s1, i-1, s2, j-1) + 1

		// 选取三个子问题的最小的
		minRes := resInsert
		if resDelete < minRes {
			minRes = resDelete
		}
		if resReplace < minRes {
			minRes = resReplace
		}

		(*memo)[i][j] = minRes
	}

	return (*memo)[i][j]
}

func editDistance2(s1, s2 string) int {
	lS1 := len(s1)
	lS2 := len(s2)

	memo := [][]int{}
	for i := 0; i < lS1; i++ {
		arr := []int{}
		for j := 0; j < lS2; j++ {
			arr = append(arr, -1)
		}
		memo = append(memo, arr)
	}
	return editDistanceDp2(&memo, s1, lS1-1, s2, lS2-1)
}

// DP 数组的解法
func editDistance3(s1, s2 string) int {
	lS1 := len(s1)
	lS2 := len(s2)

	dpArr := [][]int{}
	for i := 0; i <= lS1; i++ { // 要存 base case 问题的解，所以要加一
		arr := []int{}
		for j := 0; j <= lS2; j++ {
			arr = append(arr, 0)
		}
		dpArr = append(dpArr, arr)
	}
	// base case
	for i := 1; i <= lS1; i++ {
		dpArr[i][0] = i
	}
	for j := 1; j <= lS2; j++ {
		dpArr[0][j] = j
	}

	for i := 1; i <= lS1; i++ {
		for j := 1; j <= lS2; j++ {
			if s1[i-1] == s2[j-1] { // 从小推出大
				dpArr[i][j] = dpArr[i-1][j-1]
			} else {
				resInsert := dpArr[i][j-1] + 1
				resDelete := dpArr[i-1][j] + 1
				resReplace := dpArr[i-1][j-1] + 1

				// 选取三个子问题的最小的
				minRes := resInsert
				if resDelete < minRes {
					minRes = resDelete
				}
				if resReplace < minRes {
					minRes = resReplace
				}
				dpArr[i][j] = minRes
			}
		}
	}

	// for i := 0; i <= lS1; i++ {
	// 	for j := 0; j <= lS2; j++ {
	// 		fmt.Printf("%d ", dpArr[i][j])
	// 	}
	// 	fmt.Println()
	// }

	return dpArr[lS1][lS2]
}

const (
	opInsert  = 0
	opDelete  = 1
	opReplace = 2
	opSkip    = 3
)

var opStr = map[int]string{
	0: "insert",
	1: "delete",
	2: "replace",
	3: "skip",
}

// 编辑路径节点，用来保存上一次操作节点
type editNode struct {
	char     byte
	charOld  byte // 保存替换前的字符
	op       int
	distance int
	prev     *editNode
}

// 求编辑路径
func editDistancePath(s1, s2 string) ([]string, int) {
	lS1 := len(s1)
	lS2 := len(s2)

	res := []string{}

	dpArr := [][]*editNode{}
	for i := 0; i <= lS1; i++ { // 要存 base case 问题的解，所以要加一
		arr := []*editNode{}
		for j := 0; j <= lS2; j++ {
			arr = append(arr, &editNode{})
		}
		dpArr = append(dpArr, arr)
	}
	// base case
	for i := 1; i <= lS1; i++ {
		dpArr[i][0].distance = i
	}
	for j := 1; j <= lS2; j++ {
		dpArr[0][j].distance = j
	}

	for i := 1; i <= lS1; i++ {
		for j := 1; j <= lS2; j++ {
			currentNode := dpArr[i][j]
			if lS1 > lS2 && i > j { // s1 长，剩下的需要都删除
				deleteNode := dpArr[i-1][j]
				resDelete := deleteNode.distance + 1
				deleteNode.op = opDelete
				deleteNode.char = s1[i-1]
				currentNode.prev = deleteNode
				currentNode.distance = resDelete
				continue
			}

			if lS1 < lS2 && i < j { // s2 长，不够的需要插入
				insertNode := dpArr[i][j-1]
				resInsert := insertNode.distance + 1
				insertNode.op = opInsert
				insertNode.char = s2[j-1]
				currentNode.prev = insertNode
				currentNode.distance = resInsert
				continue
			}

			if s1[i-1] == s2[j-1] { // 从小推出大
				prevNode := dpArr[i-1][j-1]
				currentNode.distance = prevNode.distance
				currentNode.prev = prevNode
				prevNode.char = s1[i-1]
				prevNode.op = opSkip
			} else {
				insertNode := dpArr[i][j-1]
				deleteNode := dpArr[i-1][j]
				replaceNode := dpArr[i-1][j-1]

				resInsert := insertNode.distance + 1
				resDelete := deleteNode.distance + 1
				resReplace := replaceNode.distance + 1

				if s1[i-1] == 'r' && s2[j-1] == 's' {
					fmt.Println(i, j, resInsert, resDelete, resReplace)
				}

				// 选取三个子问题的最小的
				minRes := resInsert
				op := opInsert
				if resDelete < minRes {
					minRes = resDelete
					op = opDelete
				}
				if resReplace < minRes {
					minRes = resReplace
					op = opReplace
				}

				switch op {
				case opInsert:
					insertNode.op = op
					insertNode.char = s2[j-1]
					currentNode.prev = insertNode
				case opDelete:
					deleteNode.op = op
					deleteNode.char = s1[i-1]
					currentNode.prev = deleteNode
				case opReplace:
					replaceNode.op = op
					replaceNode.char = s2[j-1]
					replaceNode.charOld = s1[i-1]
					currentNode.prev = replaceNode
				}

				currentNode.distance = minRes
			}
		}
	}

	lastNode := dpArr[lS1][lS2]
	pN := lastNode.prev
	for pN != nil {
		detail := []string{fmt.Sprintf("%s char %c", opStr[pN.op], pN.char)}
		if pN.charOld != 0 {
			detail = []string{fmt.Sprintf("%s char %c to %c", opStr[pN.op], pN.charOld, pN.char)}
		}
		res = append(detail, res...)
		pN = pN.prev
	}

	for i := 0; i <= lS1; i++ {
		for j := 0; j <= lS2; j++ {
			fmt.Printf("%d-", dpArr[i][j].distance)
		}
		fmt.Println()
	}

	return res, dpArr[lS1][lS2].distance
}
