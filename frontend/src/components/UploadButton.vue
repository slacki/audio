<template>
    <div class="upload">
        <form id="upload-form">
            <input
                ref="fileUploadField"
                name="file-upload-field"
                class="file-upload-field"
                type="file"
                value
                multiple
            />
        </form>

        <button @click="uploadButtonClick" class="button-large upload-button">
            <span>Choose files</span>
        </button>
    </div>
</template>

<script>
export default {
    name: "UploadButton",
    methods: {
        uploadButtonClick() {
            this.$refs.fileUploadField.click();
        }
    }
};
</script>

<style scoped lang="scss">
@import "@/scss/variables";

.upload {
    margin-top: 30px;

    .file-upload-field {
        display: none;
    }

    /* Custom color */
    .upload-button {
        background-color: $color-buttons;
        border-color: $color-buttons;
    }
}
</style>
