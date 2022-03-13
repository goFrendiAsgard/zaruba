<!--startTocHeader-->
[🏠](../README.md)
# 🏃 Run task
<!--endTocHeader-->

To run any task, you can perform:

```
zaruba please -i
```

__Example:__

```bash
cd example/tasks
zaruba please -i
```
 
<details>
<summary>Output</summary>
 
```````
 Task Name
✔  printHelloWorld
 Action
✔ 🏁 Run
 Load additional value file
✔ 🏁 No
 Load additional env
✔ 🏁 No
 Auto terminate
✔ 🏁 No
  Job Starting...
 Elapsed Time: 2.656µs
 Current Time: 17:01:11
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloWorld       17:01:11.15  hello world
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.228176ms
 Current Time: 17:01:11

  Job Complete!!!
  Terminating
  Job Ended...
 Elapsed Time: 213.587264ms
 Current Time: 17:01:11
zaruba please printHelloWorld
```````
</details>
 



<!--startTocSubTopic-->
# Sub-topics
* [🍺 Run a Single Task](run-a-single-task.md)
* [🍻 Run Many Tasks in Parallel](run-many-tasks-in-parallel.md)
* [🏝️ Run Task with Custom Environments](run-task-with-custom-environments.md)
* [🔤 Run task with custom values](run-task-with-custom-values.md)
* [🏓 Run task interactively](run-task-interactively.md)
<!--endTocSubTopic-->