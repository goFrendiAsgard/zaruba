from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, DateTime, String
from sqlalchemy.orm import declarative_mixin

import uuid
import datetime

Base = declarative_base()


def generate_primary_key() -> str:
    return str(uuid.uuid4())


@declarative_mixin
class BaseMixin():
    id = Column(
        String(36), primary_key=True, index=True, default=generate_primary_key
    )
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
    created_by = Column(String(36), nullable=True)
    updated_at = Column(DateTime, nullable=True)
    updated_by = Column(String(36), nullable=True)
