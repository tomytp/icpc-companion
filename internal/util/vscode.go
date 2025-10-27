package util

import (
    "os/exec"
)

func OpenVSCode(path string) {
    if path == "" { return }
    if _, err := exec.LookPath("code"); err != nil {
        return
    }
    cmd := exec.Command("code", "--reuse-window", path)
    _ = cmd.Start() // run asynchronously; ignore errors
}

