from typing import List, Optional
from pydantic import BaseModel
import datetime

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class ContentAttributeData(BaseModel):
    content_id: str
    key: str
    value: str
    created_at: Optional[datetime.datetime] # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
    created_by: Optional[str]
    updated_at: Optional[datetime.datetime]
    updated_by: Optional[str]


class ContentAttribute(ContentAttributeData):
    id: str
    class Config:
        orm_mode = True


class ContentAttributeResult(BaseModel):
    count: int
    rows: List[ContentAttribute]