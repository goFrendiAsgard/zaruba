<template>

    <!-- Filter -->
    <div class="row g-3 align-items-center">
        <div class="col-auto">
            <label for="input-keyword" class="col-form-label">Keyword</label>
        </div>
        <div class="col-auto">
            <input type="text" id="input-keyword" class="form-control" v-model="keyword" @keyup.enter="applyFilter" @change="resetFilterApplied" @keyup="resetFilterApplied" />
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
                <th>Type</th>
                <th>Title</th>
                <th>Attributes</th>
                <th>Description</th>
                <!-- Put column header here, Note: 💀 Don't delete this line, Zaruba use it for pattern matching --> 
                <th id="th-action">Actions</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(row, index) in result.rows">
                <td>{{ index+1 }}</td>
                <td>{{ row.id }}</td>
                <td>{{ row.type_id }}</td>
                <td>{{ row.title }}</td>
                <td>{{ JSON.stringify(row.attributes) }}</td>
                <td>{{ row.description }}</td>
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
            <input type="number" id="input-page" class="form-control" min="1" :max="Math.ceil(result.count/limit)" v-model="page" @change="applyFilter" />
        </div>
        <div class="col-auto">
            <label for="input-limit" class="col-form-label">Result/Page</label>
        </div>
        <div class="col-auto">
            <input type="number" id="input-limit" class="form-control" min="1" v-model="limit" @change="applyFilter" />
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
                        <label for="form-input-type" class="col-form-label">Type:</label>
                        <input type="text" class="form-control" id="form-input-type" v-model="formData.type_id" />
                    </div>
                    <div class="mb-3">
                        <label for="form-input-title" class="col-form-label">Title:</label>
                        <input type="text" class="form-control" id="form-input-title" v-model="formData.title" />
                    </div>
                    <div class="mb-3">
                        <label for="form-input-attributes" class="col-form-label">Attributes:</label>
                        <Json class="form-control" id="form-input-attributes" v-model="formData.attributes" />
                    </div>
                    <div class="mb-3">
                        <label for="form-input-description" class="col-form-label">Description:</label>
                        <input type="text" class="form-control" id="form-input-description" v-model="formData.description" />
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
    import Json from '../../../components/jsonInput.vue';
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
        entityName: 'content',
        apiUrl: props.apiUrl,
        formComponentId: 'form-crud',
    });

    applyFilter().then(() => {
        setInterval(() => applyFilter(), 3 * 60 * 1000);
    });
</script>