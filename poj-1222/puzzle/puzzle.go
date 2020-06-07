package main

import "fmt"

func puzzleLight(lightArr [][]int) [][][]int {
	result := make([][][]int, 0)
	row := len(lightArr)
	col := len(lightArr[0])
	for i:=0; i<(1<<col); i++ {
		// 初始化按钮棋盘
		pressMark := preGenerator(row, col, i)
		if ok := preValidation(pressMark, lightArr); ok {
			var ret [][]int
			for i:=1; i<=row; i++ {
				ret = append(ret, pressMark[i][1:col+1])
			}
			result = append(result, ret)
		}
	}
	return result
}

func preGenerator(row, col, p int) [][]int {
	pressMark := make([][]int, row+1)
	for i:=0; i<len(pressMark); i++ {
		placeHolder := make([]int, col+2)
		pressMark[i] = placeHolder
	}
	i := col
	for ;p>0; p/=2 {
		if p % 2 == 1 {
			pressMark[1][i] = 1
		}
		i --
	}
	return pressMark
}

func preValidation(press, lightArr [][]int) bool {
	// 根据press的第一行和lightArr的状态确定是否能够使得灯全灭
	// 根据press的第一行状态确定整个press矩阵的状态
	row := len(lightArr)
	col := len(lightArr[0])
	for i:=1; i<row; i++ {
		for j:=1; j<=col; j++ {
			press[i+1][j] = (lightArr[i-1][j-1] + press[i][j] + press[i][j-1] + press[i][j+1] + press[i-1][j]) % 2
		}
	}
	// 验证是否能使灯全灭,即判断是否能使最后一行的灯全灭
	for i:=0; i<col; i ++ {
		if (press[row][i+1] + press[row-1][i+1] + press[row][i+2] +
			press[row][i]) % 2 != lightArr[row-1][i] {
			return false
		}
	}
	return true
}

func printRet(ret [][]int) {
	for i:=0; i<len(ret); i++ {
		for j:=0; j<len(ret[0]); j++ {
			fmt.Print(ret[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	puzzle:= [][]int{{0, 1, 1, 0, 1, 0},{1, 0, 0, 1, 1, 1}, {0, 0, 1, 0, 0, 1}, {1, 0, 0, 1, 0, 1}, {0, 1, 1, 1, 0, 0}}
	//puzzle2 := [][]int{{0, 0, 1, 0, 1, 0}, {1, 0, 1, 0, 1, 1}, {0, 0, 1, 0, 1, 1}, {1, 0, 1, 1, 0, 0}, {0, 1, 0, 1, 0, 0}}
	rets := puzzleLight(puzzle)
	for i, ret := range rets {
		fmt.Printf("PUZZLE #%d\n", i+1)
		printRet(ret)
	}
}

