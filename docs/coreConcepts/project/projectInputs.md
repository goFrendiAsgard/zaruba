<!--startTocHeader-->
[🏠](../../README.md) > [🧠 Core Concepts](../README.md) > [🏗️ Project](README.md)
# 🔤 Project Inputs
<!--endTocHeader-->

Project inputs make your tasks more configurable and interactive.

Once defined, a project input will be available globally. That means several tasks might share the same input.

Defining a project input is pretty simple. Here is the anatomy of project input:

```yaml
inputs:

  inputName:                         # your input name
    default: defaultValue            # default value of the input
    description: input description   # description of the input
    options: [option1, option2]      # options available for the input (will be shown in interactive mode)
    prompt: inputPrompt              # input prompt (will be shown in interactive mode)
    allowCustom: true                # if set to true, user will be able to put any values in interactive mode (even if the value is not in the `options`)
    secret: false                    # if set to true, the value will be treated as secret and will not be stored anywhere
```

To use project input in your task, please see [task's input](../tasks/task-inputs.md).

<!--startTocSubtopic-->

<!--endTocSubtopic-->