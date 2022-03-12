<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🔨 Task](README.md)
# Simple Command
<!--endTocHeader-->


Simple command is something you run and considered `completed` once the process has been ended.

For example, `python -c "print('hello')"` is a command:

```
❯ python -c "print('hello')"
hello
```

We can see that once the process has been ended, the command is completed. When you compile your Go/Typescript/Java application you are definitely running a command (even if you don't really open a terminal).

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

## Higher Level Approach

For the sake of readability and avoid typos, you might want to [extend](./extend-task.md) [zrbRunShellScript](../../../core-tasks/zrbRunShellScript.md) instead:

```yaml
tasks:

  printHello:
    extend: zrbRunShellScript
    configs:
      start: python -c "print('hello')"
```

Another way to do this is by extend [zrbRunPythonScript](../../../core-tasks/zrbRunPythonScript.md)

```yaml
tasks:

  printHello:
    extend: zrbRunPythonScript
    configs:
      start: print('hello')
```

Here are some of the tasks you can extend when you want to run simple commands:

* [zrbRunScript](../../../core-tasks/zrbRunShellScript.md): Lowest level
* [zrbRunShellScript](../../../core-tasks/zrbRunShellScript.md): Preferable for common use cases
* [zrbRunPythonScript](../../../core-tasks/zrbRunPythonScript.md): Run Python script instead of shell script
* [zrbRunNodeJsScript](../../../core-tasks/zrbRunNodeJsScript.md): Run Node.Js script instead of shell script

<!--startTocSubTopic-->
<!--endTocSubTopic-->
