package cmd

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func testCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Run an interface api test",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			logrus.Info("Start Running Test")
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {

			testCase := ""
			if len(args) > 0 {
				testCase = strings.ToLower(args[0])
			}

			switch testCase {
			default:
				logrus.Info("Running Default Test Case")
			}

		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			logrus.Info("End Running Test")
			return nil
		},
	}
}
