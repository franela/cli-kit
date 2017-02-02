package main

import (
	"log"

	"github.com/spf13/cobra"
)

func DescribeCommands() []*cobra.Command {
	var cmdDeploy = &cobra.Command{
		Use:   "deploy",
		Short: "Deploy shit",
		Long:  `Deploy lots of shit`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Deploying")
		},
	}
	var cmdHealth = &cobra.Command{
		Use:   "health",
		Short: "Health shit",
		Long:  `Health lots of shit`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Health")
		},
	}

	return []*cobra.Command{cmdDeploy, cmdHealth}
}
