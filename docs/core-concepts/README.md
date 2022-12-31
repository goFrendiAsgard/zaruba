<!--startTocHeader-->
[🏠](../README.md)
# 🧠 Core Concepts
<!--endTocHeader-->

At its core, Zaruba is a [task](task/README.md) runner.


Some tasks (like a [builtin ones](../core-tasks/README.md)) are globally accessible. But most of the tasks are [project-specific](project/README.md). Project-specific tasks are only accessible from the project directory.

You can think of a project as a container of your tasks and any related resources. In most cases, a project is also a git repository or a mono-repo.

Tasks and configurations are written in [YAML](https://en.wikipedia.org/wiki/YAML) format.

Please visit the subtopics to learn more about tasks and anything related to it.

<!--startTocSubtopic-->
# Subtopics
- [🏗️ Project](project/README.md)
  - [🧬 Project Anatomy](project/project-anatomy.md)
  - [🧳 Includes](project/includes.md)
  - [🔤 Project Inputs](project/project-inputs.md)
  - [⚙️ Project Configs](project/project-configs.md)
  - [🏝️ Project Envs](project/project-envs.md)
- [🔨 Task](task/README.md)
  - [🧬 Task Anatomy](task/task-anatomy.md)
  - [🥛 Simple Command](task/simple-command.md)
  - [🍹 Long Running Service](task/long-running-service.md)
  - [⚙️ Task Configs](task/task-configs/README.md)
    - [Shared Configs](task/task-configs/shared-configs.md)
  - [🏝️ Task Envs](task/task-envs/README.md)
    - [Shared Envs](task/task-envs/shared-envs.md)
  - [🔤 Task Inputs](task/task-inputs.md)
  - [🧒 Extend task](task/extend-task.md)
  - [🍲 Define task dependencies](task/define-task-dependencies.md)
- [🐹 Use Go Template](use-go-template.md)
<!--endTocSubtopic-->