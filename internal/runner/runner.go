package runner

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/tomytp/icpc-companion/internal/config"
    fsutil "github.com/tomytp/icpc-companion/internal/fs"
    "github.com/tomytp/icpc-companion/internal/platform"
    "github.com/tomytp/icpc-companion/internal/server"
    "github.com/tomytp/icpc-companion/internal/tester"
    "github.com/tomytp/icpc-companion/internal/util"
)

func Setup(cfgPath string) error {
    cfg, _ := config.Load(cfgPath)
    in := bufio.NewReader(os.Stdin)
    fmt.Printf("Starting setup! This can be changed at any point at: %s\n\n", cfg.ConfigPath)
    fmt.Print("Base directory to save problems: ")
    base, _ := in.ReadString('\n')
    base = strings.TrimSpace(base)
    if base == "" {
        return errors.New("base path is required")
    }
    fmt.Print("Path to cpp template file (leave empty to skip): ")
    tpl, _ := in.ReadString('\n')
    tpl = strings.TrimSpace(tpl)
    fmt.Print("Path to makefile template (leave empty to skip): ")
    mk, _ := in.ReadString('\n')
    mk = strings.TrimSpace(mk)

    cfg.BasePath = base
    if tpl != "" { cfg.TemplatePath = tpl }
    if mk != "" { cfg.MakefilePath = mk }
    return config.Save(cfg)
}

func Solve() error {
    cfg, _ := config.Load("")
    if cfg.BasePath == "" {
        return errors.New("configuration not found! Run `comp setup` to create one")
    }

    srv := server.NewTimedServer(":10043")
    fmt.Println("Waiting for problems on port 10043")
    batch := srv.ServeWithTimeout()

    if len(batch) == 0 {
        fmt.Println("\nNo contest directory created.")
        return nil
    }

    mgr := platform.NewManager()
    var firstFolder string
    // Load optional templates
    var tplCode, mkCode string
    if cfg.TemplatePath != "" {
        if b, err := os.ReadFile(cfg.TemplatePath); err == nil { tplCode = string(b) }
    }
    if cfg.MakefilePath != "" {
        if b, err := os.ReadFile(cfg.MakefilePath); err == nil { mkCode = string(b) }
    }

    for _, p := range batch {
        info, err := mgr.Resolve(cfg.BasePath, p.URL)
        if err != nil {
            fmt.Println("resolve error:", err)
            continue
        }
        if err := fsutil.EnsureDir(info.FolderPath); err != nil {
            fmt.Println("mkdir error:", err)
            continue
        }
        if firstFolder == "" { firstFolder = info.FolderPath }
        // Solution template
        sol := filepath.Join(info.FolderPath, info.FileName+".cpp")
        if _, err := os.Stat(sol); errors.Is(err, os.ErrNotExist) {
            _ = fsutil.WriteFile(sol, tplCode)
        }
        // Makefile
        if mkCode != "" {
            _ = fsutil.WriteFile(filepath.Join(info.FolderPath, "makefile"), mkCode)
        }
        // Tests
        tests := make([]struct{Input, Output string}, len(p.Tests))
        for i, t := range p.Tests {
            tests[i].Input, tests[i].Output = t.Input, t.Output
        }
        if len(tests) > 0 {
            _ = fsutil.CreateTestCases(info.FolderPath, info.FileName, tests)
        }
    }

    if firstFolder != "" {
        fmt.Printf("\nOpening VS Code in directory: %s\n", firstFolder)
        util.OpenVSCode(firstFolder)
    }
    return nil
}

func Test(debug bool) error {
    return tester.Run(debug)
}
