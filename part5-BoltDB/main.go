package main

import (
	"github.com/boltdb/bolt"
	"log"
)

// 数据库名字
const dbFile = "blockchain.db"

// 仓库
const blocksBucket = "blocks"

func main()  {
	// ---------------数据库创建-------------
	//如果数据库存在，则打开；如果数据库不存在，则创建一个数据库
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
}