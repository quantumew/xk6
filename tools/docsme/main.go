// Package main contains a build-time documentation generation tool.
// This tool generates the README.md file based on the CLI help.
package main

import (
	"github.com/quantumew/xk6/internal/cmd"

	"github.com/spf13/cobra"
	"github.com/szkiba/docsme"
)

func main() {
	cobra.EnableCommandSorting = false

	cobra.CheckErr(docsme.For(cmd.New(nil)).Execute())
}
