package main

import (
	"encoding/json"
	"fmt"
	"orm"
)

func main() {
	defer test()
	var t1 = orm.Model{Id: 1}
	fmt.Printf("Hello t1's Id: %d\r\n", t1.Id)
	var t2 = t1
	t3 := &t1
	t1.Id = 2
	fmt.Printf("Hello t2's Id:  %d\r\n", t2.Id)
	fmt.Printf("Hello t3's Id:  %d\r\n", t3.Id)

	str, err := json.Marshal(t1)
	if err == nil {
		fmt.Printf("The t1 json : %s\r\n", str)
	}
}

func test() {
	fmt.Printf("the fun end \r\n")
}
