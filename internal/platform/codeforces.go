package platform

import (
    "fmt"
    "net/url"
    "path"
    "regexp"
    "strings"
)

var (
    cfContestRe   = regexp.MustCompile(`(?i)codeforces\.com/(?:contest|gym)/(\d+)/problem/([A-Z]\d*)`)
    cfProblemsetRe = regexp.MustCompile(`(?i)codeforces\.com/problemset/problem/(\d+)/([A-Z]\d*)`)
)

type CodeforcesResolver struct{}

func (c CodeforcesResolver) Matches(rawURL string) bool {
    u, _ := url.Parse(rawURL)
    if u == nil { return false }
    host := strings.ToLower(u.Hostname())
    return strings.Contains(host, "codeforces.com") || strings.Contains(host, "codeforces")
}

func (c CodeforcesResolver) Resolve(basePath, rawURL string) (ProblemInfo, error) {
    if m := cfContestRe.FindStringSubmatch(rawURL); len(m) == 3 {
        contestID, letter := m[1], strings.ToLower(m[2])
        folder := path.Join(basePath, "codeforces", contestID)
        return ProblemInfo{Platform: "codeforces", ProblemID: contestID+letter, FolderPath: folder, FileName: letter}, nil
    }
    if m := cfProblemsetRe.FindStringSubmatch(rawURL); len(m) == 3 {
        contestID, letter := m[1], strings.ToLower(m[2])
        folder := path.Join(basePath, "codeforces", contestID)
        return ProblemInfo{Platform: "codeforces", ProblemID: contestID+letter, FolderPath: folder, FileName: letter}, nil
    }
    return ProblemInfo{}, fmt.Errorf("unrecognized codeforces url: %s", rawURL)
}

