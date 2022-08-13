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
         Elapsed Time: 2.023µs
         Current Time: 00:24:50
💀 🏁 Running 🍎 'printHelloWorld' runner on /home/gofrendi/zaruba/docs/examples/run-tasks (Attempt: 1/3)
💀 🏁 Running 🍏 'printHelloHuman' runner on /home/gofrendi/zaruba/docs/examples/run-tasks (Attempt: 1/3)
💀    🚀 printHelloWorld      🍎 hello world
💀 🎉 Successfully running 🍎 'printHelloWorld' runner (Attempt: 1/3)
💀    🚀 printHelloHuman      🍏 hello human
💀 🎉 Successfully running 🍏 'printHelloHuman' runner (Attempt: 1/3)
💀 🔎 Job Running...
         Elapsed Time: 1.99339ms
         Current Time: 00:24:50
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 405.76486ms
         Current Time: 00:24:50
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->