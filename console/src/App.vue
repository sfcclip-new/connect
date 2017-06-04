<template>
  <v-app>
    <v-navigation-drawer persistent light v-model="drawer">
      <v-list　subheader>

        <v-subheader>ユニット</v-subheader>
        <v-list-item v-for="unit in units" :key="unit.ID">
          <v-list-tile avatar>
            <v-list-tile-avatar>
              <v-icon>list</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title v-text="unit.ID"></v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list-item>

        <v-list-item>
          <v-list-tile avatar>
            <v-list-tile-avatar>
              <v-icon>add</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title>新規追加</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list-item>

        <v-divider />

        <v-subheader>グループ</v-subheader>
        <v-list-item v-for="group in groups" :key="group.ID">
          <v-list-tile avatar @click.native.stop="$router.push(`/group/${group.ID}`)">
            <v-list-tile-avatar>
              <v-icon>list</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title v-text="group.ID"></v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list-item>

        <v-list-item>
          <v-list-tile avatar @click.native.stop="$router.push('/group')">
            <v-list-tile-avatar>
              <v-icon>add</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title>新規追加</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list-item>

        <v-divider />

        <v-subheader>メニュー</v-subheader>
        <v-list-item>
          <v-list-tile avatar @click.native.stop="drawer = !drawer">
            <v-list-tile-avatar>
              <v-icon>close</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title>閉じる</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list-item>

      </v-list>
    </v-navigation-drawer>

    <v-toolbar fixed light>
      <v-toolbar-side-icon light @click.native.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title v-text="title"></v-toolbar-title>
    </v-toolbar>
    <main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </main>
  </v-app>
</template>

<script>
import axios from 'axios';

export default {
  data () {
    return {
      drawer: true,
      title: 'connect.sfcclip.net',
      units: [],
      groups: [],
    }
  },
  created () {
    axios.get('/unit')
    .then(res => this.units = res.data || [])
    .catch(err => console.error(err));

    axios.get('/group')
    .then(res => this.groups = res.data || [])
    .catch(err => console.error(err));
  },
}
</script>

<style lang="stylus">
@import './stylus/main'
</style>
