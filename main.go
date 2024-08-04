package main

import (
	"fmt"
	"os"
	"strings"
)

type Dockerfile struct {
	Instructions []Instruction
}
type Instruction struct {
	Command string
	Args    []string
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a dockerfile to lint.")
		os.Exit(1)
	}

	dockerfilePath := os.Args[1]

	byteFile, err := os.ReadFile(dockerfilePath)
	if err != nil {
		fmt.Printf("Error reading dockerfile: %w", err)
		os.Exit(1)
	}

	dockerfile, pErr := DockerfileParser(byteFile)
	if pErr != nil {
		fmt.Printf("Error parsing dockerfile %w", pErr)
		os.Exit(1)
	}

	issues := LintDockerfile(dockerfile)

	if len(issues) == 0 {
		fmt.Println("No issues found in Dockerfile.")
	} else {
		fmt.Println("Issues found in Dockerfile:")

		for _, issue := range issues {
			fmt.Println("\t", issue, "\n")
		}
	}
}

func DockerfileParser(content []byte) (Dockerfile, error) {
	lines := strings.Split(string(content), "\n")

	var instructions []Instruction

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		instruction := Instruction{
			Command: parts[0],
			Args:    parts[1:],
		}

		instructions = append(instructions, instruction)
	}

	return Dockerfile{Instructions: instructions}, nil
}

func LintDockerfile(dockerfile Dockerfile) []string {

	var issues []string
	var previousCommand string

	for _, instruction := range dockerfile.Instructions {
		switch instruction.Command {
		case "FROM":
			if len(instruction.Args) < 1 {
				issues = append(issues, "FROM instruction should have at least one argument")
			}

		case "RUN":
			if previousCommand == instruction.Command {
				issues = append(issues, "Multiple consecutive RUN instructions. do this 'RUN download_a_really_big_file && \\ \n \t remove_the_really_big_file'")
			}

			for _, arg := range instruction.Args {
				if arg == "cd" {
					issues = append(issues, "Make use of WORKDIR instead of RUN cd 'some-path'")
				}
			}

		case "ENTRYPOINT":
			if instruction.Args[0][:1] != "[" {
				issues = append(issues, "Use JSON notation for ENTRYPOINT arguments. example 'ENTRYPOINT ['foo', 'run-server']'")
			}

		case "CMD":
			if instruction.Args[0][:1] != "[" {
				issues = append(issues, "Use JSON notation for CMD arguments. example 'CMD ['foo', 'run-server']'")
			}

		default:
			issues = append(issues, fmt.Sprintf("Unknown instruction or rule not provided, %s", instruction.Command))
		}

		previousCommand = instruction.Command
	}

	return issues
}
