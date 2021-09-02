mkdir -p /sources/logs /sources/dags /sources/plugins
chown -R "${AIRFLOW_UID}:${AIRFLOW_GID}" /sources/{logs,dags,plugins}
exec /entrypoint airflow version
echo "Init zarubaServiceName configurator"