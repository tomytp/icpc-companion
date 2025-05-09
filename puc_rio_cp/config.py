import json
from json import JSONDecodeError
from pathlib import Path
from typing import Any

import click

CONFIG_DIR = Path.home() / ".config" / "puc_rio_cp"
CONFIG_FILE = CONFIG_DIR / "config.json"


def load_config(config_path: Path = CONFIG_FILE):
    default_cfg = {'config_path': str(config_path)}
    if not config_path.exists():
        return default_cfg
    with open(config_path, "r") as file:
        try:
            return json.load(file)
        except JSONDecodeError:
            click.echo(f'Invalid Configuration File! Check {config_path} or run `comp setup` to generate a new one.')
            return default_cfg


def save_config(config: dict[str, Any]):
    if not CONFIG_DIR.exists():
        CONFIG_DIR.mkdir(parents=True, exist_ok=True)
    with open(CONFIG_FILE, "w") as file:
        json.dump(config, file, indent=4)