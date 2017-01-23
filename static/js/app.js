import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.css'
import '../css/font-awesome.css'
import '../css/theme.css'
import 'babel-polyfill'

import Vue from 'vue'
import App from '../components/App'
import store from '../store'

$(document).ready(function() {
    new Vue({
        el: '#app',
        store,
        render: h => h(App),
        mounted() {
            // Load first data set here.
            this.loadTasks();
        },
        methods: {
            loadTasks() {
                const ajaxPromise = Promise.resolve($.ajax({
                    method: 'GET',
                    url: '/api/v1/tasks',
                }));
                ajaxPromise
                    .then((response) => {
                        this.$store.dispatch('loadTasks', response)
                    })
                    .catch((jqXHR) => {
                        const response = jqXHR.responseJSON;
                        console.log('Error!', response);
                        // in case of chain like then({}).then{} the `return new Promise(() => reject());` is required
                    });
            }
        },
    });
});
