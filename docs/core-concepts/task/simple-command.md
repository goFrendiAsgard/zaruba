<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🔨 Task](README.md)
# 🥛 Simple Command
<!--endTocHeader-->

Simple commands are considered `completed` once the process has been ended.

For example, `python -c "print('hello')"` is a command:

<!--startCode-->
```bash
python -c "print('hello')"
```
 
<details>
<summary>Output</summary>
 
```````
hello
```````
</details>
<!--endCode-->

We can see that once the process has been ended, the command is completed.

When you compile your Go/Typescript/Java application you are definitely running a command (even if you don't really open a terminal).

# Running Simple Command with Zaruba

Running simple command is quite trivial. You can run it directly or create a shell script to run it.

Zaruba offers two advantages when running simple command:

* You can run multiple tasks at once (i.e: `zaruba please task-1 task-2... task-n`)
* You can use re-use the task by [extending](./extend-task.md) it, or make it [pre-requisites](./define-task-dependencies.md) of other tasks.

Let's see how you can define simple command with Zaruba.

## Lower Level Approach

On lower level approach, you can make use of `start` property:

```yaml
tasks:

  printHello:
    start: [python, -c, "print('hello')"]
```

Once defined, you can run the task by invoking `zaruba please printHello`.

__Example:__

<!--startCode-->
```bash
cd examples/core-concepts/task/simple-command/low-level
zaruba please printHello
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.398µs
 Current Time: 22:13:34
  Run  'printHello' command on /home/gofrendi/zaruba/docs/examples/core-concepts/task/simple-command/low-level
   printHello            22:13:34.162 hello
  Successfully running  'printHello' command
  Job Running...
 Elapsed Time: 114.325055ms
 Current Time: 22:13:34
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 225.295275ms
 Current Time: 22:13:34
zaruba please printHello
```````
</details>
<!--endCode-->

## Higher Level Approach

For the sake of readability and avoid typos, you might want to [extend](./extend-task.md) [zrbRunShellScript](../../core-tasks/zrb-run-shell-script.md) instead:

```yaml
tasks:

  printHello:
    extend: zrbRunShellScript
    configs:
      start: python -c "print('hello')"
```

__Example:__

<!--startCode-->
```bash
cd examples/core-concepts/task/simple-command/high-level-shell
zaruba please printHello
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.177µs
 Current Time: 22:13:34
  Run  'printHello' command on /home/gofrendi/zaruba/docs/examples/core-concepts/task/simple-command/high-level-shell
   printHello            22:13:34.56  hello
  Successfully running  'printHello' command
  Job Running...
 Elapsed Time: 119.055858ms
 Current Time: 22:13:34
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 320.186303ms
 Current Time: 22:13:34
zaruba please printHello
```````
</details>
<!--endCode-->


Another way to do this is by extend [zrbRunPythonScript](../../core-tasks/zrb-run-python-script.md)

```yaml
tasks:

  printHello:
    extend: zrbRunPythonScript
    configs:
      start: print('hello')
```

__Example:__

<!--startCode-->
```bash
cd examples/core-concepts/task/simple-command/high-level-python
zaruba please printHello
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.411µs
 Current Time: 22:13:35
  Run  'printHello' command on /home/gofrendi/zaruba/docs/examples/core-concepts/task/simple-command/high-level-python
   printHello            22:13:35.045 hello
  Successfully running  'printHello' command
  Job Running...
 Elapsed Time: 115.911578ms
 Current Time: 22:13:35
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 317.451259ms
 Current Time: 22:13:35
zaruba please printHello
```````
</details>
<!--endCode-->


Here are some of the tasks you can extend when you want to run simple commands:

* [zrbRunScript](../../core-tasks/zrb-run-script.md): Lowest level
* [zrbRunShellScript](../../core-tasks/zrb-run-shell-script.md): Preferable for common use cases
* [zrbRunPythonScript](../../core-tasks/zrb-run-python-script.md): Run Python script instead of shell script
* [zrbRunNodeJsScript](../../core-tasks/zrb-run-node-js-script.md): Run Node.Js script instead of shell script

<!--startTocSubTopic-->
<!--endTocSubTopic-->
