package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/akashrtd/mime/pkg/mime"
	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/spf13/cobra"
)

const version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "mime",
	Short: "MIME - Browser automation with MCP support",
	Long:  `MIME (MCP-Integrated Modern Executor) is a high-performance browser automation tool with native Model Context Protocol support.`,
}

var visible bool
var stealth bool

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start MCP server for browser automation",
	Long: `Start the MIME MCP server. This allows AI assistants like Claude to control browsers through the Model Context Protocol.

Configure in Claude Desktop:
{
  "mcpServers": {
    "mime": {
      "command": "/path/to/mime",
      "args": ["serve", "--visible"]
    }
  }
}`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Handle shutdown gracefully
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		go func() {
			<-sigChan
			fmt.Fprintln(os.Stderr, "\nShutting down MIME MCP server...")
			cancel()
		}()

		// Create MCP server
		fmt.Fprintf(os.Stderr, "Starting MIME MCP server (visible=%v)...\n", visible)

		opts := &mime.BrowserOptions{
			Headless: !visible,
		}
		if stealth {
			fmt.Fprintln(os.Stderr, "Stealth mode enabled")
			opts.Stealth = mime.DefaultStealthOptions()
		}

		server, err := mime.NewMCPServer(ctx, opts)
		if err != nil {
			return fmt.Errorf("failed to create MCP server: %w", err)
		}
		defer server.Close()

		fmt.Fprintln(os.Stderr, "MIME MCP server ready. Accepting connections via stdio.")

		// Create StdIO transport and run
		transport := &mcpsdk.StdioTransport{}
		if err := server.Run(transport); err != nil {
			return fmt.Errorf("server error: %w", err)
		}

		return nil
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("MIME v%s\n", version)
		fmt.Println("Browser automation with MCP support")
		fmt.Println("https://github.com/akashrtd/mime")
	},
}

func init() {
	serveCmd.Flags().BoolVar(&visible, "visible", false, "Show browser window (debug mode)")
	serveCmd.Flags().BoolVar(&stealth, "stealth", false, "Enable stealth mode (anti-detection)")
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
