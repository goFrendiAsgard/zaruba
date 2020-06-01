# 💀 Zaruba 

```
      _____
     /     \        My name is 💀Zaruba.
    | () () |       I came to be when Garo came to be.
     \  ^  /        And I am forever with him.
      |||||

Zaruba is an agnostic scaffolding tool, service runner, as well as monorepo management tool.
Zaruba will help you build project faster.

Usage:
  zaruba <action> [...args] [flags]
  zaruba [command]

Available Commands:
  create            Create component.
  help              Help about any command
  init              Init project.
  install-template  Install template.
  organize          Organize project.
  pull              Pull project.
  push              Push project.
  remove-containers Remove all containers.
  run               Run project.
  stop-containers   Stop all containers.
  version           Zaruba's version.

Flags:
  -h, --help   help for zaruba

```

**What Zaruba is**

Zaruba is an agnostic scaffolding tool, service runner, as well as monorepo management tool. Zaruba will help you build and test your project faster.

**Who Zaruba is for**

Do you have to run several tmux panels in order to test your app? Do you need to copy-pasting codes from one app to another one during development phase? Do you want to publish some parts of your monorepo as open-source projects?

If you answer "yes" for any question above, then Zaruba is for you.

**How to use Zaruba**

* 👨‍💻 [Installing Zaruba](#install)
* 🔨 [Using Zaruba as scaffolding tool](#scaffolding)
* ✈️ [Using Zaruba as service runner](#service-runner)
* 🐙 [Using Zaruba as monorepo management tool](#monorepo-management)

# <a name="install"></a> 👨‍💻 Installing Zaruba

In order to install and use zaruba, you need to set couple of things:

* WSL/Linux/Mac
* Git
* Docker
* Golang
* Curl or Wget

To install Zaruba using `curl`, you should invoke:

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

To install Zaruba using `wget`, you should invoke:

```bash
sh -c "$(wget -O- https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
```

## Configuring Zaruba

Zaruba injects several environment variables to your `.bashrc` and `.zshrc`. You can find Zaruba's environment definition at `~/.zaruba/zaruba.env`:

```bash
if [ -z ${HOME} ]
then
    HOME=~
fi
export PATH="$(go env GOPATH)/bin:${PATH}"
export ZARUBA_SHELL="/bin/bash"
export ZARUBA_SHELL_ARG="-c"
export ZARUBA_TEMPLATE_DIR="${HOME}/.zaruba/templates"
```


* `ZARUBA_TEMPLATE_DIR` is the location of your templates. You can install template by performing `zaruba install-template <git-url>`
* `ZARUBA_SHELL` define shell used by zaruba (for running command). It doesn't have to be the same as your default one. For example, you might use `zsh` as your default terminal and `bash` as zaruba's terminal.
* `ZARUBA_SHELL_ARG` default to `-c`

つづく


# <a name="scaffolding"></a> 🔨 Using Zaruba as scaffolding tool

# <a name="service-runner"></a> ✈️ Using Zaruba as service runner

# <a name="monorepo-management"></a> 🐙 Using Zaruba as monorepo management tool
