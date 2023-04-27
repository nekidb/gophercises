package cmd

import (
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("task.db", 0600, nil)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		taskNumber, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))

			task := b.Get(itob64(uint64(taskNumber)))
			if task == nil {
				fmt.Printf("There is no such task.\n")
				return nil
			}

			if err := b.Delete(itob64(uint64(taskNumber))); err != nil {
				return fmt.Errorf("delete task: %s", err)
			}

			fmt.Printf("You have completed the \"%s\" task.\n", task)

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
