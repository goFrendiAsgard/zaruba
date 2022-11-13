from typing import List, Optional, Mapping
from pydantic import BaseModel
import datetime


DEFAULT_MARKDOWN_TEMPLATE =  '\n'.join([
    '# {{ title }}',
    '',
    '{{ description }}',
    '',
    '<table class="table">',
    '  <tbody>',
    '  {% for cta in content_type.attributes %}',
    '    <tr>',
    '      <th>{{ cta.caption }}</th>',
    '      <td>{{ attributes[cta.name] if attributes[cta.name] else cta.default_value }}</td>',
    '    </tr>',
    '  {% endfor %}',
    '  </tbody>',
    '</table>'
])


class OptionConfig(BaseModel):
    app_api_url: str = ''
    value_key: str = 'id'
    caption_key: str = 'id'
    option_list: List[str] = []
    option_map: Mapping[str, str] = {}


class ContentTypeAttribute(BaseModel):
    name: str
    caption: str
    default_value: str = ''
    input_type: str = 'markdown'
    option_config: OptionConfig = OptionConfig()


# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class ContentTypeData(BaseModel):
    name: str
    template: str = DEFAULT_MARKDOWN_TEMPLATE
    attributes: List[ContentTypeAttribute] = [ContentTypeAttribute(
        name='content',
        caption='Content',
        default_value='',
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