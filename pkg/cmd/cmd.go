package fluence

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"k8s.io/client-go/tools/clientcmd/api"

	pb "github.com/converged-computing/fluence-kubectl/pkg/fluence/service-grpc"
	"github.com/converged-computing/fluence-kubectl/pkg/graph"
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
	server    string
	port      string
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

// listGroups lists all group names and sizes known by the Fluence scheduler
func (o *FluenceOptions) listGroups() error {
	conn, err := grpc.Dial(o.address(), grpc.WithTransportCredentials(insecure.NewCredentials()))

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
	cmd.Flags().StringVar(&o.port, "port", "4242", "port for GRPC server")
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

	// This would be the name of a group provided by the user
	o.group, err = cmd.Flags().GetString("group")
	if err != nil {
		return err
	}
	o.port, err = cmd.Flags().GetString("port")
	if err != nil {
		return err
	}

	currentContext, exists := o.rawConfig.Contexts[o.rawConfig.CurrentContext]
	if !exists {
		return errNoContext
	}

	// Get the control plane url, and split the port from it
	controlPlane := o.rawConfig.Clusters[o.rawConfig.CurrentContext].Server
	u, err := url.Parse(controlPlane)
	if err != nil {
		return err
	}
	// Often the host doesn't have a port
	if strings.Contains(u.Host, ":") {
		o.server = u.Host[0:strings.LastIndex(u.Host, ":")]
	} else {
		o.server = u.Host
	}
	o.resultingContext = api.NewContext()
	o.resultingContext.Cluster = currentContext.Cluster
	o.resultingContext.AuthInfo = currentContext.AuthInfo
	return nil
}

// Validate ensures that all required arguments and flag values are provided
func (o *FluenceOptions) Validate() error {
	return nil
}

// address returns the grpc address toping
func (o *FluenceOptions) address() string {
	return fmt.Sprintf("%s:%s", o.server, o.port)
}

// listResources lists cluster resources
func (o *FluenceOptions) listResources() error {
	conn, err := grpc.Dial(o.address(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Errorf("[Fluence] Error connecting to server: %v\n", err)
		return err
	}
	defer conn.Close()

	grpcclient := pb.NewExternalPluginServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	request := &pb.ResourceRequest{}
	r, err := grpcclient.GetResources(context.Background(), request)
	if err != nil {
		fmt.Errorf("[Fluence] did not get resources: %v\n", err)
		return err
	}
	// Create new Json Graph
	g := graph.ResourceGraph{}
	err = json.Unmarshal([]byte(r.Graph), &g)
	if err != nil {
		fmt.Errorf("[Fluence] could not read graph into struct: %v\n", err)
		return err
	}
	g.Graph.Summary()
	return nil
}

// Run lists all available namespaces on a user's KUBECONFIG or updates the
// current context based on a provided namespace.
func (o *FluenceOptions) Run() error {

	// Case 1: no groups flag and no group
	if o.listGroupsFlag || o.groupFlag == "" {
		err := o.listResources()
		if err != nil {
			fmt.Printf("There was an error %s", err)
		}
	} else if o.listGroupsFlag {
		fmt.Println("Listing fluence groups")
		o.listGroups()
	} else if o.groupFlag != "" {
		fmt.Printf("Listing group %s\n", o.group)
	}
	return nil
}
