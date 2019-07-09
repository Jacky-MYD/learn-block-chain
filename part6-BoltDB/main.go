package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

// 数据库名字
const dbFile = "blockchain.db"

// 仓库--blocks（表名）
const blocksBucket = "blocks"

func main()  {
	// ---------------数据库创建-------------
	//如果数据库存在，则打开；如果数据库不存在，则创建一个数据库
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// 插入更新数据库
	err = db.Update(func(tx *bolt.Tx) error {
		// 判断这一张表是否存在于数据库中
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")

			// CreateBucket 创建blocks表
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			// 储存数据
			// key []byte, value []byte
			err = b.Put([]byte("Jacky"), []byte("https://github.com/Jacky-MYD")) // 存储一条数据
			if err != nil {
				log.Panic(err)
			}
		}

		return  nil
	})

	if err != nil {
		log.Panic(err)
	}
}