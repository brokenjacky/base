package database

import (
    "fmt"
    "github.com/boltdb/bolt"
)

func InitBolt(dbName string, bucket string) (*bolt.DB, error) {
    db, err := bolt.Open(dbName, 0600, nil)
    if err != nil {
        return nil, err
    }

    err = db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte(bucket))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        return nil
    })
    if err != nil {
        return nil, err
    }
    return db, nil
}
