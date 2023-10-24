/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/just-arun/micro-session/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		env, err := cmd.Flags().GetString("environment")
		if err != nil {
			panic("env is not provided")
		}
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			panic("port is not provided")
		}
		ctx, err := cmd.Flags().GetString("environment-context")
		if err != nil {
			panic("environment-context is not provided")
		}
		server.Run(env, ctx, port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().StringP("environment", "e", "dev", "env refers to the file which we call on run time ie if we pass env as `dev` then it will look for `.env.dev.yml` in case of `prod` it is `.env.prod.yml`")
	serveCmd.Flags().StringP("port", "p", ":4200", "provides port for app to run")
	serveCmd.Flags().StringP("environment-context", "c", ".", "environment context to find app configuration path")
}
