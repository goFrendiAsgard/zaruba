[⬅️ Table of Content](../../README.md)

# Simple Typescript Application

To add simple typeScript application, you can invoke [addSimpleTypeScriptApp](../../core-tasks/addSimpleTypeScriptApp)


# How to

```bash
zaruba please addSimpleTypeScriptApp \
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
zaruba please addSimpleTypeScriptApp \
  appDirectory=myApp \
  [appEnvs='{"APP_HTTP_PORT":"3000"}']
```