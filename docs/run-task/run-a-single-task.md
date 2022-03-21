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
 Elapsed Time: 1.307µs
 Current Time: 06:43:31
  Run  'clearLog' command on /home/gofrendi/zaruba/docs
   clearLog              06:43:31.36  Log removed
  Successfully running  'clearLog' command
  Job Running...
 Elapsed Time: 104.530156ms
 Current Time: 06:43:31
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 305.488319ms
 Current Time: 06:43:31
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
 Elapsed Time: 965ns
 Current Time: 06:43:31
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloWorld       06:43:31.824 hello world
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 101.866546ms
 Current Time: 06:43:31
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.474327ms
 Current Time: 06:43:32
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->