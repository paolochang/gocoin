package db

import (
	"github.com/boltdb/bolt"
	"github.com/paolochang/gocoin/utils"
)

const (
	dbName = "blockchain.db"
	dataBucket = "data"
	blocksBucket = "blocks"
)
var db *bolt.DB

func DB() *bolt.DB {
 if db == nil {
	 // init db
	 dbPoniter, err := bolt.Open(dbName, 0600, nil)
	 db = dbPoniter
	 utils.HandleErr(err)
	 err = db.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
		utils.HandleErr(err)
		_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
		return err
	 })
	 utils.HandleErr(err)
 }
 return db
}