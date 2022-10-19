/*
 * @Author: Luenci
 * @Date: 2022-10-19 11:49:42
 * @LastEditors: Luenci
 * @LastEditTime: 2022-10-19 16:18:49
 * @FilePath: /golang-programmingFAQ/init_value/channels/buffer_chan/init_bufferchan.go
 * @Description:
 *
 * Copyright (c) 2022 by Luenci, All Rights Reserved.
 */
package main

import (
	"fmt"

	"github.com/Lucareful/golang-programmingFAQ/pkg"
)

func main() {
	bufferChan()
	pkg.Block()
}

func bufferChan() {
	ch := make(chan int, 10)

	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()

	ch <- 0
	ch <- 1

	// 用进程读取
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
}

func bufferChanErr() {
	ch := make(chan int, 10)

	ch <- 0
	ch <- 1

	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()
}
