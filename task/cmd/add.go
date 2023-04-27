package cmd

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No task provided.")
			return
		}

		var task string

		for _, arg := range args {
			task += arg
			task += " "
		}
		task = strings.TrimSuffix(task, " ")

		db, err := bolt.Open("task.db", 0600, nil)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))

			// if v := b.Get([]byte(task)); v != nil {
			// 	fmt.Printf("Task \"%s\" already exists!\n", task)
			// 	return nil
			// }
			id, _ := b.NextSequence()

			if err := b.Put([]byte(itob64(id)), []byte(task)); err != nil {
				return fmt.Errorf("add task: %s", err)
			}

			fmt.Printf("Added \"%s\" to your task list.\n", task)

			return nil
		})

		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func itob64(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
