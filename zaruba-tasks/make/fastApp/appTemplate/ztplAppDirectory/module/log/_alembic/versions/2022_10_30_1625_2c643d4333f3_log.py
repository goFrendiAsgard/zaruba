"""20221030162538 log

Revision ID: 2c643d4333f3
Revises: 7dda1641a129
Create Date: 2022-10-30 16:25:38.976399

"""
from alembic import op
import sqlalchemy as sa
import os


# revision identifiers, used by Alembic.
revision = '2c643d4333f3'
down_revision = None
branch_labels = None
depends_on = None



def upgrade() -> None:
    # ### commands auto generated by Alembic - please adjust! ###
    op.create_table('activities',
        sa.Column('id', sa.String(length=36), nullable=False),
        sa.Column('user_id', sa.String(length=255), nullable=False),
        sa.Column('activity', sa.String(length=255), nullable=False),
        sa.Column('object', sa.String(length=255), nullable=True),
        sa.Column('row_id', sa.String(length=255), nullable=True),
        sa.Column('json_row', sa.Text(), nullable=True),
        sa.Column('created_at', sa.DateTime(), nullable=True),
        sa.Column('created_by', sa.String(length=36), nullable=True),
        sa.Column('updated_at', sa.DateTime(), nullable=True),
        sa.Column('updated_by', sa.String(length=36), nullable=True),
        sa.PrimaryKeyConstraint('id')
    )
    op.create_index(op.f('ix_activities_activity'), 'activities', ['activity'], unique=False)
    op.create_index(op.f('ix_activities_id'), 'activities', ['id'], unique=False)
    op.create_index(op.f('ix_activities_object'), 'activities', ['object'], unique=False)
    op.create_index(op.f('ix_activities_row_id'), 'activities', ['row_id'], unique=False)
    op.create_index(op.f('ix_activities_user_id'), 'activities', ['user_id'], unique=False)
    # ### end Alembic commands ###


def downgrade() -> None:
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_index(op.f('ix_activities_user_id'), table_name='activities')
    op.drop_index(op.f('ix_activities_row_id'), table_name='activities')
    op.drop_index(op.f('ix_activities_object'), table_name='activities')
    op.drop_index(op.f('ix_activities_id'), table_name='activities')
    op.drop_index(op.f('ix_activities_activity'), table_name='activities')
    op.drop_table('activities')
    # ### end Alembic commands ###
