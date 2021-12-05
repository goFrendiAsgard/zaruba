if [ "$("${ZARUBA_BIN}" list length "${_ZRB_APP_CRUD_FIELDS}")" = 0 ]
then
    _ZRB_APP_CRUD_FIRST_FIELD="id"
else
    _ZRB_APP_CRUD_FIRST_FIELD="$("${ZARUBA_BIN}" list get "${_ZRB_APP_CRUD_FIELDS}" 0)"
fi

_setReplacementMap "ztplAppCrudFirstField" "${_ZRB_APP_CRUD_FIRST_FIELD}"
