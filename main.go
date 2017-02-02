package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"plugin"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "foo",
	Short: "Foo is a nice cli",
	Long: `Foo is a nice cli
bla bla bla`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func main() {
	loadPlugins(RootCmd)

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func loadPlugins(cmd *cobra.Command) {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	homedir := u.HomeDir

	// Remove this hack
	homedir = "/home/jonas"

	files, err := ioutil.ReadDir(fmt.Sprintf("%s/.%s/plugins", homedir, cmd.Use))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		p, err := plugin.Open(fmt.Sprintf("%s/.%s/plugins/%s", homedir, cmd.Use, file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		s, err := p.Lookup("DescribeCommands")
		if err != nil {
			log.Fatal(err)
		}
		cmds := s.(func() []*cobra.Command)()
		for _, c := range cmds {
			cmd.AddCommand(c)
		}
	}
}
