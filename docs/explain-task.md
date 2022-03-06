<!--startTocHeader-->
[🏠](README.md)
# 🔎 Explain task
<!--endTocHeader-->

To explain a task, you can invoke `zaruba please` with `-x` flag:

```bash
zaruba please <task-name> -x
```

__Example:__

<!--startCode-->
```bash
cd examples/tasks
zaruba please printHelloHuman -x
```
 
<details>
<summary>Output</summary>
 
```````
## Information

File Location:

    /home/gofrendi/zaruba/docs/examples/tasks/index.zaruba.yaml

Should Sync Env:

    false

Type:

    command


## Start

* `bash`
* `-c`
* `echo ${GREETINGS} ${ZARUBA_INPUT_HUMAN_NAME}`


## Inputs


### Inputs.humanName

Prompt:

    Your name

Default Value:

    human

Secret:

    false


## Envs


### Envs.GREETINGS

From:

    GREETINGS

Default:

    hello
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->