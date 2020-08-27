<template>
  <v-app dark>
    <v-app-bar fixed app >
      <!-- <v-toolbar-title v-text="title" @click="$router.push('/')"/> -->
      <img src="@/static/logo.png" height="110px" style="position: relative; right: 30px; bottom: 5px" />
      <v-spacer />
      <template v-if="$store.state.user.name">
        <v-toolbar-title class="mr-2" v-text="$store.state.user.name" />
        <v-btn color="red" @click="logout">
          logout
        </v-btn>
      </template>
      <template v-else>
        <v-btn class="mr-2" color="primary" to="/signup">
          signup
        </v-btn>
        <v-btn color="primary" to="/login">
          login
        </v-btn>
      </template>
    </v-app-bar>
    <v-main>
      <v-container>
        <nuxt />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      title: 'Vuetify.js',
    }
  },
  methods: {
    logout() {
      this.$store.dispatch('user/logout')
      this.$store.dispatch('tasks/logout')
      this.$store.dispatch('game/logout')
    },
  },
}
</script>
