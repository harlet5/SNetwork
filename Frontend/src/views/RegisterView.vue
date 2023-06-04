<script setup lang="ts">
import Ws from "@/Websocket";
import { onUnmounted, ref, computed } from "vue";
import { useErrorsStore } from "@/stores/errors";
const store = useErrorsStore()
const ageError = ref();
const emailError = computed(() => store.registerErrorEmail)
const usernameError = computed(() => store.registerErrorUsername)
const firstnameError = ref();
const lastnameError = ref();
const passwordError = ref();
const username = ref();
const age = ref();
const gender = ref();
const firstname = ref();
const lastname = ref();
const email = ref();
const password = ref();
const password2 = ref();
const aboutme = ref();
const nickname = ref();
let avatarpic: any = ref();
const previewpic = ref("");
interface ustruct {
  uname?: string;
  upass?: string;
  umail?: string;
  uage?: string;
  ugender?: string;
  ufname?: string;
  ulname?: string;
  upic?: File;
  unickname?: string;
  uaboutme?: string;
}
const onFileSelected = (event: any) => {
  if (event.target.files[0].size > 1048576 || event.target.files[0].type.substring(0, 5) !== "image") {
    return;
  }
  const reader = new FileReader();
  reader.addEventListener("load", (e) => {
    avatarpic = reader.result;
  });
  reader.readAsDataURL(event.target.files[0]);
  previewpic.value = URL.createObjectURL(event.target.files[0]);
};

const registerUser = async () => {
  passwordError.value =
    password.value === password2.value ? "" : "Passwords don't match";
  ageError.value =
    new Date() > new Date(age.value) ? "" : "You from the future?";
  firstnameError.value = /^[A-Za-z]+$/.test(firstname.value)
    ? ""
    : "Letters only";
  lastnameError.value = /^[A-Za-z]+$/.test(lastname.value)
    ? ""
    : "Letters only";
  if (
    passwordError.value ||
    ageError.value ||
    lastnameError.value ||
    firstnameError.value
  ) {
    return;
  }
  const user: ustruct = {
    uname: username.value,
    upass: password.value,
    umail: email.value,
    uage: age.value,
    ugender: gender.value,
    ufname: firstname.value,
    ulname: lastname.value,
    upic: (typeof avatarpic === "object") ? "" : avatarpic,
    unickname: nickname.value ? nickname.value : "",
    uaboutme: aboutme.value ? aboutme.value : "",
  };
  console.log(user);
  Ws.send({
      Page: "signup",
      Data: user,
    }
  );
};
onUnmounted(() => {
  username.value = "";
  age.value = "";
  gender.value = "";
  firstname.value = "";
  lastname.value = "";
  email.value = "";
  password.value = "";
  password2.value = "";
  store.registerErrorEmail = "";
  store.registerErrorUsername = "";
});
</script>
<template>
  <div class="bg">
    <div class="wrapper">
      <h1>Sign up</h1>
      <form @submit.prevent="registerUser">
        <div class="form">
          <label for="username" class="form__label">Username</label>

          <input
            id="username"
            type="text"
            required
            v-model="username"
            class="form__input"
            autocomplete="off"
            placeholder=" "
          />
        </div>
        <div v-if="usernameError" class="error">{{ usernameError }}</div>

        <div class="form">
          <label for="age" class="form__label">Date of birth</label>

          <input
            id="age"
            class="form__input"
            type="date"
            required
            v-model="age"
            autocomplete="off"
            placeholder=" "
          />
        </div>
        <div v-if="ageError" class="error">{{ ageError }}</div>

        <div class="form genderindput">
          <label for="gender" class="form__label">Gender</label>

          <select v-model="gender" id="gender" class="form__input">
            <option value="male">Male</option>
            <option value="female">Female</option>
          </select>
        </div>

        <div class="form">
          <label for="firstname" class="form__label">First name:</label>

          <input
            id="firstname"
            class="form__input"
            type="text"
            required
            v-model="firstname"
            autocomplete="off"
            placeholder=" "
          />
        </div>
        <div v-if="firstnameError" class="error">{{ firstnameError }}</div>

        <div class="form">
          <label for="lastname" class="form__label">Last name</label>

          <input
            id="lastname"
            class="form__input"
            type="text"
            required
            v-model="lastname"
          />
        </div>
        <div v-if="lastnameError" class="error">{{ lastnameError }}</div>

        <div class="form">
          <label for="email" class="form__label">Email</label>

          <input
            id="email"
            class="form__input"
            type="email"
            required
            v-model="email"
          />
        </div>
        <div v-if="emailError" class="error">{{ emailError }}</div>

        <div class="form">
          <label for="avatar" class="form__label"
            >Avatar (optional 1MB max)</label
          ><br />
          <input type="file" accept="image/*" @change="onFileSelected" />
          <div class="image-preview" id="image-preview">
            <img
              :src="previewpic"
              alt="Image Preview"
              class="image-preview__image"
              style="height: 100%; width: 100%; object-fit: contain"
            />
          </div>
        </div>
        <div class="form">
          <label for="nickname" class="form__label">Nickname (optional)</label>
          <input
            id="nickname"
            class="form__input"
            type="text"
            v-model="nickname"
          />
        </div>

        <div class="form">
          <label for="password" class="form__label">Password</label>

          <input
            id="password"
            class="form__input"
            type="password"
            required
            v-model="password"
          />
        </div>

        <div class="form">
          <label for="password2" class="form__label">Retype password</label>
          <input
            id="password2"
            class="form__input"
            type="password"
            required
            v-model="password2"
          />
        </div>
        <div class="form">
          <label for="aboutme" class="form__label">About me (optional)</label>
          <input
            id="aboutme"
            class="form__input"
            type="textarea"
            v-model="aboutme"
          />
        </div>
        <div v-if="passwordError" class="error">{{ passwordError }}</div>

        <div class="submit">
          <button>Register</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.image-preview {
  width: 300px;
  min-height: 100px;
  border: 2px solid #dddddd;
  margin-top: 15px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  color: #cccccc;
}

.wrapper {
  width: 330px;
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
.form__input {
  width: 92%;
  outline: none;
  border: 1px solid #fff;
  padding: 12px 20px;
  margin-bottom: 10px;
  border-radius: 20px;
  background: #e4e4e4;
}
button {
  font-size: 1rem;
  margin-top: 1.8rem;
  padding: 10px 0;
  border-radius: 20px;
  outline: none;
  border: none;
  width: 90%;
  color: #fff;
  cursor: pointer;
  background: rgb(17, 107, 143);
}
button:hover {
  background: rgba(17, 107, 143, 0.877);
}
input:focus {
  border: 1px solid rgb(192, 192, 192);
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

.error {
  color: red;
}
</style>
