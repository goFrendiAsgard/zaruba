from typing import List, Mapping, Optional
from schemas.activity import Activity, ActivityData

import abc

class ActivityRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> Optional[Activity]:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[Activity]:
        pass

    @abc.abstractmethod
    def count(self, keyword: str) -> int:
        pass

    @abc.abstractmethod
    def insert(self, activity_data: ActivityData) -> Optional[Activity]:
        pass

    @abc.abstractmethod
    def update(self, id: str, activity_data: ActivityData) -> Optional[Activity]:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> Optional[Activity]:
        pass