<!--startTocHeader-->
[🏠](../../README.md) > [🔧 Utilities](../README.md) > [🏝️ Env](README.md)
# Print
<!--endTocHeader-->

# Usage

<!--startCode-->
```bash
zaruba env print --help
```

````
Print environment

Usage:
  zaruba env print <envMap> [prefix] [flags]

Flags:
  -h, --help   help for print
````
<!--endCode-->

# Examples

## Print Environment

<!--startCode-->
```bash
zaruba env print \
    '{"APP_CONTEXT": "aws", "APP_VERSION": "1.0.0"}'
```

````
APP_CONTEXT="aws"
APP_VERSION="1.0.0"
````
<!--endCode-->

## Print Cascaded Environment

Cascaded environment is very useful if you manage several platforms with similar environments.

For example, you have `dev`, `stag`, and `prod`.

By adding `DEV` prefix (i.e: `zaruba env print <map> DEV`), your `DEV` environment will be assigned to the original ones.

If prefixed environment is not exists, Zaruba will use original values. Thus you can treat original environment as fallback/default values.


<!--startCode-->
```bash
zaruba env print \
    '{"APP_CONTEXT": "aws", "APP_VERSION": "1.0.0", "STAG_CONTEXT": "azure", "STAG_APP_VERSION": "1.1.0", "DEV_APP_CONTEXT": "gcp", "DEV_APP_VERSION": "1.1.1"}' \
    DEV
```

````
APP_CONTEXT="gcp"
APP_VERSION="1.1.1"
DEV_APP_CONTEXT="gcp"
DEV_APP_VERSION="1.1.1"
STAG_APP_VERSION="1.1.0"
STAG_CONTEXT="azure"
````
<!--endCode-->


<!--startTocSubTopic-->
<!--endTocSubTopic-->