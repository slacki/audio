<template>
    <div id="app">
        <div class="wrapper">
            <router-view />
        </div>
        <footer>
            <div class="container">
                <span class="beta">beta</span>
                <router-link :to="{ name: 'home' }">Home</router-link>
                <router-link :to="{ name: 'privacyPolicy' }">Privacy</router-link>
                <router-link :to="{ name: 'terms' }">Terms</router-link>
                <router-link :to="{ name: 'about' }">About</router-link>
                <router-link :to="{ name: 'contact' }">Contact</router-link>
            </div>
        </footer>
    </div>
</template>

<style lang="scss" scoped>
@import "@/scss/variables.scss";

.beta {
    background-color: $color-tertiary;
    color: white;
    display: inline-block;
    padding: 0.25em 0.4em;
    font-size: 75%;
    font-weight: 700;
    line-height: 1;
    text-align: center;
    white-space: nowrap;
    vertical-align: baseline;
    border-radius: 0.25rem;
}
</style>