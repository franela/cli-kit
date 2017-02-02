package main

import (
	"C"
	"log"
	_ "net/http"

	"github.com/spf13/cobra"
)

func DescribeCommands() []*cobra.Command {
	var cmdDeploy = &cobra.Command{
		Use:   "mtk",
		Short: "mtk",
		Long:  `Deploy lots of shit`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("dssddsds")
		},
	}

	return []*cobra.Command{cmdDeploy}
}
