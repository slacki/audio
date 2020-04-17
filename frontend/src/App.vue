<template>
    <div class="wrapper" id="app">
        <router-view />
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
