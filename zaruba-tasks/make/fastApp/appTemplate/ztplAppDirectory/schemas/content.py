from typing import List, Optional, Mapping
from pydantic import BaseModel
import datetime

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class ContentData(BaseModel):
    title: str
    attributes: Optional[Mapping[str, str]] = {}
    type_id: str
    description: str
    created_at: Optional[datetime.datetime] # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
    created_by: Optional[str]
    updated_at: Optional[datetime.datetime]
    updated_by: Optional[str]


class Content(ContentData):
    id: str
    class Config:
        orm_mode = True


class ContentResult(BaseModel):
    count: int
    rows: List[Content]