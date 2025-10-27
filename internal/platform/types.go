package platform

type TestCase struct {
    Input  string `json:"input"`
    Output string `json:"output"`
}

type ParsedProblem struct {
    Name  string     `json:"name"`
    URL   string     `json:"url"`
    Tests []TestCase `json:"tests"`
}

type ProblemInfo struct {
    Platform   string
    ProblemID  string
    FolderPath string
    FileName   string
}

type Resolver interface {
    Matches(url string) bool
    Resolve(basePath, url string) (ProblemInfo, error)
}

