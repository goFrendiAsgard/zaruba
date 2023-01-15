from typing import Any, List, Optional, Mapping
from pydantic import BaseModel
import datetime


# Note: 🤖 Don't delete the following statement
class ActivityData(BaseModel):
    user_id: str
    activity: str
    object: Optional[str]
    row_id: Optional[str]
    row: Optional[Mapping[str, Any]]
    created_at: Optional[datetime.datetime]  # Note: 🤖 Don't delete this line
    created_by: Optional[str]
    updated_at: Optional[datetime.datetime]
    updated_by: Optional[str]


class Activity(ActivityData):
    id: str

    class Config:
        orm_mode = True


class ActivityResult(BaseModel):
    count: int
    rows: List[Activity]