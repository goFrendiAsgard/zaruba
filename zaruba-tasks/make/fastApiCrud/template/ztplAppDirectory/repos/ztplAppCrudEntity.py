from typing import List
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData

import abc

class ZtplAppCrudEntityRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> ZtplAppCrudEntity:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[ZtplAppCrudEntity]:
        pass

    @abc.abstractmethod
    def insert(self, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> ZtplAppCrudEntity:
        pass

    @abc.abstractmethod
    def update(self, id: str, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> ZtplAppCrudEntity:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> ZtplAppCrudEntity:
        pass