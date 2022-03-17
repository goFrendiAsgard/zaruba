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
 Elapsed Time: 1.145µs
 Current Time: 17:12:33
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloHuman       17:12:33.33  はじめまして human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.373604ms
 Current Time: 17:12:33
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.363434ms
 Current Time: 17:12:33
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
 Elapsed Time: 1.176µs
 Current Time: 17:12:33
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
   printHelloHuman       17:12:33.716 Hola human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.037796ms
 Current Time: 17:12:33
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.123954ms
 Current Time: 17:12:33
zaruba please printHelloHuman -e 'sample.env'
```````
</details>
<!--endCode-->

>  ⚠️ __WARNING:__  If you don't define any environment, Zaruba will load `.env` as default value.

<!--startTocSubTopic-->
<!--endTocSubTopic-->