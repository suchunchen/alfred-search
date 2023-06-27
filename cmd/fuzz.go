package cmd

import (
	"alred-search/structs"
	"encoding/json"
	"fmt"
	"github.com/sahilm/fuzzy"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "fuzz",
		Short: "Fuzzy search your alfred items",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				return
			}

			var alfredItems structs.AlfredItems
			if err := json.Unmarshal([]byte(args[1]), &alfredItems); err != nil {
				os.Stderr.WriteString(fmt.Sprintf("fuzz Unmarshal error: %s", err.Error()))
				return
			}
			results := fuzzy.FindFrom(args[0], alfredItems)

			var searchedItems structs.AlfredItems
			searchedItems.Items = make([]structs.Alfred, 0, len(results))

			for _, result := range results {
				searchedItems.Items = append(searchedItems.Items, alfredItems.Items[result.Index])
			}

			b, _ := json.Marshal(searchedItems)
			fmt.Println(string(b))
			return
		},
	})
}
