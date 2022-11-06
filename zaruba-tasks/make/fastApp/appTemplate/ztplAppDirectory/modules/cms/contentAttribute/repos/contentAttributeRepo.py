from typing import List, Mapping, Optional
from schemas.contentAttribute import ContentAttribute, ContentAttributeData

import abc

class ContentAttributeRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> Optional[ContentAttribute]:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[ContentAttribute]:
        pass

    @abc.abstractmethod
    def count(self, keyword: str) -> int:
        pass

    @abc.abstractmethod
    def insert(self, content_attribute_data: ContentAttributeData) -> Optional[ContentAttribute]:
        pass

    @abc.abstractmethod
    def update(self, id: str, content_attribute_data: ContentAttributeData) -> Optional[ContentAttribute]:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> Optional[ContentAttribute]:
        pass