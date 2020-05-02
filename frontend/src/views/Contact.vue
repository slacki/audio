<template>
    <div class="container contact">
        <h1>Contact</h1>
        <p>Please fill out the form if you'd like to get in touch.</p>
        <form @submit.prevent="send">
            <label for="name">Your name</label>
            <input
                v-model="fd.name"
                type="text"
                id="name"
                name="name"
                placeholder="e.g. Nancy L. Jones"
            />
            <label for="from">How do I reach back to you?</label>
            <input
                v-model="fd.from"
                type="text"
                id="from"
                name="from"
                placeholder="e.g. you@domain.com"
            />
            <div>
                <label for="text">Your message</label>
                <textarea v-model="fd.text" name="text" id="text" cols="30" rows="20"></textarea>
            </div>
            <button>Send</button>
        </form>
    </div>
</template>

<script>
export default {
    name: "Contact",
    data() {
        return {
            fd: {
                name: "",
                from: "",
                text: ""
            }
        };
    },
    methods: {
        send() {
            this.axios
                .post("/contact", this.fd)
                .then(r => {
                    this.$notify({
                        title: "Message sent!",
                        text:
                            "I'll try to get back to you as soon as possible.",
                        type: "Success",
                        duration: 3000
                    });
                })
                .catch(e => {
                    this.$notify({
                        type: "error",
                        title: "Message couldn't be sent.",
                        duration: 3000
                    });
                })
                .finally(() => {
                    this.progress = 0;
                    this.uploading = false;
                });
        }
    }
};
</script>

<style lang="scss" scoped>
textarea {
    max-width: 100%;
    height: 250px;
}
</style>