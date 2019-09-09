package lib

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	proc "github.com/shirou/gopsutil/process"
)

// createProcess and add it to a pull
func createProcess(commandName string, args []string) int {
	cmd := exec.Command(commandName, args...)
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("Just ran subprocess %d.\n", a.Cyan(cmd.Process.Pid))

	return int(cmd.Process.Pid)
}

// stopProcess returns nil if stop or already stopped
func stopProcess(pid int) error {
	if pid == -1 {
		return nil
	}

	p, err := proc.NewProcess(int32(pid))
	if err != nil {
		log.Fatalf("[ERROR] Trying to get the process PID.\n%v", err)
	}
	return p.Kill()
}

// SimpleCommandInvoker and add it to a pull
func SimpleCommandInvoker(commandName string, args []string) int {
	dateCmd := exec.Command(commandName, args...)
	out, err := dateCmd.Output()

	if err != nil {
		fmt.Printf("%v", string(out))
		log.Fatalf("[ERROR] During the command invoke: '%v %v'\n%v\n", commandName, args, err)
	}

	fmt.Printf("%v", string(out))

	return 0
}
