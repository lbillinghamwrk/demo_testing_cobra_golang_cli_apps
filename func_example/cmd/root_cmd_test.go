package cmd_test

import (
  "example/cmd"
  "testing"
)

func TestCallingWithShortTFlag(t *testing.T) {
  root := cmd.RootCmd()
  root.SetArgs([]string{"-t"})

  err := cmd.Execute(root)

  if err != nil {
		t.Fatalf("Error reading running cmd.Execute(cmd.RootCmd()) %v", err)
	}
}

func TestCallingWithLongToggleArg(t *testing.T) {
  root := cmd.RootCmd()
  root.SetArgs([]string{"--toggle"})

  err := cmd.Execute(root)

  if err != nil {
		t.Fatalf("Error reading running cmd.Execute(cmd.RootCmd()) %v", err)
	}
}

func TestCallingWithArgNoArgs(t *testing.T) {
  root := cmd.RootCmd()

  err := cmd.Execute(root)

  if err == nil {
    t.Fatalf("Expected calling RootCmd with no args to error; it did not")
  }
}

func TestCallingWithUnregisteredArgs(t *testing.T) {
  root := cmd.RootCmd()
  root.SetArgs([]string{"-q", "--wombles"})

  err := cmd.Execute(root)

  if err == nil {
    t.Fatalf("Expected calling RootCmd with no args to error; it did not")
  }
}
