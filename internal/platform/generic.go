package platform

import (
    "errors"
    "net/url"
    "path"
    "regexp"
    "strings"
)

var validName = regexp.MustCompile(`[^a-z0-9_-]`)
var numberSeg = regexp.MustCompile(`^\d+$`)

type GenericResolver struct{}

func (g GenericResolver) Matches(_ string) bool { return true }

func (g GenericResolver) Resolve(basePath, rawURL string) (ProblemInfo, error) {
    u, err := url.Parse(rawURL)
    if err != nil {
        return ProblemInfo{}, err
    }
    host := strings.TrimPrefix(strings.ToLower(u.Hostname()), "www.")
    if host == "" {
        return ProblemInfo{}, errors.New("invalid host")
    }
    segs := strings.FieldsFunc(u.EscapedPath(), func(r rune) bool { return r == '/' })
    var group string
    for _, s := range segs {
        if numberSeg.MatchString(s) {
            group = s
            break
        }
    }
    if group == "" && len(segs) > 0 {
        group = segs[0]
    }
    // filename is last meaningful segment
    var fname string
    if len(segs) > 0 {
        fname = segs[len(segs)-1]
    } else {
        fname = host
    }
    platform := sanitize(host)
    group = sanitize(group)
    fname = sanitize(fname)
    folder := path.Join(basePath, platform)
    if group != "" {
        folder = path.Join(folder, group)
    }
    return ProblemInfo{
        Platform:   platform,
        ProblemID:  fname,
        FolderPath: folder,
        FileName:   strings.ToLower(fname),
    }, nil
}

func sanitize(s string) string {
    s = strings.ToLower(s)
    s = validName.ReplaceAllString(s, "-")
    s = strings.Trim(s, "-")
    if len(s) > 64 {
        s = s[:64]
    }
    if s == "" {
        s = "problem"
    }
    return s
}

