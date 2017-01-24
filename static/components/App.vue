<template>
    <div>
        <TasksFilter></TasksFilter>
        <br/>
        <TasksTable></TasksTable>
    </div>
</template>

<script>
import TasksTable from './TasksTable'
import TasksFilter from './TasksFilter'

export default {
  components: {
    TasksTable,
    TasksFilter
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
