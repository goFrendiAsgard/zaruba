from typing import List, Optional, Mapping
from pydantic import BaseModel
import datetime

class ContentTypeAttribute(BaseModel):
    name: str
    caption: str
    input_type: str
    options: Mapping[str, str]


# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class ContentTypeData(BaseModel):
    name: str
    template: str
    attributes: List[ContentTypeAttribute] = []
    created_at: Optional[datetime.datetime] # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
    created_by: Optional[str]
    updated_at: Optional[datetime.datetime]
    updated_by: Optional[str]


class ContentType(ContentTypeData):
    id: str
    class Config:
        orm_mode = True


class ContentTypeResult(BaseModel):
    count: int
    rows: List[ContentType]