/*
 * @Author: Luenci
 * @Date: 2022-10-19 11:16:09
 * @LastEditors: Luenci
 * @LastEditTime: 2022-10-19 17:34:30
 * @FilePath: /golang-programmingFAQ/init_value/init_other.go
 * @Description:
 *
 * Copyright (c) 2022 by Luenci, All Rights Reserved.
 */
package main

import "fmt"

func main() {

	var slc []string

	slc = append(slc, "luenci")

	fmt.Printf("%#v\n", slc)

}
