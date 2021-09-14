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

type TfJson struct {
	Name string
	Val  []byte
	// 是否是文件夹
	IsDir bool
	// 如果是文件夹则内
	DirVal []TfJson
}

// NewTfJson 实例化
func NewTfJson() TfJson {
	t := TfJson{}
	return t
}
