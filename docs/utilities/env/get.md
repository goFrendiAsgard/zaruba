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
Get current environment variables as jsonMap.

You can cascade the environment variable using --prefix flag.
This is useful if you have multiple environments (e.g., dev, staging, prod)

Usage:
  zaruba env get [flags]

Examples:

> export SERVER=localhost
> export PORT=3306

> zaruba env get
{"SERVER": "localhost", "PORT": "3306"}

> export SERVER=localhost
> export STG_SERVER=stg.stalchmst.com
> export PROD_SERVER=stalchmst.com
> export PORT=3306

> zaruba env get --prefix=STG
{"SERVER": "stg.stalchmst.com", "PORT": "3306", "STG_SERVER": "stg.stalchmst.com", "PROD_SERVER": "stalchmst.com"}

> zaruba env get --prefix=PROD
{"SERVER": "stalchmst.com", "PORT": "3306", "STG_SERVER": "stg.stalchmst.com", "PROD_SERVER": "stalchmst.com"}

> zaruba env get --prefix=DEV
{"SERVER": "localhost", "PORT": "3306", "STG_SERVER": "stg.stalchmst.com", "PROD_SERVER": "stalchmst.com"}


Flags:
  -h, --help            help for get
  -p, --prefix string   environment prefix
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
{"APP_CONTEXT":"aws","APP_PLATFORM":"kubernetes","APP_VERSION":"1.0.0","PATH":"/home/gofrendi/zaruba","ZARUBA_BIN":"/home/gofrendi/zaruba/zaruba","ZARUBA_DECORATION":"default","ZARUBA_ENV":"","ZARUBA_HOME":"/home/gofrendi/zaruba","ZARUBA_LOG_STATUS_LINE_INTERVAL":"40","ZARUBA_LOG_STATUS_TIME_INTERVAL":"5m","ZARUBA_LOG_TIME":"true","ZARUBA_MAX_LOG_FILE_SIZE":"5242880","ZARUBA_SCRIPTS":"","ZARUBA_SHELL":"bash"}

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

By adding `DEV` prefix (i,e., `zaruba env get DEV`), your `DEV` environment will be assigned to the original ones.

If prefixed environment does not exists, Zaruba will use original values. Thus you can treat original environment as fallback/default values.

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
{"APP_CONTEXT":"aws","APP_PLATFORM":"kubernetes","APP_VERSION":"1.0.0","DEV_APP_CONTEXT":"gcp","DEV_APP_VERSION":"1.1.1","PATH":"/home/gofrendi/zaruba","STAG_APP_CONTEXT":"azure","STAG_APP_VERSION":"1.1.0","ZARUBA_BIN":"/home/gofrendi/zaruba/zaruba","ZARUBA_DECORATION":"default","ZARUBA_ENV":"","ZARUBA_HOME":"/home/gofrendi/zaruba","ZARUBA_LOG_STATUS_LINE_INTERVAL":"40","ZARUBA_LOG_STATUS_TIME_INTERVAL":"5m","ZARUBA_LOG_TIME":"true","ZARUBA_MAX_LOG_FILE_SIZE":"5242880","ZARUBA_SCRIPTS":"","ZARUBA_SHELL":"bash"}

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