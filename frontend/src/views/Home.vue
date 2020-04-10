<template>
    <div class="container home">
        <h1>Upload your audio</h1>
        <p class="upload-info">
            mp3, wav and ogg
            <br />maximum of 15 megabytes
        </p>

        <upload-button></upload-button>
    </div>
</template>

<script>
// @ is an alias to /src
import UploadButton from "@/components/UploadButton.vue";

export default {
    name: "Home",
    components: {
        UploadButton
    }
};
</script>

<style lang="scss" scoped>
</style>