<template>
    <select :value="modelValue" @input="updateValue">
        <option v-for="(value, key) in options" :value="key" :key="key" :selected="key == modelValue">
            {{ value }}
        </option>
    </select>
</template>
<script setup>
    import {defineProps, defineEmits, onMounted, ref} from 'vue';

    const props = defineProps({
        modelValue: {
            type: String
        },
        appApiUrl: {
            type: String,
            default: '',
        },
        optionValueKey: {
            type: String,
            default: 'id',
        },
        optionCaptionKey: {
            type: String,
            default: 'id',
        },
        optionList: {
            type: Array,
            default: () => {
                return [];
            },
        },
        optionMap: {
            type: Object,
            default: () => {
                return {};
            }
        },
        getOptions: {
            type: Function,
            default: async () => {
                return {}
            },
        }
    });

    const emit = defineEmits(['update:modelValue'])
    function updateValue(event) {
        console.log(event.target.value);
        const newModelValue = event.target.value;
        emit('update:modelValue', newModelValue);
    }

    const options = ref({});
    onMounted(async () => {

        // using list as options
        if (props.optionList.length > 0) {
            const optionValue = {};
            props.optionList.forEach((option) => {
                optionValue[option] = option;
            });
            options.value = optionValue;
            return
        }

        // using object as options
        if (Object.keys(props.optionMap).length > 0) {
            options.value = props.optionMap;
            return
        }

        // using appApiUrl
        if (props.appApiUrl != '') {
            const response = await appHelper.axios.get(props.appApiUrl, appHelper.getConfigAuthHeader());
            if (response && response.status == 200 && response.data && typeof(response.data.count) == 'number' && response.data.rows) {
                const rows = response.data.rows;
                const optionValue = {};
                rows.forEach((row) => {
                    const value = row[props.optionValueKey];
                    const caption = row[props.optionValueKey];
                    optionValue[value] = caption;
                });
                options.value = optionValue;
            }
            return
        }

        // using getOptions
        if (typeof(props.getOptions) == 'function') {
            options.value = await props.getOptions();
            return
        }

    });
</script>