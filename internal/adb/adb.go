package adb

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func PixelUploading() bool {
	cmd := exec.Command("adb", "shell", "dumpsys notification | grep -E 'Uploading|Syncing|Backing up'")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		if strings.Contains(err.Error(), "exit status 1") {
			return false
		}
		fmt.Println(err)
	}
	return len(strings.TrimSpace(out.String())) > 0
}

func ClearPixel(targetDir string) {
	fmt.Println("Clearing Pixel")
	cmd := exec.Command("adb", "shell", fmt.Sprintf("rm -rf %s*", targetDir))
	err := cmd.Run()
	if err != nil {
		fmt.Println("wipe error", err)
	}
}

func PushToPixel(localPath, remoteDir string) {
	cmd := exec.Command("adb", "push", localPath, remoteDir)
	err := cmd.Run()
	if err != nil {
		fmt.Println("push error", err)
	}
}
