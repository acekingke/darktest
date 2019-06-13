package main

import (
	"testing"
	"fmt"
)

func TestConnect(t *testing.T) {
	db := Connect();
	fmt.Println(db)
}