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


```bash
cd examples/run-tasks
zaruba please printHelloWorld printHelloHuman
```
 
<details>
<summary>Output</summary>
 
```````
🤖 🔎 Job Starting...
         Elapsed Time: 1.19µs
         Current Time: 09:10:19
🤖 🏁 Running 🍎 printHelloWorld runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
🤖 🏁 Running 🍏 printHelloHuman runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
🤖    🚀 🍎 printHelloWorld      hello world
🤖 🎉 Successfully running 🍎 printHelloWorld runner (Attempt 1 of 3)
🤖 🎉 Successfully running 🍏 printHelloHuman runner (Attempt 1 of 3)
🤖    🚀 🍏 printHelloHuman      hello human
🤖 🔎 Job Running...
         Elapsed Time: 2.56427ms
         Current Time: 09:10:19
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 406.988857ms
         Current Time: 09:10:20
zaruba please printHelloWorld printHelloHuman
```````
</details>


<!--startTocSubtopic--><!--endTocSubtopic-->