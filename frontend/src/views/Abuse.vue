<template>
    <div class="container contact">
        <h1>Takedown requests / DCMA</h1>
        <p>You can request a takedown of uploads violating Terms of Service or copyright laws here.</p>
        <form @submit.prevent="send">
            <label for="from">Email address (optional)</label>
            <p class="hint">Provide an email address if you'd like me to get back to you.</p>
            <input
                v-model="fd.from"
                type="text"
                id="from"
                name="from"
                placeholder="you@example.com"
            />
            <div>
                <label for="urls">URLs</label>
                <p
                    class="hint"
                >Provide URLs to the uploads you want to report. Please put each URL in a separate line.</p>
                <textarea
                    v-model="fd.urls"
                    name="urls"
                    id="urls"
                    cols="30"
                    rows="20"
                    placeholder="e.g. https://shareaudio.cc/qsCR9R2MJFyLdFEpUYupeS"
                ></textarea>
            </div>
            <div>
                <label for="text">Message</label>
                <p
                    class="hint"
                >Provide descriptive detail why the uploads you're reporting are abusive.</p>
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
                from: "",
                urls: "",
                text: ""
            }
        };
    },
    methods: {
        send() {
            this.axios
                .post("/abuse", this.fd)
                .then(r => {
                    this.$notify({
                        title: "Report sent!",
                        type: "Success",
                        duration: 3000
                    });
                })
                .catch(e => {
                    this.$notify({
                        type: "error",
                        title: "Report couldn't be sent.",
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
p.hint {
    padding: 0;
    margin: 0;
    font-size: 0.9em;
}
</style>