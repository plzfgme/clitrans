package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	from          string
	to            string
	translator    string
	preprocessors []string

	rootCmd = &cobra.Command{
		Use:                   "clitrans [flags] [text]",
		DisableFlagsInUseLine: true,
		Short:                 "A command-line translator",
		Long: `A command-line translator.
It translate text from one language to another. When no text is provided, stdin will be used instead.

Currently supported translators: google.
Please refer to each translator's documentation for language codes used in 'from' and 'to' flags.

Currently supported preprocessors: remove_newlines (useful for text copied from a pdf).
    `,
		Run: func(cmd *cobra.Command, args []string) {
			translate, ok := translatorMap[translator]
			if !ok {
				fmt.Fprintf(os.Stderr, "Unknown translator: %s\n", translator)
				return
			}

			var text string
			if len(args) > 0 {
				text = args[0]
			} else {
				if raw, err := io.ReadAll(os.Stdin); err != nil {
					fmt.Println(err.Error())
					return
				} else {
					text = string(raw)
				}
			}

			for _, preprocessor := range preprocessors {
				preprocess, ok := preprocessorMap[preprocessor]
				if !ok {
					fmt.Fprintf(os.Stderr, "Unknown preprocessor: %s\n", preprocessor)
					return
				}

				text = preprocess(text)
			}

			translation, err := translate(text, from, to)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				return
			}

			fmt.Println(translation)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&from, "from", "f", "auto", "The language to translate from, or 'auto'")
	rootCmd.PersistentFlags().StringVarP(&to, "to", "t", "", "The language to translate to")
	rootCmd.MarkPersistentFlagRequired("to")
	rootCmd.PersistentFlags().StringVarP(&translator, "translator", "l", "google", "The translator to use")
	rootCmd.PersistentFlags().StringSliceVarP(&preprocessors, "preprocessors", "p", []string{}, "The preprocessors to use")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
