<template>
    <div v-if="loaded" class="uploaded-file">
        <h1 class="title">{{ file.original_name }}</h1>
        <player :url="url"></player>
    </div>
</template>

<script>
import GreenAudioPlayer from "green-audio-player";
import Player from "@/components/Player.vue";

export default {
    name: "UploadedFile",
    components: {
        Player
    },
    props: {
        id: {
            type: String,
            required: true
        }
    },
    computed: {
        url() {
            return process.env.VUE_APP_STATIC_ENDPOINT + "/" + this.file.file;
        }
    },
    data() {
        return {
            file: {},
            loaded: false
        };
    },
    created() {
        this.axios.get("/info/" + this.id).then(r => {
            this.loaded = true;
            this.file = r.data;
        });
    }
};
</script>

<style scoped lang="scss">
@import "@/scss/variables";

.title {
    padding-bottom: 60px;
}
</style>
