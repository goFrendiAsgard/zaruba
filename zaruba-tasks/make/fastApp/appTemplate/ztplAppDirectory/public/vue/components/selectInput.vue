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
        fetchOptions: {
            type: Function,
            default: async () => {return {}}
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
        options.value = await props.fetchOptions();
    });
</script>