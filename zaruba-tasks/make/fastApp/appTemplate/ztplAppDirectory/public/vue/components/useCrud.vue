<script>
    import { ref } from 'vue'

    function useCrud(config) {
        //////////////////////////////////////////////////////////////////
        // Configurations
        //////////////////////////////////////////////////////////////////

        const configEntityName = config.entityName || 'entity'
        const configApiUrl = config.apiUrl || '';
        const configInsertFormTitle = config.insertFormTitle || 'New ${Entity}';
        const configUpdateFormTitle = config.updateFormTitle || 'Edit ${Entity} ${id}';
        const configConfirmDeleteMessage = config.updateFormTitle || 'Are you sure to delete ${entity} ${id}';
        const configFormComponentId = config.formId || 'form-crud';
        const configGetRowErrorMessage = config.getRowErrorMessage || 'Cannot get ${entity} ${id}';
        const configInsertRowErrorMessage = config.getRowErrorMessage || 'Cannot create ${entity}';
        const configUpdateRowErrorMessage = config.updateRowErrorMessage || 'Cannot update ${entity} ${id}';
        const configDeleteRowErrorMessage = config.deleteRowErrorMessage || 'Cannot delete ${entity} ${id}';
        const configFormDataToRow = config.formDataToRow || ((obj) => obj);
        const configRowToFormData = config.rowToFormData || ((obj) => obj);

        //////////////////////////////////////////////////////////////////
        // Private functions
        //////////////////////////////////////////////////////////////////

        async function _formDataToRow(formData) {
            let copyFormData = {};
            Object.assign(copyFormData, formData);
            const row = await configFormDataToRow(copyFormData)
            return row;
        }

        async function _rowToFormData(row) {
            let copyRow = {};
            Object.assign(copyRow, row);
            const formData = await configRowToFormData(copyRow)
            return formData;
        }

        function _parse(template, slot) {
            let result = template;
            for (let key in slot) {
                result = result.replaceAll('${' + key + '}', slot[key]);
            }
            result = result.replaceAll(/\${.*}/g, '');
            return result;
        }

        function _capitalize(sentence) {
            const words = sentence.split(" ");
            for (let i = 0; i < words.length; i++) {
                words[i] = words[i][0].toUpperCase() + words[i].substr(1);
            }
            return words.join(" ");
        }

        //////////////////////////////////////////////////////////////////
        // Row manipulation
        //////////////////////////////////////////////////////////////////

        async function getRow(id) {
            const response =  await axios.get(`${configApiUrl}/${id}`, appHelper.getConfigAuthHeader());
            if (response && response.status == 200 && response.data) {
                return response.data;
            }
            const errorMessage = _parse(configGetRowErrorMessage, {
                id,
                entity: configEntityName,
            });
            throw new Error(appHelper.getResponseErrorMessage(response, errorMessage));
        }

        async function insertRow(row) {
            const response = await axios.post(configApiUrl, row, appHelper.getConfigAuthHeader());
            if (response && response.status == 200 && response.data) {
                await fetchResult();
                return response.data;
            }
            const errorMessage = _parse(configInsertRowErrorMessage, {
                entity: configEntityName,
            });
            throw new Error(appHelper.getResponseErrorMessage(response, errorMessage));
        }

        async function updateRow(id, row) {
            const response = await axios.put(`${configApiUrl}/${id}`, row, appHelper.getConfigAuthHeader());
            if (response && response.status == 200 && response.data) {
                await fetchResult();
                return response.data;
            }
            const errorMessage = _parse(configUpdateRowErrorMessage, {
                id,
                entity: configEntityName,
            });
            throw new Error(appHelper.getResponseErrorMessage(response, errorMessage));
        }

        async function deleteRow(id) {
            const response = await axios.delete(`${configApiUrl}/${id}`, appHelper.getConfigAuthHeader());
            if (response && response.status == 200 && response.data) {
                await fetchResult();
                return response.data;
            }
            const errorMessage = _parse(configDeleteRowErrorMessage, {
                id,
                entity: configEntityName,
            });
            throw new Error(appHelper.getResponseErrorMessage(response, errorMessage));
        }

        //////////////////////////////////////////////////////////////////
        // Table
        //////////////////////////////////////////////////////////////////

        const keyword = ref('');
        const limit = ref(100)
        const page = ref(1);
        const result = ref({count: 0, rows: []});

        async function fetchResult() {
            const response = await axios.get(configApiUrl, {
                params: {
                    keyword: keyword.value,
                    limit: limit.value,
                    offset: limit.value * (page.value-1)
                },
                ...appHelper.getConfigAuthHeader(),
            });
            if (response && response.status == 200 && response.data && typeof response.data.count == 'number' && response.data.rows) {
                result.value = response.data;
                return;
            }
            throw new Error(appHelper.getResponseErrorMessage(response, 'Cannot fetch result'));
        }

        //////////////////////////////////////////////////////////////////
        // Filter status
        //////////////////////////////////////////////////////////////////

        const isFilterApplied = ref(false);

        function setFilterApplied(state) {
            isFilterApplied.value = state;
        }

        function resetFilterApplied() {
            setFilterApplied(false);
        }

        async function applyFilter() {
            try {
                await fetchResult();
                setFilterApplied(true);
            } catch (error) {
                appHelper.alertError(error);
            }
        }

        const formTitle = ref('');
        const formData = ref({});
        const rowId = ref('');
        const formMode = ref('');

        //////////////////////////////////////////////////////////////////
        // Form
        //////////////////////////////////////////////////////////////////

        async function showInsertForm() {
            formTitle.value = _parse(configInsertFormTitle, {
                entity: configEntityName,
                Entity: _capitalize(configEntityName),
            });
            formMode.value = 'insert';
            formData.value = {};
        }

        async function showUpdateForm(id) {
            try {
                const row = await getRow(id);
                formTitle.value = _parse(configUpdateFormTitle, {
                    entity: configEntityName,
                    Entity: _capitalize(configEntityName),
                    id,
                });
                formMode.value = 'update';
                rowId.value = id;
                formData.value = {};
                formData.value = await _rowToFormData(row);
            } catch (error) {
                appHelper.alertError(error);
            }
        }

        async function confirmDelete(id) {
            try {
                const confirmationMessage = _parse(configConfirmDeleteMessage, {
                    entity: configEntityName,
                    Entity: _capitalize(configEntityName),
                    id,
                });
                const confirmation = await appHelper.confirm(confirmationMessage);
                if (!confirmation) {
                    return;
                }
                await deleteRow(id);
            } catch (error) {
                appHelper.alertError(error);
            }
        }

        async function save() {
            try {
                const form = bootstrap.Modal.getInstance(document.getElementById(configFormComponentId));
                const row = await _formDataToRow(formData.value);
                if (formMode.value == 'insert') {
                    await insertRow(row);
                    form.hide();
                    return fetchResult();
                }
                if (formMode.value == 'update') {
                    await updateRow(rowId.value, row);
                    form.hide();
                    return fetchResult();
                }
                throw new Error(`Invalid formMode: ${formMode.value}`);
            } catch (error) {
                appHelper.alertError(error);
            }
        }

        //////////////////////////////////////////////////////////////////
        // Exposed
        //////////////////////////////////////////////////////////////////

        return {
            keyword,
            limit,
            page,
            result,
            getRow,
            insertRow,
            updateRow,
            deleteRow,
            applyFilter,
            fetchResult,
            isFilterApplied,
            setFilterApplied,
            resetFilterApplied,
            formTitle,
            formData,
            rowId,
            formMode,
            showInsertForm,
            showUpdateForm,
            confirmDelete,
            save,
        } 
    }

    export default {
        useCrud
    };
</script>