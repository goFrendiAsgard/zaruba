<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🍺 Run a Single Task
<!--endTocHeader-->

You can run a single task by providing it's name.

There are two types of task:

* __Globally-accessible task__: You can execute these tasks from anywhere
* __Project-specific task__: You can only execute these tasks the [project](../core-concepts/project/README.md)'s directory.

# Run a Globally Accessible Task

To execute globally accessible task, you can invoke `zaruba please` from anywhere:

```bash
zaruba please <task-name>
```

__Example:__

<!--startCode-->
```bash
zaruba please showSolidPrinciple
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.41µs
         Current Time: 07:53:20
💀 🏁 Running 🦉 showSolidPrinciple runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs
💀    🚀 🦉 showSolidPrinciple    S  Single Responsibility Principle 
💀    🚀 🦉 showSolidPrinciple    O  Open/Closed Principle 
💀    🚀 🦉 showSolidPrinciple    L  Liskov's Substitution Principle 
💀    🚀 🦉 showSolidPrinciple    I  Interface Segregation Principle 
💀    🚀 🦉 showSolidPrinciple    D  Dependency Inversion Principle 
💀 🎉 Successfully running 🦉 showSolidPrinciple runner (Attempt 1 of 3)
💀 🔎 Job Running...
         Elapsed Time: 7.39661ms
         Current Time: 07:53:20
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 310.570471ms
         Current Time: 07:53:21
zaruba please showSolidPrinciple
```````
</details>
<!--endCode-->

 By default, Zaruba provides some [builtin core tasks](../core-tasks/README.md) that are globally accessible.
 
 To make your tasks globally accessible, you need to add them to `ZARUBA_SCRIPTS` environment variable. Please refer to [zaruba configuration](../configuration.md) for more information.

# Run a Project Specific Task

To execute any project-specific tasks, you need to be in the project's directory first:

```bash
cd <project-directory>
zaruba please <task-name>
```

Please note that the command will not work from the project's subdirectory.

__Example:__

<!--startCode-->
```bash
cd examples/run-tasks
zaruba please printHelloWorld
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.678µs
         Current Time: 07:53:21
💀 🏁 Running 🍎 printHelloWorld runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
💀 🎉 Successfully running 🍎 printHelloWorld runner (Attempt 1 of 3)
💀 🔎 Job Running...
         Elapsed Time: 1.917541ms
         Current Time: 07:53:21
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀    🚀 🍎 printHelloWorld      hello world
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 404.632191ms
         Current Time: 07:53:21
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->