package main

//一个用来统计英文语句中出现单词数量的程序
import (
	"fmt"
	"strings"
)

func main() {
	var str = "how do you do you do you do"
	var strSlice = strings.Split(str, " ")
	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)
}
