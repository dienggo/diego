package helper

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// CaptureStdout not thread safe
func CaptureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// GetModuleName : getter module name from app
func GetModuleName() string {
	// Run the "go list -m" command as a sub-process
	cmd := exec.Command("go", "list", "-m")

	// Capture the standard output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// Convert the output to a string
	return strings.TrimSpace(string(output))
}
