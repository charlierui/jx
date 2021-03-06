package cmd

import (
	"github.com/spf13/cobra"

	"github.com/pkg/browser"
)

const (
	docsURL = "https://jenkins-x.io/documentation/"
)

/* open the docs - Jenkins X docs by default */
func NewCmdDocs(commonOpts *CommonOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docs",
		Short: "Open the documentation in a browser",
		Run: func(cmd *cobra.Command, args []string) {
			err := browser.OpenURL(docsURL)
			CheckErr(err)
		},
	}
	return cmd
}
