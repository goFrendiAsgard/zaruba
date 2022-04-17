<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🏝️ Env](README.md)
# Get
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba env get --help
```
 
<details>
<summary>Output</summary>
 
```````
Get envmap from currently loaded environment variables

Usage:
  zaruba env get [prefix] [flags]

Flags:
  -h, --help   help for get
```````
</details>
<!--endCode-->

# Examples

> __Note:__ For demo purpose, we ignore `global environments` by performing `env –i [Var=Value]… command args…`. For real use case, you can do `zaruba env get` directly.

## Get Environment as Map

<!--startCode-->
```bash
ENV_DICT=$(env -i \
    PATH="${ZARUBA_HOME}" \
    APP_CONTEXT=aws \
    APP_VERSION=1.0.0 \
    APP_PLATFORM=kubernetes \
    zaruba env get )

echo "💀 Environment:"
echo "${ENV_DICT}"

echo ""
echo "💀 APP_CONTEXT:"
zaruba map get "${ENV_DICT}" APP_CONTEXT

echo ""
echo "💀 APP_VERSION:"
zaruba map get "${ENV_DICT}" APP_VERSION

echo ""
echo "💀 APP_PLATFORM:"
zaruba map get "${ENV_DICT}" APP_PLATFORM
```
 
<details>
<summary>Output</summary>
 
```````
💀 Environment:
{"APP_CONTEXT":"aws","APP_PLATFORM":"kubernetes","APP_VERSION":"1.0.0","PATH":"/home/gofrendi/zaruba","ZARUBA_BIN":"/home/gofrendi/zaruba/zaruba","ZARUBA_DECORATION":"default","ZARUBA_HOME":"/home/gofrendi/zaruba","ZARUBA_SCRIPTS":"","ZARUBA_SHELL":"bash"}

💀 APP_CONTEXT:
aws

💀 APP_VERSION:
1.0.0

💀 APP_PLATFORM:
kubernetes
```````
</details>
<!--endCode-->

## Get Cascaded Environment as Map

Cascaded environment is very useful if you manage several platforms with similar environments.

For example, you have `dev`, `stag`, and `prod`.

By adding `DEV` prefix (i.e: `zaruba env get DEV`), your `DEV` environment will be assigned to the original ones.

If prefixed environment is not exists, Zaruba will use original values. Thus you can treat original environment as fallback/default values.

<!--startCode-->
```bash
ENV_DICT=$(env -i \
    PATH="${ZARUBA_HOME}" \
    APP_CONTEXT=aws \
    APP_VERSION=1.0.0 \
    APP_PLATFORM=kubernetes \
    STAG_APP_CONTEXT=azure \
    STAG_APP_VERSION=1.1.0 \
    DEV_APP_CONTEXT=gcp \
    DEV_APP_VERSION=1.1.1 \
    zaruba env get DEV )

echo "💀 Environment:"
echo "${ENV_DICT}"

echo ""
echo "💀 APP_CONTEXT:"
zaruba map get "${ENV_DICT}" APP_CONTEXT

echo ""
echo "💀 APP_VERSION:"
zaruba map get "${ENV_DICT}" APP_VERSION

echo ""
echo "💀 APP_PLATFORM:"
zaruba map get "${ENV_DICT}" APP_PLATFORM
```
 
<details>
<summary>Output</summary>
 
```````
💀 Environment:
{"APP_CONTEXT":"aws","APP_PLATFORM":"kubernetes","APP_VERSION":"1.0.0","DEV_APP_CONTEXT":"gcp","DEV_APP_VERSION":"1.1.1","PATH":"/home/gofrendi/zaruba","STAG_APP_CONTEXT":"azure","STAG_APP_VERSION":"1.1.0","ZARUBA_BIN":"/home/gofrendi/zaruba/zaruba","ZARUBA_DECORATION":"default","ZARUBA_HOME":"/home/gofrendi/zaruba","ZARUBA_SCRIPTS":"","ZARUBA_SHELL":"bash"}

💀 APP_CONTEXT:
aws

💀 APP_VERSION:
1.0.0

💀 APP_PLATFORM:
kubernetes
```````
</details>
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->