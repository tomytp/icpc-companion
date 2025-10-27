package platform

import (
    "fmt"
    "net/url"
    "path"
    "regexp"
    "strings"
)

var atcoderRe = regexp.MustCompile(`(?i)atcoder\.jp/contests/([^/]+)/tasks/([^/?#]+)`) // contest and task id

type AtCoderResolver struct{}

func (a AtCoderResolver) Matches(rawURL string) bool {
    u, _ := url.Parse(rawURL)
    if u == nil { return false }
    return strings.Contains(strings.ToLower(u.Hostname()), "atcoder.jp")
}

func (a AtCoderResolver) Resolve(basePath, rawURL string) (ProblemInfo, error) {
    if m := atcoderRe.FindStringSubmatch(rawURL); len(m) == 3 {
        contest := strings.ToLower(m[1])
        task := strings.ToLower(m[2])
        // Filename: just the suffix after last underscore, e.g., abc299_a -> a, arc100_b -> b, ex stays ex
        letter := task
        if idx := strings.LastIndex(task, "_"); idx != -1 && idx+1 < len(task) {
            letter = task[idx+1:]
        }
        // Fallback safety
        if letter == "" { letter = task }
        folder := path.Join(basePath, "atcoder", contest)
        return ProblemInfo{Platform: "atcoder", ProblemID: letter, FolderPath: folder, FileName: letter}, nil
    }
    return ProblemInfo{}, fmt.Errorf("unrecognized atcoder url: %s", rawURL)
}

