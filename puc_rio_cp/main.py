import click
import sys, os

from .config import load_config, save_config
from .downloader.problem_downloader import solve
from .tester.test import test

@click.group()
@click.pass_context
def cli(ctx: click.Context) -> None:
    config = load_config()
    ctx.ensure_object(dict)
    ctx.obj["config"] = config



@click.command()
@click.pass_context
def setup(ctx: click.Context) -> None:
    config = ctx.obj['config']
    click.echo(f"Starting setup! This can be changed at any point at: {config.get('config_path')}\n")
    config['base_path'] = input('Base directory to save problems: ')
    template_path = input('Path to cpp template file (leave empty to skip): ')
    if template_path != '':
        config['template_path'] = template_path

    makefile_path = input('Path to makefile template (leave empty to skip): ')
    if makefile_path != '':
        config['makefile_path'] = makefile_path

    save_config(config)

cli.add_command(solve)
cli.add_command(test)
cli.add_command(setup)
