<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🍻 Run Many Tasks in Parallel
<!--endTocHeader-->

# Run Many Tasks in Parallel

Zaruba allows you to run several tasks in parallel. To do this, you need to invoke:

```bash
zaruba please <first-task-name> <second-task-name>... <last-task-name>
```

__Example:__

<!--startCode-->
```bash
cd examples/run-tasks
zaruba please printHelloWorld printHelloHuman
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.107µs
 Current Time: 22:13:05
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloWorld       22:13:05.204 hello world
   printHelloHuman       22:13:05.204 hello human
  Successfully running  'printHelloWorld' command
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.577494ms
 Current Time: 22:13:05
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 214.787469ms
 Current Time: 22:13:05
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->