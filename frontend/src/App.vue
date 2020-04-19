<template>
    <div id="app">
        <div class="wrapper">
            <router-view />
        </div>
        <footer>
            <div class="container">
                <router-link :to="{ name: 'home' }">Home</router-link>
                <router-link :to="{ name: 'privacyPolicy' }">Privacy</router-link>
                <router-link :to="{ name: 'terms' }">Terms</router-link>
                <router-link :to="{ name: 'about' }">About</router-link>
                <router-link :to="{ name: 'contact' }">Contact</router-link>
            </div>
        </footer>
    </div>
</template>
