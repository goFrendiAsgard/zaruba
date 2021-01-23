# 💀 Zaruba 

Declarative Task Runner Framework

# 👨‍💻 Installation

> **⚠️NOTE** You need to have `go 1.13` or higher to install zaruba. To install `go` quickly you can visit it's [official website](https://golang.org/doc/install)

Using curl

```sh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

Using wget

```sh
sh -c "$(wget -O- https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

Zaruba doesn't have any other dependency. However some of Zaruba's task depends on third party (yet commonly used) libraries/programming language.

Some of the dependencies are:

* wget, curl, and git
* netcat
* docker.io
* python3, pip, and pipenv
* nodejs, npm, and typescript

For updated list, please have a look on [devbox/Dockerfile](./devbox/Dockerfile).


# 🏁 Getting Started

```sh
# setting up
sudo -E zaruba please setupUbuntu
zaruba please setupPyenv

# create project
mkdir myProject
cd myProject
zaruba please initProject

# add external repo
zaruba please addSubrepo url="https://github.com/state-alchemists/fibonacci-clock"
zaruba please initSubrepos

cd fibonacci-clock
zaruba please serveHttp
```

While Zaruba shows you what happened in the background:

![zaruba-serveHttp](screenshots/zaruba-serveHttp.png)

You can enjoy a pretty fibonacci clock:

![fibo-clock](screenshots/fibo-clock.png)

# 🔨 Create Custom Task

Zaruba can also do a lot of other tasks. Please type `zaruba please` to see what Zaruba is capable of.

Now, to make things even more interesting, you can define custom tasks. Open up `main.zaruba.yaml` and perform this modification:

```yaml
includes:
  - "${ZARUBA_HOME}/scripts/core.zaruba.yaml"

tasks:

  runFiboClock:
    extend: core.startService
    location: fibonacci-clock
    lconfig:
      ports: [3031]
    config:
      start: "python -m http.server {{ index .LConfig.ports 0 }}"
  
  runNginx:
    extend: core.startDockerContainer
    config:
      containerName: myNginx
      imageName: nginx
      port::3030: 80
  
  runAll:
    dependencies:
      - runFiboClock
      - runNginx
```

By extending `core.startService` and `core.startDockerContainer`, you can run multiple services (either dockerized or not) with a single command.

Try to run

```sh
zaruba please runFiboClock runNginx
# or even better:
zaruba please runAll
```

![zaruba-serveHttp](screenshots/zaruba-runAll.png)

Perfect !!!

Now you can run micro-services (dockerized or not) in a single computer.

> **💡NOTE** You can add `autostop=true` argument in case of you want zaruba to kill all processes once the task is finished. E.g: `zaruba please testMyApp autostop=true`

# 🐳 Zaruba In Docker

You can use Zaruba docker-container as part of your CI/CD pipeline or as your development machine.

To run Zaruba docker-container you can run:

```sh
docker run --name zaruba -p 2810:8080 -v ${HOME}/your-project-location:/project -d stalchmst/zaruba:latest
```

This will expose Zaruba to port `2810` of your host, as well as mount `${HOME}/your-project-location` as Zaruba's current working directory.

Next time you want to re-run Zaruba's container you can simply perform:

```sh
docker start zaruba
```

# 📁 Documentation

For more comprehensive explanation, please read the [documentation](docs/Documentation.md).

# 🗺️ Roadmap

To see the future of Zaruba, plese visit our [Roadmap](Roadmap.md).

# 🐞 Bug, Feature Request and Contribution

Open [issue](https://github.com/state-alchemists/zaruba/issues) or [pull request](https://github.com/state-alchemists/zaruba/pulls).


# 🎉 Fun Fact

> Madou Ring Zaruba (魔導輪ザルバ, Madōrin Zaruba?) is the Madou Ring for Golden Knight Garo's duties as a Makai Knight. He is a recurring character in the series, acting as a guide for the wearers of the Garo armor and being the narrator of the series in some episodes. [(Garo Wiki | Fandom)](https://garoseries.fandom.com/wiki/Zaruba)
