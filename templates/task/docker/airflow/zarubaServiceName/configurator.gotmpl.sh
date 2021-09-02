mkdir -p /opt/airflow/logs /opt/airflow/dags /opt/airflow/plugins
chown -R "${AIRFLOW_UID}:${AIRFLOW_GID}" /opt/airflow/{logs,dags,plugins}
exec /entrypoint airflow version
echo "Init zarubaServiceName configurator"