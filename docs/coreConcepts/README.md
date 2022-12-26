<!--startTocHeader-->
[🏠](../README.md)
# 🧠 Core Concepts
<!--endTocHeader-->

At its core, Zaruba is a [task](task/README.md) runner.

Some tasks (like a [builtin ones](../coreTasks/README.md)) are globally accessible. But most of the tasks are [project-specific](project/README.md). Project-specific tasks are only accessible from the project directory.

You can think of a project as a container of your tasks and any related resources. In most cases, a project is also a git repository or a mono-repo.

Tasks and configurations are written in [YAML](https://en.wikipedia.org/wiki/YAML) format.

Please visit the subtopics to learn more about tasks and anything related to it.

<!--startTocSubtopic-->
- [🏗️ Project](project/README.md)
  - [🧬 Project Anatomy](project/projectAnatomy.md)
  - [🧳 Includes](project/includes.md)
  - [🔤 Project Inputs](project/projectInputs.md)
  - [⚙️ Project Configs](project/projectConfigs.md)
  - [🏝️ Project Envs](project/projectEnvs.md)
- [🔨 Task](task/README.md)
  - [🧬 Task Anatomy](task/taskAnatomy.md)
  - [🥛 Simple Command](task/simpleCommand.md)
  - [🍹 Long Running Service](task/longRunningService.md)
  - [⚙️ Task Configs](task/taskConfigs/README.md)
    - [Shared Configs](task/taskConfigs/sharedConfigs.md)
  - [🏝️ Task Envs](task/taskEnvs/README.md)
    - [Shared Envs](task/taskEnvs/sharedEnvs.md)
  - [🔤 Task Inputs](task/taskInputs.md)
  - [🧒 Extend task](task/extendTask.md)
  - [🍲 Define task dependencies](task/defineTaskDependencies.md)
- [🐹 Use Go Template](useGoTemplate.md)
<!--endTocSubtopic-->