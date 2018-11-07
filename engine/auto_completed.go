package engine

import (
	"log"
	"strings"
)

/**
* @DESCRIPTION
* @AUTHOR didi
* @CREATE 2018/11/6.
*
 */

func (engine *Engine) suggestion(key string) []string {
	if !engine.initialized {
		log.Fatal("请初始化引擎")
	}
	res := make([]string, 10)
	var i = 0
	if engine.initOptions.UsePersistentStorage {
		for _, db := range engine.dbs {
			db[getDB("index")].ForEach(func(k, v []byte) error {
				str := string(k)
				if strings.HasPrefix(str, key) {
					res[i] = str
					i = i + 1
				}
				return nil
			})
		}
	}
	if len(res) >= 10 {
		return res[0:10]
	}
	return res
}
