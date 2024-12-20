package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	AppName = "Buher Go"
	Version = "v0.0.1"
)

var (
	root = &cobra.Command{
		Use:     AppName,
		Version: Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			logrus.SetFormatter(&logrus.JSONFormatter{})
			logrus.Info(fmt.Sprintf("Starting Application " + AppName + " With Version " + Version + " ..."))
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			logrus.Info(fmt.Sprintf("Shutting Down Application " + AppName + " With Version " + Version + " ..."))
			return nil
		},
	}
)

func Execute() error {
	root.AddCommand(
		testCommand(),
	)
	return root.Execute()
}
