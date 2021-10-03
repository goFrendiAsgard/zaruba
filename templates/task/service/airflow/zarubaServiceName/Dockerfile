FROM apache/airflow:2.1.2

ENV DOCKER_HOST="tcp://host.docker.internal:2375"

USER root
RUN apt-get update \
  && apt-get install -y --no-install-recommends docker.io \
  && apt-get autoremove -yqq --purge \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* \
  && usermod -aG docker airflow

USER airflow

# Because of this: https://stackoverflow.com/a/68646041
RUN pip install --no-cache-dir apache-airflow-providers-docker==2.1.1