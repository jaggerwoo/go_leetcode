package main

import (
	"fmt"
	"time"
)

func main() {
	info := []map[string]interface{}{}
	info = append(info, map[string]interface{}{"id": 1, "name": "name1"})
	info = append(info, map[string]interface{}{"id": 2, "name": "name1"})
	info = append(info, map[string]interface{}{"id": 3, "name": "name1"})

	for i := range info {
		if info[i]["id"] == 1 {
			info[i]["id"] = 6666
			fmt.Println("insert --------")
		}
	}

	fmt.Println("info ---", time.Unix(1574754780, 0).UTC())
}
