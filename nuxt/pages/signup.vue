<template>
  <v-container>
    <v-row justify="center">
      <v-col md="4" sm="8" xs="10">
          <v-form ref="form" v-model="valid" >
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
              label="パスワード"
              required
              outlined
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
import axios from 'axios'
export default {
  data() {
    return {
      valid: false,
      name: '',
      email: '',
      pass: '',
    }
  },
  methods: {
    create_user() {
      const post_json = {
        name: this.name,
        email: this.email,
        pass: this.pass,
      }
      console.log(post_json)
      axios.post('localhost:8080/users/signup', post_json).then((res) => {
        if (res.status == 200) {
          console.log('ok!')
        }
      })
      this.$router.push('/')
    },
  },
}
</script>

<style></style>
