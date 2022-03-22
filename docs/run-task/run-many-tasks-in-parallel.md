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
 Elapsed Time: 1.244µs
 Current Time: 07:56:52
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloHuman       07:56:52.02  hello human
   printHelloWorld       07:56:52.02  hello world
  Successfully running  'printHelloWorld' command
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.545541ms
 Current Time: 07:56:52
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.317122ms
 Current Time: 07:56:52
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->