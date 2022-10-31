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

    <!-- Table -->
    <table class="table">
        <thead>
            <tr>
                <th>#</th>
                <th>Id</th>
                <th>User</th>
                <th>Activity</th>
                <th>Object</th>
                <th>RowId</th>
                <th>Row</th>
                <!-- Put column header here, Note: 💀 Don't delete this line, Zaruba use it for pattern matching --> 
            </tr>
        </thead>
        <tbody>
            <tr v-for="(row, index) in result.rows">
                <td>{{ index+1 }}</td>
                <td>{{ row.id }}</td>
                <td>{{ row.user_id }}</td>
                <td>{{ row.activity }}</td>
                <td>{{ row.object }}</td>
                <td>{{ row.row_id }}</td>
                <td>{{ JSON.stringify(row.row) }}</td>
                <!-- Put column value here, Note: 💀 Don't delete this line, Zaruba use it for pattern matching -->
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
        showInsertForm,
    } = useCrud({
        formDataToRow,
        rowToFormData,
        entityName: 'activity',
        apiUrl: props.apiUrl,
        formComponentId: 'form-crud',
    });

    applyFilter().then(() => {
        setInterval(() => applyFilter(), 3 * 60 * 1000);
    });
</script>