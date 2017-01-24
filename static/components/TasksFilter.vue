<template>
        <div id="tasks-filter">
            <h3>Filter</h3>
            <input :value="search" @input="setSearch($event.target.value)" class="form-control" placeholder="Filter title by name">
        </div>
</template>

<script>
import {mapState, mapActions} from 'vuex'

export default{
    methods: {
        ...mapActions([
          'setSearch',
        ]),
    },
    computed: {
        ...mapState({
            search: state => state.search,
        }),
    },
}
</script>
