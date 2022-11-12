<template>
    <textarea :value="modelValue" @input="updateValue" rows="5" />
</template>
<script setup>
    import {defineProps, defineEmits} from 'vue';

    defineProps({
        modelValue: {
            type: [Object, Array, Boolean, String, Number],
        },
    });

    const emit = defineEmits(['update:modelValue'])
    function updateValue(event) {
        try {
            const newModelValue = event.target.value;
            emit('update:modelValue', newModelValue);
        } catch(error) {
            // do nothing
        }
    }
</script>