<!--startTocHeader-->
[🏠](../README.md) > [👷🏽 Use Cases](README.md)
# 🏗️ Create a Project
<!--endTocHeader-->

The recommended way to create a project is by invoking `zaruba please initProject`:

__Example:__

<!--startCode-->
```bash
mkdir -p examples/playground/use-cases/newProject
cd examples/playground/use-cases/newProject
zaruba please initProject

tree
```
 
<details>
<summary>Output</summary>
 
```````
🤖 🔎 Job Starting...
         Elapsed Time: 1.53µs
         Current Time: 07:25:07
🤖 🏁 Running 🚧 initProject runner (Attempt 1 of 3) on /home/gofrendi/zaruba/docs/examples/playground/use-cases/newProject
🤖    🚀 🚧 initProject          Initialized empty Git repository in /home/gofrendi/zaruba/docs/examples/playground/use-cases/newProject/.git/
🤖 🎉 Successfully running 🚧 initProject runner (Attempt 1 of 3)
🤖 🔎 Job Running...
         Elapsed Time: 17.113521ms
         Current Time: 07:25:07
🤖 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
🤖 🎉 Job Complete!!! 🎉🎉🎉
🤖 🔥 Terminating
🤖 🔎 Job Ended...
         Elapsed Time: 320.787972ms
         Current Time: 07:25:07
zaruba please initProject  
.
├── default.values.yaml
└── index.zaruba.yaml

0 directories, 2 files
```````
</details>
<!--endCode-->

Aside from generating a project using `zaruba please initProject`, you can also clone/fork existing project from Github or other Git server. Please see [../core-concepts/projects/README.md] for more information.

# Initial Project Structure

Typically, a new project contains of two files:

* `default.values.yaml`: The default project value
* `index.zaruba.yaml`: The entry point of project's zaruba script.

# What's Next

Once you created an empty project, you can start [adding resources to your project](add-resources/README.md), [run some tasks](../run-task/README.md), and [syncrhonize task environments](syncrhonize-task-environments.md)

<!--startTocSubTopic-->
<!--endTocSubTopic-->
