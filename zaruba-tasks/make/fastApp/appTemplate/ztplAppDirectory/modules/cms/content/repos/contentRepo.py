from typing import List, Mapping, Optional
from schemas.content import Content, ContentData

import abc

class ContentRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> Optional[Content]:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[Content]:
        pass

    @abc.abstractmethod
    def count(self, keyword: str) -> int:
        pass

    @abc.abstractmethod
    def insert(self, content_data: ContentData) -> Optional[Content]:
        pass

    @abc.abstractmethod
    def update(self, id: str, content_data: ContentData) -> Optional[Content]:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> Optional[Content]:
        pass