package scipipe

import (
	// "github.com/go-errors/errors"
	//"os"
	"os"
	"os/exec"
	re "regexp"
)

func ExecCmd(cmd string) string {
	Info.Println("Executing command: ", cmd)
	combOutput, err := exec.Command("bash", "-lc", cmd).CombinedOutput()
	if err != nil {
		Error.Println("Could not execute command `" + cmd + "`: " + string(combOutput))
		os.Exit(1)
	}
	return string(combOutput)
}

func Check(err error, errMsg string) {
	if err != nil {
		Error.Println(errMsg)
		panic(err)
	}
}

// Return the regular expression used to parse the place-holder syntax for in-, out- and
// parameter ports, that can be used to instantiate a SciProcess.
func getShellCommandPlaceHolderRegex() *re.Regexp {
	regex := "{(o|os|i|is|p):([^{}:]+)}"
	r, err := re.Compile(regex)
	Check(err, "Could not compile regex: "+regex)
	return r
}
