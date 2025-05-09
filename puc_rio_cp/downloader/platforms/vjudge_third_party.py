import re

from .platform import Platform
from ..problem_info import ProblemInfo


class VJudgeThirdPartyHandler(Platform):
    URL_PATTERNS = [r'.*vjudge\.net/problem/([^-]+)-(.+)']

    @staticmethod
    def _parse_codeforces_id(problem_id: str) -> tuple[str, str]:
        """Parse CodeForces problem ID into contest ID and problem letter."""
        match = re.match(r'(\d+)([A-Z]\d*)', problem_id)
        if not match:
            raise ValueError(f"Invalid CodeForces problem ID format: {problem_id}")
        return match.groups()

    def get_info_from_url(self, url: str) -> ProblemInfo:
        """Extract problem information from URL"""
        for pattern in self.URL_PATTERNS:
            if match := re.match(pattern, url):     
                platform, problem_id = match.groups()
                platform = platform.lower()
                
                if platform == 'codeforces':
                    contest_id, letter = self._parse_codeforces_id(problem_id)
                    problem_id = letter
                    folder_path = self._get_problem_directory(platform, contest_id)

                
                folder_path = self._get_problem_directory(platform)
                return ProblemInfo(
                    platform_name=platform,
                    problem_id=problem_id,
                    folder_path=folder_path,
                    file_name=problem_id.lower()
                )
        raise ValueError(f"Invalid VJudge URL: {url}")
