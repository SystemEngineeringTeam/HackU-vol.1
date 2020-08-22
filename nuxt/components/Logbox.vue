<template>
  <v-card :outlined="true">
    <v-card-title class="justify-center">
      ステータス情報
    </v-card-title>
    <v-textarea solo readonly v-model="$store.state.game.log"></v-textarea>
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

  data: () => ({}),

  methods: {},
}
</script>
