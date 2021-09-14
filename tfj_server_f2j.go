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
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func Save(path string, save ...string) {
	file := "save.json"
	if save != nil && len(save) >= 1 && save[0] != "" {
		file = save[0]
	}
	err := ioutil.WriteFile(file, code2Json(path), 0666)
	if err != nil {
		log.Panicln(err)
	}
}

func code2Json(path string) []byte {
	tf := codeFile(path)
	jsons, err := json.Marshal(tf)
	if err != nil {
		log.Panicln(err)
	}
	return jsons
}

// codeFile 读取文件
func codeFile(path string) TfJson {
	log.Printf("[SAVE]Path: %s", path)
	tf := NewTfJson()
	stat, err := os.Stat(path)
	if err != nil {
		log.Panicln(err)
	}
	sep := string(os.PathSeparator)
	tf.Name = stat.Name()
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)

	defer func() {
		errf := file.Close()
		if err != nil {
			log.Panicln(errf)
		}
	}()

	if stat.IsDir() {
		// dir
		tf.IsDir = true
		dir, err := file.ReadDir(-1)
		if err != nil {
			log.Panicln(err)
		}
		for _, info := range dir {
			infoPath := path + sep + info.Name()
			tf.DirVal = append(tf.DirVal, codeFile(infoPath))
		}
	} else {
		if err != nil {
			log.Panicln(err)
		}
		all, _ := ioutil.ReadAll(file)
		if err == io.EOF {
			log.Panicln(err)
		}
		tf.Val = all
	}
	return tf
}
