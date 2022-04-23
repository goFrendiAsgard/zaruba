<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🍺 Run a Single Task
<!--endTocHeader-->

You can run a single task by providing it's name.

There are two types of task:

* __Globally accessible task__: Can be executed from anywhere
* __Project specific task__: Can only be executed from project's top level directory.

# Run a Globally Accessible Task

To execute globally accessible task, you can invoke `zaruba please` from anywhere:

```bash
zaruba please <task-name>
```

__Example:__

<!--startCode-->
```bash
zaruba please clearLog
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.411µs
         Current Time: 13:17:30
💀 🏁 Run 🔥 'clearLog' command on /home/gofrendi/zaruba/docs
💀    🚀 clearLog             🔥 13:17:30.251 Log removed
💀 🎉 Successfully running 🔥 'clearLog' command
💀 🔎 Job Running...
         Elapsed Time: 104.143398ms
         Current Time: 13:17:30
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 305.342627ms
         Current Time: 13:17:30
zaruba please clearLog
```````
</details>
<!--endCode-->

 By default, Zaruba provide some [builtin core tasks](../core-tasks/README.md) that are globally accessible.
 
 If you want to make your tasks globally accessible, you can add them to `ZARUBA_SCRIPTS` environment variable. Please refer to [zaruba configuration](../configuration.md) for more information.

# Run a Project Specific Task

To execute any [project](./project/README.md) specific tasks, you need to be in the project's directory first:

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
         Elapsed Time: 1.169µs
         Current Time: 13:17:30
💀 🏁 Run 🍎 'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
💀    🚀 printHelloWorld      🍎 13:17:30.742 hello world
💀 🎉 Successfully running 🍎 'printHelloWorld' command
💀 🔎 Job Running...
         Elapsed Time: 102.146426ms
         Current Time: 13:17:30
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 213.05067ms
         Current Time: 13:17:30
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->