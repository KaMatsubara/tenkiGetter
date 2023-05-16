package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

type options struct {
	day     bool
	week    bool
	help    bool
	version bool
}

func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := &options{}
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.BoolVarP(&opts.day, "day", "d", false, "短期天気概況の取得")
	flags.BoolVarP(&opts.week, "week", "w", false, "週間天気概況の取得")
	flags.BoolVarP(&opts.help, "help", "h", false, "ヘルプを表示")
	flags.BoolVarP(&opts.version, "version", "v", false, "バージョンを表示")
	return opts, flags
}

func perform(opts *options, args []string) *tenkiGetterError {
	fmt.Println("Hello World")
	return nil
}

func parseOptions(args []string) (*options, []string, *tenkiGetterError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.help {
		fmt.Println(helpMessage(args[0]))
		return nil, nil, &tenkiGetterError{statusCode: 0, message: ""}
	}
	return opts, flag.Args(), nil
}
func goMain(args []string) int {
	opts, args, err := parseOptions(args)
	if err != nil {
		if err.statusCode != 0 {
			fmt.Println(err.Error())
		}
		return err.statusCode
	}
	if err := perform(opts, args); err != nil {
		fmt.Println(err.Error())
		return err.statusCode
	}
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
