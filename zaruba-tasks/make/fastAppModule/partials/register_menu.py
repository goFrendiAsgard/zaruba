enable_ztpl_app_module_name_module = os.getenv('APP_ENABLE_ZTPL_APP_MODULE_NAME_MODULE', '1') != '0'
# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
if enable_ui and enable_ztpl_app_module_name_module:
    menu_service.add_menu(name='ztplAppModuleName', title='ZtplAppModuleName', url='#', auth_type=AuthType.ANYONE)

