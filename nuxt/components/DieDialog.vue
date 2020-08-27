<template>
  <v-row>
    <v-btn color="primary" dark @click="$store.commit('game/setHP', 25)">
      HP減らす
    </v-btn>
    <v-dialog v-model="dialog" max-width="400">
      <v-card>
        <v-img
          src="https://4.bp.blogspot.com/-TalVsZX9oJU/U0pS4v8f_mI/AAAAAAAAe-w/YA6n6klmR6s/s800/kamisama.png"
        />
        <v-card-text>
          君は死んでしまった！
          あまり無理をするんじゃないぞ…
          今あるタスクをしっかり終わらせてしばらく休むが良い。
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" text @click="dialog = false"
            >Disagree</v-btn
          >
          <v-btn color="green darken-1" text @click="dialog = false"
            >Agree</v-btn
          >
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
