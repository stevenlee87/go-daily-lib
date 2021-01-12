package main

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

func main() {
	//拼凑json   body为map数组
	var rbody []map[string]interface{}
	t1 := make(map[string]interface{})
	t1["device_id"] = "t1_id"
	t1["device_hid"] = "t1_hid"

	rbody = append(rbody, t1)
	t2 := make(map[string]interface{})
	t2["device_id"] = "t2_id"
	t2["device_hid"] = "t2_hid"

	rbody = append(rbody, t2)

	cnnJson := make(map[string]interface{})
	cnnJson["code"] = 0
	cnnJson["request_id"] = 123
	cnnJson["code_msg"] = ""
	cnnJson["body"] = rbody
	cnnJson["page"] = 0
	cnnJson["page_size"] = 0

	b, _ := json.Marshal(cnnJson)
	cnnn := string(b)
	fmt.Println("cnnn:", cnnn)
	cn_json, _ := simplejson.NewJson([]byte(cnnn))
	cn_body, _ := cn_json.Get("body").Array()

	for _, di := range cn_body {
		//就在这里对di进行类型判断
		newdi, _ := di.(map[string]interface{})
		device_id := newdi["device_id"]
		device_hid := newdi["device_hid"]
		fmt.Println(device_hid, device_id)
	}

}
