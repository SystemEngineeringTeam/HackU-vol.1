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
      <v-card :outlined="true">
        <v-card-title>
          タスクの作成
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  label="title"
                  v-model="title"
                  :rules="[required]"
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-btn
                  v-if="!detailPicker"
                  @click="detailPicker = true"
                  block
                  text
                  color="primary"
                  >詳細設定</v-btn
                >
                <v-btn
                  v-if="detailPicker"
                  @click="detailPicker = false"
                  block
                  text
                  color="primary"
                  >詳細設定を閉じる</v-btn
                >
              </v-col>
              <div v-if="detailPicker">
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
                    :items="$store.state.tasks.weights"
                    :menu-props="{ maxHeight: '400' }"
                    label="Select"
                    hint="tasks weight"
                    persistent-hint
                  ></v-select>
                </v-col>
              </div>
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
    required: (value) => !!value || 'please input title',
    detailPicker: false,
  }),

  methods: {
    postCancel() {
      this.postDialogBool = false
      this.$store.dispatch('tasks/postAllReset')
      this.detailPicker = false
    },
    doPost() {
      if (this.title === '') {
        return
      }
      this.postDialogBool = false
      this.$store.dispatch('tasks/postTask')
      this.$store.dispatch('game/writeEnterLog', this.title)
      this.detailPicker = false
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
        return this.$store.state.tasks.post.weight
      },
      set(value) {
        this.$store.commit('tasks/setPostWeight', value)
      },
    },
  },

  created() {
    this.$store.dispatch('tasks/getWeights')
  },
}
</script>
