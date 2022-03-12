<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# Run a single task
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
 Elapsed Time: 1.03µs
 Current Time: 13:00:38
  Run  'clearLog' command on /home/gofrendi/zaruba/docs
   clearLog              13:00:38.705 Log removed
  Successfully running  'clearLog' command
  Job Running...
 Elapsed Time: 104.554526ms
 Current Time: 13:00:38
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 305.936092ms
 Current Time: 13:00:39
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
 Elapsed Time: 1.382µs
 Current Time: 13:00:39
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloWorld       13:00:39.151 hello world
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.523453ms
 Current Time: 13:00:39
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 214.397111ms
 Current Time: 13:00:39
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->