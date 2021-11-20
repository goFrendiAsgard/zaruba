
# PushSubrepos

File Location:

    /zaruba-tasks/chore/subrepo/task.pushSubrepos.yaml

Should Sync Env:

    true

Type:

    command

Description:

    Publish subrepositories.
    ARGUMENTS:
      subrepo::<name>::prefix   : Prefix (directory name) of the subrepo
      subrepo::<name>::url      : Remote url of the subrepo



## Extends

* `zrbRunShellScript`


## Dependencies

* `initSubrepos`
* `updateProjectLinks`
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


### Configs.cmdArg

Value:

    -c


### Configs.finish


### Configs.setup


### Configs._setup

Value:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


### Configs._start


### Configs.afterStart


### Configs.beforeStart


### Configs.start

Value:

    set -e
    {{ $d := .Decoration -}}
    {{ $names := .GetSubValueKeys "subrepo" -}}
    {{ $this := . -}}
    BRANCH="{{ if .GetValue "defaultBranch" }}{{ .GetValue "defaultBranch" }}{{ else }}main{{ end }}"
    ORIGINS=$("{{ .ZarubaBin }}" str split "$(git remote)")
    {{ range $index, $name := $names -}}
      PREFIX="{{ $this.GetValue "subrepo" $name "prefix" }}"
      URL="{{ $this.GetValue "subrepo" $name "url" }}"
      NAME="{{ $name }}"
      ORIGIN_EXISTS=$("{{ .ZarubaBin }}" list contain "${ORIGINS}" "${NAME}")
      if [ $ORIGIN_EXISTS = 1 ]
      then
        gitSave.sh" "Save works before p
        git subtree push --prefix="${PREFIX}" "${NAME}" "${BRANCH}"
      fi
    {{ end -}}
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Subrepos pushed{{ $d.Normal }}"



### Configs.strictMode

Value:

    true


### Configs._finish


### Configs._initShell

Value:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



### Configs.cmd

Value:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


### Configs.includeShellUtil

Value:

    true


## Envs


### Envs.PYTHONUNBUFFERED

From:

    PYTHONUNBUFFERED

Default:

    1