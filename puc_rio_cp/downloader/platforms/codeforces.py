import re

from .platform import Platform
from ..problem_info import ProblemInfo


class CodeforcesHandler(Platform):
    URL_PATTERNS = [
        r'.*codeforces\.com/contest/(\d+)/problem/([A-Z]\d*)',
        r'.*codeforces\.com/problemset/problem/(\d+)/([A-Z]\d*)',
        r'.*codeforces\.com/gym/(\d+)/problem/([A-Z]\d*)'
    ]
    
    def get_info_from_url(self, url: str) -> ProblemInfo:
        for pattern in self.URL_PATTERNS:
            if match := re.match(pattern, url):
                contest_id, letter = match.groups()
                folder_path = self._get_problem_directory('codeforces', contest_id)
                return ProblemInfo(
                    platform_name='codeforces',
                    problem_id=f"{contest_id}{letter}",
                    folder_path=folder_path,
                    file_name=letter.lower()
                )
        raise ValueError(f"Invalid Codeforces URL: {url}")