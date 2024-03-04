package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// upload_type='ios'
	// build_id='com.example.xx'
	// path='build/app/outputs/runner.ipa'
	// api_token='sdadadasdas1231'
	// desc='测试上传'
	// version='1.0.0'
	// build_number='1'
	// 从命令行获取upload_type和build_id 以及path api_token

	upload_type := strings.Split(os.Args[1], "=")[1] // ios or android
	build_id := strings.Split(os.Args[2], "=")[1]    // com.example.xx
	path := strings.Split(os.Args[3], "=")[1]        // build/app/outputs/runner.ipa
	api_token := strings.Split(os.Args[4], "=")[1]   // sdadadasdas1231
	// desc := strings.Split(
	GetTokenAndKey(upload_type, build_id, path, api_token)
}

// 使用gin框架 从接口获取token和key

func GetTokenAndKey(upload_type string, buildId string, path string, apiToken string) (string, string) {
	// 从接口获取token和key
	// ...
	fmt.Println(upload_type)
	fmt.Println(buildId)
	fmt.Println(path)
	return "token", "key"
}
