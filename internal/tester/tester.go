package tester

import (
    "bufio"
    "errors"
    "fmt"
    "io/fs"
    "os"
    "os/exec"
    "path/filepath"
    "sort"
    "strings"
)

const (
    green = "\u001b[0;32m"
    red   = "\u001b[0;31m"
    cyan  = "\u001b[0;36m"
    yellow= "\u001b[0;33m"
    nc    = "\u001b[0m"
)

func Run(debug bool) error {
    latest, err := latestCpp(".")
    if err != nil {
        return err
    }
    target := strings.TrimSuffix(filepath.Base(latest), ".cpp")
    makeArgs := []string{"make", target}
    if debug {
        makeArgs = append(makeArgs, "CPPFLAGS=\"-DDEBUG\"")
    }
    cmd := exec.Command(makeArgs[0], makeArgs[1:]...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("make failed: %w", err)
    }

    ins, outs, err := testsFor(target)
    if err != nil { return err }
    for i := 0; i < len(ins); i++ {
        if i >= len(outs) {
            fmt.Printf("%sTest %d: Expected output not found%s\n", yellow, i+1, nc)
            if out, _ := runBinary("./"+target, ins[i]); out != "" {
                fmt.Println(cyan+"Got:"+nc)
                fmt.Print(out)
            }
            continue
        }
        out, _ := runBinary("./"+target, ins[i])
        expected, _ := os.ReadFile(outs[i])
        ao := trimTrailing(string(out))
        eo := trimTrailing(string(expected))
        if ao == eo {
            fmt.Printf("%sTest %d: Passed%s\n", green, i+1, nc)
        } else {
            fmt.Printf("%sTest %d: Failed%s\n", red, i+1, nc)
            fmt.Println(cyan+"Expected:"+nc)
            fmt.Println(eo)
            fmt.Println(cyan+"Got:"+nc)
            fmt.Println(ao)
        }
    }
    _ = os.Remove(target)
    return nil
}

func latestCpp(dir string) (string, error) {
    var files []fs.DirEntry
    entries, err := os.ReadDir(dir)
    if err != nil { return "", err }
    for _, e := range entries {
        if e.IsDir() { continue }
        name := e.Name()
        if strings.HasSuffix(name, ".cpp") { files = append(files, e) }
    }
    if len(files) == 0 { return "", errors.New("no .cpp files found in current directory") }
    sort.Slice(files, func(i, j int) bool {
        fi, _ := files[i].Info()
        fj, _ := files[j].Info()
        return fi.ModTime().After(fj.ModTime())
    })
    return files[0].Name(), nil
}

func testsFor(base string) ([]string, []string, error) {
    inDir := filepath.Join("in")
    outDir := filepath.Join("out")
    var ins, outs []string

    entries, _ := os.ReadDir(inDir)
    for _, e := range entries {
        if e.IsDir() { continue }
        name := e.Name()
        if strings.HasPrefix(name, base) {
            ins = append(ins, filepath.Join(inDir, name))
        }
    }
    entries, _ = os.ReadDir(outDir)
    for _, e := range entries {
        if e.IsDir() { continue }
        name := e.Name()
        if strings.HasPrefix(name, base) {
            outs = append(outs, filepath.Join(outDir, name))
        }
    }
    sort.Strings(ins)
    sort.Strings(outs)
    return ins, outs, nil
}

func runBinary(bin, inputPath string) (string, error) {
    f, err := os.Open(inputPath)
    if err != nil { return "", err }
    defer f.Close()
    cmd := exec.Command(bin)
    cmd.Stdin = bufio.NewReader(f)
    out, err := cmd.Output()
    return string(out), err
}

func trimTrailing(s string) string {
    s = strings.ReplaceAll(s, "\r\n", "\n")
    s = strings.TrimRight(s, "\n")
    return s
}

