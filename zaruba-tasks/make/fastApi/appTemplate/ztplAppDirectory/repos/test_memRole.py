from schemas.role import RoleData, Role
from repos.role import MemRoleRepo
import datetime

def test_find_by_id():
    mem_role_repo = MemRoleRepo()
    mem_role_repo.set_storage({
        'test_exist_role_id': Role(
            id='test_exist_role_id',
            name='test_exist_role_name',
            permissions=['test_exist_role_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    role = mem_role_repo.find_by_id('test_exist_role_id')
    assert role.id == 'test_exist_role_id'
    assert role.name == 'test_exist_role_name'
    assert len(role.permissions) == 1
    assert role.permissions[0] == 'test_exist_role_permission'

def test_find_inexist_by_id():
    mem_role_repo = MemRoleRepo()
    mem_role_repo.set_storage({
        'test_exist_role_id': Role(
            id='test_exist_role_id',
            name='test_exist_role_name',
            permissions=['test_exist_role_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    role = mem_role_repo.find_by_id('test_inexist_role_id')
    assert role is None

def test_find():
    mem_role_repo = MemRoleRepo()
    mem_role_repo.set_storage({
        'test_exist_role_id': Role(
            id='test_exist_role_id',
            name='test_exist_role_name',
            permissions=['test_exist_role_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    roles = mem_role_repo.find('', 10, 0)
    assert len(roles) == 1
    role = roles[0]
    assert role.id == 'test_exist_role_id'
    assert role.name == 'test_exist_role_name'
    assert len(role.permissions) == 1
    assert role.permissions[0] == 'test_exist_role_permission'

def test_insert():
    mem_role_repo = MemRoleRepo()
    mem_role_repo.set_storage({
        'test_exist_role_id': Role(
            id='test_exist_role_id',
            name='test_exist_role_name',
            permissions=['test_exist_role_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    new_role = mem_role_repo.insert(RoleData(
        name='test_new_role_name',
        permissions=['test_new_role_permission']
    ))
    assert new_role.id != ''
    assert new_role.name == 'test_new_role_name'
    assert len(new_role.permissions) == 1
    assert new_role.permissions[0] == 'test_new_role_permission'
    roles = mem_role_repo.find('', 10, 0)
    assert len(roles) == 2

def test_update():
    mem_role_repo = MemRoleRepo()
    mem_role_repo.set_storage({
        'test_exist_role_id': Role(
            id='test_exist_role_id',
            name='test_exist_role_name',
            permissions=['test_exist_role_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    updated_role = mem_role_repo.update('test_exist_role_id', RoleData(
        name='test_new_role_name',
        permissions=['test_new_role_permission']
    ))
    assert updated_role.id == 'test_exist_role_id'
    assert updated_role.name == 'test_new_role_name'
    assert len(updated_role.permissions) == 1
    assert updated_role.permissions[0] == 'test_new_role_permission'
    roles = mem_role_repo.find('', 10, 0)
    assert len(roles) == 1

def test_update_inexist():
    mem_role_repo = MemRoleRepo()
    mem_role_repo.set_storage({
        'test_exist_role_id': Role(
            id='test_exist_role_id',
            name='test_exist_role_name',
            permissions=['test_exist_role_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    updated_role = mem_role_repo.update('test_inexist_role_id', RoleData(
        name='test_new_role_name',
        permissions=['test_new_role_permission']
    ))
    assert updated_role is None
    roles = mem_role_repo.find('', 10, 0)
    assert len(roles) == 1
    role = roles[0]
    assert role.id == 'test_exist_role_id'
    assert role.name == 'test_exist_role_name'
    assert len(role.permissions) == 1
    assert role.permissions[0] == 'test_exist_role_permission'

def test_delete():
    mem_role_repo = MemRoleRepo()
    mem_role_repo.set_storage({
        'test_exist_role_id': Role(
            id='test_exist_role_id',
            name='test_exist_role_name',
            permissions=['test_exist_role_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    deleted_role = mem_role_repo.delete('test_exist_role_id')
    assert deleted_role.id == 'test_exist_role_id'
    assert deleted_role.name == 'test_exist_role_name'
    assert len(deleted_role.permissions) == 1
    assert deleted_role.permissions[0] == 'test_exist_role_permission'
    roles = mem_role_repo.find('', 10, 0)
    assert len(roles) == 0

def test_delete_inexist():
    mem_role_repo = MemRoleRepo()
    mem_role_repo.set_storage({
        'test_exist_role_id': Role(
            id='test_exist_role_id',
            name='test_exist_role_name',
            permissions=['test_exist_role_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    deleted_role = mem_role_repo.delete('test_inexist_role_id')
    assert deleted_role is None
    roles = mem_role_repo.find('', 10, 0)
    assert len(roles) == 1
    role = roles[0]
    assert role.id == 'test_exist_role_id'
    assert role.name == 'test_exist_role_name'
    assert len(role.permissions) == 1
    assert role.permissions[0] == 'test_exist_role_permission'