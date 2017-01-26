import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

export default new Vuex.Store({
    strict: process.env.NODE_ENV !== 'production',
    state: {
        sortKey: 'title',
        reverse: false,
        search: '',
        columns: ['title', 'body', 'date'],
        tasks: [],
    },
    mutations: {
        // In Component:
        // methods: {
        //     updateSearch(e) {
        //         this.$store.commit('SET_SEARCH', e.target.value)
        //     }
        // },
        LOAD_TASKS(state, tasks) {
            state.tasks = tasks
        },
        REMOVE_TASK(state, task){
            const tasks = state.tasks;
            tasks.splice(tasks.indexOf(task), 1)
        },
        SET_SEARCH(state, search) {
            state.search = search
        },
        SET_REVERSE(state, reverse) {
            state.reverse = reverse
        },
        SET_SORTKEY(state, sortKey) {
            state.sortKey = sortKey
        }
    },
    actions: {
        // In Component:
        // methods: {
        //     remove(todo){
        //         this.$store.dispatch('removeTodo', todo)
        //     }
        // },
        loadTasks({commit}, tasks) {
            commit('LOAD_TASKS', tasks)
        },
        async removeTask({commit}, task) {
            try {
                await $.ajax({
                    method: 'DELETE',
                    url: `/api/v1/tasks/${task.id}`,
                });
                commit('REMOVE_TASK', task)
            } catch (jqXHR) {
                const response = jqXHR.responseJSON;
                console.log('Error!', response);
            }
        },
        setSearch({commit}, search) {
            commit('SET_SEARCH', search)
        }
    },
    getters: {
        // In Component:
        // computed: {
        //     tasks(){
        //         return this.$store.getters.tasks
        //     }
        // }
        tasks: state => state.tasks.filter(item => item.title.indexOf(state.search) > -1),
    }
});
