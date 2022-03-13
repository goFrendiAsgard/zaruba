<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🏓 Run task interactively
<!--endTocHeader-->

When you run tasks in interactive mode, Zaruba will ask you to fill some [inputs](../core-concepts/task/task-inputs.md) and [environments](../core-concepts/task/task-envs/README.md).

To run a task in interactive mode you can invoke:

```bash
zaruba please <task-name> -i
```

or

```bash
zaruba please <first-task-name> <second-task-name> -i
```

__Example:__

```bash
cd example/tasks
zaruba please printHelloHuman -i
```
 
<details>
<summary>Output</summary>
 
```````
 Load additional value file
✔ 🏁 No
 Load additional env
✔ 🏁 No
 1 of 1) humanName
✔ Let me type it!
Your name: Robert Boyle
  Job Starting...
 Elapsed Time: 1.362µs
 Current Time: 16:53:55
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloHuman       16:53:55.483 hello Robert Boyle
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.91258ms
 Current Time: 16:53:55

  Job Complete!!!
  Terminating
  Job Ended...
 Elapsed Time: 212.914199ms
 Current Time: 16:53:55
zaruba please printHelloHuman  -v 'humanName=Robert Boyle'
```````
</details>



<!--startTocSubTopic-->
<!--endTocSubTopic-->