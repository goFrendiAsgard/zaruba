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
 Elapsed Time: 1.044µs
 Current Time: 14:18:57
  Run  'clearLog' command on /home/gofrendi/zaruba/docs
   clearLog              14:18:57.647 Log removed
  Successfully running  'clearLog' command
  Job Running...
 Elapsed Time: 104.450931ms
 Current Time: 14:18:57
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 305.442002ms
 Current Time: 14:18:57
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
 Elapsed Time: 958ns
 Current Time: 14:18:58
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloWorld       14:18:58.125 hello world
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.028932ms
 Current Time: 14:18:58
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.313517ms
 Current Time: 14:18:58
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->