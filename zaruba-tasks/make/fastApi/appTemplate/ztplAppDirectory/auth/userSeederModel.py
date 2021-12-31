from auth.userModel import UserModel
from schemas.user import UserData

class UserSeederModel():

    def __init__(self, user_model: UserModel):
        self.user_model = user_model


    def seed(self, user_data: UserData):
        existing_root_user = self.user_model.find_by_username(user_data.username)
        if existing_root_user:
            return
        self.user_model.insert(user_data)