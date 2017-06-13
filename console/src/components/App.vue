<template>
  <v-app>
    <v-navigation-drawer persistent light v-model="drawer">
      <v-list subheader>

        <v-subheader>ユニット</v-subheader>
        <v-list-item v-for="unit in units" :key="unit.id">
          <v-list-tile avatar @click.native.stop="$router.push(`/unit/${unit.id}`)">
            <v-list-tile-avatar>
              <v-icon>list</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title v-text="unit.id"></v-list-tile-title>
              <v-list-tile-sub-title v-html="unit.attributes.name"></v-list-tile-sub-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list-item>

        <v-list-item>
          <v-list-tile avatar @click.native.stop="$router.push('/unit')">
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
        <v-list-item v-for="group in groups" :key="group.id">
          <v-list-tile avatar @click.native.stop="$router.push(`/group/${group.id}`)">
            <v-list-tile-avatar>
              <v-icon>list</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title v-text="group.id"></v-list-tile-title>
              <v-list-tile-sub-title v-html="group.attributes.name"></v-list-tile-sub-title>
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
        <router-view :units="units" @update="reload"></router-view>
      </v-container>
    </main>
  </v-app>
</template>

<script>
import api from '../api'
import { Group, Unit } from '../models'

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
    this.reload()
  },
  methods: {
    reload () {
      api.get('/units')
      .then(res => this.units = Unit.list(res.data) || [])
      .catch(err => console.error(err));

      api.get('/groups')
      .then(res => this.groups = Group.list(res.data) || [])
      .catch(err => console.error(err));
    }
  },
  events: {
    updated() {
      this.reload()
    }
  }
}
</script>

<style lang="stylus">
@import '../stylus/main'
</style>
