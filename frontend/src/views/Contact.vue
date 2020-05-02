<template>
    <div class="container contact">
        <h1>Contact</h1>
        <p>Please fill out the form if you'd like to get in touch.</p>
        <form @submit.prevent="sendContact">
            <label for="name">Your name</label>
            <input type="text" id="name" name="name" placeholder="e.g. Nancy L. Jones" />
            <label for="email">How do I reach back to you?</label>
            <input type="text" id="email" name="email" placeholder="e.g. you@domain.com" />
            <div>
                <label for="text">Your message</label>
                <textarea name="text" id="text" cols="30" rows="20"></textarea>
            </div>
            <button>Send</button>
        </form>
    </div>
</template>

<style lang="scss" scoped>
textarea {
    max-width: 100%;
    height: 250px;
}
</style>