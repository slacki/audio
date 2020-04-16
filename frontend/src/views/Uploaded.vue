<template>
    <div v-if="loaded" class="container uploaded">
        <h1 class="title">{{ file.original_name }}</h1>
        <player :url="url"></player>
    </div>
</template>

<script>
import GreenAudioPlayer from "green-audio-player";
import Player from "@/components/Player.vue";

export default {
    name: "Home",
    components: {
        Player,
        GreenAudioPlayer
    },
    data() {
        return {
            id: null,
            file: {},
            loaded: false
        };
    },
    computed: {
        url() {
            return process.env.VUE_APP_STATIC_ENDPOINT + "/" + this.file.file;
        }
    },
    mounted() {
        this.id = this.$route.params.id;

        this.axios.get("/info/" + this.id).then(r => {
            this.loaded = true;
            this.file = r.data;
        });
    }
};
</script>

<style lang="scss" scoped>
@import "@/scss/variables";

.title {
    padding-bottom: 60px;
}
</style>