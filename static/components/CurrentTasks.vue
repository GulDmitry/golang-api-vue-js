<template>
    <div class="container">
        <h3 v-if="tasks.length > 0">All Tasks ({{tasks.length}})</h3>
        <input :value="search" @input="setSearch($event.target.value)" class="form-control" placeholder="Filter title by name">

        <table class="table table-striped">
            <thead>
            <tr>
                <th v-for="column in columns" v-bind:class="classSort(column)">
                    <a href="#" @click="sortBy(column)">
                        {{ column | capitalize }}
                    </a>
                </th>
            </tr>
            </thead>

            <tbody>
            <tr v-for="item in filteredItems">
                <td>{{ item.title }}</td>
                <td>{{ item.body }}</td>
                <td>{{ item.date }}</td>
                <td>
                    <button type="button" @click="removeTask(item)" class="btn btn-danger btn-sm">
                        <span class="glyphicon glyphicon-remove-circle" /> Remove
                    </button>
                </td>
            </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
import {mapState, mapActions} from 'vuex'

export default{
    /*
    computed: {
        todos(){
            return this.$store.getters.todos
        }
    }
    */
    methods: {
        ...mapActions([
          'setSearch',
          'removeTask'
          // {updateSearch: 'setSearch'}
        ]),
        sortBy: function(sortKey) {
            this.$store.commit('SET_REVERSE', this.$store.state.sortKey == sortKey ? !this.$store.state.reverse : false)
            this.$store.commit('SET_SORTKEY', sortKey)
        },
        classSort(column) {
            const res = {}
            if (this.$store.state.sortKey == column) {
                const className = this.$store.state.reverse ? 'header-sort-asc' : 'header-sort-desc'
                res[className] = true
            }
            return res
        },
    },
    computed: {
        ...mapState({
            tasks: state => state.tasks,
            search: state => state.search,
            sortKey: state => state.sortKey,
            newUser: state => state.newUser,
            columns: state => state.columns,
            reverse: state => state.reverse
        }),
        filteredItems() {
            return this.$store.state.tasks.filter(item => item.title.indexOf(this.$store.state.search) > -1) // filterBy
                .sort((a, b) => {
                    let v1 = a[this.$store.state.sortKey], v2 = b[this.$store.state.sortKey];
                    if (this.$store.state.reverse) {
                        v2 = [v1, v1 = v2][0];
                    }
                    return v1 > v2 ? 1 : -1
                }); // orderBy
        }
    },
    filters: {
        capitalize(value) {
            if (!value) return '';
            value = value.toString();
            return value.charAt(0).toUpperCase() + value.slice(1)
        }
    }
}
</script>
