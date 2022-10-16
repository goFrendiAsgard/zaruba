<template>
    <select :value="modelValue ? 'true' : 'false'" @input="updateValue">
        <option value="true">true</option>
        <option value="false">false</option>
    </select>
</template>
<script setup>
    import {defineProps, defineEmits} from 'vue';

    defineProps({
        modelValue: Boolean,
    });

    const emit = defineEmits(['update:modelValue'])
    function updateValue(event) {
        console.log(event)
        const newModelValue = event.target.value == 'true';
        emit('update:modelValue', newModelValue);
    }
</script>