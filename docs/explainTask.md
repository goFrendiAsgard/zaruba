<!--startTocHeader-->
[🏠](README.md)
# 🔎 Explain task
<!--endTocHeader-->

To see detailed information about a specific task, you can invoke `zaruba please` with `-x` flag:

```bash
zaruba please <task-name> -x
```

__Example:__


```bash
cd examples/run-tasks
zaruba please printHelloHuman -x
```
 
<details>
<summary>Output</summary>
 
```````
## Information

File Location:

    /home/gofrendi/zaruba/docs/examples/run-tasks/index.zaruba.yaml

Should Sync Env:

    false

Type:

    simple


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



<!--startTocSubtopic-->

<!--endTocSubtopic-->