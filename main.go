package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/mpetavy/common"
	"os/exec"
	"strings"
	"time"
)

var (
	committed *bool
)


func init() {
	common.Init("1.0.0", "2020", "Utility for simple GIT interaction", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, true, nil, nil, run, 0)

	committed = flag.Bool("committed",false,"check if current directory containes uncommitted changes")
}

func run() error {
	if *committed {
		cmd := exec.Command("git","diff")

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := common.Watchdog(cmd, time.Millisecond*time.Duration(time.Second))
		if common.Error(err) {
			return err
		}

		output := string(stdout.Bytes())

		if strings.TrimSpace(output) == "" {
			fmt.Printf("Current repo is committed")
			common.Exit(0)
		} else {
			fmt.Printf("Current repo containes uncommitted changes")
			common.Exit(1)
		}

		return nil
	}

	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
