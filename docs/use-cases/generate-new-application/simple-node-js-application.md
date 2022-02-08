<!--startTocHeader-->
[🏠](../../README.md) > [👷🏽 Use Cases](../README.md) > [Generate New Application](README.md)
# Simple NodeJs Application
<!--endTocHeader-->

To add simple node.js application, you can invoke [addSimpleNodeJsApp](../../core-tasks/addSimpleNodeJsApp)


# How to

```bash
zaruba please addSimpleNodeJsApp \
  appDirectory=<directory-name> \             # Location of your application. Must be provided
  [appName=<app-name>] \                      # application name
  [appContainerName=<app-container-name>] \   # application's container name
  [appImageName=<app-image-name>] \           # application's image name
  [appDependencies=<app-dependencies>] \      # JSON list containing names of other applications
  [appEnvs=<app-envs>]                        # JSON map containing custom environments
  [appPorts=<app-ports>]                      # JSON list containing application's ports
```

# Structure

# Use Case

```bash
zaruba please addSimpleNodeJsApp \
  appDirectory=myApp \
  [appEnvs='{"APP_HTTP_PORT":"3000"}']
```

<!--startTocSubtopic-->

<!--endTocSubtopic-->