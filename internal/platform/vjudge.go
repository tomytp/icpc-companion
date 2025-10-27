package platform

import (
    "fmt"
    "net/url"
    "path"
    "regexp"
    "strings"
)

var vjContestRe = regexp.MustCompile(`(?i)vjudge\.net/contest/(\d+)#problem/(.+)`) 
var vjThirdPartyRe = regexp.MustCompile(`(?i)vjudge\.net/problem/([^-]+)-(.+)`)

type VJudgeResolver struct{}

func (v VJudgeResolver) Matches(rawURL string) bool {
    u, _ := url.Parse(rawURL)
    if u == nil { return false }
    return strings.Contains(strings.ToLower(u.Hostname()), "vjudge")
}

func (v VJudgeResolver) Resolve(basePath, rawURL string) (ProblemInfo, error) {
    if m := vjContestRe.FindStringSubmatch(rawURL); len(m) == 3 {
        contestID, name := m[1], sanitize(m[2])
        folder := path.Join(basePath, "vjudge", contestID)
        return ProblemInfo{Platform: "vjudge", ProblemID: contestID+name, FolderPath: folder, FileName: strings.ToLower(name)}, nil
    }
    if m := vjThirdPartyRe.FindStringSubmatch(rawURL); len(m) == 3 {
        platformName := strings.ToLower(m[1])
        prob := sanitize(m[2])
        // Special case for codeforces-like IDs embedded
        if platformName == "codeforces" {
            // Attempt to split like 1234A or 1234B1
            re := regexp.MustCompile(`(?i)^(\d+)([A-Z]\d*)$`)
            if sub := re.FindStringSubmatch(prob); len(sub) == 3 {
                contestID, letter := sub[1], strings.ToLower(sub[2])
                folder := path.Join(basePath, platformName, contestID)
                return ProblemInfo{Platform: platformName, ProblemID: letter, FolderPath: folder, FileName: letter}, nil
            }
        }
        folder := path.Join(basePath, platformName)
        return ProblemInfo{Platform: platformName, ProblemID: prob, FolderPath: folder, FileName: strings.ToLower(prob)}, nil
    }
    return ProblemInfo{}, fmt.Errorf("unrecognized vjudge url: %s", rawURL)
}

