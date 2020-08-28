<template>
  <v-row>
    <!-- <v-btn color="primary" dark @click="$store.commit('game/setHP', 25)">
      HP減らす
    </v-btn> -->
    <v-dialog v-model="dialog" max-width="450">
      <v-card>
        <v-row justify="center">
          <v-col cols="6">
            <v-img
              src="https://4.bp.blogspot.com/-TalVsZX9oJU/U0pS4v8f_mI/AAAAAAAAe-w/YA6n6klmR6s/s800/kamisama.png"
            />
          </v-col>
        </v-row>
        <v-card-text>
          <p>{{ this.$store.state.user.name }}、君は死んでしまった！</p>
          <p>タスクを抱えすぎてしまったようじゃな…</p>
          <p>時には休息も必要じゃ</p>
          <p>今あるタスクをしっかり終わらせてしばらく休むが良い。</p>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn outlined @click="dialog = false"
            >はい</v-btn
          >
          <v-spacer></v-spacer>
          <v-btn outlined @click="dialog = false"
            >いいえ</v-btn
          >
          <v-spacer></v-spacer>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  data: () => ({
    intervalID: null,
    dialog: false,
  }),

  computed: {
    dieFlag: function () {
      return this.$store.getters['game/dieFlag']
    },
  },

  watch: {
    dieFlag(val) {
      this.dialog = val
    },
  },
}
</script>
