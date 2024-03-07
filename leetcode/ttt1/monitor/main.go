package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func isAlive(pid string) bool {
	p, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println("bad pid")
		return false
	}
	process, err := os.FindProcess(p)
	if err != nil {
		fmt.Println("can not find")
		return false
	}
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

func main() {
	for {
		informations := ""

		file, err := os.OpenFile("../forever/pid.txt", os.O_RDONLY, 0666)
		if err != nil {
			fmt.Println("open pid file error")
			time.Sleep(5 * time.Second)
			continue
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			ss := strings.Fields(text)
			pid := ss[0]
			info := ss[1]

			if isAlive(pid) {
				informations = informations + fmt.Sprintf("%s %s\n", pid, info)
				continue
			} else {
				pid, err := start(info)
				if err != nil {
					continue
				} else {
					informations = informations + fmt.Sprintf("%d %s\n", pid, info)
				}
			}
		}

		os.WriteFile("../forever/pid.txt", []byte(informations), 0666)
		fmt.Println(informations)
		time.Sleep(5 * time.Second)
	}
}

func start(input string) (int, error) {
	cmd := exec.Command("sleep", "60")
	err := cmd.Start()
	if err != nil {
		fmt.Println("cmd start error", err)
		return -1, err
	}

	pid := cmd.Process.Pid

	fmt.Printf("forever %d restart\n", pid)
	return pid, nil
}
