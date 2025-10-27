package server

import (
    "encoding/json"
    "log"
    "net/http"
    "sync"
    "time"

    "github.com/tomytp/icpc-companion/internal/platform"
)

type TimedServer struct {
    Addr              string
    batch             []platform.ParsedProblem
    lastProblemTime   time.Time
    mu                sync.Mutex
    firstProblemWait  time.Duration
    idleAfterLast     time.Duration
}

func NewTimedServer(addr string) *TimedServer {
    return &TimedServer{
        Addr:            addr,
        firstProblemWait: 3 * time.Minute,
        idleAfterLast:    5 * time.Second,
    }
}

func (ts *TimedServer) handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }
    defer r.Body.Close()
    var p platform.ParsedProblem
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        _, _ = w.Write([]byte("invalid payload"))
        return
    }
    ts.mu.Lock()
    ts.batch = append(ts.batch, p)
    ts.lastProblemTime = time.Now()
    ts.mu.Unlock()
    w.WriteHeader(http.StatusOK)
}

func (ts *TimedServer) ServeWithTimeout() []platform.ParsedProblem {
    mux := http.NewServeMux()
    mux.HandleFunc("/", ts.handler)
    srv := &http.Server{Addr: ts.Addr, Handler: mux, ReadHeaderTimeout: 5 * time.Second}

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
        select {
        case <-ticker.C:
            ts.mu.Lock()
            have := len(ts.batch) > 0
            last := ts.lastProblemTime
            ts.mu.Unlock()

            if !have && time.Since(start) > ts.firstProblemWait {
                _ = srv.Close()
                <-done
                return ts.batch
            }
            if have && time.Since(last) > ts.idleAfterLast {
                _ = srv.Close()
                <-done
                return ts.batch
            }
        }
    }
}
