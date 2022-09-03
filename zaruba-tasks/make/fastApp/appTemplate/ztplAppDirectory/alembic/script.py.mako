"""${message}

Revision ID: ${up_revision}
Revises: ${down_revision | comma,n}
Create Date: ${create_date}

"""
from alembic import op
import sqlalchemy as sa
import os
${imports if imports else ""}

# revision identifiers, used by Alembic.
revision = ${repr(up_revision)}
down_revision = ${repr(down_revision)}
branch_labels = ${repr(branch_labels)}
depends_on = ${repr(depends_on)}


def run_migration() -> bool:
    return os.getenv('MIGRATION_RUN_ALL', '0') != '0'


def upgrade() -> None:
    if not run_migration():
        return None
    ${upgrades if upgrades else "pass"}


def downgrade() -> None:
    if not run_migration():
        return None
    ${downgrades if downgrades else "pass"}
