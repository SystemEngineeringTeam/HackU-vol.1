<template>
  <v-card :outlined="true">
    <v-card-title class="justify-center">
      {{ task.title }}
    </v-card-title>
    <v-card-subtitle v-if="subtitleBool">
      <span>{{ this.task.deadlineDate }}</span>
      <span v-if="deadlineSpaceBool">{{ space }}</span>
      <span>{{ this.task.deadlineTime }}</span>
      <span v-if="weightSlashBool">/</span>
      <span>{{ this.task.weight }}</span>
    </v-card-subtitle>
    <v-card-text v-if="task.description !== ''">
      {{ task.description }}
    </v-card-text>
    <v-card-actions class="justify-center">
      <v-btn @click="success">達成</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'Task',

  props: ['task'],

  data: () => ({
    space: ' ',
  }),

  methods: {
    success: function () {
      let index = this.$store.state.tasks.tasks.findIndex(
        (element) => element.id === this.task.id
      )
      this.$store.dispatch('tasks/successTask', index)
      this.$store.commit('tasks/setTaskIDbyDeleteAnimation', index)
      this.$store.dispatch('game/writeSuccessLog', this.task.title)
    },
  },

  computed: {
    subtitleBool() {
      if (
        this.task.deadlineDate === '' &&
        this.task.deadlineTime === '' &&
        this.task.weight === ''
      ) {
        return false
      } else {
        return true
      }
    },
    deadlineSpaceBool() {
      if (this.task.deadlineDate === '' || this.task.deadlineTime === '') {
        return false
      } else {
        return true
      }
    },
    weightSlashBool() {
      if (
        this.task.weight === '' ||
        (this.task.deadlineDate === '' &&
          this.task.deadlineTime === '' &&
          this.task.weight !== '')
      ) {
        return false
      } else {
        return true
      }
    },
  },
}
</script>
