from pydantic import BaseModel
import datetime

class ZarubaEntityNameData(BaseModel):
    pass


class ZarubaEntityName(ZarubaEntityNameData):
    id: str
    created_at: datetime.datetime
    updated_at: datetime.datetime
    class Config:
        orm_mode = True
