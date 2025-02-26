package cmd

import (
	"net/http"
	"os"
	"time"

	"github.com/ajabep/icsFilter/internal"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Args: cobra.ExactArgs(1),

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		verbosity, err := cmd.Flags().GetCount("verbose")
		if err != nil {
			return err
		}

		zerolog.TimeFieldFormat = time.DateTime
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        zerolog.SyncWriter(os.Stderr),
			TimeFormat: time.DateTime,
		}).With().Timestamp().Logger()

		// Default level for this example is info, unless debug flag is present
		switch verbosity {
		case 0:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case 1:
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Stack().Logger()
		default:
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Stack().Logger()
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ruleFilePath := args[0]
		rf := internal.RulesFile{}
		err := rf.Load(ruleFilePath)
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to parse the rule file.")
		}

		rf.InitHttp()
		err = http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to start the HTTP server.")
		}
		log.Warn().Msg("The HTTP server is stopped.")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().CountP("verbose", "v", "Increase verbosity")
}
