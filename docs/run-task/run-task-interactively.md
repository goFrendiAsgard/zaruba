<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🏓 Run task interactively
<!--endTocHeader-->

When you run tasks in interactive mode, Zaruba will ask you to fill in some [inputs](../core-concepts/task/task-inputs.md) and [environments](../core-concepts/task/task-envs/README.md).

To run a task in interactive mode you can invoke:

```bash
zaruba please <task-name> -i
```

or

```bash
zaruba please <first-task-name> <second-task-name> -i
```

__Example:__

> __💡NOTE__ We use `|` operator to simulate interactive session.


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
🤖 Load additional value file
Search: █
? Do you want to load additional value file?: 
  ▸ 🏁 No
✔ 🏁 No
🤖 Load additional env
Search: █
? Do you want to load additional env?: 
  ▸ 🏁 No
    📝 Yes, from file
✔ 🏁 No
🤖 1 of 1) humanName
Search: █
? Your name: 
  ▸ human
✔ Let me type it!
Your name: Robert Boyle
🤖 🔎 Job Starting...
         Elapsed Time: 1.478µs
         Current Time: 09:10:25
🤖 🏁 Running 🍏 printHelloHuman runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
🤖    🚀 🍏 printHelloHuman      hello Robert Boyle
🤖 🎉 Successfully running 🍏 printHelloHuman runner (Attempt 1 of 3)
🤖 🔎 Job Running...
         Elapsed Time: 1.5956ms
         Current Time: 09:10:25
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 407.815983ms
         Current Time: 09:10:26
zaruba please printHelloHuman  -v 'humanName=Robert Boyle'
```````
</details>


# Subtopics
<!--startTocSubtopic-->
<!--endTocSubtopic-->