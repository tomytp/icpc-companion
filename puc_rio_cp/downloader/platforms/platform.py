import re
from abc import ABC, abstractmethod
from pathlib import Path

from ..problem_info import ProblemInfo


class Platform(ABC):
    def __init__(self, base_dir: Path):
        self.base_directory = base_dir

    @property
    @abstractmethod
    def URL_PATTERNS(self):
        pass

    @abstractmethod
    def get_info_from_url(self, url: str) -> ProblemInfo:
        """Extract problem information from URL"""
        pass

    def matches_url(self, url: str) -> bool:
        """Check if URL belongs to this platform"""
        return any(re.match(pattern, url) for pattern in self.URL_PATTERNS)

    def _get_problem_directory(self, platform: str, contest_id: str = None) -> Path:
        """Internal method to generate problem directory path"""
        base_path = self.base_directory / platform
        if contest_id:
            return base_path / contest_id
        return base_path