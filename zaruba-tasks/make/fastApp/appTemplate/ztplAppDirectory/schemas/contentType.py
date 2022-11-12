from typing import List, Optional, Mapping
from pydantic import BaseModel
import datetime

class OptionConfig(BaseModel):
    app_api_url: str = ''
    value_key: str = 'id'
    caption_key: str = 'id'
    option_list: List[str] = []
    option_map: Mapping[str, str] = {}


class ContentTypeAttribute(BaseModel):
    name: str
    caption: str
    input_type: str
    option_config: OptionConfig = OptionConfig()


# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class ContentTypeData(BaseModel):
    name: str
    template: str = '\n'.join([
        '# {{ title }}',
        '{{ description }}'
    ])
    attributes: List[ContentTypeAttribute] = [ContentTypeAttribute(
        name='content',
        caption='Content',
        input_type='markdown'
    )]
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