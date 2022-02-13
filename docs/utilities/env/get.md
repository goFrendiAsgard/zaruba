<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [Env](README.md)
# Get
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba env get --help
```

````
Get envmap from currently loaded environment variables

Usage:
  zaruba env get [prefix] [flags]

Flags:
  -h, --help   help for get
````
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
    zaruba env get )

echo "🐶 Environment:"
echo "${ENV_DICT}"

echo ""
echo "🐶 APP_CONTEXT:"
zaruba map get "${ENV_DICT}" APP_CONTEXT

echo ""
echo "🐶 APP_VERSION:"
zaruba map get "${ENV_DICT}" APP_VERSION
```

````
🐶 Environment:
{"APP_CONTEXT":"aws","APP_VERSION":"1.0.0","PATH":"/home/gofrendi/zaruba","ZARUBA_BIN":"/home/gofrendi/zaruba/zaruba","ZARUBA_HOME":"/home/gofrendi/zaruba","ZARUBA_SHELL":"bash"}

🐶 APP_CONTEXT:
aws

🐶 APP_VERSION:
1.0.0
````
<!--endCode-->

## Get Cascaded Environment as Map

<!--startCode-->
```bash
ENV_DICT=$(env -i \
    PATH="${ZARUBA_HOME}" \
    APP_CONTEXT=aws \
    APP_VERSION=1.0.0 \
    DEV_APP_CONTEXT=gcp \
    DEV_APP_VERSION=1.1.0 \
    zaruba env get DEV )

echo "🐶 Environment:"
echo "${ENV_DICT}"

echo ""
echo "🐶 APP_CONTEXT:"
zaruba map get "${ENV_DICT}" APP_CONTEXT

echo ""
echo "🐶 APP_VERSION:"
zaruba map get "${ENV_DICT}" APP_VERSION
```

````
🐶 Environment:
{"APP_CONTEXT":"gcp","APP_VERSION":"1.0.0","DEV_APP_CONTEXT":"gcp","DEV_APP_VERSION":"1.1.0","PATH":"/home/gofrendi/zaruba","ZARUBA_BIN":"/home/gofrendi/zaruba/zaruba","ZARUBA_HOME":"/home/gofrendi/zaruba","ZARUBA_SHELL":"bash"}

🐶 APP_CONTEXT:
gcp

🐶 APP_VERSION:
1.0.0
````
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->