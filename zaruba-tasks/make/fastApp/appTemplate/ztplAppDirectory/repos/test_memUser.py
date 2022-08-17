from schemas.user import UserData, User
from repos.user import MemUserRepo
import datetime

def test_find_by_id():
    mem_user_repo = MemUserRepo()
    mem_user_repo.set_storage({
        'test_exist_user_id': User(
            id='test_exist_user_id',
            username='test_exist_username',
            full_name='test_exist_full_name',
            email='test_exist_email@domain.com',
            phone_number='+62812345678',
            active=True,
            password='test_exist_password',
            role_ids=[],
            permissions=['test_exist_user_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    user = mem_user_repo.find_by_id('test_exist_user_id')
    assert user.id == 'test_exist_user_id'
    assert user.username == 'test_exist_username'
    assert user.full_name == 'test_exist_full_name'
    assert user.email == 'test_exist_email@domain.com'
    assert user.phone_number == '+62812345678'
    assert user.active == True
    assert user.password == 'test_exist_password'
    assert len(user.role_ids) == 0
    assert len(user.permissions) == 1
    assert user.permissions[0] == 'test_exist_user_permission'

def test_find_inexist_by_id():
    mem_user_repo = MemUserRepo()
    mem_user_repo.set_storage({
        'test_exist_user_id': User(
            id='test_exist_user_id',
            username='test_exist_username',
            full_name='test_exist_full_name',
            email='test_exist_email@domain.com',
            phone_number='+62812345678',
            active=True,
            password='test_exist_password',
            role_ids=[],
            permissions=['test_exist_user_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    user = mem_user_repo.find_by_id('test_inexist_user_id')
    assert user is None

def test_find():
    mem_user_repo = MemUserRepo()
    mem_user_repo.set_storage({
        'test_exist_user_id': User(
            id='test_exist_user_id',
            username='test_exist_username',
            full_name='test_exist_full_name',
            email='test_exist_email@domain.com',
            phone_number='+62812345678',
            active=True,
            password='test_exist_password',
            role_ids=[],
            permissions=['test_exist_user_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    users = mem_user_repo.find('', 10, 0)
    assert len(users) == 1
    user = users[0]
    assert user.id == 'test_exist_user_id'
    assert user.username == 'test_exist_username'
    assert user.full_name == 'test_exist_full_name'
    assert user.email == 'test_exist_email@domain.com'
    assert user.phone_number == '+62812345678'
    assert user.active == True
    assert user.password == 'test_exist_password'
    assert len(user.role_ids) == 0
    assert len(user.permissions) == 1
    assert user.permissions[0] == 'test_exist_user_permission'

def test_insert():
    mem_user_repo = MemUserRepo()
    mem_user_repo.set_storage({
        'test_exist_user_id': User(
            id='test_exist_user_id',
            username='test_exist_username',
            full_name='test_exist_full_name',
            email='test_exist_email@domain.com',
            phone_number='+62812345678',
            active=True,
            password='test_exist_password',
            role_ids=[],
            permissions=['test_exist_user_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    new_user = mem_user_repo.insert(UserData(
        username='test_new_username',
        full_name='test_new_full_name',
        email='test_new_email@domain.com',
        phone_number='+62812345678',
        active=True,
        password='test_new_password',
        role_ids=[],
        permissions=['test_new_user_permission']
    ))
    assert new_user.id != ''
    assert new_user.username == 'test_new_username'
    assert new_user.full_name == 'test_new_full_name'
    assert new_user.email == 'test_new_email@domain.com'
    assert new_user.phone_number == '+62812345678'
    assert new_user.active == True
    assert new_user.password == 'test_new_password'
    assert len(new_user.role_ids) == 0
    assert len(new_user.permissions) == 1
    assert new_user.permissions[0] == 'test_new_user_permission'
    users = mem_user_repo.find('', 10, 0)
    assert len(users) == 2

def test_update():
    mem_user_repo = MemUserRepo()
    mem_user_repo.set_storage({
        'test_exist_user_id': User(
            id='test_exist_user_id',
            username='test_exist_username',
            full_name='test_exist_full_name',
            email='test_exist_email@domain.com',
            phone_number='+62812345678',
            active=True,
            password='test_exist_password',
            role_ids=[],
            permissions=['test_exist_user_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    updated_user = mem_user_repo.update('test_exist_user_id', UserData(
        username='test_new_username',
        full_name='test_new_full_name',
        email='test_new_email@domain.com',
        phone_number='+62812345678',
        active=True,
        password='test_new_password',
        role_ids=[],
        permissions=['test_new_user_permission']
    ))
    assert updated_user.id == 'test_exist_user_id'
    assert updated_user.username == 'test_new_username'
    assert updated_user.full_name == 'test_new_full_name'
    assert updated_user.email == 'test_new_email@domain.com'
    assert updated_user.phone_number == '+62812345678'
    assert updated_user.active == True
    assert updated_user.password == 'test_new_password'
    assert len(updated_user.role_ids) == 0
    assert len(updated_user.permissions) == 1
    assert updated_user.permissions[0] == 'test_new_user_permission'
    users = mem_user_repo.find('', 10, 0)
    assert len(users) == 1

def test_update_inexist():
    mem_user_repo = MemUserRepo()
    mem_user_repo.set_storage({
        'test_exist_user_id': User(
            id='test_exist_user_id',
            username='test_exist_username',
            full_name='test_exist_full_name',
            email='test_exist_email@domain.com',
            phone_number='+62812345678',
            active=True,
            password='test_exist_password',
            role_ids=[],
            permissions=['test_exist_user_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    updated_user = mem_user_repo.update('test_inexist_user_id', UserData(
        username='test_new_username',
        full_name='test_new_full_name',
        email='test_new_email@domain.com',
        phone_number='+62812345678',
        active=True,
        password='test_new_password',
        role_ids=[],
        permissions=['test_new_user_permission']
    ))
    assert updated_user is None
    users = mem_user_repo.find('', 10, 0)
    assert len(users) == 1
    user = users[0]
    assert user.id == 'test_exist_user_id'
    assert user.username == 'test_exist_username'
    assert user.full_name == 'test_exist_full_name'
    assert user.email == 'test_exist_email@domain.com'
    assert user.phone_number == '+62812345678'
    assert user.active == True
    assert user.password == 'test_exist_password'
    assert len(user.role_ids) == 0
    assert len(user.permissions) == 1
    assert user.permissions[0] == 'test_exist_user_permission'

def test_delete():
    mem_user_repo = MemUserRepo()
    mem_user_repo.set_storage({
        'test_exist_user_id': User(
            id='test_exist_user_id',
            username='test_exist_username',
            full_name='test_exist_full_name',
            email='test_exist_email@domain.com',
            phone_number='+62812345678',
            active=True,
            password='test_exist_password',
            role_ids=[],
            permissions=['test_exist_user_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    deleted_user = mem_user_repo.delete('test_exist_user_id')
    assert deleted_user.id == 'test_exist_user_id'
    assert deleted_user.username == 'test_exist_username'
    assert deleted_user.full_name == 'test_exist_full_name'
    assert deleted_user.email == 'test_exist_email@domain.com'
    assert deleted_user.phone_number == '+62812345678'
    assert deleted_user.active == True
    assert deleted_user.password == 'test_exist_password'
    assert len(deleted_user.role_ids) == 0
    assert len(deleted_user.permissions) == 1
    users = mem_user_repo.find('', 10, 0)
    assert len(users) == 0

def test_delete_inexist():
    mem_user_repo = MemUserRepo()
    mem_user_repo.set_storage({
        'test_exist_user_id': User(
            id='test_exist_user_id',
            username='test_exist_username',
            full_name='test_exist_full_name',
            email='test_exist_email@domain.com',
            phone_number='+62812345678',
            active=True,
            password='test_exist_password',
            role_ids=[],
            permissions=['test_exist_user_permission'],
            created_at=datetime.datetime.min,
            updated_at=datetime.datetime.min
        )
    })
    deleted_user = mem_user_repo.delete('test_inexist_user_id')
    assert deleted_user is None
    users = mem_user_repo.find('', 10, 0)
    assert len(users) == 1
    user = users[0]
    assert user.id == 'test_exist_user_id'
    assert user.username == 'test_exist_username'
    assert user.full_name == 'test_exist_full_name'
    assert user.email == 'test_exist_email@domain.com'
    assert user.phone_number == '+62812345678'
    assert user.active == True
    assert user.password == 'test_exist_password'
    assert len(user.role_ids) == 0
    assert len(user.permissions) == 1
    assert user.permissions[0] == 'test_exist_user_permission'