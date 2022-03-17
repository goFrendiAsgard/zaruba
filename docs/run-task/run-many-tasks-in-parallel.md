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
 Elapsed Time: 1.56µs
 Current Time: 17:12:32
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloWorld       17:12:32.943 hello world
   printHelloHuman       17:12:32.943 hello human
  Successfully running  'printHelloWorld' command
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.725919ms
 Current Time: 17:12:33
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.314574ms
 Current Time: 17:12:33
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->