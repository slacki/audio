import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import VueAxios from 'vue-axios'
import Notifications from 'vue-notification'

import './scss/main.scss'

// axios
axios.defaults.baseURL = process.env.VUE_APP_API_ENDPOINT;
Vue.use(VueAxios, axios);

Vue.use(Notifications);

Vue.config.productionTip = false;
new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')
