<!--startTocHeader-->
[🏠](../README.md) > [🥝 Core Tasks](README.md)
# 📦 initSubrepos
<!--endTocHeader-->

[1m[33m## Information[0m

[1m[34mFile Location[0m:

    ~/.zaruba/zaruba-tasks/chore/subrepo/task.initSubrepos.yaml

[1m[34mShould Sync Env[0m:

    true

[1m[34mType[0m:

    command

[1m[34mDescription[0m:

    Init subrepositories.
    ARGUMENTS:
      subrepo::<name>::prefix   : Prefix (directory name) of the subrepo
      subrepo::<name>::url      : Remote url of the subrepo
      subrepo::<name>::name     : Origin name of the subrepo
    TIPS:
      It is recommended to put `subrepo` arguments in `default.values.yaml`.
      In order to do that, you can invoke `zaruba please addSubrepo <subrepoUrl=remote-url>`



[1m[33m## Extends[0m

* [zrbRunShellScript](zrb-run-shell-script.md)


[1m[33m## Dependencies[0m

* [zrbIsProject](zrb-is-project.md)
* [zrbIsValidSubrepos](zrb-is-valid-subrepos.md)


[1m[33m## Start[0m

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


[1m[33m## Configs[0m


[1m[33m### Configs._finish[0m


[1m[33m### Configs._initShell[0m

[1m[34mValue[0m:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ else }}{{ "" -}}{{ end }}
    {{ $d := .Decoration -}}
    {{ $d.ToShellVariables }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitConfigMapVariable") }}_ZRB_CONFIG_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetConfigs "^[^_].*$")) }}{{ else }}{{ "" -}}{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "shouldInitEnvMapVariable") }}_ZRB_ENV_MAP={{ .Util.Str.EscapeShellValue (.Util.Json.FromStringDict (.GetEnvs)) }}{{ else }}{{ "" -}}{{ end }}



[1m[33m### Configs._setup[0m

[1m[34mValue[0m:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


[1m[33m### Configs._start[0m


[1m[33m### Configs.afterStart[0m


[1m[33m### Configs.beforeStart[0m


[1m[33m### Configs.cmd[0m

[1m[34mValue[0m:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


[1m[33m### Configs.cmdArg[0m

[1m[34mValue[0m:

    -c


[1m[33m### Configs.finish[0m


[1m[33m### Configs.setup[0m


[1m[33m### Configs.shouldInitConfigMapVariable[0m

[1m[34mValue[0m:

    false


[1m[33m### Configs.shouldInitEnvMapVariable[0m

[1m[34mValue[0m:

    false


[1m[33m### Configs.shouldInitUtil[0m

[1m[34mValue[0m:

    true


[1m[33m### Configs.start[0m

[1m[34mValue[0m:

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



[1m[33m### Configs.strictMode[0m

[1m[34mValue[0m:

    true


[1m[33m## Envs[0m


[1m[33m### Envs.PYTHONUNBUFFERED[0m

[1m[34mFrom[0m:

    PYTHONUNBUFFERED

[1m[34mDefault[0m:

    1