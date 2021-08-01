. ${ZARUBA_HOME}/scripts/bash/util.sh

show_version() {
    if [ -z "$(get_latest_git_tag)" ]
    then
        echo "Dev - $(get_latest_git_commit)"
    elif [ "$(get_latest_git_tag_commit)" = "$(get_latest_git_commit)" ]
    then
        echo "$(get_latest_git_tag) - $(get_latest_git_commit)"
    else
        echo "Dev - $(get_latest_git_tag) - $(get_latest_git_commit)"
    fi
}