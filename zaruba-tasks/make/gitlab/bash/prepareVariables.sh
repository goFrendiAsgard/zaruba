# gitlabWeb envs
_ZRB_APP_GITLAB_WEB_ENVS="$("${ZARUBA_BIN}" path getEnv "${_ZRB_APP_DIRECTORY}/gitlabWeb")"
# yaml
_ZRB_APP_YAML_GITLAB_WEB_ENVS="$(_getYamlEnvs "${_ZRB_APP_GITLAB_WEB_ENVS}" "${_ZRB_APP_ENV_PREFIX}")"

# gitlabRunner envs
_ZRB_APP_GITLAB_RUNNER_ENVS="$("${ZARUBA_BIN}" path getEnv "${_ZRB_APP_DIRECTORY}/gitlabRunner")"
# yaml
_ZRB_APP_YAML_GITLAB_RUNNER_ENVS="$(_getYamlEnvs "${_ZRB_APP_GITLAB_RUNNER_ENVS}" "${_ZRB_APP_ENV_PREFIX}_GITLAB_RUNNER")"

