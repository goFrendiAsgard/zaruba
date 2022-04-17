<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🏓 Run task interactively
<!--endTocHeader-->

When you run tasks in interactive mode, Zaruba will ask you to fill some [inputs](../core-concepts/task/task-inputs.md) and [environments](../core-concepts/task/task-envs/README.md).

To run a task in interactive mode you can invoke:

```bash
zaruba please <task-name> -i
```

or

```bash
zaruba please <first-task-name> <second-task-name> -i
```

__Example:__

<!--startCode-->
```bash
cd examples/run-tasks
( \
  echo "" && \
  sleep 1 && \
  echo "" && \
  sleep 1 && \
  echo "let" && \
  echo "" && \
  sleep 1 && \
  echo "Robert Boyle" \
) | zaruba please printHelloHuman -i
```
 
<details>
<summary>Output</summary>
 
```````
💀 Load additional value file
Search: █
? Do you want to load additional value file?: 
  ▸ 🏁 No
    📝 Yes
Search: █? Do you want to load additional value file?:   ▸ 🏁 No    📝 Yes✔ 🏁 No
💀 Load additional env
Search: █
? Do you want to load additional env?: 
  ▸ 🏁 No
    📝 Yes, from file
    📝 Yes, manually
Search: █? Do you want to load additional env?:   ▸ 🏁 No    📝 Yes, from file    📝 Yes, manually✔ 🏁 No
💀 1 of 1) humanName
Search: █
? Your name: 
  ▸ human
    Let me type it!
Search: l█? Your name:   ▸ Let me type it!Search: le█? Your name:   ▸ Let me type it!Search: let█? Your name:   ▸ Let me type it!Search: let█? Your name:   ▸ Let me type it!✔ Let me type it!
✔ Your name: █
✔ Your name: R█
✔ Your name: Ro█
✔ Your name: Rob█
✔ Your name: Robe█
✔ Your name: Rober█
✔ Your name: Robert█
✔ Your name: Robert █
✔ Your name: Robert B█
✔ Your name: Robert Bo█
✔ Your name: Robert Boy█
✔ Your name: Robert Boyl█
✔ Your name: Robert Boyle█
✔ Your name: Robert Boyle█
Your name: Robert Boyle
💀 🔎 Job Starting...
         Elapsed Time: 1.4µs
         Current Time: 11:19:03
💀 🏁 Run 🍏 'printHelloHuman' command on /home/gofrendi/zaruba/docs/examples/run-tasks
💀    🚀 printHelloHuman      🍏 11:19:03.465 hello Robert Boyle
💀 🎉 Successfully running 🍏 'printHelloHuman' command
💀 🔎 Job Running...
         Elapsed Time: 101.827861ms
         Current Time: 11:19:03
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 212.18079ms
         Current Time: 11:19:03
zaruba please printHelloHuman  -v 'humanName=Robert Boyle'
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->