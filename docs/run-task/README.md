<!--startTocHeader-->
[🏠](../README.md)
# 🏃 Run task
<!--endTocHeader-->

To see a list of available tasks, you can run `zaruba please` [interactively](run-task-interactively):

```
zaruba please -i
```

Once you perform the command, Zaruba will show you list of available tasks that you can select.

__Example:__

> __💡NOTE__ We use `|` operator to simulate interactive session.

<!--startCode-->
```bash
cd examples/run-tasks
( \
  echo "printHelloWorld" && \
  sleep 1 && \
  echo "" && \
  sleep 1 && \
  echo "" && \
  sleep 1 && \
  echo "" && \
  sleep 1 && \
  echo "" \
) | zaruba please -i
```
 
<details>
<summary>Output</summary>
 
```````
💀 Task Name
Search: █
? 💀 Please select task: 
  ▸ 🪂 addAirbyte
    🎐 addAirflow
    🚢 addAppHelmDeployment
    🐶 addAppRunner
    👀 addCassandra
    🟨 addClickhouse
    🧰 addContainerRegistry
    🐳 addDockerAppRunner
    🐳 addDockerComposeAppRunner
✔ 🍎 printHelloWorld
💀 Action
Search: █
? 💀 What do you want to do with printHelloWorld?: 
  ▸ 🏁 Run
✔ 🏁 Run
💀 Load additional value file
Search: █
? Do you want to load additional value file?: 
  ▸ 🏁 No
✔ 🏁 No
💀 Load additional env
Search: █
? Do you want to load additional env?: 
  ▸ 🏁 No
    📝 Yes, from file
✔ 🏁 No
    📝 Yes, from file
    📝 Yes, manually
💀 Auto terminate
Search: █
? 💀 Do you want to terminate tasks once completed?: 
  ▸ 🏁 No
✔ 🏁 No
💀 🔎 Job Starting...
         Elapsed Time: 1.407µs
         Current Time: 13:05:25
💀 🏁 Running 🍎 printHelloWorld runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
💀    🚀 🍎 printHelloWorld      hello world
💀 🎉 Successfully running 🍎 printHelloWorld runner (Attempt 1 of 3)
💀 🔎 Job Running...
         Elapsed Time: 1.615018ms
         Current Time: 13:05:25
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 404.650024ms
         Current Time: 13:05:25
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
# Sub-topics
* [🍺 Run a Single Task](run-a-single-task.md)
* [🍻 Run Many Tasks in Parallel](run-many-tasks-in-parallel.md)
* [🏝️ Run Task with Custom Environments](run-task-with-custom-environments.md)
* [🔤 Run task with custom values](run-task-with-custom-values.md)
* [🏓 Run task interactively](run-task-interactively.md)
<!--endTocSubTopic-->