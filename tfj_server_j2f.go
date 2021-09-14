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
	"io/ioutil"
	"log"
	"os"
)

func Load(file string, path ...string) {
	load := ""
	if path != nil && len(path) >= 1 && path[0] != "" {
		load = path[0]
	}
	unCodeFile(LoadJsons(file), load)
}

func LoadJsons(path string) TfJson {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln(err)
	}
	return json2code(data)
}

func json2code(jsons []byte) TfJson {
	var tf TfJson
	if err := json.Unmarshal(jsons, &tf); err != nil {
		log.Panicln(err)
	}
	return tf
}

// to
func unCodeFile(tf TfJson, path string) {
	sep := string(os.PathSeparator)
	log.Printf("[Load]name : %s,isDir %v", path+tf.Name, tf.IsDir)
	if tf.IsDir {
		// 如果是文件夹
		pathdir := path + sep + tf.Name
		// 不存在就创建
		if _, err := os.Stat(pathdir); err != nil {
			// 不存在就创建
			if os.IsNotExist(err) {
				err = os.Mkdir(pathdir, os.ModePerm)
				if err != nil {
					log.Panicln(err)
				}
			} else {
				log.Panicln(err)
			}
		}
		// 递归
		for _, t := range tf.DirVal {
			unCodeFile(t, pathdir)
		}
	} else {
		var name string
		// 创建文件
		if path != "" {
			name = path + sep + tf.Name
		} else {
			name = tf.Name
		}
		_, err := os.Create(name)
		if err != nil {
			log.Panicln(err)
		}
		err = ioutil.WriteFile(name, tf.Val, os.ModePerm)
		if err != nil {
			log.Panicln(err)
		}
	}

}
