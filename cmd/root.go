package cmd

import (
	"github.com/x0f5c3/md-codedocs/pkg"
	"os"
	"os/signal"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "md-codedocs",
	Short:   "This cli tool will parse your markdown document and extract all the codeblocks in it",
	Version: "v0.0.1", // <---VERSION---> Updating this version, will also create a new GitHub release.
	// Uncomment the following lines if your bare application has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		// Your code here
		n, err := pkg.ParseFile("official_docs.md")
		if err != nil {
			return err
		}
		n.Walk(pkg.TestWalker)
		if len(pkg.CodeNodes) == 0 {
			pterm.Warning.Println("No code nodes found")
		} else {
			pterm.Success.Printfln("Found %d code nodes", len(pkg.CodeNodes))
			pterm.Success.Printfln("Example code node:\n%+v", pkg.CodeNodes[14].Parent)
		}
		if len(pkg.CodeBlockNodes) == 0 {
			pterm.Warning.Println("No code block nodes found")
		} else {
			pterm.Success.Printfln("Found %d code block nodes", len(pkg.CodeNodes))
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Fetch user interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		pterm.Warning.Println("user interrupt")
		pcli.CheckForUpdates()
		os.Exit(0)
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		pcli.CheckForUpdates()
		os.Exit(1)
	}

	pcli.CheckForUpdates()
}

func checkForUpdates() {
	err := pcli.CheckForUpdates()
	if err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Adds global flags for PTerm settings.
	// Fill the empty strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")
	rootCmd.PersistentFlags().BoolVarP(&pcli.DisableUpdateChecking, "disable-update-checks", "", false, "disables update checks")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRepo("x0f5c3/md-codedocs")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
