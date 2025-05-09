from pathlib import Path

from ..problem_info import ParsedProblem

from .codeforces import CodeforcesHandler
from .cses import CsesHandler
from .kattis import KattisHandler
from .vjudge import VJudgeHandler
from .vjudge_third_party import VJudgeThirdPartyHandler
from .yosupo import YosupoHandler

HANDLERS = [
    CodeforcesHandler,
    CsesHandler,
    KattisHandler,
    VJudgeHandler,
    VJudgeThirdPartyHandler,
    YosupoHandler,
]

class PlatformManager:
    def __init__(self, base_dir: str):
        self.platforms = [handler_class(Path(base_dir)) for handler_class in HANDLERS]

    def fill_problem_info(self, problem: ParsedProblem) -> None:
        for platform in self.platforms:
            if platform.matches_url(problem.url):
                problem.problem_info = platform.get_info_from_url(problem.url)
                return
        raise ValueError(f"No matching platform found for URL: {problem.url}")