<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🏝️ Run Task with Custom Environments
<!--endTocHeader-->

You can set [task's environment](../core-concepts/task/task-envs/README.md) by performing:

```bash
zaruba please <task-name> -e <first-key=first-value> -e <second-key=second-value>
```

__Example:__

<!--startCode-->
```bash
cd examples/run-tasks
zaruba please printHelloHuman -e GREETINGS=はじめまして
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 850ns
 Current Time: 22:13:05
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloHuman       22:13:05.583 はじめまして human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.927935ms
 Current Time: 22:13:05
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.538734ms
 Current Time: 22:13:05
zaruba please printHelloHuman -e 'GREETINGS=はじめまして'
```````
</details>
<!--endCode-->


> ⚠️ __WARNING:__ Parameter order matters, if you set an environment twice, Zaruba will only use the __last__ one.

# Load Environment File

You can also load `environments` from files:

```bash
zaruba please <task-name> -e <first-file.env> -e <second-file.env>
```

__Example:__

<!--startCode-->
```bash
cd examples/run-tasks
zaruba please printHelloHuman -e sample.env
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.342µs
 Current Time: 22:13:05
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloHuman       22:13:05.958 Hola human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.933036ms
 Current Time: 22:13:06
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.778343ms
 Current Time: 22:13:06
zaruba please printHelloHuman -e 'sample.env'
```````
</details>
<!--endCode-->

>  ⚠️ __WARNING:__  If you don't define any environment, Zaruba will load `.env` as default value.

<!--startTocSubTopic-->
<!--endTocSubTopic-->