<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# 🍺 Run a Single Task
<!--endTocHeader-->

You can run a single task by providing it's name.

There are two types of task:

* __Globally-accessible task__: You can execute these tasks from anywhere
* __Project-specific task__: You can only execute these tasks the [project](../coreConcepts/project/README.md)'s directory.

# Run a Globally Accessible Task

To execute globally accessible task, you can invoke `zaruba please` from anywhere:

```bash
zaruba please <task-name>
```

__Example:__


```bash
zaruba please showSolidPrinciple
```
 
<details>
<summary>Output</summary>
 
```````
🤖 🔎 Job Starting...
         Elapsed Time: 1.287µs
         Current Time: 09:10:18
🤖 🏁 Running 🦉 showSolidPrinciple runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs
🤖 🎉 Successfully running 🦉 showSolidPrinciple runner (Attempt 1 of 3)
🤖    🚀 🦉 showSolidPrinciple    S  Single Responsibility Principle 
🤖    🚀 🦉 showSolidPrinciple    O  Open/Closed Principle 
🤖    🚀 🦉 showSolidPrinciple    L  Liskov's Substitution Principle 
🤖 🔎 Job Running...
         Elapsed Time: 6.011442ms
         Current Time: 09:10:18
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖    🚀 🦉 showSolidPrinciple    I  Interface Segregation Principle 
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖    🚀 🦉 showSolidPrinciple    D  Dependency Inversion Principle 
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 309.079912ms
         Current Time: 09:10:18
zaruba please showSolidPrinciple
```````
</details>


 By default, Zaruba provides some [builtin core tasks](../coreTasks/README.md) that are globally accessible.
 
 To make your tasks globally accessible, you need to add them to `ZARUBA_SCRIPTS` environment variable. Please refer to [zaruba configuration](../configuration.md) for more information.

# Run a Project Specific Task

To execute any project-specific tasks, you need to be in the project's directory first:

```bash
cd <project-directory>
zaruba please <task-name>
```

Please note that the command will not work from the project's subdirectory.

__Example:__


```bash
cd examples/run-tasks
zaruba please printHelloWorld
```
 
<details>
<summary>Output</summary>
 
```````
🤖 🔎 Job Starting...
         Elapsed Time: 1.39µs
         Current Time: 09:10:19
🤖 🏁 Running 🍎 printHelloWorld runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/run-tasks
🤖 🎉 Successfully running 🍎 printHelloWorld runner (Attempt 1 of 3)
🤖    🚀 🍎 printHelloWorld      hello world
🤖 🔎 Job Running...
         Elapsed Time: 1.602969ms
         Current Time: 09:10:19
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 407.092577ms
         Current Time: 09:10:19
zaruba please printHelloWorld
```````
</details>



<!--startTocSubtopic-->

<!--endTocSubtopic-->