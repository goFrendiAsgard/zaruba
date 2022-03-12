<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# Run task with custom environments
<!--endTocHeader-->

You can set `environments` for your tasks by performing:

```bash
zaruba please <task-name> -e <first-key=first-value> -e <second-key=second-value>
```

__Example:__

<!--startCode-->
```bash
cd examples/tasks
zaruba please printHelloHuman -e GREETINGS=はじめまして
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.236µs
 Current Time: 13:00:39
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       13:00:39.881 はじめまして human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.741759ms
 Current Time: 13:00:39
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.688571ms
 Current Time: 13:00:40
zaruba please printHelloHuman -e 'GREETINGS=はじめまして'
```````
</details>
<!--endCode-->

You can also load those `environments` from files:

```bash
zaruba please <task-name> -e <first-file.env> -e <second-file.env>
```

__Example:__

<!--startCode-->
```bash
cd examples/tasks
zaruba please printHelloHuman -e sample.env
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 2.149µs
 Current Time: 13:00:40
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       13:00:40.251 Hola human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.747077ms
 Current Time: 13:00:40
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.390943ms
 Current Time: 13:00:40
zaruba please printHelloHuman -e 'sample.env'
```````
</details>
<!--endCode-->


If you don't define any environment, Zaruba will load `.env` as default value.

> ⚠️ __WARNING:__ Parameter order matters, if you set an environment twice, Zaruba will only use the __last__ one.



<!--startTocSubTopic-->
<!--endTocSubTopic-->