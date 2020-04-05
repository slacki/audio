<template>
    <div class="wrapper" id="app">
        <router-view />
        <footer>
            <div class="container">
                <router-link to="/privacy">Privacy</router-link>
                <router-link to="/terms">Terms</router-link>
                <router-link to="/about">About</router-link>
                <router-link to="/contact">Contact</router-link>
            </div>
        </footer>
    </div>
</template>
