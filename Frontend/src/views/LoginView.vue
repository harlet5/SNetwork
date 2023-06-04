<script setup lang="ts">
import Ws from "@/Websocket"
import { onUnmounted, ref, computed } from "vue";
import { useErrorsStore } from "@/stores/errors";
const store = useErrorsStore()

const username = ref("");
const password = ref();
//const chats: Ref<Chat[]> = computed<Chat[]>(() => chatStore.chats as Chat[]);
const loginError = computed(() => 
  store.loginError 
)
interface ustruct {
  uname: string;
  upass: string;
}
const loginUser = async () => {
  const user: ustruct = {
    uname: username.value,
    upass: password.value,
  };
  Ws.send({
      Page: "login",
      Data: user,
    }
  );
};


onUnmounted(() => {
  username.value = "";
  password.value = "";
 // inputError.value = "";
  store.setLoginError("")
});
</script>

<template>
      <div class="wrapper">
      <h1>Login</h1>
  <div class="grid-login">
    <form @submit.prevent="loginUser()" class="loginform">
      <div>
        <div class="form">
          <label for="email" class="form__label">Username or email</label>

          <input
            required
            v-model="username"
            type="text"
            id="email"
            class="form__input"
            autocomplete="off"
            placeholder=" "
          />
        </div>
      </div>
      <div>
        <div class="form">
          <label for="password2" class="form__label">Password</label>

          <input
            required
            v-model="password"
            type="password"
            id="password2"
            class="form__input"
            autocomplete="off"
            placeholder=" "
          />
        </div>
        <div class="error">{{ loginError }}</div>
      </div>

      <div class="submit">
        <button>Log in</button>
      </div>
    </form>
  </div>
</div>
</template>

<style scoped>
*{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
.wrapper{
  width:330px;
  padding: 2rem 1rem;
  margin: 50px auto;
  background-color: #fff;
  border-radius: 10px;
  text-align: center;
  box-shadow: 0 20px 35px rgba(0, 0, 0, 0.1);
}
h1 {
  font-size: 2rem;
  color: #07001f;
  margin-bottom: 1.2rem;

}
.form__input{
  width: 92%;
  outline: none;
  border: 1px solid #fff;
  padding: 12px 20px;
  margin-bottom: 10px;
  border-radius: 20px;
  background: #e4e4e4;

}
button{
  font-size: 1rem;
  margin-top: 1.8rem;
  padding: 10px 0;
  border-radius: 20px;
  outline: none;
  border: none;
  width: 90%;
  color: #fff;
  cursor: pointer;
  background: rgb(17, 107, 143)
}
button:hover{
  background: rgba(17, 107, 143, 0.877)
}
input:focus {
  border: 1px solid rgb(192, 192, 192)
}
.member {
  font-size: 0.8rem;
  margin-top: 1.4rem;
  color: #636363;
  
}
.member a {
  color: (17, 107, 143);
  text-decoration: none;
}
</style>
