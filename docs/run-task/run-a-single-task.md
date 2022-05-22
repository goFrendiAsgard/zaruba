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
zaruba please clearLog
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 14.535µs
         Current Time: 07:20:13
💀 🏁 Run 🔥 'clearLog' command on /home/gofrendi/zaruba/docs
💀    🚀 clearLog             🔥 Log removed
💀 🎉 Successfully running 🔥 'clearLog' command
💀 🔎 Job Running...
         Elapsed Time: 111.669855ms
         Current Time: 07:20:13
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 312.500551ms
         Current Time: 07:20:13
zaruba please clearLog
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
         Elapsed Time: 2.48µs
         Current Time: 07:20:13
💀 🏁 Run 🍎 'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
💀    🚀 printHelloWorld      🍎 hello world
💀 🎉 Successfully running 🍎 'printHelloWorld' command
💀 🔎 Job Running...
         Elapsed Time: 103.331235ms
         Current Time: 07:20:13
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 214.040333ms
         Current Time: 07:20:14
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->