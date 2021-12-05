from datetime import datetime, timedelta

from airflow import DAG
from airflow.operators.bash import BashOperator
from airflow.operators.dummy import DummyOperator

with DAG(
    dag_id='my_dag',
    schedule_interval='0 0 * * *',
    start_date=datetime(2021, 1, 1),
    catchup=False,
    dagrun_timeout=timedelta(minutes=60),
    tags=['example'],
    params={"example_key": "example_value"},
) as dag:

    run_this = BashOperator(
        task_id='show_hello_world',
        bash_command='echo "💀 Hello world"',
    )

    run_this

if __name__ == "__main__":
    dag.cli()