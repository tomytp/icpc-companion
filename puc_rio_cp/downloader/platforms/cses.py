import re

from .platform import Platform
from ..problem_info import ProblemInfo


class CsesHandler(Platform):
    URL_PATTERNS = [r'.*cses\.fi/problemset/task/(\d+)']

    def get_info_from_url(self, url: str) -> ProblemInfo:
        for pattern in self.URL_PATTERNS:
            if match := re.match(pattern, url):
                problem_id = match.group(1)
                folder_path = self._get_problem_directory('cses')
                return ProblemInfo(
                    platform_name='cses',
                    problem_id=problem_id,
                    folder_path=folder_path,
                    file_name=problem_id.lower()
                )
        raise ValueError(f"Invalid CSES URL: {url}")
