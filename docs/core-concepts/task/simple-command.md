<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🔨 Task](README.md)
# 🥛 Simple Command
<!--endTocHeader-->

Simple commands like `cat`, `ls`, `grep`, and `echo` are not meant to run forever. Once `completed` a simple-command will return an exit status.

Any tasks with similar behavior are called `simple-command`.

For example, `python -c "print('hello')"` is a simple command:

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

# Running Simple Command with Zaruba

Running simple command is trivial. You can invoke it from the CLI or you can create a shell script to run it.

Running simple command with Zaruba gives you several advantages:

* You can run many simple commands in parallel (i,e., `zaruba please task-1 task-2... task-n`)
* You can use re-use the command by [extending](./extend-task.md) it
* You can define some [pre-requisites](./define-task-dependencies.md) for your command.
* If you run many commands in parallel, you can see their logs in real time.

Let's see how you can define simple command with Zaruba.

## Lower Level Approach

In lower-level approach, you can make use of `start` property:

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
💀 🔎 Job Starting...
         Elapsed Time: 1.74µs
         Current Time: 21:01:30
💀 🏁 Running 🍏 printHello runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/core-concepts/task/simple-command/low-level
💀    🚀 🍏 printHello           hello
💀 🎉 Successfully running 🍏 printHello runner (Attempt 1 of 3)
💀 🔎 Job Running...
         Elapsed Time: 27.927428ms
         Current Time: 21:01:30
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 431.514954ms
         Current Time: 21:01:30
zaruba please printHello
```````
</details>
<!--endCode-->

## Higher Level Approach

Instead of accessing `start` property directly, you can [extend](./extend-task.md) [zrbRunShellScript](../../core-tasks/zrb-run-shell-script.md) as follows:

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
💀 🔎 Job Starting...
         Elapsed Time: 2.485µs
         Current Time: 21:01:30
💀 🏁 Running 🍏 printHello runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/core-concepts/task/simple-command/high-level-shell
💀    🚀 🍏 printHello           hello
💀 🎉 Successfully running 🍏 printHello runner (Attempt 1 of 3)
💀 🔎 Job Running...
         Elapsed Time: 30.240225ms
         Current Time: 21:01:30
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 333.018577ms
         Current Time: 21:01:31
zaruba please printHello
```````
</details>
<!--endCode-->


Another way to do this is by extending [zrbRunPythonScript](../../core-tasks/zrb-run-python-script.md)

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
💀 🔎 Job Starting...
         Elapsed Time: 1.276µs
         Current Time: 21:01:31
💀 🏁 Running 🍏 printHello runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/core-concepts/task/simple-command/high-level-python
💀    🚀 🍏 printHello           hello
💀 🎉 Successfully running 🍏 printHello runner (Attempt 1 of 3)
💀 🔎 Job Running...
         Elapsed Time: 20.399784ms
         Current Time: 21:01:31
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 324.698789ms
         Current Time: 21:01:31
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
