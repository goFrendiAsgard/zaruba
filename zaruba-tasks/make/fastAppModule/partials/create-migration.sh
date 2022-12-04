
if [ "${APP_ENABLE_ZTPL_APP_MODULE_NAME_MODULE}" != "0" ]
then
    alembic --name ztpl_app_module_name upgrade head
    python ./helper/alembic/check_migration.py ./alembic.ini ztpl_app_module_name || alembic --name ztpl_app_module_name revision --autogenerate -m "${SLUG}"
fi
