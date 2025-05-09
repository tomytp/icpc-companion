import re

from .platform import Platform
from ..problem_info import ProblemInfo


class YosupoHandler(Platform):
    URL_PATTERNS = [r'.*judge.yosupo.jp/problem/(.+)']

    def get_info_from_url(self, url: str) -> ProblemInfo:
        for pattern in self.URL_PATTERNS:
            if match := re.match(pattern, url):
                problem_id = match.group(1)
                folder_path = self._get_problem_directory('yosupo')
                return ProblemInfo(
                    platform_name='yosupo',
                    problem_id=problem_id,
                    folder_path=folder_path,
                    file_name=problem_id.lower()
                )
        raise ValueError(f"Invalid YOSUPO URL: {url}")
