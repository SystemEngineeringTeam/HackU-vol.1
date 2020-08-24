<template>
  <v-card :outlined="true">
    <v-card-title class="justify-center">
      ステータス情報
    </v-card-title>
    <v-textarea v-if="getHP <= 25" solo readonly v-model="$store.state.game.log" STYLE="color:#F44336"></v-textarea>
    <v-textarea v-else-if="getHP <= 50" solo readonly v-model="$store.state.game.log" STYLE="color:#FFEB3B"></v-textarea>
    <v-textarea v-else solo readonly v-model="$store.state.game.log"></v-textarea>
    <HPGauge />
  </v-card>
</template>

<script>
import HPGauge from '../components/HPGauge'

export default {
  name: 'Logbox',

  components: {
    HPGauge,
  },

  computed: {
    getHP() {
      return (this.$store.state.game.HP / this.$store.state.game.maxHP) * 100
    },
  },

  data: () => ({}),

  methods: {},
}
</script>
