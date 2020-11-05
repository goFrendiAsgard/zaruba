# 💀 Zaruba 

Declarative Task Runner Framework

# Installation

Using curl

```sh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

Using wget

```sh
sh -c "$(wget -O- https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

> **⚠️NOTE** You need to have `go 1.13` or higher to install zaruba. To install `go` quickly you can visit it's [official website](https://golang.org/doc/install)

# Getting Started

```sh
mkdir myProject
cd myProject

zaruba please initProject

zaruba please addSubrepo name=fiboclock prefix=fibo-clock url="https://github.com/therealvasanth/fibonacci-clock"
zaruba please initSubrepos

cd fibo-clock
zaruba please serveHttp
```

While Zaruba shows you what happened in the background:

![zaruba-serveHttp](screenshots/zaruba-serveHttp.png)

You can enjoy a pretty fibonacci clock:

![fibo-clock](screenshots/fibo-clock.png)

# Create Custom Task

Zaruba can also do a lot of other tasks. Please type `zaruba please` to see what Zaruba is capable of.

Now, to make things even more interesting, you can define custom tasks. Open up `main.zaruba.yaml` and perform this modification:

```yaml
includes:
  - "${ZARUBA_HOME}/scripts/core.zaruba.yaml"

tasks:

  runFiboClock:
    extend: core.startService
    lconfig:
      ports: [3031]
    start: ["python", "-m", "http.server", "{{ index .LConfig.ports 0 }}"]
  
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


# Documentation

For more comprehensive explanation, please read the [documentation](docs/Documentation.md).

# Roadmap

To see the future of Zaruba, plese visit our [Roadmap](roadmap.md).


# Fun Fact

> Madou Ring Zaruba (魔導輪ザルバ, Madōrin Zaruba?) is the Madou Ring for Golden Knight Garo's duties as a Makai Knight. He is a recurring character in the series, acting as a guide for the wearers of the Garo armor and being the narrator of the series in some episodes. [(Garo Wiki | Fandom)](https://garoseries.fandom.com/wiki/Zaruba)
