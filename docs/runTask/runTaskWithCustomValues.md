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


```bash
cd examples/run-tasks
zaruba please printHelloHuman humanName="Go Frendi"
```
 
<details>
<summary>Output</summary>
 
```````
🤖 🔎 Job Starting...
         Elapsed Time: 1.231µs
         Current Time: 09:10:21
🤖 🏁 Running 🍏 printHelloHuman runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
🤖    🚀 🍏 printHelloHuman      hello Go Frendi
🤖 🎉 Successfully running 🍏 printHelloHuman runner (Attempt 1 of 3)
🤖 🔎 Job Running...
         Elapsed Time: 2.210287ms
         Current Time: 09:10:21
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 408.803047ms
         Current Time: 09:10:21
zaruba please printHelloHuman  -v 'humanName=Go Frendi'
```````
</details>


> ⚠️ __WARNING:__ Parameter order matters, if you set an input value twice, Zaruba will only use the __last__ value.

# Load Value File

You can load `value files` by performing.

```bash
zaruba please <task-name> -v <first-file.value.yaml> -v <second-file.value.yaml>
```

__Example:__


```bash
cd examples/run-tasks
zaruba please printHelloHuman -v sample.values.yaml
```
 
<details>
<summary>Output</summary>
 
```````
🤖 🔎 Job Starting...
         Elapsed Time: 1.265µs
         Current Time: 09:10:22
🤖 🏁 Running 🍏 printHelloHuman runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
🤖    🚀 🍏 printHelloHuman      hello Avogadro
🤖 🎉 Successfully running 🍏 printHelloHuman runner (Attempt 1 of 3)
🤖 🔎 Job Running...
         Elapsed Time: 1.786066ms
         Current Time: 09:10:22
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 406.711553ms
         Current Time: 09:10:22
zaruba please printHelloHuman  -v 'sample.values.yaml'
```````
</details>


If you don't define any value, Zaruba will load `default.values.yaml` as default value.

<!--startTocSubtopic-->

<!--endTocSubtopic-->