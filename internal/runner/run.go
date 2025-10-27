package runner

import (
    "errors"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "sort"
    "strings"
)

// RunInteractive compiles the latest .cpp using make and runs the produced
// binary with stdin/stdout attached, suitable for manual input or shell redirection.
func RunInteractive(debug bool) error {
    latest, err := latestCpp(".")
    if err != nil {
        return err
    }
    target := strings.TrimSuffix(filepath.Base(latest), ".cpp")
    makeArgs := []string{"make", target}
    if debug {
        makeArgs = append(makeArgs, "CPPFLAGS=\"-DDEBUG\"")
    }
    mk := exec.Command(makeArgs[0], makeArgs[1:]...)
    mk.Stdout = os.Stdout
    mk.Stderr = os.Stderr
    if err := mk.Run(); err != nil {
        return fmt.Errorf("make failed: %w", err)
    }

    bin := exec.Command("./" + target)
    bin.Stdin = os.Stdin
    bin.Stdout = os.Stdout
    bin.Stderr = os.Stderr
    runErr := bin.Run()
    _ = os.Remove(target)
    return runErr
}

// latestCpp is copied from tester with minimal deps to avoid circular imports.
func latestCpp(dir string) (string, error) {
    entries, err := os.ReadDir(dir)
    if err != nil { return "", err }
    var files []string
    for _, e := range entries {
        if e.IsDir() { continue }
        name := e.Name()
        if strings.HasSuffix(name, ".cpp") { files = append(files, name) }
    }
    if len(files) == 0 { return "", errors.New("no .cpp files found in current directory") }
    sort.Slice(files, func(i, j int) bool {
        fi, _ := os.Stat(files[i])
        fj, _ := os.Stat(files[j])
        return fi.ModTime().After(fj.ModTime())
    })
    return files[0], nil
}

