from typing import Any, List, Optional, Mapping
from sqlalchemy import Column, String, Text
from sqlalchemy.orm import InstrumentedAttribute
from schema.role import Role, RoleData
from module.auth.role.repo.role_repo import RoleRepo
from repo import Base, BaseMixin, DBRepo

import jsons


class DBRoleEntity(Base, BaseMixin):
    __tablename__ = "roles"  # Note: 🤖 Don't delete this line
    name = Column(String(20), unique=True, index=True)
    json_permissions = Column(Text(), index=False, default='[]')


class DBRoleRepo(
    DBRepo[DBRoleEntity, Role, RoleData],
    RoleRepo
):
    schema_class = Role
    db_entity_class = DBRoleEntity

    def find_by_name(self, name: str) -> Optional[Role]:
        db = self.create_db_sesion()
        try:
            search_filter = DBRoleEntity.name == name
            return self.fetch_one_by_filter(db, search_filter)
        finally:
            db.close()

    def get_keyword_fields(self) -> List[InstrumentedAttribute]:
        return [
            DBRoleEntity.name
        ]

    def from_schema_data_to_db_entity_dict(
        self, role_data: RoleData
    ) -> Mapping[str, Any]:
        role_dict = super().from_schema_data_to_db_entity_dict(role_data)
        role_dict['json_permissions'] = jsons.dumps(role_data.permissions)
        return role_dict

    def from_db_entity_to_schema(self, db_role: DBRoleEntity) -> Role:
        role = super().from_db_entity_to_schema(db_role)
        role.permissions = jsons.loads(db_role.json_permissions)
        return role
