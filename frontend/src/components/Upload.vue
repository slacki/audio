<template>
    <div class="upload">
        <div v-if="!uploading" class="upload-form">
            <h1>Upload your audio</h1>
            <p class="upload-info">
                mp3, wav and ogg
                <br />maximum of 15 megabytes
            </p>
            <form id="upload-form" enctype="multipart/form-data" novalidate>
                <input
                    ref="fileUploadField"
                    name="files"
                    class="file-upload-field"
                    type="file"
                    accept=".wav, .ogg, .mp3"
                    @change="filesChange($event.target.name, $event.target.files); fileCount = $event.target.files.length"
                />
            </form>

            <button @click="uploadButtonClick" class="button-large upload-button">
                <span>Choose files</span>
            </button>
        </div>
        <div v-if="uploading" class="upload-items">
            <h1>
                Uploading...
                <span class="progress">{{ uploadProgress }}%</span>
            </h1>
        </div>
    </div>
</template>

<script>
export default {
    name: "Upload",
    data() {
        return {
            uploading: false,
            uploadProgress: 0
        };
    },
    methods: {
        uploadButtonClick() {
            this.$refs.fileUploadField.click();
        },
        filesChange(fieldName, fileList) {
            const formData = new FormData();

            if (!fileList.length) return;

            // append the files to FormData
            Array.from(Array(fileList.length).keys()).map(x => {
                formData.append(fieldName, fileList[x], fileList[x].name);
            });

            // upload
            this.save(formData);
        },
        save(formData) {
            this.uploading = true;
            this.uploadProgress = 0;

            const config = {
                onUploadProgress: progressEvent => {
                    this.onUploadProgress(progressEvent);
                }
            };

            this.axios
                .post("/upload", formData, config)
                .then(r => {
                    this.$emit("uploaded", r.data);
                })
                .catch(e => {
                    console.log(e);
                })
                .finally(() => {
                    this.progress = 0;
                    this.uploading = false;
                });
        },
        onUploadProgress(progressEvent) {
            this.uploadProgress = Math.round(
                (progressEvent.loaded * 100) / progressEvent.total
            );
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
