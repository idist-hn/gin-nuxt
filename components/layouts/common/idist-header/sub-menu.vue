<template>
  <div class="sub-menu">
    <div v-if="!loading" v-for="entry in entries" :key="entry.id" class="sub-menu-item">
      <nuxt-link :to="entry.url">
        {{ entry.title }}
      </nuxt-link>
    </div>
    <div class="sub-menu-item" v-else>
      <a-spin :size="`small`"/>
    </div>
    <div class="sub-menu-item">
      <a href="#">
        <idist-feather-icons type="mail"/>
      </a>
    </div>
    <div class="sub-menu-item">
      <nuxt-link :to="{ name: 'tim-kiem' }">
        <idist-feather-icons type="search"/>
      </nuxt-link>
    </div>
  </div>
</template>

<script>
import IdistFeatherIcons from "~/components/commons/IdistFeatherIcons.vue";
import {useArticleStore} from "~/stores/articles.js";
import {useMenuStore} from "~/stores/menus.js";

export default {
  name: `SubMenu`,
  components: {IdistFeatherIcons},
  data: () => ({
    loading: true,
    entries: [],
  }),
  computed: {},
  methods: {
    async getData() {
      this.loading = true
      let response = await useMenuStore().GetSubMenuEntry()
      let menu = response.data?.entry

      if (menu) {
        this.entries = menu.items
      }

      console.log(this.entries)
      this.loading = false
    },
  },
  created() {
    this.getData()
  },
  mounted() {
  },
};
</script>


