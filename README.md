# go-blockchain

## part1: 创世区块及新增block 
## part2: pow工作量证明，模拟挖矿
## part3、part4: 区块的序列化和反序列
    安装库：go get github.com/boltdb/bolt/...
    参考链接：https://github.com/boltdb/bolt
## part5: **boltdb**创建数据库
    bolt.Open(数据库名, 0600, nil)
    参考链接：https://github.com/boltdb/bolt
    引入："github.com/boltdb/bolt"
    // 如果数据库存在，则打开；如果数据库不存在，则创建一个数据库
    bolt.Open(数据库名, 0600, nil)
    blockchain.db生成路径会在根目录，因此我们可以通过终端命令进入当前文件路径执行
    go build -o "文件名" main.go，此时生成一个"文件名"的文件，再输入"./文件名"将生成blockchain.db文件
    例如："go build -o bc main.go", "./bc"
## part6: **boltdb**插入或更新数据库
    bolt.Update(func(tx *bolt.Tx) error) {}
    
   
    