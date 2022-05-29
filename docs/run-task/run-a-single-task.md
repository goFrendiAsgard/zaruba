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
         Elapsed Time: 2.039µs
         Current Time: 11:28:54
💀 🏁 Run 🔥 'clearLog' command on /home/gofrendi/zaruba/docs
💀    🚀 clearLog             🔥 Log removed
💀 🎉 Successfully running 🔥 'clearLog' command
💀 🔎 Job Running...
         Elapsed Time: 108.304542ms
         Current Time: 11:28:54
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 409.637984ms
         Current Time: 11:28:54
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
         Elapsed Time: 1.653µs
         Current Time: 11:28:54
💀 🏁 Run 🍎 'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
💀 🎉 Successfully running 🍎 'printHelloWorld' command
💀 🔎 Job Running...
         Elapsed Time: 101.881347ms
         Current Time: 11:28:55
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 312.576215ms
         Current Time: 11:28:55
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->