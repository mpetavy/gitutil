package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/mpetavy/common"
)

var (
	committed *bool
	rmdir     *string
)

func init() {
	common.Init(false, "1.0.0", "2020", "Utility for simple GIT interaction", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, nil, run, 0)

	committed = flag.Bool("committed", false, "check if current directory containes uncommitted changes")
	rmdir = flag.String("rmdir", "", "remove directory")
}

func run() error {
	if *committed {
		cmd := exec.Command("git", "diff")

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := common.WatchdogCmd(cmd, time.Millisecond*time.Duration(time.Second))
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

	if *rmdir != "" {
		return os.RemoveAll(*rmdir)
	}

	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
