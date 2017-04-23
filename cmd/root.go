// Package cmd contains CLI commands called through cobra.
package cmd

import (
	"fmt"
	"net/http"

	"github.com/russross/blackfriday"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-learn [-p <port>]",
	Short: "Example Application Server",
	Long: `The example application server will run a local web server for you
to access.`,
	// Running the root command
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetString("port")
		http.HandleFunc("/status", Status)
		http.HandleFunc("/markdown", GenerateMarkdown)
		http.Handle("/", http.FileServer(http.Dir("public")))
		fmt.Printf("Server started under port %v...\n", port)
		http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	},
}

// GenerateMarkdown will convert the markdown from the request and return it as html.
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

// Status will give 'OK' if server is running.
func Status(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "OK")
}
