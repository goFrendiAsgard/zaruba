[⬅️ README](../README.md)

# 🏠 Zaruba Documentation

Welcome to Zaruba Documentation.

Zaruba is a [task](core-concepts/task/README.md) runner and [CLI utility](utilities/README.md). It helps you __prepare__, __run__, __deploy__, and __debug__ your applications.

To start working with Zaruba, you need to [create a project](use-cases/create-a-project.md) and [add some resources](use-cases/add-resources/README.md). Once your project is ready, you can start [running some tasks](run-tasks/README.md). 

In Zaruba, you can define your tasks as YAML.

![](images/solution-example.png)

You can [visit the core concepts](core-concepts/README.md) to know more about the detail.

Finally, you can visit our [end-to-end tutorial](use-cases/from-zero-to-cloud.md) to get a better glimpse.


# Table of Content

<!--startToc-->
- [⚙️ Zaruba Configuration](zaruba-configuration.md)
- [🏃 Run task](run-task/README.md)
  - [🍺 Run a Single Task](run-task/run-a-single-task.md)
  - [🍻 Run Many Tasks in Parallel](run-task/run-many-tasks-in-parallel.md)
  - [🏝️ Run Task with Custom Environments](run-task/run-task-with-custom-environments.md)
  - [🔤 Run task with custom values](run-task/run-task-with-custom-values.md)
  - [🏓 Run task interactively](run-task/run-task-interactively.md)
- [🔎 Explain task](explain-task.md)
- [🧠 Core Concepts](core-concepts/README.md)
  - [🏗️ Project](core-concepts/project/README.md)
    - [🧬 Project Anatomy](core-concepts/project/project-anatomy.md)
    - [🧳 Includes](core-concepts/project/includes.md)
    - [🔤 Project Inputs](core-concepts/project/project-inputs.md)
    - [⚙️ Project Configs](core-concepts/project/project-configs.md)
    - [🏝️ Project Envs](core-concepts/project/project-envs.md)
  - [🔨 Task](core-concepts/task/README.md)
    - [🧬 Task Anatomy](core-concepts/task/task-anatomy.md)
    - [🥛 Simple Command](core-concepts/task/simple-command.md)
    - [🍹 Long Running Service](core-concepts/task/long-running-service.md)
    - [⚙️ Task Configs](core-concepts/task/task-configs/README.md)
      - [Shared Configs](core-concepts/task/task-configs/shared-configs.md)
    - [🏝️ Task Envs](core-concepts/task/task-envs/README.md)
      - [Shared Envs](core-concepts/task/task-envs/shared-envs.md)
    - [🔤 Task Inputs](core-concepts/task/task-inputs.md)
    - [🧒 Extend task](core-concepts/task/extend-task.md)
    - [🍲 Define task dependencies](core-concepts/task/define-task-dependencies.md)
  - [🐹 Use Go Template](core-concepts/use-go-template.md)
- [👷🏽 Use Cases](use-cases/README.md)
  - [❇️ From Zero to Cloud](use-cases/from-zero-to-cloud.md)
  - [🏗️ Create a Project](use-cases/create-a-project.md)
  - [📦 Add Resources](use-cases/add-resources/README.md)
    - [🧩 Integration](use-cases/add-resources/integration/README.md)
      - [📦 External Repository](use-cases/add-resources/integration/external-repository.md)
      - [🐳 Docker Container](use-cases/add-resources/integration/docker-container.md)
      - [🐳 Docker Compose](use-cases/add-resources/integration/docker-compose.md)
      - [☸️ Helm Chart](use-cases/add-resources/integration/helm-chart.md)
    - [✨ From Scratch](use-cases/add-resources/from-scratch/README.md)
      - [🏃 Add Runner for Existing Application](use-cases/add-resources/from-scratch/add-runner-for-existing-application/README.md)
        - [Go Application Runner](use-cases/add-resources/from-scratch/add-runner-for-existing-application/go-application-runner.md)
        - [NodeJs Application Runner](use-cases/add-resources/from-scratch/add-runner-for-existing-application/node-js-application-runner.md)
        - [Python Application Runner](use-cases/add-resources/from-scratch/add-runner-for-existing-application/python-application-runner.md)
      - [✨ Generate New Application](use-cases/add-resources/from-scratch/generate-new-application/README.md)
        - [Simple Go Application](use-cases/add-resources/from-scratch/generate-new-application/simple-go-application.md)
        - [Simple NodeJs Application](use-cases/add-resources/from-scratch/generate-new-application/simple-node-js-application.md)
        - [Simple Python Application](use-cases/add-resources/from-scratch/generate-new-application/simple-python-application.md)
        - [Simple TypeScript Application](use-cases/add-resources/from-scratch/generate-new-application/simple-type-script-application.md)
        - [FastApi Application](use-cases/add-resources/from-scratch/generate-new-application/fast-api-application/README.md)
          - [Route](use-cases/add-resources/from-scratch/generate-new-application/fast-api-application/route.md)
          - [Event Handler](use-cases/add-resources/from-scratch/generate-new-application/fast-api-application/event-handler.md)
          - [Rpc Handler](use-cases/add-resources/from-scratch/generate-new-application/fast-api-application/rpc-handler.md)
          - [Crud](use-cases/add-resources/from-scratch/generate-new-application/fast-api-application/crud.md)
      - [🚢 Add Application Deployment](use-cases/add-resources/from-scratch/add-application-deployment.md)
      - [🥉 Add Third Party Service](use-cases/add-resources/from-scratch/add-third-party-service.md)
    - [⚙️ Resource Configurations](use-cases/add-resources/resource-configurations.md)
  - [🏭 Add Generator](use-cases/add-generator.md)
  - [🏝️ Synchronize task environments](use-cases/synchronize-task-environments.md)
  - [🚌 Run Applications Locally](use-cases/run-applications-locally.md)
  - [🏃‍♂️ Run Some Applications Locally](use-cases/run-some-applications-locally.md)
  - [🚀 Deploy Applications](use-cases/deploy-applications.md)
- Built-in
<!--endToc-->