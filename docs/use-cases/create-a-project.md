[⬅️ Table of Content](../README.md)
# Create a Project

The recommended way to create a project is by invoking `zaruba please initProject`:

```bash
~/playground on ☁️  (ap-southeast-1) on ☁️  gofrendi@kata.ai
❯ mkdir myproject

~/playground on ☁️  (ap-southeast-1) on ☁️  gofrendi@kata.ai
❯ cd myproject

~/playground/myproject on ☁️  (ap-southeast-1) on ☁️  gofrendi@kata.ai
❯ zaruba please initProject
💀 🔎 Job Starting...
         Elapsed Time: 1.2µs
         Current Time: 07:10:25
💀 🏁 Run 🚧 'initProject' command on /home/gofrendi/playground/myproject
💀    🚀 initProject          🚧 07:10:25.647 Initialized empty Git repository in /home/gofrendi/playground/myproject/.git/
💀    🚀 initProject          🚧 07:10:25.654 🎉🎉🎉
💀    🚀 initProject          🚧 07:10:25.654 Project created
💀 🎉 Successfully running 🚧 'initProject' command
💀 🔎 Job Running...
         Elapsed Time: 311.8279ms
         Current Time: 07:10:25
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 512.9968ms
         Current Time: 07:10:25
```

# Initial Project Structure

Once created, you will have two files:

```
myproject on  master [?] on ☁️  (ap-southeast-1) on ☁️  gofrendi@kata.ai
❯ tree
.
├── default.values.yaml
└── index.zaruba.yaml

0 directories, 2 files
```

* `default.values.yaml` is your default project value
* `index.zaruba.yaml` is the entry point of your zaruba script.

# Next Step

Depending on your use case, you can:

* [generate new application](./generate-new-application.md)
* [add third party services](./add-third-party-service.md)
* [add subrepo to your project](./add-subrepo.md)
* [add runner for existing application](./add-runner-for-existing-application/README.md)
