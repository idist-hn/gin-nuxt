<template>
  <div class='nav-author'>
    <a-popover placement='bottomRight' trigger='click'>
      <a-avatar src='/no-image.png' :size='32'/>
      <template v-slot:content>
        <div class='user-dropdown'>
          <figure class='user-dropdown__info'>
            <img v-if='!profile || !profile.avatar' src='~/assets/images/no-image.png' style='width: 45px;' alt=''/>
            <img v-else :src='user.avatar' alt=''/>
            <figcaption>
              <h4>{{ profile ? profile.name : '' }}</h4>
              <p>{{ (profile && profile.role) ? user.role.name : 'đang cập nhật' }}</p>
            </figcaption>
          </figure>
          <ul class='user-dropdown__links'>
            <li>
              <nuxt-link :to="{name :'admin-profile'}">
                <IdistFeatherIcons type='user'/>
                Thông tin
              </nuxt-link>
            </li>
          </ul>
          <a @click='SignOut' class='user-dropdown__bottomAction' href='#'>
            <IdistFeatherIcons v-if='!loadingSignOut' type='log-out'/>
            <a-spin v-else/>
            Đăng xuất </a>
        </div>
      </template>
    </a-popover>
  </div>
</template>

<script>

import IdistFeatherIcons from "~/components/commons/IdistFeatherIcons.vue";
import {useProfileStore} from '~/stores/profile'
import {mapState} from 'pinia'
export default {
  name: 'header-user',
  components: {IdistFeatherIcons},
  data: () => ({
    loadingSignOut: false
  }),
  computed: {
    ...mapState(useProfileStore, ['profile']),
  },
  methods: {
    async SignOut() {
      this.loadingSignOut = true
      let response = await useProfileStore().PostLogout()
      if (response.code === 200) {
        this.$toast('Đăng xuất thành công', {type: 'success', autoClose: '3000'})
        this.$router.push({name: 'auth-login'})
      }
      this.loadingSignOut = false
    }
  },
  created() {
  },
  mounted() {
  }
}
</script>

<style scoped>

</style>
