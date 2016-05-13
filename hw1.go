package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	src_json, _ := ioutil.ReadFile("NSYSU.json")

	info1 := map[string]interface{}{}
	json.Unmarshal(src_json, &info1)

	info2 := info1["current_observation"].(map[string]interface{})

	fmt.Println("temp_c : ", info2["temp_c"])
	fmt.Println("icon_url : ", info2["icon_url"])

}
