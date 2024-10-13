package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type OS struct {
	currentDir string
}

func (o *OS) jobs(...[]string) {

}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (o *OS) StartEventLoop() {
	text := ""
	scanner := bufio.NewScanner(os.Stdin)
	var err error
	var output string
	for text != "\\quit" {
		fmt.Printf("> ")
		scanner.Scan()
		text = scanner.Text()
		cmd := strings.Split(text, " ")
		log.Printf("Scanned text <%s>\n", text)
		log.Println("Splitted text:", cmd)
		switch cmd[0] {
		case "ls":
			output, err = o.Ls(cmd[1:]...)
		case "cd":
			output, err = o.ChangeDirectory(cmd[1:]...)
		case "pwd":
			output, err = o.Pwd(cmd[1:]...)
		case "echo":
			output, err = o.Echo(cmd[1:]...)
		case "kill":
			output, err = o.Kill(cmd[1:]...)
		case "ps":
			output, err = o.Ps(cmd[1:]...)
		default:
			output = "Unknown command..."
		}
		if err != nil {
			output = err.Error()
		}
		fmt.Println(output)
	}
}

func (o *OS) ChangeDirectory(args ...string) (string, error) {
	cmd := exec.Command("cd", args...)
	data, err := cmd.Output()
	return string(data), err
}
func (o *OS) Pwd(args ...string) (string, error) {
	cmd := exec.Command("pwd", args...)
	data, err := cmd.Output()
	return string(data), err
}

func (o *OS) Echo(args ...string) (string, error) {
	cmd := exec.Command("echo", args...)
	data, err := cmd.Output()
	return string(data), err
}
func (o *OS) Kill(args ...string) (string, error) {
	cmd := exec.Command("kill", args...)
	data, err := cmd.Output()
	return string(data), err
}
func (o *OS) Ps(args ...string) (string, error) {
	cmd := exec.Command("ps", args...)
	data, err := cmd.Output()
	return string(data), err
}

func (o *OS) Ls(args ...string) (string, error) {
	cmd := exec.Command("ls", args...)
	data, err := cmd.Output()
	return string(data), err
}

func main() {
	dir, _ := os.Getwd()
	OS := &OS{currentDir: dir}
	log.Println("Starting event loop")
	OS.StartEventLoop()
}
