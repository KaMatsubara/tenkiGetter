package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/KaMatsubara/tenkiGetter"
	flag "github.com/spf13/pflag"
)

const VERSION = "0.1.3"

func versionString(args []string) string {
	prog := "tenkiGetter"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf("%s version %s", prog, VERSION)
}

func helpMessage(args []string) string {
	prog := "tenkiGetter"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf(`%s tenkiGetter [オプション] <場所>
オプション
-h --help ヘルプを表示
-v --version　バージョンを表示
-d --day　短期天気概況の取得
-w --week　週間天気概況の取得`, prog)
}

type tenkiGetterError struct {
	statusCode int
	message    string
}

func (e tenkiGetterError) Error() string {
	return e.message
}

type flags struct {
	dayFlag     bool
	weekFlag    bool
	helpFlag    bool
	versionFlag bool
}

type runOpts struct {
	config string
}

type options struct {
	runOpt  *runOpts
	flagSet *flags
}

func newOptions() *options {
	return &options{runOpt: &runOpts{}, flagSet: &flags{}}
}

func (opts *options) mode(args []string) tenkiGetter.Mode {
	switch {
	case opts.flagSet.dayFlag:
		return tenkiGetter.Day
	default:
		return tenkiGetter.Week
	}
}

func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := newOptions()
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.BoolVarP(&opts.flagSet.dayFlag, "day", "d", false, "短期天気概況の取得")
	flags.BoolVarP(&opts.flagSet.weekFlag, "week", "w", false, "週間天気概況の取得")
	flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "ヘルプを表示")
	flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "バージョンを表示")
	return opts, flags
}

func performImpl(args []string, executor func(url string) error) *tenkiGetterError {
	for _, url := range args {
		err := executor(url)
		if err != nil {
			return makeError(err, 3)
		}
	}
	return nil
}

func perform(opts *options, args []string) *tenkiGetterError {
	forecast := tenkiGetter.NewForecast()
	config := tenkiGetter.NewConfig(opts.mode(args), args[0])
	if config.OfficeCode == "" {
		return &tenkiGetterError{statusCode: 0, message: "地域が存在しません"}
	}
	switch config.RunMode {
	case tenkiGetter.Day:
		return performImpl(args, func(url string) error {
			return getTenki(forecast, config, url)
		})
	case tenkiGetter.Week:
		return performImpl(args, func(url string) error {
			return getTenki(forecast, config, url)
		})
	}
	return nil
}

func parseOptions(args []string) (*options, []string, *tenkiGetterError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.flagSet.helpFlag {
		fmt.Println(helpMessage(args))
		return nil, nil, &tenkiGetterError{statusCode: 0, message: ""}
	}
	if opts.flagSet.versionFlag {
		fmt.Println(versionString(args))
		return nil, nil, &tenkiGetterError{statusCode: 0, message: ""}
	}
	return opts, flags.Args(), nil
}

func getTenki(forecast *tenkiGetter.Forecast, config *tenkiGetter.Config, url string) error {
	result, err := forecast.GetForecast(config)
	if err != nil {
		return err
	}
	fmt.Println(result.GetData("text"))
	return nil
}

func makeError(err error, status int) *tenkiGetterError {
	if err == nil {
		return nil
	}
	ue, ok := err.(*tenkiGetterError)
	if ok {
		return ue
	}
	return &tenkiGetterError{statusCode: status, message: err.Error()}
}

func goMain(args []string) int {
	opts, args, err := parseOptions(args)
	fmt.Println(args)
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
