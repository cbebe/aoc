package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bitfield/script"
)

func main() {
	input, err := script.File("input.txt").String()
	if err != nil {
		log.Fatal(err)
	}

	dis := 0
	dep := 0
	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			continue
		}
		arr := strings.Split(v, " ")
		num, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatal(err)
		}
		if arr[0][0] == 'f' {
			dis += num
		} else if arr[0][0] == 'd' {
			dep += num
		} else if arr[0][0] == 'u' {
			dep -= num
		}
	}

	fmt.Println(dis * dep)
}
