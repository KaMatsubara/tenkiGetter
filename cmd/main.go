package tenkiGetter

import (
	"fmt"
	"os"
	"path/filepath"

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

type options struct {
	day     bool
	week    bool
	help    bool
	version bool
}

func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := &options{}
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
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
		fmt.Println(helpMessage(args))
		return nil, nil, &tenkiGetterError{statusCode: 0, message: ""}
	}
	if opts.version {
		fmt.Println(versionString(args))
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
