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
         Elapsed Time: 1.424µs
         Current Time: 07:49:51
💀 🏁 Run 🦉 'showSolidPrinciple' command on /home/gofrendi/zaruba/docs
💀    🚀 showSolidPrinciple   🦉  S  Single Responsibility Principle 
💀    🚀 showSolidPrinciple   🦉  O  Open/Closed Principle 
💀    🚀 showSolidPrinciple   🦉  L  Liskov's Substitution Principle 
💀    🚀 showSolidPrinciple   🦉  I  Interface Segregation Principle 
💀    🚀 showSolidPrinciple   🦉  D  Dependency Inversion Principle 
💀 🎉 Successfully running 🦉 'showSolidPrinciple' command
💀 🔎 Job Running...
         Elapsed Time: 107.552834ms
         Current Time: 07:49:51
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 409.817841ms
         Current Time: 07:49:51
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
         Elapsed Time: 1.359µs
         Current Time: 07:49:52
💀 🏁 Run 🍎 'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
💀    🚀 printHelloWorld      🍎 hello world
💀 🎉 Successfully running 🍎 'printHelloWorld' command
💀 🔎 Job Running...
         Elapsed Time: 101.780318ms
         Current Time: 07:49:52
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 505.008896ms
         Current Time: 07:49:52
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->