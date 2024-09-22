/* 测试一下json的格式化 */
package model_test

import (
	"encoding/json"
	"fmt"
	"os"
	. "patchouli_lib/model"
	"testing"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

// json使用简单示例
func TestJsonSimply(t *testing.T) {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	fmt.Println(bolB)
}

// 各种基本类型的json化
func TestJson(t *testing.T) {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

// Game结构的json化
func TestGameEncode(t *testing.T) {
	// ns:=[]string{'1sda','2sda'}
	g := NewGame([]string{"reimu", "marisa"}, nil, nil)
	ge, err := json.Marshal(g)
	// ge, err := json.MarshalIndent(g, " ", "")

	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(string(ge))
	fmt.Println(ge)

}

// 能不能将Game结构实例当做匿名键值对，似乎不能
// 想要给结构体添加额外的字段，只能先用json编码再转为字典，给字典添加额外字段
// 然后再对字典编码，
func TestGameToMap(t *testing.T) {
	game := NewGame([]string{"reimu"}, nil, nil)

	// 将结构体编码为 JSON 字节流
	jsonBytes, err := json.Marshal(game)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	// 解析 JSON 字节流为 map[string]interface{}
	var dataMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &dataMap)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	// 向 map 中添加新的键值对
	dataMap["yes"] = true

	// 将更新后的 map 编码为 JSON
	updatedJson, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println("Error marshaling updated JSON:", err)
		return
	}

	fmt.Println(string(updatedJson))
}

// 也可以定义一个新的以Game为匿名结构体的结构体，对其json编码
func TestGameInfoJson(t *testing.T) {
	g := NewGame([]string{"reimu"}, nil, nil)
	gi := GameInfo{Game: g, LocalPlayer: 0}
	jsonBytes, err := json.Marshal(gi)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
