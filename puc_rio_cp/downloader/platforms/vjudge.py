import re

from .platform import Platform
from ..problem_info import ProblemInfo


class VJudgeHandler(Platform):
    URL_PATTERNS = [
        r'.*vjudge\.net/contest/(\d+)#problem/(.+)'
    ]

    def get_info_from_url(self, url: str) -> ProblemInfo:
        for pattern in self.URL_PATTERNS:
            if match := re.match(pattern, url):
                contest_id, problem_name = match.groups()
                folder_path = self._get_problem_directory('vjudge', contest_id)
                return ProblemInfo(
                    platform_name='vjudge',
                    problem_id=f"{contest_id}{problem_name}",
                    folder_path=folder_path,
                    file_name=problem_name.lower()
                )
        raise ValueError(f"Invalid Vjudge URL: {url}")