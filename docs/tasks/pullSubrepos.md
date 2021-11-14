
# PullSubrepos

`File Location`:

    /zaruba-tasks/chore/subrepo/task.pullSubrepos.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:

    Pull subrepositories.
    ARGUMENTS:
      subrepo::<name>::prefix   : Prefix (directory name) of the subrepo
      subrepo::<name>::url      : Remote url of the subrepo




## Extends

* `zrbRunShellScript`


## Dependencies

* `initSubrepos`
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


## Check




## Inputs


## Configs

`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`_start`:




`finish`:




`includeShellUtil`:

    true


`setup`:




`start`:

    set -e
    {{ $d := .Decoration -}}
    {{ $names := .GetSubValueKeys "subrepo" -}}
    {{ $this := . -}}
    ORIGINS=$("{{ .ZarubaBin }}" str split "$(git remote)")
    BRANCH="{{ if .GetValue "defaultBranch" }}{{ .GetValue "defaultBranch" }}{{ else }}main{{ end }}"
    {{ range $index, $name := $names -}}
      PREFIX="{{ $this.GetValue "subrepo" $name "prefix" }}"
      URL="{{ $this.GetValue "subrepo" $name "url" }}"
      NAME="{{ $name }}"
      ORIGIN_EXISTS=$("{{ $this.ZarubaBin }}" list contain "${ORIGINS}" "${NAME}")
      if [ $ORIGIN_EXISTS = 1 ]
      then
        gitSave "Save works before pull"
        git subtree pull --prefix="${PREFIX}" "${NAME}" "${BRANCH}"
      fi
    {{ end -}}
    echo 🎉🎉🎉
    echo "{{ $d.Bold }}{{ $d.Yellow }}Subrepos pulled{{ $d.Normal }}"



`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`afterStart`:




`beforeStart`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`cmdArg`:

    -c


`strictMode`:

    true


`_finish`:





## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1