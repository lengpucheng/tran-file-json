/*
--------------------------------------------------
@Create 2021-09-14 9:52
@Author lengpucheng
@Program my-loved.cn/tran-file-json
@Describe
--------------------------------------------------
@Version 1.0 2021-09-14
@Memo create this file
*/

package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

var (
	cmd, path, target string
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	file, _ := os.OpenFile("./tfj.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	writer := io.MultiWriter(file, os.Stdout)
	log.SetOutput(writer)
	flag.StringVar(&cmd, "o", "", "执行的命令\n [save]编码保存指定目录或文件\n [load]加载并解码文件到指定目录")
	flag.StringVar(&path, "f", "", "【必填】保存或者加载的文件路径\n 需要保存或加载的文件路径")
	flag.StringVar(&target, "t", "", "保存或加载后的目标文件路径\n 若为保存则会将指定目录文件编码成一个文本文件保存到指定的目录"+
		"\n 若为加载则会将-f指定的文件加载并还原到-t指定的目录")
}

func main() {
	flag.Parse()
	if strings.ToUpper(cmd) == "SAVE" {
		Save(path, target)
	} else if strings.ToUpper(cmd) == "LOAD" {
		Load(path, target)
	} else {
		log.Printf("指定的-o命令有误，请使用 -help 查看用法")
	}
}
