package download

import (
	"github.com/Benbentwo-Sandbox/BeatSaverDownloader/pkg/cmd/common"
	"github.com/Benbentwo/utils/util"
	"github.com/spf13/cobra"
)

// options for the command
type SongOptions struct {
	*common.CommonOptions
	batch bool
}

var (
	downloadSongsLong    = `Utiltiy to help download BeatSaber Songs`
	downloadSongsExample = `bsd download songs`
)

func NewCmdDownloadSongs(commonOpts *common.CommonOptions) *cobra.Command {
	options := &SongOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "songs",
		Short:   "Download BeatSaber Songs",
		Long:    downloadSongsLong,
		Example: downloadSongsExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}

	return cmd
}

// Run implements this command
func (o *SongOptions) Run() error {
	util.Logger().Infof("Congratulations generating %s", o.Cmd.Use)
	return nil
}
