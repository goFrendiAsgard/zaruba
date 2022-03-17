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
 Elapsed Time: 1.814µs
 Current Time: 07:56:26
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloHuman       07:56:26.49  hello human
   printHelloWorld       07:56:26.49  hello world
  Successfully running  'printHelloHuman' command
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.337336ms
 Current Time: 07:56:26
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.621343ms
 Current Time: 07:56:26
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->