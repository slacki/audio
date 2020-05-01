<template>
    <div v-if="loaded" class="container uploaded">
        <h1 class="title">{{ file.original_name }}</h1>
        <p class="hash">{{ file.views }} {{ viewOrViews }}</p>

        <div class="player-wrapper">
            <player :url="fileUrl"></player>
        </div>

        <label for="uploadUrl">Upload URL</label>
        <input id="uploadUrl" type="text" :value="url" />
        <label for="directUrl">Direct audio file URL</label>
        <input id="directUrl" type="text" :value="fileUrl" />
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
        fileUrl() {
            return process.env.VUE_APP_STATIC_ENDPOINT + "/" + this.file.file;
        },
        url() {
            return window.location.href;
        },
        viewOrViews() {
            return this.file.views > 1 ? "views" : "view";
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

.hash {
    font-size: 80%;
}

.player-wrapper {
    padding: 65px 0;
}
</style>