# Meltano

Your meltano project is located at `app` folder. Feel free to modify it. Also, don't forget to edit the command on `app/start.sh`. The command will be used whether you choose to run meltano locally (both natively or as container), and in DockerOperator (in case you want to run it with local airflow)

To run meltano locally you can invoke:

```sh
zaruba please runZarubaServiceName
zaruba please runZarubaServiceNameContainer
```

To run meltano ETL in your local airflow, you will need DockerOperator. The DAG file is located at `airflow/dags`.

If you need to install something/set variable in your airflow worker, please modify `airflow/initWorker.gotmpl.sh`.

Once you are ready, you can invoke:

```sh
zaruba please registerZarubaServiceName
```

# Known Issue

* [Installing meltano using pipenv is problematic](https://gitlab.com/meltano/meltano/-/issues/141). Currently meltano should be installed globally (i.e: `pip install meltano`)
* Meltano require some prerequisites like `psycopg2-binary`. If you are using ubuntu/debian, please consider running `~/.zaruba/setup/init.sh` and choose `setup ubuntu` to make sure that all necessary packages are installed.