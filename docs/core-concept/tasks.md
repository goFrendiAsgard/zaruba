[⬆️](./README.md)

# Tasks

Tasks are the core of your zaruba tasks. A task define what Zaruba can do and how to do it.

Let's start by creating a very simple script.

```
gofrendi@sanctuary [17:07:47] [~/playground/example]
-> % cat > index.zaruba.yaml << EOF
heredoc> tasks:
heredoc>   sayHello:
heredoc>     start: [figlet, hello]
heredoc> EOF
```

and executing it:

```
gofrendi@sanctuary [17:12:33] [~/playground/example]
-> % zaruba please sayHello
💀 🔎 Job Starting...
         Elapsed Time: 1.3µs
         Current Time: 17:12:49
💀 🏁 Run 🍏 'sayHello' command on /home/gofrendi/playground/example
💀    🚀 sayHello             🍏 17:12:49.475  _          _ _
💀    🚀 sayHello             🍏 17:12:49.475 | |__   ___| | | ___
💀    🚀 sayHello             🍏 17:12:49.475 | '_ \ / _ \ | |/ _ \
💀    🚀 sayHello             🍏 17:12:49.475 | | | |  __/ | | (_) |
💀    🚀 sayHello             🍏 17:12:49.475 |_| |_|\___|_|_|\___/
💀    🚀 sayHello             🍏 17:12:49.475
💀 🎉 Successfully running 🍏 'sayHello' command
💀 🔎 Job Running...
         Elapsed Time: 106.3051ms
         Current Time: 17:12:49
💀 🎉 🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉
💀 🎉 Job Complete!!! 🎉🎉🎉
💀 🔥 Terminating
💀 🔎 Job Ended...
         Elapsed Time: 216.7833ms
         Current Time: 17:12:49
```

Perfect.

Now let's see script a little bit:
