package runner

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "sync"
    "time"
)

// DryListen starts a temporary server on :10043, prints raw POST payloads,
// and exits after 5s of inactivity or 3 minutes without any payload.
func DryListen() error {
    var (
        mu              sync.Mutex
        lastProblemTime time.Time
        gotAny          bool
    )

    handler := func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }
        defer r.Body.Close()
        b, err := io.ReadAll(r.Body)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        fmt.Println("--- Received Payload ---")
        var buf bytes.Buffer
        if err := json.Indent(&buf, b, "", "  "); err == nil {
            fmt.Println(buf.String())
        } else {
            // Fallback to raw if not valid JSON
            fmt.Println(string(b))
        }
        fmt.Println("------------------------")
        mu.Lock()
        lastProblemTime = time.Now()
        gotAny = true
        mu.Unlock()
        w.WriteHeader(http.StatusOK)
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)
    srv := &http.Server{Addr: ":10043", Handler: mux, ReadHeaderTimeout: 5 * time.Second}

    fmt.Println("Listening on :10043 (dry mode). Press Ctrl+C to stop.")
    done := make(chan struct{})
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Printf("server error: %v", err)
        }
        close(done)
    }()

    start := time.Now()
    ticker := time.NewTicker(500 * time.Millisecond)
    defer ticker.Stop()
    for {
        <-ticker.C
        mu.Lock()
        have := gotAny
        last := lastProblemTime
        mu.Unlock()
        if !have && time.Since(start) > 3*time.Minute {
            _ = srv.Close()
            <-done
            return nil
        }
        if have && time.Since(last) > 5*time.Second {
            _ = srv.Close()
            <-done
            return nil
        }
    }
}
