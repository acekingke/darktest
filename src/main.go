package main

import (
	"fmt"
	"flag"
)
const CONN_STR="root:root@tcp(127.0.0.1:3306)/data"
func choose_two_lines(lines_0 []string, lines_1 []string) {
	total := len(lines_0) + len(lines_1)
	choose_1 := len(lines_1)
    db_0 := Connect(CONN_STR)
    db_1 := Connect(CONN_STR)
    defer db_0.Close()
    defer db_1.Close()
	comb(total, choose_1, func(c []int){
		//抽签方式,0号签从第一组选,1号签从第二组选
		tokens := make([]uint8, total)
		for i:=0; i < len(c); i++{
			tokens[c[i]]=1
		}
		var i int = 0
		var j int = 0
		fmt.Println("tokens :", tokens)
		for k:=0; k < total; k++ {
			if tokens[k] == 0 {
				fmt.Println(lines_0[i])
				err:= Execute(db_0, lines_0[i])
				if err != nil {
					panic(err)
				}
				i++
			}else{
				fmt.Println(lines_1[j])
				err:=Execute(db_1, lines_1[j])
				if err != nil {
					panic(err)
				}
				j++
			}
		}
		fmt.Println("-----Over One Case------")
	})
}

func main() {

	fisrt :=flag.String("file0", "sql1.txt", "-file0=sql1.txt")
	second :=flag.String("file1", "sql2.txt", "-file1=sql2.txt")
	lines_0, _ := ScanLines(*fisrt)
	lines_1, _ := ScanLines(*second)
	choose_two_lines(lines_0, lines_1)
}

func comb(n, m int, emit func([]int)) {
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
}
