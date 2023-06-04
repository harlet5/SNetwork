<script setup lang="ts">
import { useRouter } from "vue-router";
const router = useRouter();
const post = defineProps(["post"]);
const navigateToPost = (id: number) => {
  router.push(`posts/${id}`);
};
const navigateToProfile = (id: number) => {
  router.replace(`/user/${id}`);
};

</script>
<template>
  <div class="post" >
    <div class="post-header">
      <img
        class="author-profile-picture"
        :src="`../../images/${post.post.TProf}`"
        alt="Author Profile Picture"
      />
      <div class="author-name" @click="navigateToProfile(post.post.TUId)">{{ post.post.TUName }}</div>
      <div class="date-posted">{{post.post.TTime}}</div>
      <div class="privacy-setting">Privacy Setting: {{ post.post.TPriv }}</div>
      
    </div>
    <div class="post-content" @click="navigateToPost(post.post.TId)">
      <p>{{ post.post.TBody }}</p>
     <!--  <button class="previous-image-button">&lt;</button> -->
    <div class="all-images">
      <div class="post-images" v-for="pic in post.post.TPics">
        <img :src=pic  alt="Image 1" />
      </div>
    </div>
      <!-- <button class="next-image-button">&gt;</button> -->
      <div class='all-post-categories'>
        <div class="post-categories" v-for="cat in post.post.TCats">
          <span class="category">{{ cat.CName }}</span>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
.post {
  min-width: 250px;
  margin: 0 auto;
  padding: 10px;
  margin-bottom: 50px;
  border-radius: 10px;
  display: grid;
  grid-template-columns: 140px auto;
  box-sizing: border-box;
  box-shadow: rgba(50, 50, 93, 0.25) 0px 13px 27px -5px, rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
}

.post-header {
  display: flex;
  align-items: center;
  margin-right: 25px;
  flex-direction: column;
  width: 80px;
  margin-left: 25px;
}

.author-profile-picture {
  width: 50px;
  height: 50px;
  object-fit: cover;
  border-radius: 50%;
  border: 1px solid black;
  font-size: 10px;
  text-align: center
}

.author-name {
  font-weight: bold;
  font-size: 1rem;
  margin-bottom: 7px;
}

.privacy-setting {
  font-style: italic;
  font-size: 0.7rem;
  text-align: center
}

.date-posted {
    font-size: 0.7rem;
    text-align: center;
}

.post-content {
  display: flex;
  flex-direction: column;
  height: auto;
  min-height: 100px;
  width: auto;
}

.post-content p {
  height: 75%
}

.all-images {
  display: flex;
  flex-wrap: wrap;
  margin-bottom: 10px;
}

.post-images img {
  max-width: 100px;
  max-height: 100px;
  object-fit: cover;
  margin-right: 10px;
  margin-bottom: 10px;
}

.next-image-button {
  position: absolute;
  top: 50%;
  right: 0;
  transform: translateY(-50%);
}
.previous-image-button {
  position: absolute;
  top: 50%;
  left: 0;
  transform: translateY(-50%);
}

.all-post-categories {
  display: flex;
  gap: 3px;
}

.category {
  background-color: gray;
  color: white;
  padding: 2px 5px;
  margin-right: 5px;
  margin-bottom: 5px;
  border-radius: 5px;
  font-size: 14px;
}

</style>
<!--

 TBody: "tere2"​​​​​
TCats: Array [ {…} ]​​​​​​
0: Object { CId: 0, CName: "teema#1" }​​​​​​
length: 1​​​​​​
<prototype>: Array []​​​​​
TId: 8​​​​​
TTime: "2.20.2023"​​​​​
TTitle: "tere"​​​​​
TUId: 8​​​​​
TUName: "123"
-->
