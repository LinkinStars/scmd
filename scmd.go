package scmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
	case "stop":
		Stop()
	}
}

func Stop() {
	processName := os.Args[0]
	if len(os.Args) >= 3 && len(os.Args[2]) > 0 {
		processName = os.Args[2]
	}
	processName = filepath.Base(processName)
	fmt.Printf("Try to find process by name [%s]...\n", processName)

	currentPid := os.Getpid()

	processes, err := process.Processes()
	if err != nil {
		panic(err)
	}

	for _, p := range processes {
		// skip current process
		if p.Pid == int32(currentPid) {
			continue
		}
		name, _ := p.Name()
		if !strings.Contains(name, processName) {
			continue
		}

		cmd, _ := p.Cmdline()
		fmt.Printf("Find process [%d][%s][%s], stop it? yes(y)/no(n)\n", p.Pid, name, cmd)

		var input string
		fmt.Scanf("%s", &input)
		if input != "yes" && input != "y" {
			continue
		}

		if err := p.Kill(); err != nil {
			fmt.Println("Stop failed", err.Error())
		} else {
			fmt.Println("The process has been stopped")
		}
	}
	os.Exit(0)
}
