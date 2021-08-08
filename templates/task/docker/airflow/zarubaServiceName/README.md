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

You can simply put your DAG on `dag` directory.

## Set up airflow

To run custom script on `zarubaServiceNameConfigurator`, you can modify `init.template.sh`.

By default, this is the content of the script:

```sh
#!/bin/bash

BANNER='
 _       _ _         _
(_)_ __ (_) |_   ___| |__
| | `_ \| | __| / __| `_ \
| | | | | | |_ _\__ \ | | |
|_|_| |_|_|\__(_)___/_| |_|
      Prepare your airflow.
        Achieve your dream.
===========================
'
echo "${BANNER}"

{{ $prefix := .GetConfig "envPrefix" }}
{{ $d := .Decoration }}

echo 'Your envPrefix is: {{ $prefix }}'
echo 'You can change your envPrefix by set up ENV_PREFIX before running the task'
echo ''
echo 'Your USER is: {{ .GetPrefixedEnv $prefix "USER" }}'
echo 'This value was taken from ${{ $prefix }}_USER'
echo 'In case of ${{ $prefix }}_USER does not exist, $USER will be used instead'
echo ''
echo 'Feel free to modify this script.'
echo 'For example, if you need to set airflow variable, you can do:'
echo '{{ $d.Yellow }}{{ $d.Bold }}airflow variables set key VALUE{{ $d.Normal }}'
```

As you can see, this is not an ordinary shell script. Rather, it is a go template that will help you generate the real `init.sh` (the one that will be executed by `zarubaServiceNameConfigurator`).

## Generating init(dot)sh

The `init.sh` script will be automatically generated whenenver you invoke `zaruba please runZarubaServiceName`.

However, you can also generate the script manually by invoking `zaruba please makeZarubaServiceNameInitScript`.

## Understanding the Go template

We use `Go template` in order to handle multiple environments. For example you might have different environment among your local computer, staging/dev server, and production server. We need a way to make the script reusable, and go-template is one of the best tool available to solve the solution.

Many other tools like `helm` use go template a lot, so why don't we use it as well?

To see every functions and variables available in zaruba's go template, you can directly jump into the [source code](https://github.com/state-alchemists/zaruba/blob/master/config/taskdata.go).

However, for typical use cases, you will only need to deal with these function:

* `.GetConfig <key>`
* `.GetPrefixedEnv <prefix> <envKey>`

You can use `.GetConfig` in order to get the configuration value of a task. Our `makeZarubaServiceNameInitScript` has `envPrefix` configuration that can be modified by changing `ENV_PREFIX` environment.

On the other hand, we also have `.GetPrefixedEnv`. By using this function, you will be able to get the value of any environment variable based on it's prefix.

Let's say you have the following environment:

```
LOCAL_HTTP_PORT=3000
PROD_HTTP_PORT=80
SERVICE_NAME=common
ENV_PREFIX=LOCAL
```

Now, let's see the following `init.template.sh`:

```sh
{{ $prefix := .GetConfig "envPrefix" }}
# this yield `LOCAL` since our `ENV_PREFIX` is set to `LOCAL`
echo "{{ $prefix }}" 

# This yield `3000` since `LOCAL_HTTP_PORT` is `3000`
echo "{{ .GetPrefixedEnv $prefix "HTTP_PORT" }}" 

# Since we don't have `LOCAL_SERVICE_NAME`, this will yield `common` instead (the value of `SERVICE_NAME`)
echo "{{ .GetPrefixedEnv $prefix "SERVICE_NAME" }}" 
```

The output should be:

```
LOCAL
3000
common
```

## Setting up your airflow variables

Your DAG might run anywhere, from your local computer (for development) or in the production.

The common approach to deal with different environment is by using [airflow variable](https://airflow.apache.org/docs/apache-airflow/stable/howto/variable.html).

You can set environment variable by:

* [Using the web interface](https://airflow.apache.org/docs/apache-airflow/stable/howto/variable.html#managing-variables)
* [Using environment variable](https://airflow.apache.org/docs/apache-airflow/stable/howto/variable.html#storing-variables-in-environment-variables)
* [Using command](https://airflow.apache.org/docs/apache-airflow/stable/cli-and-env-variables-ref.html#set_repeat1)

Obviously, for our example we will use the third one.

Let's say you have the following DAG:

```python

with airflow.DAG(
    dag_id,
    tags=tags,
    catchup=False,
    default_args=args,
    schedule_interval='0 0 0 * 0',
    max_active_runs=1,
) as dag:

    # see: https://airflow.apache.org/docs/apache-airflow/1.10.12/_api/airflow/contrib/operators/kubernetes_pod_operator/index.html
    myTask = KubernetesPodOperator(
        dag=dag,
        namespace=airflow.model.Variable.get('k8s_namespace'),
        image=airflow.model.Variable.get('image'),
        arguments=['--log-level=debug',  'elt',  'tap-hubspot',  'target-postgres'],
        env_vars = meltano_env_vars,
        labels={'foo': 'bar'},
        name='my-task',
        is_delete_operator_pod=True,
        task_id='my-task',
        get_logs=True,
    )
```

The `namespace` and `image` for each environment might be different (e.g: you might use `latest` image in your local computer and `v0.5.7` on production)

First you need to prepare your environment variable:

```
LOCAL_K8S_NAMESPACE=default
LOCAL_IMAGE=image-name:latest

PROD_K8S_NAMESPACE=prod
PROD_IMAGE=image-name:v0.5.7
```

Finally, you add this to your `init.template.sh`:

```sh
{{ $prefix := .GetConfig "envPrefix" }}
airflow variables set k8s_namespace {{ .GetPrefixedEnv "K8S_NAMESPACE" }}
airflow variables set image {{ .GetPrefixedEnv "IMAGE" }}
```
