package version_controller

import (
	"fmt"
	"os"
	"os/exec"
)

func UpdateCommitPush() {
	// Command 1: Add all files recursively to git repo
	addCmd := exec.Command("git", "add", "-A")
	addCmd.Stderr = os.Stderr
	addCmd.Stdout = os.Stdout
	err := addCmd.Run()
	if err != nil {
		fmt.Println("Error: Failed to execute git add command")
		os.Exit(1)
	}

	// Prompt the user for a commit message to
	fmt.Print("Enter a commit message: ")
	var commitMsg string
	fmt.Scanln(&commitMsg)

	// Command 2: Commit all changes with the user's provided message
	commitCmd := exec.Command("git", "commit", "-m", commitMsg)
	commitCmd.Stderr = os.Stderr
	commitCmd.Stdout = os.Stdout
	err = commitCmd.Run()
	if err != nil {
		fmt.Println("Error: Failed to execute git commit command")
		os.Exit(1)
	}

	// Command 3: Push to remote (origin master)
	pushCmd := exec.Command("git", "push")
	pushCmd.Stderr = os.Stderr
	pushCmd.Stdout = os.Stdout
	err = pushCmd.Run()
	if err != nil {
		fmt.Println("Error: Failed to execute git push command")
		os.Exit(1)
	}

	fmt.Println("Successfully added, committed, and pushed changes!")
}
