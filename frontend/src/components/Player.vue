<template>
    <div class="player">
        <audio>
            <source :src="url" />
        </audio>
    </div>
</template>

<script>
import GreenAudioPlayer from "green-audio-player";

export default {
    name: "Player",
    props: {
        url: {
            type: String,
            required: true
        }
    },
    mounted() {
        GreenAudioPlayer.init({
            selector: ".player"
        });
    }
};
</script>

<style lang="scss">
@import "@/scss/player/main";

.player {
    margin: 0 auto;
}
</style>