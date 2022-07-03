from helpers.transport import RPC
from auth.accountService import AccountService

def register_account_rpc(rpc: RPC, account_service: AccountService):

    @rpc.handle('create_access_token')
    def create_access_token(identity: str, password: str) -> str:
        return account_service.create_access_token(identity, password)

    @rpc.handle('refresh_access_token')
    def refresh_access_token(token: str) -> str:
        return account_service.refresh_access_token(token)

    print('Handle RPC for auth.Account')