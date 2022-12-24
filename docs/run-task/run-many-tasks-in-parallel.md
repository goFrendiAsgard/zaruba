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
🤖 🔎 Job Starting...
         Elapsed Time: 1.985µs
         Current Time: 07:51:53
🤖 🏁 Running 🍎 printHelloWorld runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
🤖 🏁 Running 🍏 printHelloHuman runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
🤖    🚀 🍏 printHelloHuman      hello human
🤖 🎉 Successfully running 🍎 printHelloWorld runner (Attempt 1 of 3)
🤖 🎉 Successfully running 🍏 printHelloHuman runner (Attempt 1 of 3)
🤖 🔎 Job Running...
         Elapsed Time: 2.048246ms
         Current Time: 07:51:53
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖    🚀 🍎 printHelloWorld      hello world
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 406.676468ms
         Current Time: 07:51:53
zaruba please printHelloWorld printHelloHuman
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->