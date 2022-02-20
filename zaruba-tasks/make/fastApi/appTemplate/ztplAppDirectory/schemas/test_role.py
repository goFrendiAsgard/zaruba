from schemas.role import RoleData

def test_role_data_add_new_permission():
    role_data = RoleData(name='test_role')
    role_data.permissions=['permission']
    role_data.add_permission('new')
    assert len(role_data.permissions) == 2
    assert role_data.permissions[0] == 'permission'
    assert role_data.permissions[1] == 'new'

def test_role_data_add_exist_permission():
    role_data = RoleData(name='test_role')
    role_data.permissions=['permission']
    role_data.add_permission('permission')
    assert len(role_data.permissions) == 1
    assert role_data.permissions[0] == 'permission'

def test_role_data_remove_permission():
    role_data = RoleData(name='test_role')
    role_data.permissions = ['permission1', 'permission2']
    role_data.remove_permission('permission1')
    assert len(role_data.permissions) == 1
    assert role_data.permissions[0] == 'permission2'

def test_role_data_remove_inexist_permission():
    role_data = RoleData(name='test_role')
    role_data.permissions = ['permission1', 'permission2']
    role_data.remove_permission('inexist')
    assert len(role_data.permissions) == 2
    assert role_data.permissions[0] == 'permission1'
    assert role_data.permissions[1] == 'permission2'

def test_role_data_has_inexist_permission():
    role_data = RoleData(name='test_role')
    role_data.permissions = ['entity.id1.add', 'entity.id2.add']
    assert role_data.has_permission('inexist') == False

def test_role_data_has_exist_permission():
    role_data = RoleData(name='test_role')
    role_data.permissions = ['entity.id1.add', 'entity.id2.add']
    assert role_data.has_permission('entity.id1.add') == True

def test_role_data_has_exist_wildcard_permission():
    role_data = RoleData(name='test_role')
    role_data.permissions = ['entity.id1.add', 'entity.id2.add']
    assert role_data.has_permission('entity.*.add') == True
