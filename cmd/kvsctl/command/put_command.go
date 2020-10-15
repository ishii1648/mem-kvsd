package command

import (
	// "fmt"
	// "os"
	// "strconv"

	"github.com/spf13/cobra"
	// "go.etcd.io/etcd/v3/clientv3"
)

func NewPutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put [options] <key> <value>",
		Short: "Puts the given key into the store",
		Run:   putCommandFunc,
	}

	return cmd
}

func putCommandFunc(cmd *cobra.Command, args []string) {

}
