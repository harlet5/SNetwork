<template>
<div class="error-div">
    <div class="error-text">
        <span style="--i:1">E</span>
        <span style="--i:2">R</span>
        <span style="--i:3">R</span>
        <span style="--i:4">O</span>
        <span style="--i:5">R</span>
        <span style="--i:6">.</span>
    </div>
</div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Kanit&display=swap');

.error-div {
    height: calc(100vh - 55px);
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 0px;
    padding: 0px;
    background-color: #292a2d;
    text-transform: uppercase;
}

.error-text span {
    position: relative;
    display: inline-block;
    font-size: 70px;
    font-family: 'Kanit', sans-serif;
    color: #fff;
    text-transform: uppercase;
    animation: flip 2.5s infinite;
    animation-delay: calc(.2s * var(--i))
}

@keyframes flip {
  0%,35% {
    transform: rotateY(360deg) 
  }
} 
</style>
