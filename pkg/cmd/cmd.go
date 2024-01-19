/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fluence

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"k8s.io/client-go/tools/clientcmd/api"

	//	"github.com/flux-framework/flux-k8s/flux-plugin/fluence"
	pb "github.com/converged-computing/fluence-kubectl/pkg/fluence/fluxcli-grpc"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

var (
	fluenceExample = `
	# view the current resource graph of your cluster
	%[1]s fluence

	# view all groups in the graph
	%[1]s fluence --groups

	# view details for a group
	%[1]s fluence --group <name>
`

	errNoContext = fmt.Errorf("no context is currently set, use %q to select a new one", "kubectl config use-context <context>")
)

// FluenceOptions provides information to interact with the fluence scheduler
type FluenceOptions struct {
	configFlags *genericclioptions.ConfigFlags

	resultingContext     *api.Context
	resultingContextName string

	group                string
	userSpecifiedContext string

	// listGroupsFlag
	listGroupsFlag bool
	groupFlag      string

	rawConfig api.Config
	args      []string

	genericiooptions.IOStreams
}

// NewNamespaceOptions provides an instance of NamespaceOptions with default values
func NewFluenceOptions(streams genericiooptions.IOStreams) *FluenceOptions {
	return &FluenceOptions{
		configFlags: genericclioptions.NewConfigFlags(true),

		IOStreams: streams,
	}
}

func (o *FluenceOptions) listGroups() error {

	// NOTE this is deprecated see
	// https://pkg.go.dev/google.golang.org/grpc@v1.60.1#WithInsecure
	conn, err := grpc.Dial("127.0.0.1:4242", grpc.WithInsecure())

	if err != nil {
		fmt.Errorf("[Fluence] Error connecting to server: %v\n", err)
		return err
	}
	defer conn.Close()

	grpcclient := pb.NewExternalPluginServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	request := &pb.GroupRequest{Group: ""}
	r, err := grpcclient.ListGroups(context.Background(), request)
	if err != nil {
		fmt.Errorf("[Fluence] did not a group: %v\n", err)
		return err
	}
	fmt.Println(r)
	return nil
}

// NewCmdFluence provides a cobra command wrapping FluenceOptions
func NewCmdFluence(streams genericiooptions.IOStreams) *cobra.Command {
	o := NewFluenceOptions(streams)

	cmd := &cobra.Command{
		Use:          "fluence [subcommand] [flags]",
		Short:        "View fluence resources or groups",
		Example:      fmt.Sprintf(fluenceExample, "kubectl"),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}
	cmd.Flags().BoolVar(&o.listGroupsFlag, "groups", o.listGroupsFlag, "if true, print groups and sizes known by fluence")
	cmd.Flags().StringVar(&o.groupFlag, "group", o.groupFlag, "get a specific group")
	o.configFlags.AddFlags(cmd.Flags())
	return cmd
}

// Complete sets all information required for updating the current context
func (o *FluenceOptions) Complete(cmd *cobra.Command, args []string) error {
	o.args = args

	var err error
	o.rawConfig, err = o.configFlags.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return err
	}

	o.group, err = cmd.Flags().GetString("group")
	if err != nil {
		return err
	}

	currentContext, exists := o.rawConfig.Contexts[o.rawConfig.CurrentContext]
	if !exists {
		return errNoContext
	}

	o.resultingContext = api.NewContext()
	o.resultingContext.Cluster = currentContext.Cluster
	o.resultingContext.AuthInfo = currentContext.AuthInfo
	return nil
}

// Validate ensures that all required arguments and flag values are provided
func (o *FluenceOptions) Validate() error {
	fmt.Println("Executing Validate()")
	return nil
}

// Run lists all available namespaces on a user's KUBECONFIG or updates the
// current context based on a provided namespace.
func (o *FluenceOptions) Run() error {

	// If we were asked to list groups (or nothing)
	if o.listGroupsFlag || o.groupFlag == "" {
		fmt.Println("Listing fluence groups")
		o.listGroups()
	}
	return nil
}
