<template>
    <div class="container home">
        <upload></upload>
    </div>
</template>

<script>
// @ is an alias to /src
import Upload from "@/components/Upload.vue";

export default {
    name: "Home",
    components: {
        Upload
    }
};
</script>

<style lang="scss" scoped>
</style>