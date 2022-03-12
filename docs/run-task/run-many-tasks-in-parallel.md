<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# Run many tasks in parallel
<!--endTocHeader-->

# Run Many Tasks in Parallel

Zaruba allows you to run several tasks in parallel. To do this, you need to invoke:

```bash
zaruba please <first-task-name> <second-task-name>... <last-task-name>
```

__Example:__

<!--startCode-->
```bash
cd examples/tasks
zaruba please printHelloWorld printHelloHuman
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.528µs
 Current Time: 13:00:39
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloWorld       13:00:39.525 hello world
   printHelloHuman       13:00:39.525 hello human
  Successfully running  'printHelloHuman' command
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.364388ms
 Current Time: 13:00:39
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.716396ms
 Current Time: 13:00:39
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->