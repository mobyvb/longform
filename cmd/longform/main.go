package main

import (
	"github.com/spf13/cobra"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"storj.io/private/cfgstruct"
	"storj.io/private/process"

	"github.com/mobyvb/longform/server"
	"github.com/mobyvb/longform/static"
)

var cfg struct {
	server.Config

	// other config options could go here
}

func main() {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the server",
		RunE:  cmdRun,
	}
	rootCmd := &cobra.Command{
		Use:   "longform",
		Short: "longform server",
	}
	rootCmd.AddCommand(runCmd)
	process.SetHardcodedApplicationName("longform")
	process.Bind(runCmd, &cfg, cfgstruct.DefaultsFlag(runCmd))
	process.Exec(rootCmd)
}

func cmdRun(cmd *cobra.Command, args []string) (err error) {
	log := zap.L()

	server, err := server.New(log, cfg.Config, static.FS)
	if err != nil {
		return errs.New("Error creating server: %+v", err)
	}

	runError := server.Serve()
	closeError := server.Close()

	return errs.Combine(runError, closeError)
}
