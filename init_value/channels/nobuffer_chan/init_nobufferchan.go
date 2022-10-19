/*
 * @Author: Luenci
 * @Date: 2022-10-19 11:15:45
 * @LastEditors: Luenci
 * @LastEditTime: 2022-10-19 16:47:10
 * @FilePath: /golang-programmingFAQ/init_value/channels/nobuffer_chan/init_nobufferchan.go
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

	noBufferChanErr1()
	noBufferChan()

	pkg.Block()

}

func noBufferChan() {
	// 无缓冲 chan ，接收 和发送必须同时（单独开go程）
	ch := make(chan int)

	go func() {
		{
			ch <- 0
		}
	}()
	go func() {
		{
			ch <- 1
		}
	}()

	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()

}

func noBufferChanErr1() {

	ch := make(chan int)

	ch <- 0
	ch <- 1

	for val := range ch {
		fmt.Println(val)
	}
	/*
		output:
				fatal error: all goroutines are asleep - deadlock!

			goroutine 1 [chan send]:
			main.noBufferChan()
					/Users/lusheng/study/golang-programmingFAQ/init_value/init_chan.go:23 +0x3c
			main.main()
					/Users/lusheng/study/golang-programmingFAQ/init_value/init_chan.go:17 +0x20
			exit status 2
	*/
}
