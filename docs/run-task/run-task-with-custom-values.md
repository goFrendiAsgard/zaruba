<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🔤 Run task with custom values
<!--endTocHeader-->

Some tasks might depend on [task inputs](../core-concepts/task/task-inputs.md).

To set input values for your tasks, you can invoke:

```bash
zaruba please <task-name> -v <first-key=first-value> -v <second-key=second-value>
```

or shorter:


```bash
zaruba please <task-name> <first-key=first-value> <second-key=second-value>
```

__Example:__

<!--startCode-->
```bash
cd examples/run-tasks
zaruba please printHelloHuman humanName="Go Frendi"
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.221µs
         Current Time: 10:40:34
💀 🏁 Run 🍏 'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
💀    🚀 printHelloHuman      🍏 hello Go Frendi
💀 🎉 Successfully running 🍏 'printHelloHuman' command
💀 🔎 Job Running...
         Elapsed Time: 103.030455ms
         Current Time: 10:40:34
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 506.093251ms
         Current Time: 10:40:34
zaruba please printHelloHuman  -v 'humanName=Go Frendi'
```````
</details>
<!--endCode-->

> ⚠️ __WARNING:__ Parameter order matters, if you set an input value twice, Zaruba will only use the __last__ value.

# Load Value File

You can load `value files` by performing.

```bash
zaruba please <task-name> -v <first-file.value.yaml> -v <second-file.value.yaml>
```

__Example:__

<!--startCode-->
```bash
cd examples/run-tasks
zaruba please printHelloHuman -v sample.values.yaml
```
 
<details>
<summary>Output</summary>
 
```````
💀 🔎 Job Starting...
         Elapsed Time: 1.889µs
         Current Time: 10:40:35
💀 🏁 Run 🍏 'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
💀    🚀 printHelloHuman      🍏 hello Avogadro
💀 🎉 Successfully running 🍏 'printHelloHuman' command
💀 🔎 Job Running...
         Elapsed Time: 102.7193ms
         Current Time: 10:40:35
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 505.171621ms
         Current Time: 10:40:35
zaruba please printHelloHuman  -v 'sample.values.yaml'
```````
</details>
<!--endCode-->

If you don't define any value, Zaruba will load `default.values.yaml` as default value.

<!--startTocSubTopic-->
<!--endTocSubTopic-->