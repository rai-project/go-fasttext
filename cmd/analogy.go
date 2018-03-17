package cmd

import (
	"fmt"

	"github.com/Unknwon/com"
	"github.com/k0kubun/pp"
	fasttext "github.com/rai-project/go-fasttext"
	"github.com/spf13/cobra"
)

var (
	unsupervisedModelPath string
)

// predictCmd represents the predict command
var analogyCmd = &cobra.Command{
	Use:   "analogy -m [path_to_model]",
	Short: "Perform word analogy on a query using an input model",
	Args:  cobra.ExactArgs(1), // make sure that there is only one argument being passed in
	Run: func(cmd *cobra.Command, args []string) {
		if !com.IsFile(unsupervisedModelPath) {
			fmt.Println("the file %s does not exist", unsupervisedModelPath)
			return
		}

		// create a model object
		model := fasttext.Open(unsupervisedModelPath)
		// close the model at the end
		defer model.Close()
		// perform the prediction
		analogies, err := model.Analogy(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		pp.Println(analogies)
	},
}

func init() {
	analogyCmd.Flags().StringVarP(&unsupervisedModelPath, "model", "m", "", "path to the fasttext model")
	rootCmd.AddCommand(analogyCmd)
}
