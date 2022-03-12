<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🍺 Run a single task
<!--endTocHeader-->

You can run a specific task by knowing it's name.

There are two types of task:

* Builtin task: You can run these tasks from anywhere
* Project specific task: You should be in project's top level directory in order to execute these tasks.

# Run a Builitin Task

To execute builtin [core-tasks](../core-tasks/README.md), you can invoke `zaruba please` from anywhere:

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
 Elapsed Time: 1.346µs
 Current Time: 14:42:59
  Run  'clearLog' command on /home/gofrendi/zaruba/docs
   clearLog              14:42:59.809 Log removed
  Successfully running  'clearLog' command
  Job Running...
 Elapsed Time: 103.634628ms
 Current Time: 14:42:59
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 303.998071ms
 Current Time: 14:43:00
zaruba please clearLog
```````
</details>
<!--endCode-->

# Run a Project Specific Task

But, to execute any [project](./project/README.md) specific tasks, you need to be in the project directory first:

```bash
cd <project-directory>
zaruba please <task-name>
```

__Example:__

<!--startCode-->
```bash
cd examples/tasks
zaruba please printHelloWorld
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.214µs
 Current Time: 14:43:00
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloWorld       14:43:00.251 hello world
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.95529ms
 Current Time: 14:43:00
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 214.77285ms
 Current Time: 14:43:00
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->