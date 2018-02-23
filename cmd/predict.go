package cmd

import (
	"fmt"

	"github.com/Unknwon/com"
	"github.com/k0kubun/pp"
	fasttext "github.com/rai-project/go-fasttext"
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
		if !com.IsFile(modelPath) {
			fmt.Println("the file %s does not exist", modelPath)
			return
		}
		model := fasttext.Open(modelPath)
		defer model.Close()
		preds, err := model.Predict(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		pp.Println(preds)
	},
}

func init() {
	predictCmd.Flags().StringVarP(&modelPath, "model", "m", "", "path to the fasttext model")
	rootCmd.AddCommand(predictCmd)
}
