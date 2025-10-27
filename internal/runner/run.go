package runner

import (
    "errors"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
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
    // Prefer sanitizer-friendly output if the binary was built with -fsanitize.
    // These env vars are harmless if ASAN/UBSAN aren’t linked.
    {
        env := os.Environ()
        if sym, err := exec.LookPath("llvm-symbolizer"); err == nil {
            env = append(env, "ASAN_SYMBOLIZER_PATH="+sym, "MSAN_SYMBOLIZER_PATH="+sym)
        }
        // Force colored, verbose sanitizer output to stderr.
        env = append(env,
            "ASAN_OPTIONS=color=always:abort_on_error=1:detect_leaks=1:strict_string_checks=1",
            "UBSAN_OPTIONS=print_stacktrace=1:color=always",
        )
        bin.Env = env
    }
    runErr := bin.Run()
    // Detect segfault to improve messaging
    if ee, ok := runErr.(*exec.ExitError); ok {
        if status, ok2 := ee.Sys().(interface{ Signaled() bool; Signal() os.Signal }); ok2 {
            if status.Signaled() && status.Signal().String() == "segmentation fault" || strings.Contains(strings.ToLower(status.Signal().String()), "segv") {
                fmt.Printf("\u001b[0;31mSegmentation fault detected.\u001b[0m\n")
                if !debug {
                    fmt.Println("Tip: run with -d to include debug symbols for better traces.")
                }
                if runtime.GOOS == "darwin" {
                    fmt.Println("Tip (macOS): install lldb and run: lldb -- ./" + target + " then use 'run' and 'bt'.")
                }
            }
        }
    }
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
