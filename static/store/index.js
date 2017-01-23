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
        // Component:
        // methods: {
        //     updateSearch(e) {
        //         this.$store.commit('SET_SEARCH', e.target.value)
        //     }
        // },
        LOAD_TASKS(state, tasks) {
            state.tasks = tasks
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
        // Component
        // methods: {
        //     remove(todo){
        //         this.$store.dispatch('removeTodo', todo)
        //     }
        // },
        loadTasks({commit}, tasks) {
            commit('LOAD_TASKS', tasks)
        },
        setSearch({commit}, search) {
            commit('SET_SEARCH', search)
        }
    },
    getters: {
        // Component:
        // computed: {
        //     todos(){
        //         return this.$store.getters.todos
        //     }
        // }
        completedTodos: state => state.todos.filter((todo) => todo.completed),
    }
});
