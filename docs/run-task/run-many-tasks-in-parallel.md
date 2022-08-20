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
         Elapsed Time: 1.295µs
         Current Time: 11:15:56
💀 🏁 Running 🍎 printHelloWorld runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
💀 🏁 Running 🍏 printHelloHuman runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
💀    🚀 🍎 printHelloWorld      hello world
💀 🎉 Successfully running 🍎 printHelloWorld runner (Attempt 1 of 3)
💀    🚀 🍏 printHelloHuman      hello human
💀 🎉 Successfully running 🍏 printHelloHuman runner (Attempt 1 of 3)
💀 🔎 Job Running...
         Elapsed Time: 2.833868ms
         Current Time: 11:15:56
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 405.180289ms
         Current Time: 11:15:56
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->