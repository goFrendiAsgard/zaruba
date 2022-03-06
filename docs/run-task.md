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
 Elapsed Time: 1.252µs
 Current Time: 16:50:56
  Run  'showVersion' command on /home/gofrendi/zaruba/docs
   showVersion           16:50:56.612 v0.9.0-alpha-2-a851fb02c9a8744f7197acef336a84f7dcc637ec
  Successfully running  'showVersion' command
  Job Running...
 Elapsed Time: 105.712061ms
 Current Time: 16:50:56
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 306.892528ms
 Current Time: 16:50:56
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
 Elapsed Time: 1.156µs
 Current Time: 16:50:57
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloWorld       16:50:57.052 hello world
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.208996ms
 Current Time: 16:50:57
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.596802ms
 Current Time: 16:50:57
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
 Elapsed Time: 1.175µs
 Current Time: 16:50:57
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       16:50:57.403 hello human
   printHelloWorld       16:50:57.403 hello world
  Successfully running  'printHelloHuman' command
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 101.804365ms
 Current Time: 16:50:57
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.946388ms
 Current Time: 16:50:57
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->

# Set Task Environments

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
 Elapsed Time: 1.163µs
 Current Time: 16:50:57
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       16:50:57.754 はじめまして human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.551876ms
 Current Time: 16:50:57
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.436521ms
 Current Time: 16:50:57
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
 Elapsed Time: 1.117µs
 Current Time: 16:50:58
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       16:50:58.114 Hola human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.396852ms
 Current Time: 16:50:58
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.94724ms
 Current Time: 16:50:58
zaruba please printHelloHuman -e 'sample.env'
```````
</details>
<!--endCode-->


If you don't define any environment, Zaruba will load `.env` as default value.

> ⚠️ __WARNING:__ Parameter order matters, if you set an environment twice, Zaruba will only use the __last__ one.

# Set Task Inputs

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
 Elapsed Time: 1.244µs
 Current Time: 16:50:58
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       16:50:58.473 hello Go Frendi
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.132822ms
 Current Time: 16:50:58
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 214.04701ms
 Current Time: 16:50:58
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
 Elapsed Time: 1.134µs
 Current Time: 16:50:58
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       16:50:58.829 hello Avogadro
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.329325ms
 Current Time: 16:50:58
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.585934ms
 Current Time: 16:50:59
zaruba please printHelloHuman  -v 'sample.values.yaml'
```````
</details>
<!--endCode-->



If you don't define any value, Zaruba will load `default.values.yaml` as default value.


> ⚠️ __WARNING:__ Parameter order matters, if you set an input value twice, Zaruba will only use the __last__ one.


# Run Tasks in Interactive Mode

When you run tasks in interactive mode, Zaruba will ask you to fill out [task inputs](./project/task/task-inputs.md) and `environments`. This is useful in case you cannot remember your `task inputs`/`environments`.

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
zaruba please showVersion -i
```
 
<details>
<summary>Output</summary>
 
```````
 Load additional value file
✔ 🏁 No
 Load additional env
✔ 🏁 No
  Job Starting...
 Elapsed Time: 2.456µs
 Current Time: 15:03:13
  Run  'showVersion' command on /home/gofrendi/zaruba/docs/examples/tasks
   showVersion           15:03:14.018 v0.9.0-alpha-2-adcf27c3ec0097d02bc4e7fff7f9906d92acea90
  Successfully running  'showVersion' command
  Job Running...
 Elapsed Time: 279.448377ms
 Current Time: 15:03:14

  Job Complete!!!
  Terminating
  Job Ended...
 Elapsed Time: 481.614881ms
 Current Time: 15:03:14
zaruba please showVersion
```````
</details>

or

```
zaruba please -i
```

When you run a task in interactive mode, Zaruba will also ask you to set environments and input values.


<!--startTocSubTopic-->
<!--endTocSubTopic-->