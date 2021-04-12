/*
Copyright © 2021 Andrew Mitchell <andrewpmitchell7@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/mitch292/janitor/internal/files"
	"github.com/mitch292/janitor/internal/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fileName string

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync this environment with the remote file or directory.",
	Long: `Sync this environment with the remote file or directory.
It will look at your .janitor.yaml file and take all source directories or files
and place them in the destination.`,
	Run: cmdRun,
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().StringVarP(&fileName, "file-name", "f", "", "The name of a specific file to sync.")
}

func cmdRun(cmd *cobra.Command, args []string) {
	configFiles := viper.GetStringMap("files")
	internalFileDir := files.NewInternalFileDirectory()

	for name := range configFiles {
		source := viper.GetString(util.GetNestedConfigValueAccessor("files", name, "source"))
		destination := viper.GetString(util.GetNestedConfigValueAccessor("files", name, "destination"))
		janitorFile := files.NewJanitorFile(name, source, destination)

		if err := internalFileDir.Sync(janitorFile); err != nil {
			internalFileDir.WriteErrorToLog(err.Error())
		}
	}
}
