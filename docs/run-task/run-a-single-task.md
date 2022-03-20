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
 Elapsed Time: 1.761µs
 Current Time: 15:58:39
  Run  'clearLog' command on /home/gofrendi/zaruba/docs
   clearLog              15:58:39.575 Log removed
  Successfully running  'clearLog' command
  Job Running...
 Elapsed Time: 104.482242ms
 Current Time: 15:58:39
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 305.853583ms
 Current Time: 15:58:39
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
 Elapsed Time: 1.255µs
 Current Time: 15:58:40
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloWorld       15:58:40.057 hello world
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.070207ms
 Current Time: 15:58:40
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.714513ms
 Current Time: 15:58:40
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->