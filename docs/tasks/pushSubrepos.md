
# PushSubrepos

`File Location`:

    /zaruba-tasks/chore/subrepo/task.pushSubrepos.yaml


`Location`:




`Should Sync Env`:

    true


`Sync Env Location`:




`Type`:

    command


`Description`:

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


## Check




## Inputs


## Configs

`setup`:




`_setup`:

    {{ .Util.Str.Trim (.GetConfig "_initShell") "\n" }}


`_start`:




`beforeStart`:




`cmd`:

    {{ if .GetValue "defaultShell" }}{{ .GetValue "defaultShell" }}{{ else }}bash{{ end }}


`cmdArg`:

    -c


`finish`:




`includeShellUtil`:

    true


`start`:

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



`_finish`:




`_initShell`:

    {{ if .Util.Bool.IsTrue (.GetConfig "strictMode") }}set -e{{ else }}set +e{{ end }}
    {{ if .Util.Bool.IsTrue (.GetConfig "includeShellUtil") }}. {{ .ZarubaHome }}/zaruba-tasks/_base/run/bash/shellUtil.sh{{ end }}



`afterStart`:




`strictMode`:

    true



## Envs


### Envs.PYTHONUNBUFFERED

`From`:

    PYTHONUNBUFFERED


`Default`:

    1