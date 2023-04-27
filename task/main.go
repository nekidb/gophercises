/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/nekidb/gophercises/task/cmd"
)

func main() {
	db, err := bolt.Open("task.db", 0600, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	db.Close()

	cmd.Execute()
}
