<template>
  <v-row justify="center">
    <v-dialog v-model="postDialogBool" persistent max-width="600px">
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          v-bind="attrs"
          v-on="on"
          color="red"
          :bottom="true"
          :fixed="true"
          :right="true"
          fab
          dark
          large
          @click="postDialogBool = !postDialogBool"
        >
          <!-- タスク作成 -->
          <v-icon dark>mdi-plus</v-icon>
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="headline">タスクの作成</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  label="title"
                  v-model="title"
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <DeadlinePicker />
              </v-col>
              <v-col cols="12">
                <v-textarea
                  outlined
                  auto-grow
                  v-model="description"
                  label="description"
                >
                </v-textarea>
              </v-col>
              <v-col cols="4">
                <v-select
                  v-model="weight"
                  :items="weights"
                  :menu-props="{ maxHeight: '400' }"
                  label="Select"
                  hint="tasks weight"
                  persistent-hint
                ></v-select>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="red darken-2" @click="postCancel">Close</v-btn>
          <v-btn color="green darken-1" @click="doPost">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import DeadlinePicker from './DeadlinePicker'

export default {
  name: 'postDialog',

  components: {
    DeadlinePicker,
  },

  data: () => ({
    postDialogBool: false,
    weights: ['ぬるい', 'ふつう', 'えぐい'],
  }),

  methods: {
    postCancel() {
      this.postDialogBool = false
      this.$store.dispatch('tasks/postAllReset')
    },
    doPost() {
      this.postDialogBool = false
      this.$store.dispatch('tasks/postTask')
    },
  },

  computed: {
    title: {
      get() {
        return this.$store.state.tasks.post.title
      },
      set(value) {
        this.$store.commit('tasks/setPostTitle', value)
      },
    },
    description: {
      get() {
        return this.$store.state.tasks.post.description
      },
      set(value) {
        this.$store.commit('tasks/setPostDescription', value)
      },
    },
    weight: {
      get() {
        return this.weights[this.$store.state.tasks.post.weight]
      },
      set(value) {
        let index = this.weights.findIndex((element) => element === value)
        this.$store.commit('tasks/setPostWeight', index)
      },
    },
  },
}
</script>
