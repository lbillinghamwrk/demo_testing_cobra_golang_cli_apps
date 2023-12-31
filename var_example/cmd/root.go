package cmd

import (
  "errors"

  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:  "example",
  RunE: RootCmdRunE,
}

func RootCmdRunE(cmd *cobra.Command, args []string) error {
    t, err := cmd.Flags().GetBool("toggle")

    if err != nil {
      return err
    }

    if t {
      cmd.Println("ok")
      return nil
    }

    return errors.New("not ok")
}

func RootCmdFlags(cmd *cobra.Command) {
  cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Execute() {
  cobra.CheckErr(rootCmd.Execute())
}

func init() {
  RootCmdFlags(rootCmd)
}
