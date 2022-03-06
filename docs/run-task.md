<!--startTocHeader-->
[🏠](README.md)
# Run task
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
 Elapsed Time: 1.653µs
 Current Time: 13:58:44
  Run  'showVersion' command on /home/gofrendi/zaruba/docs
   showVersion           13:58:44.637 v0.9.0-alpha-2-adcf27c3ec0097d02bc4e7fff7f9906d92acea90
  Successfully running  'showVersion' command
  Job Running...
 Elapsed Time: 104.817825ms
 Current Time: 13:58:44
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 305.956929ms
 Current Time: 13:58:44
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
 Elapsed Time: 1.127µs
 Current Time: 13:58:45
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloWorld       13:58:45.103 hello world
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 101.976849ms
 Current Time: 13:58:45
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.472565ms
 Current Time: 13:58:45
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
 Elapsed Time: 1.156µs
 Current Time: 13:58:45
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       13:58:45.495 hello human
   printHelloWorld       13:58:45.495 hello world
  Successfully running  'printHelloHuman' command
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.400434ms
 Current Time: 13:58:45
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.543467ms
 Current Time: 13:58:45
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
 Elapsed Time: 1.471µs
 Current Time: 13:58:45
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       13:58:45.921 はじめまして human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.326906ms
 Current Time: 13:58:46
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.374285ms
 Current Time: 13:58:46
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
 Elapsed Time: 857ns
 Current Time: 13:58:46
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       13:58:46.32  Hola human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.9461ms
 Current Time: 13:58:46
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 212.603385ms
 Current Time: 13:58:46
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
zaruba please printHelloHuman name="Go Frendi"
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.062µs
 Current Time: 13:58:46
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       13:58:46.745 hello human
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 102.232181ms
 Current Time: 13:58:46
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.519261ms
 Current Time: 13:58:46
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
 Elapsed Time: 1.027µs
 Current Time: 13:58:47
  Run  'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloHuman       13:58:47.121 hello Avogadro
  Successfully running  'printHelloHuman' command
  Job Running...
 Elapsed Time: 101.709104ms
 Current Time: 13:58:47
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 213.459457ms
 Current Time: 13:58:47
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

or

```
zaruba please -i
```

When you run a task in interactive mode, Zaruba will also ask you to set environments and input values.


<!--startTocSubTopic-->
<!--endTocSubTopic-->