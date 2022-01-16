[⬅️](../README.md)

# Creating a project

A Zaruba project is a directory containing at least a file named `index.zaruba.yaml`.

To create a Zaruba project, you can:

* Simply create a file named `index.zaruba.yaml` in any directory.
* Use command: `zaruba please initProject` command.

In most cases, you want to use the command instead of creating the file by yourself.

When you init a project using a command, Zaruba will also turn your directory into git repository. Furtherm, Zaruba will also generate `default.value.yaml` that contains default values for you project:

```
gofrendi@sanctuary [12:14:44] [~/playground]
-> % mkdir myProject
gofrendi@sanctuary [12:14:49] [~/playground]
-> % cd myProject
gofrendi@sanctuary [12:14:53] [~/playground/myProject]
-> % zaruba please initProject
💀 🔎 Job Starting...
         Elapsed Time: 35µs
         Current Time: 12:14:58
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/playground/myProject
💀    🚀 initProject          🚧 12:14:58.849 Initialized empty Git repository in /home/gofrendi/playground/myProject/.git/
💀    🚀 initProject          🚧 12:14:58.857 🎉🎉🎉
💀    🚀 initProject          🚧 12:14:58.857 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 331.9442ms
         Current Time: 12:14:58
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 533.8407ms
         Current Time: 12:14:59
gofrendi@sanctuary [12:14:59] [~/playground/myProject] [master *]
-> % tree
.
├── default.values.yaml
└── index.zaruba.yaml

0 directories, 2 files
gofrendi@sanctuary [12:15:00] [~/playground/myProject] [master *]
-> % cat default.values.yaml
defaultBranch: master
defaultImagePrefix: ''
defaultImageTag: latest
defaultKubeContext: docker-desktop
defaultKubeNamespace: default
hostDockerInternal: host.docker.internal
pulumiUseLocalBackend: false
gofrendi@sanctuary [12:16:06] [~/playground/myProject] [master *]
-> %
```