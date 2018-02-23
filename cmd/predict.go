package cmd

import (
	"fmt"

	"github.com/Unknwon/com"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

var (
	modelPath string
)

// predictCmd represents the predict command
var predictCmd = &cobra.Command{
	Use:   "predict",
	Short: "Perform prediction",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pp.Sprintf("fdafsda")
		if !com.IsFile(modelPath) {
			fmt.Sprintf("the file %s does not exist", modelPath)
			return
		}
		// model := fasttext.Open(modelPath)
		query := args[1]
		pp.Println(query)
	},
}

func init() {
	predictCmd.Flags().StringVarP(&modelPath, "model", "m", "", "path to the fasttext model")
	rootCmd.AddCommand(predictCmd)
}
