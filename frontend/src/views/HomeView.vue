<script setup lang="ts">
import { useI18n } from "vue-i18n";

import HelloWorld from "@/components/HelloWorld.vue";
import { ref } from 'vue'
import { WriteToLeader } from "../../wailsjs/go/main/App";
import { ReadFromLeader } from "../../wailsjs/go/main/App";

const msg = ref('')
const blocks = ref([
  { command: "ls -al", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "cat ./me.txt", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: 'jq "". | ."', output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "find me", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "ls -al", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "cat ./me.txt", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: 'jq "". | ."', output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "find me", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "ls -al", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "cat ./me.txt", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: 'jq "". | ."', output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "find me", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "ls -al", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "cat ./me.txt", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: 'jq "". | ."', output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "find me", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "ls -al", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "cat ./me.txt", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: 'jq "". | ."', output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "find me", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "ls -al", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "cat ./me.txt", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: 'jq "". | ."', output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "find me", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "ls -al", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "cat ./me.txt", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: 'jq "". | ."', output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "find me", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "ls -al", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "cat ./me.txt", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: 'jq "". | ."', output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "find me", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "ls -al", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "cat ./me.txt", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: 'jq "". | ."', output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },
  { command: "find me", output: "dlsks kdksldskldsk dslkdskdkdskld kdskldskldksdldk" },

])


const { t } = useI18n();
</script>

<script lang="ts">

export default {
  data() {
    return {
      msg2: ''
    }
  },
  methods: {
    onEnter(msg: string) {
      WriteToLeader(msg)
      msg = ""
      setTimeout(this.continueExecution, 200);
    },
    continueExecution() {
      ReadFromLeader().then(
        response => {
          this.msg2 = response
          console.log(this.msg2)
        }
      )
    },
  }
}

function sleep(ms: number) {
  return new Promise(resolve => setTimeout(resolve, ms));
}


</script>

<template>
  <div class="home">
    <div style="height: 95vh; overflow-y: scroll">
      <div class="card" v-for="c in blocks">
        <div style="padding: 20px 20px 20px 20px;">
          <h1 style="margin-bottom: 10px;">> {{ c.command }}</h1>
          <h3 style="color: rgb(89, 89, 89);">{{ msg2 }}</h3>
        </div>
      </div>
    </div>
    <div class="command">
      <input placeholder="~" v-model="msg" v-on:keyup.enter="onEnter(msg)">
    </div>
  </div>
</template>

<style lang="scss">
.home {
  display: flex;
  flex-direction: column;


  .command {
    height: 5vh;
  }

  input {
    color: white;
    background-color: transparent;
    margin: 0px;
    padding: 0px 5px 0px 15px;
    width: 100%;
    height: 100%;

    :focus {
      border: none
    }

    :hover {
      border: none;
    }
  }

  .card {
    color: silver;
    display: flex;
    width: 100vw;
    flex-direction: column;
    border: 0.1px solid transparent;
    border-left: 3px solid transparent;
  }

  .card:hover {
    border: 0.1px solid grey;
    border-left: 3px solid grey;

  }

  .link {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    justify-content: center;
    margin: 18px auto;

    .btn {
      display: block;
      width: 150px;
      height: 50px;
      line-height: 50px;
      padding: 0 5px;
      margin: 0 30px;
      border-radius: 8px;
      text-align: center;
      font-weight: 700;
      font-size: 16px;
      white-space: nowrap;
      text-decoration: none;
      cursor: pointer;

      &.start {
        background-color: #fd0404;
        color: #ffffff;

        &:hover {
          background-color: #ec2e2e;
        }
      }

      &.star {
        background-color: #ffffff;
        color: #fd0404;

        &:hover {
          background-color: #f3f3f3;
        }
      }
    }
  }
}
</style>
