<script setup lang="ts">
import { useRouter } from "vue-router";
const router = useRouter();
const uGroup = defineProps(["uGroup"]);
const navigateToGroup = (id: number) => {
  router.push(`groups/${id}`);
};
</script>

<template>
  <div class="group-container">
    <div class="group" @click="navigateToGroup(uGroup.uGroup.Id)">
      <div class="group-content">
        <h2 class="group-title">{{ uGroup.uGroup.Name }}</h2>
      </div>
    </div>
  </div>
</template>

<style scoped>
.heading {
  font-size: 20px;
  font-weight: bold;
  margin-bottom: 15px;
}

.group-container {
  display: flex;
  justify-content: center;
  width: 100%;
}

.group {
  max-width: 50ch;
  width: 80%;
  background-color: #f0f0f0;
  border-radius: 4px;
  padding: 10px;
  margin-bottom: 10px;
  cursor: pointer;
  transition: background-color 0.2s;
  box-sizing: border-box;
}

.group:hover {
  background-color: #e0e0e0;
}

.group-content {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.group-title {
  font-size: 18px;
  font-weight: bold;
  margin: 0;
  text-align: center;
}
</style>
