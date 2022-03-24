<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🍺 Run a Single Task
<!--endTocHeader-->

You can run a specific task by providing it's name.

There are two types of task:

* __Globally available task__: Can be executed from anywhere
* __Project specific task__: Can only be executed from project's top level directory.

# Run a Globally Available Task

To execute globally available task, you can invoke `zaruba please` from anywhere:

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
Job Starting...
 Elapsed Time: 1.158µs
 Current Time: 21:06:22
  Run  'clearLog' command on /home/gofrendi/zaruba/docs
   clearLog              21:06:22.777 Log removed
  Successfully running  'clearLog' command
  Job Running...
 Elapsed Time: 104.337629ms
 Current Time: 21:06:22
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 305.759707ms
 Current Time: 21:06:23
zaruba please clearLog
```````
</details>
<!--endCode-->

 There are special [builtin core tasks](../core-tasks/README.md) that can be executed from anywhere.
 
 If you want to make your tasks globally available, you can add it's script definition to `ZARUBA_SCRIPTS` environment variable. Please refer to [zaruba configuration](../configuration.md).

# Run a Project Specific Task

To execute any [project](./project/README.md) specific tasks, you need to be in the project's top level directory:

```bash
cd <project-directory>
zaruba please <task-name>
```

__Example:__

<!--startCode-->
```bash
cd examples/run-tasks
zaruba please printHelloWorld
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.039µs
 Current Time: 21:06:23
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloWorld       21:06:23.231 hello world
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.361075ms
 Current Time: 21:06:23
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.128795ms
 Current Time: 21:06:23
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->