package kubectlexecutor

import (
	"bytes"
	"io"
	"log"
	"os/exec"
	"pain/pin8s/asker"
	"strings"
)

type ResultLine struct {
	Selected string
	Names    []string
}

func (c *commander) SelectConfigContext() {

	cmd := exec.Command("kubectl", "config", "get-contexts")
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatalf("RunCommand: cmd.Output(): %v", err)
	}

	buf := make([]byte, 1024)
	reader := bytes.NewReader(stdout)
	n, err := reader.Read(buf)

	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			log.Fatal("an error occurred while reading bytes")
		}
	}

	text := strings.TrimSpace(string(buf[:n]))

	var cleanLine []string
	var suscriptionNames []string
	//fields := []ResultLine{}
	resultLine := ResultLine{}

	for {
		n := strings.IndexAny(text, "\n")
		if n == -1 {
			if len(text) > 0 {
				n = len(text)
			} else {
				break
			}
		}

		line := text[:n]

		s := []byte(line)
		if string(s[0]) != "*" && len(cleanLine) > 0 {
			s[0] = 'X'
			line = string(s)
		}

		lineCols := strings.Fields(line)

		//esultLine.Selected = false
		if lineCols[0] == "*" {
			resultLine.Selected = lineCols[1]
		}

		if len(cleanLine) > 0 {
			suscriptionNames = append(suscriptionNames, lineCols[1])
		}

		cleanLine = append(cleanLine, line)

		if n == len(text) {
			break
		}

		text = text[n+1:]
	}

	resultLine.Names = suscriptionNames

	answer := asker.ShowContextsSurvey(resultLine.Names, resultLine.Selected)
	executeSelectContext(answer)

}

func executeSelectContext(context string) {
	cmd := exec.Command("kubectl", "config", "use-context", context)
	err := cmd.Run()
	if err != nil {
		log.Fatal("an error occurred while selecting context")
	}
}
