from pydantic import BaseModel
import datetime

class ZtplAppCrudEntityData(BaseModel):
    pass


class ZtplAppCrudEntity(ZtplAppCrudEntityData):
    id: str
    created_at: datetime.datetime
    updated_at: datetime.datetime
    class Config:
        orm_mode = True
