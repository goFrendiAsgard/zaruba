if [ "$("${ZARUBA_HOME}/zaruba" list length "${_ZRB_APP_CRUD_FIELDS}")" = 0 ]
then
    _ZRB_APP_CRUD_FIRST_FIELD="id"
else
    _ZRB_APP_CRUD_FIRST_FIELD="$("${ZARUBA_HOME}/zaruba" list get "${_ZRB_APP_CRUD_FIELDS}" 0)"
fi

_setReplacementMap "ztplAppCrudFirstField" "${_ZRB_APP_CRUD_FIRST_FIELD}"
