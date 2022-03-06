<!--startTocHeader-->
[🏠](README.md)
# 🏃 Run task
<!--endTocHeader-->

# Run a Task

## Run a Builitin Task

To execute builtin [core-tasks](../core-tasks/README.md), you can invoke `zaruba please` from anywhere:

```bash
zaruba please <task-name>
```

__Example:__

<!--startCode-->
```bash
zaruba please showVersion
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.822µs
 Current Time: 17:14:35
  Run  'showVersion' command on /home/gofrendi/zaruba/docs
   showVersion           17:14:35.641 v0.9.0-alpha-2-a851fb02c9a8744f7197acef336a84f7dcc637ec
  Successfully running  'showVersion' command
  Job Running...
 Elapsed Time: 107.377017ms
 Current Time: 17:14:35
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 308.898853ms
 Current Time: 17:14:35
zaruba please showVersion
```````
</details>
<!--endCode-->

## Run a Project Specific Task

But, to execute any [project](./project/README.md) specific tasks, you need to be in the project directory first:

```bash
cd <project-directory>
zaruba please <task-name>
```

__Example:__

<!--startCode-->
```bash
cd examples/tasks
zaruba please printHelloWorld
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.629µs
 Current Time: 17:14:36
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloWorld       17:14:36.1   hello world
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.513823ms
 Current Time: 17:14:36
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.938856ms
 Current Time: 17:14:36
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


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
 Elapsed Time: 1.897µs
 Current Time: 17:14:36
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       17:14:36.463 hello human
   printHelloWorld       17:14:36.463 hello world
  Successfully running  'printHelloHuman' command
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.580506ms
 Current Time: 17:14:36
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.312821ms
 Current Time: 17:14:36
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->

# Run Task with Custom Environments

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
 Elapsed Time: 1.295µs
 Current Time: 17:14:36
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       17:14:36.842 はじめまして human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.680833ms
 Current Time: 17:14:36
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.716443ms
 Current Time: 17:14:37
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
 Elapsed Time: 1.649µs
 Current Time: 17:14:37
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       17:14:37.202 Hola human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.788677ms
 Current Time: 17:14:37
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.680611ms
 Current Time: 17:14:37
zaruba please printHelloHuman -e 'sample.env'
```````
</details>
<!--endCode-->


If you don't define any environment, Zaruba will load `.env` as default value.

> ⚠️ __WARNING:__ Parameter order matters, if you set an environment twice, Zaruba will only use the __last__ one.

# Run Task with Custom Values

Some tasks might dependes on [task inputs](./project/task/task-inputs.md). To set input values for your tasks, you can invoke:

```bash
zaruba please <task-name> -v <first-key=first-value> -v <second-key=second-value>
```

or shorter:


```bash
zaruba please <task-name> <first-key=first-value> <second-key=second-value>
```

__Example:__

<!--startCode-->
```bash
cd examples/tasks
zaruba please printHelloHuman humanName="Go Frendi"
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.567µs
 Current Time: 17:14:37
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       17:14:37.56  hello Go Frendi
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.363895ms
 Current Time: 17:14:37
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.638197ms
 Current Time: 17:14:37
zaruba please printHelloHuman  -v 'humanName=Go Frendi'
```````
</details>
<!--endCode-->


You can also load `value files`.

```bash
zaruba please <task-name> -v <first-file.value.yaml> -v <second-file.value.yaml>
```

__Example:__

<!--startCode-->
```bash
cd examples/tasks
zaruba please printHelloHuman -v sample.values.yaml
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.83µs
 Current Time: 17:14:37
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       17:14:37.914 hello Avogadro
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.607486ms
 Current Time: 17:14:38
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.616932ms
 Current Time: 17:14:38
zaruba please printHelloHuman  -v 'sample.values.yaml'
```````
</details>
<!--endCode-->



If you don't define any value, Zaruba will load `default.values.yaml` as default value.


> ⚠️ __WARNING:__ Parameter order matters, if you set an input value twice, Zaruba will only use the __last__ one.


# Run Task Interactively

When you run tasks in interactive mode, Zaruba will ask you to fill out [task inputs](./project/task/task-inputs.md) and `environments`.

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
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
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


# Run Any Task

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
<!--endTocSubTopic-->