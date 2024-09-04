<template>
  <a-card title="Thông tin" class="mb-10">
    <div v-for="item in value.extras" :key="item.id">
      <a-row :gutter="16" class="mb-10">
        <a-col :xl="10" :sm="24">
          <a-input v-model:value="item.title"/>
        </a-col>
        <a-col :xl="12" :sm="24">
          <a-input v-model:value="item.value"/>
        </a-col>
        <a-col :xl="2" :sm="24">
          <a-button type="danger" @click="deleteData(item)">
            <idist-feather-icons type="delete"/>
          </a-button>
        </a-col>
      </a-row>
    </div>
    <div class="mt-20">

      <a-button tyle="success" @click="addData()">Thêm thông tin</a-button>
    </div>
  </a-card>
</template>

<script>
import {now} from "lodash";
import IdistFeatherIcons from "/components/commons/IdistFeatherIcons.vue";

export default {
  name: "DepartmentExtra",
  components: {IdistFeatherIcons},
  props: {
    value: Object,
  },
  methods: {
    addData() {
      this.value.extras.push({
        id: now(),
        title: "",
        value: "",
      });
    },
    deleteData(item) {
      this.value.extras = this.value.extras.filter((i) => i.id !== item.id);
    }
  },
  created() {
    if (!this.value.extras) {
      this.value.extras = [];
    }
  }
}
</script>
<style scoped>

</style>
