# Creating Docker Task

To create a docker task, you can perform: `zaruba please makeDockerTask -i`.

Zaruba also provide several preset docker-tasks, namely:

* Cassandra
* ElasticSearch
* MongoDB
* MySQL
* RabbitMQ
* Redis


## Example

```sh
# run interactively
zaruba please makeDockerTask -i
zaruba please makeCassandraDockerTask -i
zaruba please makeMySqlDockerTask -i

# run with parameters
zaruba please makeDockerTask generator.docker.image.name=nginx generator.docker.container.name=myNginx generator.service.name=myNginx generator.service.envs="MY_ENV=MY_VALUE" ports=80
```

## Involved tasks:

* [makeDockerTask](tasks/makeDockerTask.md)
* [makeCassandraDockerTask](tasks/makeCassandraDockerTask.md)
* [makeElasticsearchDockerTask](tasks/makeElasticsearchDockerTask.md)
* [makeMongoDockerTask](tasks/makeMongoDockerTask.md)
* [makeMysqlDockerTask](tasks/makeMysqlDockerTask.md)
* [makeRabbitmqDockerTask](tasks/makeRabbitmqDockerTask.md)
* [makeRedisDockerTask](tasks/makeRedisDockerTask.md)


## What's next

* [Running tasks](running-task.md)
* [Creating service task](creating-service-task.md)
* [Creating task manually](understanding-task.md)
* [Creating Fast API service](creating-fast-api-service.md)
