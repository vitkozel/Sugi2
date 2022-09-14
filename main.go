package sugi2

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var currentPrefix string = ""
var doDebug bool = false

// convertToString
// converts any to string
func convertToString(data any) string {
	return fmt.Sprintf("%v", data)
}

// setPrefix
// sets the current prefix
func setPrefix(prefix string) {
	currentPrefix = prefix
}

// announce
// announces a message with the current prefix
func announce(message any) {
	fmt.Println(convertToString(currentPrefix) + convertToString(message))
}

// tryDisplayError
// displays an error if it is not nil (but do not exit)
func tryDisplayError(err error) {
	if err != nil {
		announce(err)
	}
}

// tryError
// panics if the error is not nil
func tryError(err error) {
	if err != nil {
		panic(err)
	}
}

// isError
// returns true if the error is not nil
func isError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

// messageDebug
// displays a message if debug mode is enabled
func messageDebug(message string) {
	if doDebug {
		announce("Debug" + message)
	}
}

// setDebug
// sets the debug mode
func setDebug(debug bool) {
	doDebug = debug
}

// getArgs
// returns the arguments of the program
func getArgs() []string {
	return os.Args[1:]
}

// getRuntimeSystem
// returns the operating system in runtime
func getRuntimeSystem() string {
	return runtime.GOOS
}

// openWebInBrowser
// opens a web page in the browser
func openWebInBrowser(url string, operatingSystem string) error {
	var command string
	if operatingSystem == "detect" || operatingSystem == "" {
		operatingSystem = getRuntimeSystem()
	}
	switch operatingSystem {
	case "windows":
		command = "start"
	case "darwin":
		command = "open"
	default:
		command = "xdg-open"
	}
	err := exec.Command(command, url).Start()
	return err
}
