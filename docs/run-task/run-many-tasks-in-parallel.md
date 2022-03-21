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
 Elapsed Time: 1.136µs
 Current Time: 06:43:32
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloWorld       06:43:32.195 hello world
   printHelloHuman       06:43:32.195 hello human
  Successfully running  'printHelloHuman' command
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.419179ms
 Current Time: 06:43:32
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.841792ms
 Current Time: 06:43:32
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->