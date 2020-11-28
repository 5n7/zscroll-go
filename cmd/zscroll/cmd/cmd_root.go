package cmd

import (
	"fmt"

	"github.com/skmatz/zscroll-go"
	"github.com/spf13/cobra"
)

var rootOptions struct {
	text string

	afterText   string
	beforeText  string
	delay       float64
	length      int
	newLine     bool
	paddingText string
	reverse     bool
	scroll      bool
	timeout     int

	version bool
}

func runRoot(cmd *cobra.Command, args []string) error {
	if rootOptions.version {
		return runVersion(cmd, args)
	}

	scroller, err := zscroll.NewScroller(
		rootOptions.text,
		zscroll.ScrollerOptions{
			AfterText:   rootOptions.afterText,
			BeforeText:  rootOptions.beforeText,
			Delay:       rootOptions.delay,
			Length:      rootOptions.length,
			NewLine:     rootOptions.newLine,
			PaddingText: rootOptions.paddingText,
			Reverse:     rootOptions.reverse,
			Scroll:      rootOptions.scroll,
			Timeout:     rootOptions.timeout,
		},
	)
	if err != nil {
		return err
	}
	return scroller.Run()
}

var rootCmd = &cobra.Command{
	Use:   "zscroll",
	Short: "A text scroller for panels or terminals",
	Long:  "A text scroller for panels or terminals.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires a text to display")
		}
		rootOptions.text = args[0]
		return nil
	},
	RunE: runRoot,
}

func init() { //nolint:gochecknoinits
	rootCmd.Flags().StringVarP(&rootOptions.afterText, "after-text", "a", "", "stationary paddin text to display to the right side of the scroll-text")
	rootCmd.Flags().StringVarP(&rootOptions.beforeText, "before-text", "b", "", "stationary paddin text to display to the left side of the scroll-text")
	rootCmd.Flags().Float64VarP(&rootOptions.delay, "delay", "d", 0.4, "delay in seconds for scrolling update")
	rootCmd.Flags().IntVarP(&rootOptions.length, "length", "l", -1, "length of the scroll-text")
	rootCmd.Flags().BoolVarP(&rootOptions.newLine, "new-line", "n", true, "whether to print a new line after each update")
	rootCmd.Flags().StringVarP(&rootOptions.paddingText, "padding-text", "p", " - ", "padding text to diplay between the end and the head of the scroll-text")
	rootCmd.Flags().BoolVarP(&rootOptions.reverse, "reverse", "r", false, "scroll from left to right")
	rootCmd.Flags().BoolVarP(&rootOptions.scroll, "scroll", "s", true, "whether to scroll")
	rootCmd.Flags().IntVarP(&rootOptions.timeout, "timeout", "t", -1, "time in seconds to exit")

	rootCmd.Flags().BoolVarP(&rootOptions.version, "version", "V", false, "show version")
}

func Execute() {
	rootCmd.Execute() //nolint:errcheck
}
