<template>
    <div class="upload">
        <div v-if="!uploading" class="upload-form">
            <h1>Upload your audio</h1>
            <p class="upload-info">
                mp3, wav, ogg
                <br />max 15 MiB
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

            <button @click="uploadButtonClick" class="button-large">
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

            if (fileList.length != 1) {
                // only single file uploads for now
                return;
            }

            if (fileList[0].size > 15 * 1024 * 1024) {
                this.$notify({
                    type: "error",
                    title: "File too big!",
                    text: "Maximum file size is 15 MiB.",
                    duration: 3000
                });
                return;
            }

            formData.append(fieldName, fileList[0], fileList[0].name);

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
                    let notify = {
                        type: "error",
                        title: "Upload failed.",
                        duration: 3000
                    };
                    if (e.message === "Network Error") {
                        notify.text =
                            "This file is too big, max size is 15 MiB.";
                    } else {
                        notify.text = "Invalid file.";
                    }

                    this.$notify(notify);
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
    .file-upload-field {
        display: none;
    }

    .upload-info {
        padding: 15px;
    }
}
</style>
