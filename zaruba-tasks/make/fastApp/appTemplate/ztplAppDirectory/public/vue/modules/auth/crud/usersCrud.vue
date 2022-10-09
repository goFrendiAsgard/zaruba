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
                <th>Username</th>
                <th>Email</th>
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
                <td>{{ row.full_name }}</td>
                <td>{{ row.role_ids.join(', ') }}</td>
                <td>{{ row.permissions.join(', ') }}</td>
                <td>{{ row.active }}</td>
                <!-- CRUD column values -->
                <!-- Put column value here -->
                <td id="td-action">
                    <div class="d-grid gap-2 d-md-flex justify-content-md-end">
                        <button class="btn btn-warning" type="button" data-bs-toggle="modal" data-bs-target="#form-crud" @click="showUpdateForm(row.id)"><i class="bi bi-pencil-square"></i></button>
                        <button class="btn btn-danger" type="button" @click="confirmDelete(row.id)"><i class="bi bi-file-minus"></i></button>
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
                        <label for="form-input-username" class="col-form-label">Username:</label>
                        <input type="text" class="form-control" id="form-input-username" v-model="formData.username">
                    </div>
                    <div class="mb-3">
                        <label for="form-input-email" class="col-form-label">Email:</label>
                        <input type="text" class="form-control" id="form-input-email" v-model="formData.email">
                    </div>
                    <div class="mb-3">
                        <label for="form-input-full-name" class="col-form-label">Full Name:</label>
                        <input type="text" class="form-control" id="form-input-full-name" v-model="formData.full_name">
                    </div>
                    <div class="mb-3">
                        <label for="form-input-password" class="col-form-label">Password:</label>
                        <input type="password" class="form-control" id="form-input-password" v-model="formData.password" placeholder="Leave blank to keep old password">
                    </div>
                    <div class="mb-3">
                        <label for="form-input-roles" class="col-form-label">Roles:</label>
                        <input type="text" class="form-control" id="form-input-roles" v-model="formData.role_ids">
                    </div>
                    <div class="mb-3">
                        <label for="form-input-permissions" class="col-form-label">Permission:</label>
                        <input type="text" class="form-control" id="form-input-permissions" v-model="formData.permissions">
                    </div>
                    <div class="mb-3">
                        <label for="form-input-active" class="col-form-label">Active:</label>
                        <select class="form-select" id="form-input-active" v-model="formData.active">
                            <option value="true">true</option>
                            <option value="false">false</option>
                        </select>
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

<script>
    export default {

        props : {
            apiUrl: String
        },

        data() {
            return {
                keyword: '',
                limit: 100,
                page: 1,
                result: {
                    count: 0,
                    rows: [],
                },
                isFilterApplied: false,
                formTitle: '',
                formData: {},
                rowId: '',
                formMode: '',
            }
        },

        methods: {

            async _fetchResult() {
                const response = await axios.get(this.apiUrl, {
                    params: {
                        keyword: this.keyword,
                        limit: this.limit,
                        offset: this.limit * (this.page-1)
                    },
                    ...appHelper.getConfigAuthHeader(),
                });
                if (response && response.status == 200 && response.data && typeof response.data.count == 'number' && response.data.rows) {
                    this.result = response.data;
                    return;
                }
                throw new Error(appHelper.getResponseErrorMessage(response, 'Cannot fetch result'));
            },

            async _getRow(id) {
                const response =  await axios.get(`${this.apiUrl}/${id}`, appHelper.getConfigAuthHeader());
                if (response && response.status == 200 && response.data) {
                    return response.data
                }
                throw new Error(appHelper.getResponseErrorMessage(response, `Cannot get user ${id}`));
            },

            async _insertRow(row) {
                const response = await axios.post(this.apiUrl, row, appHelper.getConfigAuthHeader());
                if (response && response.status == 200 && response.data) {
                    await this._fetchResult();
                    return;
                }
                throw new Error(appHelper.getResponseErrorMessage(response, `Cannot create user ${id}`));
            },

            async _updateRow(id, row) {
                const response = await axios.put(`${this.apiUrl}/${id}`, row, appHelper.getConfigAuthHeader());
                if (response && response.status == 200 && response.data) {
                    await this._fetchResult();
                    return;
                }
                throw new Error(appHelper.getResponseErrorMessage(response, `Cannot update user ${id}`));
            },

            async _deleteRow(id) {
                const response = await axios.delete(`${this.apiUrl}/${id}`, appHelper.getConfigAuthHeader());
                if (response && response.status == 200 && response.data) {
                    await this._fetchResult();
                    return;
                }
                throw new Error(appHelper.getResponseErrorMessage(response, `Cannot delete user ${id}`));
            },

            _formDataToRow(formData) {
                let row = {};
                Object.assign(row, formData);
                if (formData.permissions) {
                    row.permissions = JSON.parse(formData.permissions);
                }
                if (formData.role_ids) {
                    row.role_ids = JSON.parse(formData.role_ids);
                }
                row.active = formData.active == 'true';
                if (!formData.password) {
                    delete row.password;
                }
                if (formData.permissions) {
                    row.permissions = JSON.parse(formData.permissions);
                }
                return row;
            },

            _rowToFormData(row) {
                let formData = {};
                Object.assign(formData, row);
                if (row.permissions) {
                    formData.permissions = JSON.stringify(row.permissions);
                }
                if (row.role_ids) {
                    formData.role_ids = JSON.stringify(row.role_ids);
                }
                formData.active = row.active? 'true': 'false';
                return formData;
            },

            setFilterApplied(state) {
                this.isFilterApplied = state;
            },

            resetFilterApplied() {
                this.setFilterApplied(false);
            },

            async applyFilter() {
                try {
                    await this._fetchResult();
                    this.setFilterApplied(true);
                } catch (error) {
                    appHelper.alertError(error);
                }
            },

            async showInsertForm() {
                this.formTitle = 'New User';
                this.formMode = 'insert';
                this.formData = {};
            },

            async showUpdateForm(id) {
                try {
                    const row = await this._getRow(id);
                    this.formTitle = `Edit User ${id}`;
                    this.formMode = 'update';
                    this.rowId = id;
                    this.formData = this._rowToFormData(row);
                    return;
                } catch (error) {
                    appHelper.alertError(error);
                }
            },

            async save() {
                try {
                    const form = bootstrap.Modal.getInstance(document.getElementById('form-crud'));
                    const row = this._formDataToRow(this.formData);
                    if (this.formMode == 'insert') {
                        await this._insertRow(row);
                        form.hide();
                        return this._fetchResult();
                    }
                    if (this.formMode == 'update') {
                        await this._updateRow(this.rowId, row);
                        form.hide();
                        return this._fetchResult();
                    }
                    throw new Error(`Invalid formMode: ${this.formMode}`);
                } catch (error) {
                    appHelper.alertError(error);
                }
            },

            async confirmDelete(id) {
                try {
                    const confirmation = await appHelper.confirm(`Are you sure to delete user ${id}?`);
                    if (!confirmation) {
                        return;
                    }
                    await this._deleteRow(id);
                } catch (error) {
                    appHelper.alertError(error);
                }
            },

        },

        async beforeMount() {
            await this.applyFilter();
            setInterval(() => this.applyFilter(), 3 * 60 * 1000);
        }
    }
</script>