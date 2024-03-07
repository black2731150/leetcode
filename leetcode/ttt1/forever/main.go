package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func showProcess() {
	b, err := os.ReadFile("./pid.txt")
	if err != nil {
		fmt.Println("read file error : ", err)
		return
	}
	fmt.Println(string(b))
}

func wirteToFile(ws string) {
	f, err := os.OpenFile("./pid.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file error :", err)
		return
	}
	_, err = f.WriteString(ws)

	fmt.Println("wirte string ", ws, err)
	defer f.Close()
}

func start(input string) {
	cmd := exec.Command("sleep", "60")
	err := cmd.Start()
	if err != nil {
		fmt.Println("cmd start error", err)
		return
	}

	pid := cmd.Process.Pid

	ws := fmt.Sprintf("%d %s\n", pid, input)
	wirteToFile(ws)

	fmt.Printf("forever %d start", pid)
}

func main() {
	show := flag.Bool("show", false, "show all processes")
	flag.Parse()

	if *show {
		showProcess()
		return
	} else {
		args := flag.Args()
		if len(args) != 1 {
			fmt.Println("args error !")
			return
		}

		start(args[0])
	}
}
