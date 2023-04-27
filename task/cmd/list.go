package cmd

import (
	"encoding/binary"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("task.db", 0600, nil)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))

			c := b.Cursor()

			fmt.Println("You have the following tasks:")
			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("%d. %s\n", btoi64(k), v)
			}

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func btoi64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}
