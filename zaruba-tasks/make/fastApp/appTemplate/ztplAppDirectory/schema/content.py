from typing import List, Optional, Mapping
from pydantic import BaseModel
from schema.content_type import ContentType
import datetime

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class ContentData(BaseModel):
    title: str
    attributes: Optional[Mapping[str, str]] = {}
    content_type_id: str
    description: str
    created_at: Optional[datetime.datetime] # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
    created_by: Optional[str]
    updated_at: Optional[datetime.datetime]
    updated_by: Optional[str]


class Content(ContentData):
    id: str
    content_type: Optional[ContentType]
    html_content: Optional[str]
    class Config:
        orm_mode = True


class ContentResult(BaseModel):
    count: int
    rows: List[Content]