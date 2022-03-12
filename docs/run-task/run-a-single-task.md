<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🍺 Run a single task
<!--endTocHeader-->

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
 Elapsed Time: 1.512µs
 Current Time: 13:25:10
  Run  'clearLog' command on /home/gofrendi/zaruba/docs
   clearLog              13:25:10.393 Log removed
  Successfully running  'clearLog' command
  Job Running...
 Elapsed Time: 104.757518ms
 Current Time: 13:25:10
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 305.496053ms
 Current Time: 13:25:10
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
 Elapsed Time: 1.253µs
 Current Time: 13:25:10
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloWorld       13:25:10.836 hello world
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 101.614954ms
 Current Time: 13:25:10
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.773288ms
 Current Time: 13:25:11
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->