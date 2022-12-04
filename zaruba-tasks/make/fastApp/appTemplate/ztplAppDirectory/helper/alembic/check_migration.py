from typing import List
from alembic.command import revision
from alembic.config import Config
from alembic.operations.ops import MigrationScript

import itertools
import sys
import os


def simulate_autogenerate(config_path: str, name: str) -> List[tuple]:
    """Simulate the `alembic revision --autogenerate` command
    and return a list of generated operations.
    Inspired from: https://github.com/4Catalyzer/alembic-autogen-check
    """
    config = Config(config_path)
    revisions: List[MigrationScript] = []

    def process_revision_directives(context, revision, directives):
        nonlocal revisions
        revisions = list(directives)
        # Prevent actually generating a migration
        directives[:] = []

    os.environ['IS_GENERATING_MIGRATION'] =  '1'
    revision(
        config=config,
        autogenerate=True,
        process_revision_directives=process_revision_directives,
        version_path=config.get_section_option(name, 'script_location', '_alembic') if name is not None else None
    )
    for script in revisions:
        print('SCRIPT', script)
        for op in script.upgrade_ops_list:
            print('OP', op)
            print(op.as_diffs())
    return list(
        itertools.chain.from_iterable(
            op.as_diffs()
            for script in revisions
            for op in script.upgrade_ops_list
        )
    )


if __name__ == '__main__':
    '''
    throw error if no diff
    '''
    config_path = sys.argv[1] if len(sys.argv) > 1 else './alembic.ini'
    name = sys.argv[2] if len(sys.argv) > 2 else None
    diff = simulate_autogenerate(config_path, name)
    if len(diff) > 0:
        raise Exception('There are changes')