<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🏝️ Run Task with Custom Environments
<!--endTocHeader-->

To set [task's environment](../core-concepts/task/task-envs/README.md) you can use `-e` flag as follow:

```bash
zaruba please <task-name> -e <first-key=first-value> -e <second-key=second-value>
```

__Example:__


```bash
cd examples/run-tasks
zaruba please printHelloHuman -e GREETINGS=はじめまして
```
 
<details>
<summary>Output</summary>
 
```````
🤖 🔎 Job Starting...
         Elapsed Time: 1.145µs
         Current Time: 09:10:20
🤖 🏁 Running 🍏 printHelloHuman runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
🤖 🎉 Successfully running 🍏 printHelloHuman runner (Attempt 1 of 3)
🤖 🔎 Job Running...
         Elapsed Time: 2.163297ms
         Current Time: 09:10:20
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖    🚀 🍏 printHelloHuman      はじめまして human
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 408.166284ms
         Current Time: 09:10:20
zaruba please printHelloHuman -e 'GREETINGS=はじめまして'
```````
</details>



> ⚠️ __WARNING:__ Parameter order matters, if you set a single environment twice, Zaruba will only use the __last__ value.

# Load Environment File

You can also load `environments` from files:

```bash
zaruba please <task-name> -e <first-file.env> -e <second-file.env>
```

__Example:__


```bash
cd examples/run-tasks
zaruba please printHelloHuman -e sample.env
```
 
<details>
<summary>Output</summary>
 
```````
🤖 🔎 Job Starting...
         Elapsed Time: 1.263µs
         Current Time: 09:10:20
🤖 🏁 Running 🍏 printHelloHuman runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
🤖 🎉 Successfully running 🍏 printHelloHuman runner (Attempt 1 of 3)
🤖 🔎 Job Running...
         Elapsed Time: 1.668896ms
         Current Time: 09:10:20
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖    🚀 🍏 printHelloHuman      Hola human
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 404.911823ms
         Current Time: 09:10:21
zaruba please printHelloHuman -e 'sample.env'
```````
</details>


>  ⚠️ __WARNING:__  If you don't define any environment, Zaruba will load `.env` in current directory.

# Subtopics
<!--startTocSubtopic-->
<!--endTocSubtopic-->