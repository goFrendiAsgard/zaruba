ENV_MAP="$({{ .ZarubaBin }} env get "{{ .GetConfig "envPrefix" }}")"

REPLACEMENT_MAP="$({{ .ZarubaBin }} map transformKey "${ENV_MAP}" "_RUNENV_")"
{{ .ZarubaBin }} util generate "initTemplates" "init" "${REPLACEMENT_MAP}"