from datetime import datetime, timedelta
from airflow import DAG
from airflow.operators.bash_operator import BashOperator
from airflow.providers.docker.operators.docker import DockerOperator
from airflow.models import Variable

default_args = {
    'owner'                 : 'airflow',
    'description'           : 'zarubaServiceName DockerOperator',
    'depend_on_past'        : False,
    'start_date'            : datetime(2018, 1, 3),
    'email_on_failure'      : False,
    'email_on_retry'        : False,
    'retries'               : 1,
    'retry_delay'           : timedelta(minutes=5)
}

docker_url = Variable.get(key='zarubaServiceNameDockerUrl')
docker_operator_image = Variable.get(key='zarubaServiceNameImage')
docker_operator_environment = Variable.get(key='zarubaServiceNameEnv', deserialize_json=True)

with DAG('zarubaServiceName_docker_dag', default_args=default_args, schedule_interval='5 * * * *', catchup=False) as dag:

    t1 = BashOperator(
        task_id='print_current_date',
        bash_command='date'
    )

    # see: https://airflow.apache.org/docs/apache-airflow-providers-docker/stable/_api/airflow/providers/docker/operators/docker/index.html#module-airflow.providers.docker.operators.docker
    t2 = DockerOperator(
        task_id='zarubaServiceName_docker_command',
        image=docker_operator_image,
        api_version='auto',
        auto_remove=True,
        environment=docker_operator_environment,
        docker_url=docker_url,
    )

    t3 = BashOperator(
        task_id='print_done',
        bash_command='echo "Done"'
    )

    t1 >> t2 >> t3