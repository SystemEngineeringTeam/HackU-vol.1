<template>
  <v-container>
    <v-row justify="center">
      <v-col md="4" sm="8" xs="10">
        <v-form ref="form" v-model="valid">
          <h1>サインアップ</h1>
          <v-text-field
            v-model="name"
            label="ユーザーネーム"
            required
            outlined
          ></v-text-field>

          <v-text-field
            v-model="email"
            label="E-mail"
            required
            outlined
          ></v-text-field>

          <v-text-field
            v-model="pass"
            :append-icon="passwdShow ? 'mdi-eye' : 'mdi-eye-off'"
            :type="passwdShow ? 'text' : 'password'"
            label="パスワード"
            required
            outlined
            @click:append="passwdShow = !passwdShow"
          ></v-text-field>

          <v-btn
            :disabled="!valid"
            color="success"
            class="mr-4"
            @click="create_user"
          >
            送信
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      valid: false,
      name: '',
      email: '',
      pass: '',
      passwdShow: false,
    }
  },
  methods: {
    create_user() {
      const post_json = {
        name: this.name,
        email: this.email,
        pass: this.pass,
      }
      this.$store.dispatch("user/signup", post_json)
    },
  },
}
</script>

<style></style>
