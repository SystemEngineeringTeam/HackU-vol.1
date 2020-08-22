const URL_BASE = 'http://localhost:8088'

// これを自分のPCのIPアドレスに書き換えると自分のパソコンで立ち上げたGoに繋がるようになるよ
// const URL_BASE = 'http://{自分のipアドレス}:8088'

module.exports = {
  URL_BASE: URL_BASE,
  URL_TASKS: URL_BASE + '/tasks',
  URL_TASKS_SUCCESS: URL_BASE + '/tasks/success',
  URL_WEIGHTS: URL_BASE + '/tasks/weights',
  URL_USERS: URL_BASE + '/users',
  URL_SIGNUP: URL_BASE + '/users/signup',
  URL_LOGIN: URL_BASE + '/users/login',
  URL_HP: URL_BASE + '/hp',
}
