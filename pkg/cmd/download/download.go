package download

import (
	"github.com/Benbentwo-Sandbox/BeatSaverDownloader/pkg/cmd/common"
	"github.com/spf13/cobra"
)

// options for the command
type DownloadOptions struct {
	*common.CommonOptions
}

func NewCmdDownload(commonOpts *common.CommonOptions) *cobra.Command {
	options := &DownloadOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use: "download",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}
	options.AddDownloadFlags(cmd)
	// the line below (Section to...) is for the generate-function command to add a template_command to.
	// the line above this and below can be deleted.
	// DO NOT DELETE THE FOLLOWING LINE:
	// Section to add commands to:

	cmd.AddCommand(NewCmdDownloadSongs(commonOpts))
	return cmd
}

// Run implements this command
func (o *DownloadOptions) Run() error {
	return o.Cmd.Help()
}

func (o *DownloadOptions) AddDownloadFlags(cmd *cobra.Command) {
	o.Cmd = cmd
}
