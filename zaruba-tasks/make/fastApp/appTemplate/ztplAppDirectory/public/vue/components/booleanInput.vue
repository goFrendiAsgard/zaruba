<template>
    <select :value="modelValue ? 'true' : 'false'" @input="updateValue">
        <option value="true">true</option>
        <option value="false">false</option>
    </select>
</template>
<script>
    export default {
        props: {
            modelValue: [],
        },

        methods: {
            updateValue(event) {
                const newModelValue = event.target.value == 'true';
                this.$emit('update:modelValue', newModelValue);
            }
        }

    }
</script>
