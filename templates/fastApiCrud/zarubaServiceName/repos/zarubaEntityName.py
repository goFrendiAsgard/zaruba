from typing import List
from schemas.zarubaEntityName import ZarubaEntityName, ZarubaEntityNameData

import abc

class ZarubaEntityNameRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> ZarubaEntityName:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[ZarubaEntityName]:
        pass

    @abc.abstractmethod
    def insert(self, zaruba_entity_name_data: ZarubaEntityNameData) -> ZarubaEntityName:
        pass

    @abc.abstractmethod
    def update(self, id: str, zaruba_entity_name_data: ZarubaEntityNameData) -> ZarubaEntityName:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> ZarubaEntityName:
        pass