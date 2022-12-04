
if [ "${APP_ENABLE_ZTPL_APP_MODULE_NAME_MODULE}" != "0" ]
then
    alembic --name ztpl_app_module_name upgrade head
fi
