from typing import List, Mapping, Optional
from schema.content_type import ContentType, ContentTypeData

import abc

class ContentTypeRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> Optional[ContentType]:
        pass

    @abc.abstractmethod
    def find_by_name(self, name: str) -> Optional[ContentType]:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[ContentType]:
        pass

    @abc.abstractmethod
    def count(self, keyword: str) -> int:
        pass

    @abc.abstractmethod
    def insert(self, content_type_data: ContentTypeData) -> Optional[ContentType]:
        pass

    @abc.abstractmethod
    def update(self, id: str, content_type_data: ContentTypeData) -> Optional[ContentType]:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> Optional[ContentType]:
        pass