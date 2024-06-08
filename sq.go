package main

import (
	"os"
	"path/filepath"

	"github.com/square/exit"
	"github.com/square/exoskeleton"
)

func main() {
	paths := []string{}
  dir, _ := os.Getwd()
  for dir != "/" {
    path := dir + "/sqbin"
    if _, err := os.Stat(path); err == nil {
      paths = append(paths, path)
    }
    dir = filepath.Dir(dir)
  }

	cli, _ := exoskeleton.New(paths,
      exoskeleton.WithModuleMetadataFilename(".sq-module"))
	cmd, args, _ := cli.Identify(os.Args[1:])
	err := cmd.Exec(cli, args, os.Environ())
	os.Exit(exit.FromError(err))
}
