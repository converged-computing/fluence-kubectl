package main

import (
	"os"

	"github.com/spf13/pflag"

	fluence "github.com/converged-computing/fluence-kubectl/pkg/cmd"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-fluence", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := fluence.NewCmdFluence(genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
