<template>
  <v-row>
    <v-col cols="12" sm="6">
      <v-menu
        v-model="datePick"
        :close-on-content-click="false"
        :nudge-right="40"
        transition="scale-transition"
        offset-y
        min-width="290px"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="date"
            label="tasks deadline date"
            readonly
            v-bind="attrs"
            v-on="on"
            @click="dateInitial"
          ></v-text-field>
        </template>
        <v-date-picker
          v-model="date"
          @input="datePick = false"
          no-title
          scrollable
        ></v-date-picker>
      </v-menu>
    </v-col>
    <v-col cols="12" sm="6">
      <v-menu
        ref="menu"
        v-model="timePick"
        :close-on-content-click="false"
        :nudge-right="40"
        :return-value.sync="time"
        transition="scale-transition"
        offset-y
        max-width="290px"
        min-width="290px"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="time"
            label="tasks deadline time"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-time-picker
          v-if="timePick"
          v-model="time"
          full-width
          @click:minute="$refs.menu.save(time)"
        ></v-time-picker>
      </v-menu>
    </v-col>
    <v-col cols="5">
      <v-btn @click="resetDeadline">日付リセット</v-btn>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'DeadlinePicker',

  data() {
    return {
      datePick: false,
      timePick: false,
    }
  },

  computed: {
    date: {
      get() {
        return this.$store.state.tasks.post.deadlineDate
      },
      set(value) {
        this.$store.commit('tasks/setPostDeadlineDate', value)
      },
    },
    time: {
      get() {
        return this.$store.state.tasks.post.deadlineTime
      },
      set(value) {
        this.$store.commit('tasks/setPostDeadlineTime', value)
      },
    },
  },
  methods: {
    resetDeadline: function () {
      this.$store.commit('tasks/setPostDeadlineDate', '')
      this.$store.commit('tasks/setPostDeadlineTime', '')
    },

    dateInitial: function () {
      if (this.date === '') {
        this.$store.commit(
          'tasks/setPostDeadlineDate',
          new Date().toISOString().substr(0, 10)
        )
      }
    },
  },
}
</script>
