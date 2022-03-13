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
cd examples/tasks
zaruba please printHelloWorld printHelloHuman
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.369µs
 Current Time: 12:34:04
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       12:34:04.47  hello human
   printHelloWorld       12:34:04.47  hello world
  Successfully running  'printHelloHuman' command
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.159621ms
 Current Time: 12:34:04
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.054161ms
 Current Time: 12:34:04
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->