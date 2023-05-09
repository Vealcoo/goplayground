package json

import (
	"encoding/json"
	"fmt"
)

func Run() {
	test := "[1,2,3,4,5,6,7,8,9,0]"
	var in []int64
	json.Unmarshal([]byte(test), &in)
	fmt.Println(in)
}
