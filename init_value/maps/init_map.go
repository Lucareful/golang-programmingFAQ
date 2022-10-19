/*
 * @Author: Luenci
 * @Date: 2022-10-19 11:11:35
 * @LastEditors: Luenci
 * @LastEditTime: 2022-10-19 17:31:19
 * @FilePath: /golang-programmingFAQ/init_value/maps/init_map.go
 * @Description:
 *
 * Copyright (c) 2022 by Luenci, All Rights Reserved.
 */
package main

import "fmt"

func main() {
	mapInit()
	// mapInitErr()
}

func mapInit() {

	demoMap := make(map[string]string)

	// var demoMap = map[string]string{}

	demoMap["a"] = "aa"
	demoMap["b"] = "bb"

	fmt.Printf("%#v\n", demoMap)

}

func mapInitErr() {
	// map 为 nil map不能赋值，未初始化零值
	var demoMap map[string]string

	demoMap["a"] = "aa"
	demoMap["b"] = "bb"

	fmt.Printf("%#v\n", demoMap)

}
