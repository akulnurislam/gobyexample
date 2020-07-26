package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	boolByte, _ := json.Marshal(true)
	fmt.Println(string(boolByte))

	intByte, _ := json.Marshal(1)
	fmt.Println(string(intByte))

	floatByte, _ := json.Marshal(2.34)
	fmt.Println(string(floatByte))

	stringByte, _ := json.Marshal("gopher")
	fmt.Println(string(stringByte))

	sliceData := []string{"apple", "peach", "pear"}
	sliceByte, _ := json.Marshal(sliceData)
	fmt.Println(string(sliceByte))

	mapData := map[string]int{"apple": 5, "lettuce": 7}
	mapByte, _ := json.Marshal(mapData)
	fmt.Println(string(mapByte))

	res1Data := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1Byte, _ := json.MarshalIndent(res1Data, "", "  ")
	fmt.Println(string(res1Byte))

	res2Data := response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2Byte, _ := json.Marshal(res2Data)
	fmt.Println(string(res2Byte))

	fmt.Println()

	b := []byte(`{"num":6.13,"strs":["a","b"]}`)
	d := make(map[string]interface{})
	if err := json.Unmarshal(b, &d); err != nil {
		panic(err)
	}
	fmt.Println(d)
	num := d["num"].(float64)
	fmt.Println(num)
	strs := d["strs"].([]interface{})
	fmt.Println(strs)
	strs0 := strs[0].(string)
	fmt.Println(strs0)

	fmt.Println()

	s := `{"paGe": 1, "fruits": ["apple", "peach"], "unexpected_field": true}`
	res := response1{}
	json.Unmarshal([]byte(s), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	fmt.Println()

	enc := json.NewEncoder(os.Stdout)
	input := map[string]int{"apple": 5, "lettuce": 7}
	enc.SetIndent("", "  ")
	enc.Encode(input)
}
