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
                <th>Username</th>
                <th>Email</th>
                <th>Phone Number</th>
                <th>Full Name</th>
                <th>Roles</th>
                <th>Permissions</th>
                <th>Active</th>
                <!-- CRUD column headers -->
                <!-- Put column header here -->
                <th id="th-action">Actions</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(row, index) in result.rows">
                <td>{{ index+1 }}</td>
                <td>{{ row.id }}</td>
                <td>{{ row.username }}</td>
                <td>{{ row.email }}</td>
                <td>{{ row.phone_number }}</td>
                <td>{{ row.full_name }}</td>
                <td>{{ row.role_ids.join(', ') }}</td>
                <td>{{ row.permissions.join(', ') }}</td>
                <td>{{ row.active }}</td>
                <!-- CRUD column values -->
                <!-- Put column value here -->
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
                        <label for="form-input-username" class="col-form-label">Username:</label>
                        <input type="text" class="form-control" id="form-input-username" v-model="formData.username" />
                    </div>
                    <div class="mb-3">
                        <label for="form-input-email" class="col-form-label">Email:</label>
                        <input type="text" class="form-control" id="form-input-email" v-model="formData.email" />
                    </div>
                    <div class="mb-3">
                        <label for="form-input-full-name" class="col-form-label">Full Name:</label>
                        <input type="text" class="form-control" id="form-input-full-name" v-model="formData.full_name" />
                    </div>
                    <div class="mb-3">
                        <label for="form-input-phone-number" class="col-form-label">Phone Number:</label>
                        <input type="text" class="form-control" id="form-input-phone-number" v-model="formData.phone_number" />
                    </div>
                    <div class="mb-3">
                        <label for="form-input-password" class="col-form-label">Password:</label>
                        <input type="password" class="form-control" id="form-input-password" v-model="formData.password" placeholder="Leave blank to keep old password" />
                    </div>
                    <div class="mb-3">
                        <label for="form-input-roles" class="col-form-label">Roles:</label>
                        <JsonInput class="form-control" id="form-input-roles" v-model="formData.role_ids" />
                    </div>
                    <div class="mb-3">
                        <label for="form-input-permissions" class="col-form-label">Permission:</label>
                        <JsonInput class="form-control" id="form-input-permissions" v-model="formData.permissions" />
                    </div>
                    <div class="mb-3">
                        <label for="form-input-active" class="col-form-label">Active:</label>
                        <BooleanInput class="form-select" id="form-input-active" v-model="formData.active" />
                    </div>
                    <!-- CRUD form inputs -->
                    <!-- Put form input here -->
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
    import JsonInput from '../../../components/jsonInput.vue';
    import BooleanInput from '../../../components/booleanInput.vue';
    import {useCrud} from '../../../components/useCrud.vue';
    import {defineProps} from 'vue';

    const props = defineProps({
        apiUrl: String,
    });

    function formDataToRow(formData) {
        const row = formData;
        if (!formData.password) {
            delete row.password;
        }
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
        entityName: 'user',
        apiUrl: props.apiUrl,
        formComponentId: 'form-crud',
    });

    applyFilter().then(() => {
        setInterval(() => applyFilter(), 3 * 60 * 1000);
    });
</script>