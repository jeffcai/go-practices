package hello

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	example = `
	# debug a container in the running pod, the first container will be picked by default
	kubectl debug POD_NAME

	# specify namespace or container
	kubectl debug --namespace foo POD_NAME -c CONTAINER_NAME

	# override the default troubleshooting image
	kubectl debug POD_NAME --image aylei/debug-jvm

	# override entrypoint of debug container
	kubectl debug POD_NAME --image aylei/debug-jvm /bin/bash

	# override the debug config file
	kubectl debug POD_NAME --debug-config ./debug-config.yml
`
	longDesc = `
Run a container in a running pod, this container will join the namespaces of an existing container of the pod.

You may set default configuration such as image and command in the config file, which locates in "~/.kube/debug-config" by default.
`
	defaultImage          = "nicolaka/netshoot:latest"
	defaultAgentPort      = 10027
	defaultConfigLocation = "/.kube/debug-config"
)

// NewDebugCmd returns a cobra command wrapping DebugOptions
func GolangCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "debug POD [-c CONTAINER] -- COMMAND [args...]",
		DisableFlagsInUseLine: true,
		Short:                 "Run a container in a running pod",
		Long:                  longDesc,
		Example:               example,
		Run: func(c *cobra.Command, args []string) {
		},
	}
	//cmd.Flags().BoolVarP(&opts.RetainContainer, "retain", "r", defaultRetain,
	//	fmt.Sprintf("Retain container after debug session closed, default to %s", defaultRetain))
	//	cmd.Flags().StringVar(&opts.Image, "image", "",
	//		fmt.Sprintf("Container Image to run the debug container, default to %s", defaultImage))
	//	cmd.Flags().StringVarP(&opts.ContainerName, "container", "c", "",
	//		"Target container to debug, default to the first container in pod")
	//	cmd.Flags().IntVarP(&opts.AgentPort, "port", "p", 0,
	//		fmt.Sprintf("Agent port for debug cli to connect, default to %d", defaultAgentPort))
	//	cmd.Flags().StringVar(&opts.ConfigLocation, "debug-config", "",
	//		fmt.Sprintf("Debug config file, default to ~%s", defaultConfigLocation))
	//	cmd.Flags().BoolVar(&opts.Fork, "fork", false,
	//		"Fork a new pod for debugging (useful if the pod status is CrashLoopBackoff)")
	//	opts.Flags.AddFlags(cmd.Flags())

	fmt.Println("Golang Cmd ...")
	return cmd
}
