<!--startTocHeader-->
[🏠](../README.md) > [Motivation and architecture](README.md)
# Feature flags
<!--endTocHeader-->

Feature flags, also known as feature toggles, are used in software development to enable or disable functionality without having to deploy new code. You can compare them with an `if` statement that checks whether a code path should be available or not.

ZtplAppDirectory use a lot of feature flags that you can toggle to active/1 or inactive/0:

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

You can see how ZtplAppDirectory handle feature flags in `main.py`. Basically, it is just simple if-else:

```python
enable_ui = os.getenv('APP_ENABLE_UI', '1') == '1'
if enable_ui:
    serve_ui()
```

# Interface and Layers

Next, you can continue to [interface and layers](interface-and-layers.md).

<!--startTocSubTopic-->
<!--endTocSubTopic-->