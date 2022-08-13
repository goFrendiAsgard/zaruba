<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🍻 Run Many Tasks in Parallel
<!--endTocHeader-->

# Run Many Tasks in Parallel

Zaruba allows you to run many tasks in parallel and see the output on a single screen. To do this, you need to invoke:

```bash
zaruba please <first-task-name> <second-task-name>... <last-task-name>
```

__Example:__

<!--startCode-->
```bash
cd examples/run-tasks
zaruba please printHelloWorld printHelloHuman
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.552µs
         Current Time: 19:26:01
💀 🏁 Run 🍎 'printHelloWorld' on /home/gofrendi/zaruba/docs/examples/run-tasks (Attempt 1)
💀 🏁 Run 🍏 'printHelloHuman' on /home/gofrendi/zaruba/docs/examples/run-tasks (Attempt 1)
💀    🚀 printHelloWorld      🍎 hello world
💀    🚀 printHelloHuman      🍏 hello human
💀 🎉 Successfully running 🍎 'printHelloWorld' runner
💀 🎉 Successfully running 🍏 'printHelloHuman' runner
💀 🔎 Job Running...
         Elapsed Time: 102.16206ms
         Current Time: 19:26:02
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 504.479227ms
         Current Time: 19:26:02
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->