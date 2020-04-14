<template>
    <div class="container uploaded">
        <uploaded-file v-if="id !== ''" :id="id"></uploaded-file>
    </div>
</template>

<script>
// @ is an alias to /src
import UploadedFile from "@/components/UploadedFile.vue";

export default {
    name: "Home",
    components: {
        UploadedFile
    },
    data() {
        return {
            id: null
        };
    },
    created() {
        this.id = this.$route.params.id;
    }
};
</script>

<style lang="scss" scoped>
</style>