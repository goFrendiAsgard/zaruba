<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🏝️ Run Task with Custom Environments
<!--endTocHeader-->

To set [task's environment](../core-concepts/task/task-envs/README.md) you can use `-e` flag as follow:

```bash
zaruba please <task-name> -e <first-key=first-value> -e <second-key=second-value>
```

__Example:__

<!--startCode-->
```bash
cd examples/run-tasks
zaruba please printHelloHuman -e GREETINGS=はじめまして
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.596µs
         Current Time: 00:24:51
💀 🏁 Running 🍏 'printHelloHuman' runner on /home/gofrendi/zaruba/docs/examples/run-tasks (Attempt: 1/3)
💀    🚀 printHelloHuman      🍏 はじめまして human
💀 🎉 Successfully running 🍏 'printHelloHuman' runner (Attempt: 1/3)
💀 🔎 Job Running...
         Elapsed Time: 1.348148ms
         Current Time: 00:24:51
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 405.600128ms
         Current Time: 00:24:51
zaruba please printHelloHuman -e 'GREETINGS=はじめまして'
```````
</details>
<!--endCode-->


> ⚠️ __WARNING:__ Parameter order matters, if you set a single environment twice, Zaruba will only use the __last__ value.

# Load Environment File

You can also load `environments` from files:

```bash
zaruba please <task-name> -e <first-file.env> -e <second-file.env>
```

__Example:__

<!--startCode-->
```bash
cd examples/run-tasks
zaruba please printHelloHuman -e sample.env
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.49µs
         Current Time: 00:24:51
💀 🏁 Running 🍏 'printHelloHuman' runner on /home/gofrendi/zaruba/docs/examples/run-tasks (Attempt: 1/3)
💀    🚀 printHelloHuman      🍏 Hola human
💀 🎉 Successfully running 🍏 'printHelloHuman' runner (Attempt: 1/3)
💀 🔎 Job Running...
         Elapsed Time: 1.659436ms
         Current Time: 00:24:51
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 403.454544ms
         Current Time: 00:24:52
zaruba please printHelloHuman -e 'sample.env'
```````
</details>
<!--endCode-->

>  ⚠️ __WARNING:__  If you don't define any environment, Zaruba will load `.env` in current directory.

<!--startTocSubTopic-->
<!--endTocSubTopic-->