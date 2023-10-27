package cmd_test

import (
	"example/cmd"
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCmdWithToggleLongArgSet(t *testing.T) {

	root := &cobra.Command{Use: "root", RunE: cmd.RootCmdRunE}
	cmd.RootCmdFlags(root)
	root.SetArgs([]string{"--toggle"})

	err := root.Execute()
	if err != nil {
		t.Fatalf("Error reading running cmd.Execute(cmd.RootCmd()) %v", err)
	}
}

func TestRootCmdWithTShortArgSet(t *testing.T) {

	root := &cobra.Command{Use: "root", RunE: cmd.RootCmdRunE}
	cmd.RootCmdFlags(root)
	root.SetArgs([]string{"-t"})

	err := root.Execute()
	if err != nil {
		t.Fatalf("Error reading running cmd.Execute(cmd.RootCmd()) %v", err)
	}
}

func TestCallingWithArgNoArgs(t *testing.T) {
	root := &cobra.Command{Use: "root", RunE: cmd.RootCmdRunE}
	cmd.RootCmdFlags(root)


	err := root.Execute()

	if err == nil {
		t.Fatalf("Expected calling RootCmd with no args to error; it did not")
	}
}



func TestCallingWithUnregisteredArgs(t *testing.T) {
	root := &cobra.Command{Use: "root", RunE: cmd.RootCmdRunE}
	root.SetArgs([]string{"-k", "--kittehs"})

	err := root.Execute()

	if err == nil {
		t.Fatalf("Expected calling RootCmd with no args to error; it did not")
	}
}
