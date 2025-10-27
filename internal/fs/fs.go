package fs

import (
    "os"
    "path/filepath"
)

func EnsureDir(p string) error {
    return os.MkdirAll(p, 0o755)
}

func WriteFile(path, content string) error {
    if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
        return err
    }
    return os.WriteFile(path, []byte(content), 0o644)
}

func CreateTestCases(problemPath, baseName string, tests []struct{Input, Output string}) error {
    inDir := filepath.Join(problemPath, "in")
    outDir := filepath.Join(problemPath, "out")
    if err := os.MkdirAll(inDir, 0o755); err != nil { return err }
    if err := os.MkdirAll(outDir, 0o755); err != nil { return err }
    for i, t := range tests {
        idx := i + 1
        if err := os.WriteFile(filepath.Join(inDir, baseName+itoa(idx)), []byte(t.Input), 0o644); err != nil { return err }
        if err := os.WriteFile(filepath.Join(outDir, baseName+itoa(idx)), []byte(t.Output), 0o644); err != nil { return err }
    }
    return nil
}

func itoa(i int) string { return fmtInt(i) }

// small inline to avoid importing strconv everywhere
func fmtInt(i int) string {
    if i == 0 { return "0" }
    buf := [20]byte{}
    b := len(buf)
    for i > 0 {
        b--
        buf[b] = byte('0' + i%10)
        i /= 10
    }
    return string(buf[b:])
}

