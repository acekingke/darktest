package main

import (
	"testing"
	"fmt"
)

func TestScanLines(t *testing.T) {
	lines, _ := ScanLines("main.go")
	fmt.Println(lines)
	fmt.Println(len(lines))
}
