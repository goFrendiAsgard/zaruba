from typing import List, Optional
from pydantic import BaseModel
import datetime

class ZtplAppCrudEntityData(BaseModel):
    created_at: Optional[datetime.datetime]
    created_by: Optional[str]
    updated_at: Optional[datetime.datetime]
    updated_by: Optional[str]


class ZtplAppCrudEntity(ZtplAppCrudEntityData):
    id: str
    class Config:
        orm_mode = True


class ZtplAppCrudEntityResult(BaseModel):
    count: int
    rows: List[ZtplAppCrudEntity]