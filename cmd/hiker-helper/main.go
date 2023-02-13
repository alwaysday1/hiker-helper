// Copyright 2023 Innkeeper lighttree <alwaysday1@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/alwaysday1/hiker-helper.

package main

import (
	_ "go.uber.org/automaxprocs"
	"os"

	"github.com/alwaysday1/hiker-helper/internal/hikerhelper"
)

func main() {
	//fmt.Println("hello world")
	command := hikerhelper.NewHikerHelperCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
