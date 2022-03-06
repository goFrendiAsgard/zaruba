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
 Elapsed Time: 1.733µs
 Current Time: 17:52:55
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       17:52:55.087 hello human
   printHelloWorld       17:52:55.087 hello world
  Successfully running  'printHelloHuman' command
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.228322ms
 Current Time: 17:52:55
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 214.35835ms
 Current Time: 17:52:55
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->