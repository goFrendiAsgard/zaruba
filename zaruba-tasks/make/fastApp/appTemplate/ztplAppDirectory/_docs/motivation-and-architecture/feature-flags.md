<!--startTocHeader-->
[🏠](../README.md) > [Motivation and architecture](README.md)
# Feature flags
<!--endTocHeader-->

Feature flags, also known as [feature toggles](https://en.wikipedia.org/wiki/Feature_toggle), are used in software development to enable or disable functionality without having to deploy new code.

`ZtplAppDirectory` use a lot of feature flags that you can set to active (1) or inactive (0):

```bash
# HTTP handler
APP_ENABLE_ROUTE_HANDLER=1 # Whether to serve HTTP request or not
APP_ENABLE_UI=1            # Whether to serve UI or not. Only works if APP_ENABLE_ROUTE_HANDLER==1
APP_ENABLE_API=1           # Whether to serve API or not. Only works if APP_ENABLE_ROUTE_HANDLER==1

# Event handler
APP_ENABLE_EVENT_HANDLER=1

# RPC
APP_ENABLE_RPC_HANDLER=1

# Auth module
APP_ENABLE_AUTH_MODULE=1

# Other module
APP_ENABLE_<OTHER>_MODULE=1
```

You can set `ZtplAppDirectory` feature flag by using [environment variable](https://en.wikipedia.org/wiki/Environment_variable).

You can see how ZtplAppDirectory handle feature flags in `main.py` or `configs/featureFlag.py`.

```python
# The value of `enable_ui` will be `False` if APP_ENABLE_UI` is set to `0`.
# Otherwise, the value of this variable will be `True`
enable_ui = os.getenv('APP_ENABLE_UI', '1') != '0'
if enable_ui:
    # The UI will only be served
    serve_ui()
```

Feature flag mechanism is very simple. It will run a feature only if the corresponding feature flag is `True`.

# Interface and Layers

Next, you can continue to [interface and layers](interface-and-layers.md).

<!--startTocSubTopic-->
<!--endTocSubTopic-->