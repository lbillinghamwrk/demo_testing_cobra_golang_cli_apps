package cmd

import (
  "errors"

  "github.com/spf13/cobra"
)



func RootCmd() *cobra.Command {
  return &cobra.Command{
    Use:  "example",
    RunE: func(cmd *cobra.Command, args []string) error {
      t, err := cmd.Flags().GetBool("toggle")

      if err != nil {
        return err
      }

      if t {
        cmd.Println("ok")
        return nil
      }

      return errors.New("not ok")
    },
  }
}


func Execute(cmd *cobra.Command) error {
  cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
  return cmd.Execute()
}
