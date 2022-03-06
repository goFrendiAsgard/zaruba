<!--startTocHeader-->
[🏠](../README.md) > [🏃 Run task](README.md)
# Run a single task
<!--endTocHeader-->

# Run a Builitin Task

To execute builtin [core-tasks](../core-tasks/README.md), you can invoke `zaruba please` from anywhere:

```bash
zaruba please <task-name>
```

__Example:__

<!--startCode-->
```bash
zaruba please showVersion
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.853µs
 Current Time: 17:52:54
  Run  'showVersion' command on /home/gofrendi/zaruba/docs
   showVersion           17:52:54.292 v0.9.0-alpha-2-a851fb02c9a8744f7197acef336a84f7dcc637ec
  Successfully running  'showVersion' command
  Job Running...
 Elapsed Time: 108.301699ms
 Current Time: 17:52:54
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 309.832165ms
 Current Time: 17:52:54
zaruba please showVersion
```````
</details>
<!--endCode-->

# Run a Project Specific Task

But, to execute any [project](./project/README.md) specific tasks, you need to be in the project directory first:

```bash
cd <project-directory>
zaruba please <task-name>
```

__Example:__

<!--startCode-->
```bash
cd examples/tasks
zaruba please printHelloWorld
```
 
<details>
<summary>Output</summary>
 
```````
Job Starting...
 Elapsed Time: 1.522µs
 Current Time: 17:52:54
  Run  'printHelloWorld' command on /home/gofrendi/zaruba/docs/examples/tasks
   printHelloWorld       17:52:54.737 hello world
  Successfully running  'printHelloWorld' command
  Job Running...
 Elapsed Time: 102.596788ms
 Current Time: 17:52:54
  
  Job Complete!!! 
  Terminating
  Job Ended...
 Elapsed Time: 214.111038ms
 Current Time: 17:52:54
zaruba please printHelloWorld
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->