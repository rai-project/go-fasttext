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
	Use:   "predict -m [path_to_model] [query]",
	Short: "Perform prediction on a query using an input model",
	Args:  cobra.ExactArgs(1), // make sure that there is only one argument being passed in
	Run: func(cmd *cobra.Command, args []string) {
		if !com.IsFile(modelPath) {
			fmt.Println("the file %s does not exist", modelPath)
			return
		}
		// create a model object
		model := fasttext.Open(modelPath)
		// close the model at the end
		defer model.Close()
		// perform the prediction
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
