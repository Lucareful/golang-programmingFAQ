/*
 * @Author: Luenci
 * @Date: 2022-10-19 11:45:43
 * @LastEditors: Luenci
 * @LastEditTime: 2022-10-19 11:45:49
 * @FilePath: /golang-programmingFAQ/pkg/block.go
 * @Description:
 *
 * Copyright (c) 2022 by Luenci, All Rights Reserved.
 */
package pkg

import (
	"os"
	"os/signal"
	"syscall"
)

func Block() {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
}
