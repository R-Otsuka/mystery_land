<template>
  <div id="app">
    <button @click="toggleBtn">表示／非表示ボタン</button>
    <transition name="simple">
      <p v-show="show" style="width:150px;height:150px;background-color:green">本日は晴天なり</p>
    </transition>
    <button @click="count++">+ju</button>
    <transition name="datachange">
      <div :key="count">{{ count }}</div>
    </transition>
    <button @click="order=!order">切り替え</button>
    <transition-group name="group" tag="ul" class="list">
      <li v-for="item in sortedList" v-bind:key="item.id">
        {{ item.name }} {{ item.price }}円
      </li>
    </transition-group>
    <button @click="change=!change">切り替え</button>
    <transition name="hook" v-on:enter="enter" v-on:after-enter="afterEnter">
      <div v-if="change">example</div>
    </transition>
  </div>
</template>

<script>
export default {

}
</script>
