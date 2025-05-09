from pathlib import Path
from typing import Tuple
from pydantic import BaseModel, Field
from pydantic.dataclasses import dataclass

class TestCase(BaseModel):
    input: str
    output: str

class ProblemInfo(BaseModel):
    platform_name: str    # Platform name (e.g., 'codeforces', 'spoj')
    problem_id: str       # Full problem ID
    folder_path: Path     # Full path to problem folder
    file_name: str        # Name of the solution file (without extension)

    def get_solution_path(self) -> Path:
        """Get path for solution file"""
        return self.folder_path / f'{self.file_name}.cpp'

    def get_test_paths(self, test_num: int) -> Tuple[Path, Path]:
        """Get paths for input and output test files"""
        return (
            self.folder_path / f'in{self.file_name}{test_num}',
            self.folder_path / f'out{self.file_name}{test_num}'
        )

@dataclass
class ParsedProblem:
    name: str
    url: str
    tests: list[TestCase]
    problem_info: ProblemInfo = Field(ProblemInfo, init=False)
