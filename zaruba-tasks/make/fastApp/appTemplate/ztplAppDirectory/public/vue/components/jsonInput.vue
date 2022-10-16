<template>
    <textarea :value="JSON.stringify(modelValue, null, 2)" @input="updateValue" />
</template>
<script setup>
    import {defineProps, defineEmits} from 'vue';

    defineProps({
        modelValue: [],
    });

    const emit = defineEmits(['update:modelValue'])
    function updateValue(event) {
        try {
            const newModelValue = JSON.parse(event.target.value);
            emit('update:modelValue', newModelValue);
        } catch(error) {
            // do nothing
        }
    }
</script>