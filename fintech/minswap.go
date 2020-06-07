package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minSwap(arr []int) (ans int) {
	for i := 0; i < len(arr); i += 2 {
		x := arr[i]
		if arr[i+1] == x^1 {
			continue
		}
		for j := i + 2; j < len(arr); j++ {
			if arr[j] == x^1 {
				ans += 1
				arr[i+1], arr[j] = arr[j], arr[i+1]
				break
			}
		}
	}
	return
}

func main() {
	//var n int
	var data []int
	//_, _ = fmt.Scanf("%d", &n)
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	split := strings.Split(string(line), " ")
	for _, numStr := range split {
		num, _ := strconv.Atoi(numStr)
		data = append(data, num)
	}
	count := minSwap(data)
	fmt.Println(count)
}
