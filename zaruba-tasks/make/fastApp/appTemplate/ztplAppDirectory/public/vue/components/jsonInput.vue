<template>
    <textarea :value="JSON.stringify(modelValue)" @input="updateValue" />
</template>
<script>
    export default {
        props: {
            modelValue: [],
        },

        methods: {
            updateValue(event) {
                try{
                    const newModelValue = JSON.parse(event.target.value);
                    this.$emit('update:modelValue', newModelValue);
                } catch(error) {
                    // do nothing
                }
            }
        }

    }
</script>
