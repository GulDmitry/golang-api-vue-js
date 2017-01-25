import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.css'
import '../css/font-awesome.css'
import '../css/theme.css'
import 'babel-polyfill'

import Vue from 'vue'
import App from '../components/App'
import store from '../store'

$(document).ready(function() {
    const appConfig = {
        el: '#app',
        store,
        render: h => h(App),
        mounted() {
            // Load first data set here.
            this.loadTasks();
        },
        methods: {
            loadTasks() {
                const ajaxPromise = $.ajax({
                    method: 'GET',
                    url: '/api/v1/tasks',
                });
                ajaxPromise
                    .then((response) => {
                        this.$store.dispatch('loadTasks', response);
                        // in case of chain like then({}).then{} the `return new Promise(() => reject());` is required
                        // or return anything.
                    })
                    .catch((jqXHR) => {
                        const response = jqXHR.responseJSON;
                        console.log('Error!', response);
                    });
            }
        },
    };

    if ($('#app').length) {
        // Or (new Vue()).$mount("#app")
        new Vue(appConfig)
    }
});
