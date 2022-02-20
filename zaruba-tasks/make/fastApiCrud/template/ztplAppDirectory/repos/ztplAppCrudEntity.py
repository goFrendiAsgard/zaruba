from typing import List, Mapping, Optional
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData

import abc
import json
import uuid
import datetime

class ZtplAppCrudEntityRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> Optional[ZtplAppCrudEntity]:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[ZtplAppCrudEntity]:
        pass

    @abc.abstractmethod
    def insert(self, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> Optional[ZtplAppCrudEntity]:
        pass

    @abc.abstractmethod
    def update(self, id: str, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> Optional[ZtplAppCrudEntity]:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> Optional[ZtplAppCrudEntity]:
        pass


class MemZtplAppCrudEntityRepo(ZtplAppCrudEntityRepo):

    def __init__(self):
        self._ztpl_app_crud_entity_map: Mapping[str, ZtplAppCrudEntity] = {}

    def set_storage(self, ztpl_app_crud_entity_map: Mapping[str, ZtplAppCrudEntity]):
        self._ztpl_app_crud_entity_map = ztpl_app_crud_entity_map

    def find_by_id(self, id: str) -> Optional[ZtplAppCrudEntity]:
        if id not in self._ztpl_app_crud_entity_map:
            return None
        return self._ztpl_app_crud_entity_map[id]

    def find(self, keyword: str, limit: int, offset: int) -> List[ZtplAppCrudEntity]:
        mem_ztpl_app_crud_entities = list(self._ztpl_app_crud_entity_map.values())
        ztpl_app_crud_entities: List[ZtplAppCrudEntity] = []
        for index in range(offset, limit+offset):
            if index >= len(mem_ztpl_app_crud_entities):
                break
            mem_ztpl_app_crud_entity = mem_ztpl_app_crud_entities[index]
            ztpl_app_crud_entities.append(mem_ztpl_app_crud_entity)
        return ztpl_app_crud_entities

    def insert(self, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> Optional[ZtplAppCrudEntity]:
        if id not in self._ztpl_app_crud_entity_map:
            return None
        new_ztpl_app_crud_entity_id = str(uuid.uuid4())
        new_ztpl_app_crud_entity = ZtplAppCrudEntity(
            id=new_ztpl_app_crud_entity_id,
            created_at=datetime.datetime.now(),
            updated_at=datetime.datetime.now()
        )
        self._ztpl_app_crud_entity_map[new_ztpl_app_crud_entity_id] = new_ztpl_app_crud_entity
        return new_ztpl_app_crud_entity

    def update(self, id: str, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> Optional[ZtplAppCrudEntity]:
        if id not in self._ztpl_app_crud_entity_map:
            return None
        mem_ztpl_app_crud_entity = self._ztpl_app_crud_entity_map[id]
        mem_ztpl_app_crud_entity.updated_at = datetime.datetime.now()
        self._ztpl_app_crud_entity_map[id] = mem_ztpl_app_crud_entity
        return mem_ztpl_app_crud_entity

    def delete(self, id: str) -> Optional[ZtplAppCrudEntity]:
        if id not in self._ztpl_app_crud_entity_map:
            return None
        mem_ztpl_app_crud_entity = self._ztpl_app_crud_entity_map.pop(id)
        return mem_ztpl_app_crud_entity