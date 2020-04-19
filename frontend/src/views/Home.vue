<template>
    <div class="container home">
        <upload @uploaded="onUploaded"></upload>
    </div>
</template>

<script>
// @ is an alias to /src
import Upload from "@/components/Upload.vue";

export default {
    name: "Home",
    components: {
        Upload
    },
    methods: {
        onUploaded(e) {
            this.$router.push({
                name: "uploadedFile",
                params: { id: e.hash }
            });
        }
    }
};
</script>

<style lang="scss" scoped>
.home {
    text-align: center;
}
</style>