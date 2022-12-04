if [ -z "${SLUG}" ]
then
    SLUG="update"
fi

export IS_GENERATING_MIGRATION=1

alembic upgrade head
python ./helper/alembic/check_migration.py || alembic revision --autogenerate -m "${SLUG}"

if [ "${APP_ENABLE_AUTH_MODULE}" != "0" ]
then
    alembic --name auth upgrade head
    python ./helper/alembic/check_migration.py ./alembic.ini auth || alembic --name auth revision --autogenerate -m "${SLUG}"
fi

if [ "${APP_ENABLE_LOG_MODULE}" != "0" ]
then
    alembic --name log upgrade head
    python ./helper/alembic/check_migration.py ./alembic.ini log || alembic --name log revision --autogenerate -m "${SLUG}"
fi

if [ "${APP_ENABLE_CMS_MODULE}" != "0" ]
then
    alembic --name cms upgrade head
    python ./helper/alembic/check_migration.py ./alembic.ini cms || alembic --name cms revision --autogenerate -m "${SLUG}"
fi