
# InitSubrepos

File Location:

    /zaruba-tasks/chore/subrepo/task.initSubrepos.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Init subrepositories.
    ARGUMENTS:
      subrepo::<name>::prefix   : Prefix (directory name) of the subrepo
      subrepo::<name>::url      : Remote url of the subrepo
      subrepo::<name>::name     : Origin name of the subrepo
    TIPS:
      It is recommended to put `subrepo` arguments in `default.values.yaml`.
      In order to do that, you can invoke `zaruba please addSubrepo <subrepoUrl=remote-url>`



## Extends

* `zrbRunShellScript`


## Dependencies

* `zrbIsProject`
* `zrbIsValidSubrepos`


## Start

* `{{ .GetConfig "cmd" }}`
* `{{ .GetConfig "cmdArg" }}`
*
    ```
    {{ .Util.Str.Trim (.GetConfig "_setup") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "setup") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "beforeStart") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "_start") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "start") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "afterStart") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "finish") "\n " }}
    {{ .Util.Str.Trim (.GetConfig "_finish") "\n " }}

    ```


## Configs


### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigVariables") }}{{ .GetConfigsAsShellVariables "^[^_].*$" "_ZRB_CFG" }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._start


### Configs.afterStart


### Configs.beforeStart


### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.cmdArg

Value:

    -c


### Configs.finish


### Configs.setup


### Configs.shouldInitConfigMapVariable

Value:

    false


### Configs.shouldInitConfigVariables

Value:

    false


### Configs.shouldInitEnvMapVariable

Value:

    false


### Configs.shouldInitUtil

Value:

    true


### Configs.start

Value:

    {{ $names := .GetSubValueKeys "subrepo" -}}
    {{ $this := . -}}
    BRANCH="{{ if .GetValue "defaultBranch" }}{{ .GetValue "defaultBranch" }}{{ else }}main{{ end }}"
    ORIGINS=$("{{ .ZarubaBin }}" str split "$(git remote)")
    {{ range $index, $name := $names -}}
      PREFIX="{{ $this.GetValue "subrepo" $name "prefix" }}"
      URL="{{ $this.GetValue "subrepo" $name "url" }}"
      NAME="{{ $name }}"
      ORIGIN_EXISTS=$("{{ $this.ZarubaBin }}" list contain "${ORIGINS}" "${NAME}")
      if [ "$ORIGIN_EXISTS" = "1" ]
      then
        git remote set-url "${NAME}" "${URL}"
      elif [ "$ORIGIN_EXISTS" = "0" ]
      then
        echo "$NAME origin is not exist"
        gitSave "Save works before pulling from ${URL}"
        PREFIX_EXISTS=0
        if [ -d "$PREFIX" ]
        then
          PREFIX_EXISTS=1
          mv "${PREFIX}" "${PREFIX}.bak"
          gitSave "Temporarily move ${PREFIX}"
        fi
        gitInitSubrepo "${NAME}" "${PREFIX}" "${URL}" "${BRANCH}"
        if [ "$PREFIX_EXISTS" = "1" ]
        then
          rm -Rf "${PREFIX}"
          mv "${PREFIX}.bak" "${PREFIX}"
          gitSave "Restore ${PREFIX}"
        fi
      fi
    {{ end -}}
    echo 🎉🎉🎉
    echo "${_BOLD}${_YELLOW}Subrepos Initialized${_NORMAL}"



### Configs.strictMode

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1