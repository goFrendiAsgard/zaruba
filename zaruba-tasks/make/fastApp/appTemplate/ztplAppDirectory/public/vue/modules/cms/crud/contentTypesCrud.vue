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
                <th>Content Type</th>
                <th>Markdown Template (Jinja)</th>
                <!-- Put column header here, Note: 🤖 Don't delete this line; Zaruba uses it for pattern matching --> 
                <th id="th-action">Actions</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(row, index) in result.rows">
                <td>{{ index+1 }}</td>
                <td>
                    <b>{{ row.name }}</b><br />
                    ({{ row.id }})<br />
                    <p>
                        <b>Attributes:</b>
                        <ul>
                            <li v-for="attribute in row.attributes">
                                <span class="badge bg-secondary">{{ attribute.input_type }}</span>
                                {{ attribute.name }} ({{ attribute.caption }})
                            </li>
                        </ul>
                    </p>
                </td>
                <td>
                    <textarea class="form-control-plaintext" readonly rows="10" cols="80">{{ row.template }}</textarea>
                </td>
                <!-- Put column value here, Note: 🤖 Don't delete this line; Zaruba uses it for pattern matching -->
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
                        <label for="form-input-name" class="col-form-label">Name:</label>
                        <input type="text" class="form-control" id="form-input-name" v-model="formData.name" />
                    </div>
                    <div class="mb-3">
                        <label for="form-input-template" class="col-form-label">Template:</label>
                        <MarkdownInput class="form-control" id="form-input-template" v-model="formData.template" />
                    </div>
                    <div class="mb-3">
                        <label for="form-input-attributes" class="col-form-label">Attributes:</label>
                        <JsonInput class="form-control" id="form-input-attributes" v-model="formData.attributes" rows="10" />
                    </div>
                    <!-- Put form input here, Note: 🤖 Don't delete this line; Zaruba uses it for pattern matching -->
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
    import MarkdownInput from '../../../components/markdownInput.vue';
    import JsonInput from '../../../components/jsonInput.vue';
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
        entityName: 'content type',
        apiUrl: props.apiUrl,
        formComponentId: 'form-crud',
    });

    applyFilter().then(() => {
        setInterval(() => applyFilter(), 3 * 60 * 1000);
    });
</script>