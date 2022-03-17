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
 Elapsed Time: 1.322µs
 Current Time: 09:26:36
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloWorld       09:26:36.71  hello world
   printHelloHuman       09:26:36.71  hello human
  Successfully running  'printHelloHuman' command
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.720623ms
 Current Time: 09:26:36
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.125079ms
 Current Time: 09:26:36
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->