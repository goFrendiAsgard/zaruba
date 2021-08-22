# Airflow Artefact

Airflow is a platform created by the community to programmatically author, schedule and monitor workflows.

With Zaruba you can manage your airflow DAG and run it in your local computer seemlessly.

## Run airflow

To run Airflow in your local computer you can invoke:

```sh
zaruba please runZarubaServiceName
```

By invoking the command, you will also run several containers:

* `zarubaAirflowRedisService`: Redis, for caching and stuffs
* `zarubaAirflowPostgreService`: Posgre, for airflow persistance storage
* `zarubaServiceNameConfigurator`: A container that run before any other airflow containers. Through this container, Zaruba should set up the database and anything else. It will also run your custom script (see: `init.template.sh`).
* `zarubaServiceNameWebServer`: The web server, by default run on port `8080`. To change this, you should edit `zarubaServiceNameWebServer`'s `port` value at `../zaruba-tasks/zarubaServiceName/config.zaruba.yaml`.
* `zarubaServiceNameScheduler`: The scheduler.
* `zarubaServiceNameWorker`: The one that really run the task.
* `zarubaServiceNameFlower`: The one that monitor `celery` messaging. By default run on port `5555`. To change this, you should edit `zarubaServiceNameFlower`'s `port` value at `../zaruba-tasks/zarubaServiceName/config.zaruba.yaml`.

> 💡 To see how airflow really works, please visit [this article](https://airflow-tutorial.readthedocs.io/en/latest/airflow-intro.html)

## Add your DAG

You can simply put your DAG on `dags` directory.

## Setup Airflow

To setup custom configuration in your airflow instance, you will need to modify:

* `zarubaServiceName/configurator.gotmpl.sh`
* `zarubaServiceName/worker.gotmpl.sh`
