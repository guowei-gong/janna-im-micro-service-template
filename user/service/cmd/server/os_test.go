package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_OS(t *testing.T) {
	stat, err := os.Stat("../../configs")
	if err != nil {

	}

	fmt.Println(stat)
}