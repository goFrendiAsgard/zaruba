from schemas.user import UserData

def test_user_data_add_new_role_id():
    user_data = UserData(username='user', email='user@mail.com', phone_number='+621345678', password='password', full_name='User Test')
    user_data.role_ids=['role']
    user_data.add_role_id('new')
    assert len(user_data.role_ids) == 2
    assert user_data.role_ids[0] == 'role'
    assert user_data.role_ids[1] == 'new'

def test_user_data_add_exist_role_id():
    user_data = UserData(username='user', email='user@mail.com', phone_number='+621345678', password='password', full_name='User Test')
    user_data.role_ids=['role']
    user_data.add_role_id('role')
    assert len(user_data.role_ids) == 1
    assert user_data.role_ids[0] == 'role'

def test_user_data_remove_role_id():
    user_data = UserData(username='user', email='user@mail.com', phone_number='+621345678', password='password', full_name='User Test')
    user_data.role_ids = ['role1', 'role2']
    user_data.remove_role_id('role1')
    assert len(user_data.role_ids) == 1
    assert user_data.role_ids[0] == 'role2'

def test_user_data_remove_inexist_role_id():
    user_data = UserData(username='user', email='user@mail.com', phone_number='+621345678', password='password', full_name='User Test')
    user_data.role_ids = ['role1', 'role2']
    user_data.remove_role_id('inexist')
    assert len(user_data.role_ids) == 2
    assert user_data.role_ids[0] == 'role1'
    assert user_data.role_ids[1] == 'role2'

def test_user_data_add_new_permission():
    user_data = UserData(username='user', email='user@mail.com', phone_number='+621345678', password='password', full_name='User Test')
    user_data.permissions=['permission']
    user_data.add_permission('new')
    assert len(user_data.permissions) == 2
    assert user_data.permissions[0] == 'permission'
    assert user_data.permissions[1] == 'new'

def test_user_data_add_exist_permission():
    user_data = UserData(username='user', email='user@mail.com', phone_number='+621345678', password='password', full_name='User Test')
    user_data.permissions=['permission']
    user_data.add_permission('permission')
    assert len(user_data.permissions) == 1
    assert user_data.permissions[0] == 'permission'

def test_user_data_remove_permission():
    user_data = UserData(username='user', email='user@mail.com', phone_number='+621345678', password='password', full_name='User Test')
    user_data.permissions = ['permission1', 'permission2']
    user_data.remove_permission('permission1')
    assert len(user_data.permissions) == 1
    assert user_data.permissions[0] == 'permission2'

def test_user_data_remove_inexist_permission():
    user_data = UserData(username='user', email='user@mail.com', phone_number='+621345678', password='password', full_name='User Test')
    user_data.permissions = ['permission1', 'permission2']
    user_data.remove_permission('inexist')
    assert len(user_data.permissions) == 2
    assert user_data.permissions[0] == 'permission1'
    assert user_data.permissions[1] == 'permission2'

def test_user_data_has_inexist_permission():
    user_data = UserData(username='user', email='user@mail.com', phone_number='+621345678', password='password', full_name='User Test')
    user_data.permissions = ['entity.id1.add', 'entity.id2.add']
    assert user_data.has_permission('inexist') == False

def test_user_data_has_exist_permission():
    user_data = UserData(username='user', email='user@mail.com', phone_number='+621345678', password='password', full_name='User Test')
    user_data.permissions = ['entity.id1.add', 'entity.id2.add']
    assert user_data.has_permission('entity.id1.add') == True

def test_user_data_has_exist_wildcard_permission():
    user_data = UserData(username='user', email='user@mail.com', phone_number='+621345678', password='password', full_name='User Test')
    user_data.permissions = ['entity.id1.add', 'entity.id2.add']
    assert user_data.has_permission('entity.*.add') == True
