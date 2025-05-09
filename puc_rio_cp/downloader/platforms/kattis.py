import re

from .platform import Platform
from ..problem_info import ProblemInfo


class KattisHandler(Platform):
    URL_PATTERNS = [
        r'.*open\.kattis\.com/problems/([\w-]+)',
        r'.*open\.kattis\.com/contests/\w+/problems/([\w-]+)'
    ]

    def get_info_from_url(self, url: str) -> ProblemInfo:
        for pattern in self.URL_PATTERNS:
            if match := re.match(pattern, url):
                problem_id = match.group(1)
                folder_path = self._get_problem_directory('kattis')
                return ProblemInfo(
                    platform_name='kattis',
                    problem_id=problem_id,
                    folder_path=folder_path,
                    file_name=problem_id.lower()
                )
        raise ValueError(f"Invalid Kattis URL: {url}")