alembic upgrade head

if [ "${APP_ENABLE_AUTH_MODULE}" != "0" ]
then
    alembic --name auth upgrade head
fi

if [ "${APP_ENABLE_LOG_MODULE}" != "0" ]
then
    alembic --name log upgrade head
fi

if [ "${APP_ENABLE_CMS_MODULE}" != "0" ]
then
    alembic --name cms upgrade head
fi