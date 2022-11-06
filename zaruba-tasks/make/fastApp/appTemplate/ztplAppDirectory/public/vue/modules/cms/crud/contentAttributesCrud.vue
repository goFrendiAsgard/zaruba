<template>

    <!-- Filter -->
    <div class="row g-3 align-items-center">
        <div class="col-auto">
            <label for="input-keyword" class="col-form-label">Keyword</label>
        </div>
        <div class="col-auto">
            <input type="text" id="input-keyword" class="form-control" v-model="keyword" @keyup.enter="applyFilter" @change="resetFilterApplied" @keyup="resetFilterApplied">
        </div>
        <div class="col-auto">
            <button id="btn-filter" class="btn btn-primary" @click="applyFilter" :disabled="isFilterApplied"><i class="bi bi-funnel"></i> Filter</button>
        </div>
    </div>

    <!-- Add button -->
    <div class="d-grid gap-2 d-md-flex justify-content-md-end">
        <button class="btn btn-success" type="button" data-bs-toggle="modal" data-bs-target="#form-crud" @click="showInsertForm"><i class="bi bi-file-plus"></i> Add</button>
    </div>

    <!-- Table -->
    <table class="table">
        <thead>
            <tr>
                <th>#</th>
                <th>Id</th>
                <th>ContentId</th>
                <th>Key</th>
                <th>Value</th>
                <!-- Put column header here, Note: 💀 Don't delete this line, Zaruba use it for pattern matching --> 
                <th id="th-action">Actions</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(row, index) in result.rows">
                <td>{{ index+1 }}</td>
                <td>{{ row.id }}</td>
                <td>{{ row.content_id }}</td>
                <td>{{ row.key }}</td>
                <td>{{ row.value }}</td>
                <!-- Put column value here, Note: 💀 Don't delete this line, Zaruba use it for pattern matching -->
                <td id="td-action">
                    <div class="d-grid gap-2 d-md-flex justify-content-md-end">
                        <button class="btn btn-warning" type="button" data-bs-toggle="modal" data-bs-target="#form-crud" @click="showUpdateForm(row.id)"><i class="bi bi-pencil-square"></i> Edit</button>
                        <button class="btn btn-danger" type="button" @click="confirmDelete(row.id)"><i class="bi bi-file-minus"></i> Delete</button>
                    </div>
                </td>
            </tr>
        </tbody>
    </table>

    <!-- Page selector -->
    <div class="row g-3 align-items-center">
        <div class="col-auto">
            <label for="input-page" class="col-form-label">Page</label>
        </div>
        <div class="col-auto">
            <input type="number" id="input-page" class="form-control" min="1" :max="Math.ceil(result.count/limit)" v-model="page" @change="applyFilter">
        </div>
        <div class="col-auto">
            <label for="input-limit" class="col-form-label">Result/Page</label>
        </div>
        <div class="col-auto">
            <input type="number" id="input-limit" class="form-control" min="1" v-model="limit" @change="applyFilter">
        </div>
    </div>

    <!-- Form -->
    <div class="modal fade" id="form-crud" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="modal-title" aria-hidden="true">
        <div class="modal-dialog modal-dialog-scrollable">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="modal-title">{{ formTitle }}</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <div class="mb-3">
                        <label for="form-input-content-id" class="col-form-label">ContentId:</label>
                        <input type="text" class="form-control" id="form-input-content-id" v-model="formData.content_id">
                    </div>
                    <div class="mb-3">
                        <label for="form-input-key" class="col-form-label">Key:</label>
                        <input type="text" class="form-control" id="form-input-key" v-model="formData.key">
                    </div>
                    <div class="mb-3">
                        <label for="form-input-value" class="col-form-label">Value:</label>
                        <input type="text" class="form-control" id="form-input-value" v-model="formData.value">
                    </div>
                    <!-- Put form input here, Note: 💀 Don't delete this line, Zaruba use it for pattern matching -->
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
                    <button type="button" class="btn btn-primary" @click="save">Save</button>
                </div>
            </div>
        </div>
    </div>

</template>

<script setup>
    import {useCrud} from '../../../components/useCrud.vue';
    import {defineProps} from 'vue';

    const props = defineProps({
        apiUrl: String,
    });

    function formDataToRow(formData) {
        const row = formData;
        return row;
    }

    function rowToFormData(row) {
        const formData = row;
        return formData;
    }

    const {
        keyword,
        limit,
        page,
        result,
        applyFilter,
        isFilterApplied,
        resetFilterApplied,
        formTitle,
        formData,
        showInsertForm,
        showUpdateForm,
        confirmDelete,
        save,
    } = useCrud({
        formDataToRow,
        rowToFormData,
        entityName: 'contentAttribute',
        apiUrl: props.apiUrl,
        formComponentId: 'form-crud',
    });

    applyFilter().then(() => {
        setInterval(() => applyFilter(), 3 * 60 * 1000);
    });
</script>