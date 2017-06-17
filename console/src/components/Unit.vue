<template>
  <v-container fluid>
    <v-text-field v-if="unit.id" name="id" label="ユニットID"  v-model="unit.id" disabled></v-text-field>
    <v-text-field name="name" label="ユニット名" v-model="unit.attributes.name"></v-text-field>

    <v-text-field name="advertizer" label="出稿元" v-model="unit.attributes.advertizer"></v-text-field>
    <v-text-field name="image_url" label="画像URL" v-model="unit.attributes.imageURL"></v-text-field>
    <v-text-field name="target_url" label="遷移先URL" v-model="unit.attributes.targetURL"></v-text-field>

    <v-text-field name="image_count" label="表示回数" v-model="unit.attributes.image_count" disabled></v-text-field>
    <v-text-field name="open_count" label="遷移回数" v-model="unit.attributes.open_count" disabled></v-text-field>

    <v-btn outline primary @click.native.stop="send">登録</v-btn>
    <v-btn v-if="unit.id" outline error @click.native.stop="remove">削除</v-btn>
  </v-container>
</template>

<script>
import api from '../api'
import { Group, Unit } from '../models'

export default {
  props: ['id'],
  data () {
    return {
      unit: new Unit(),
      menu: {
        startDate: false,
        endDate: false,
      }
    };
  },
  created () {
    if (this.id) this.reload(this.id);
  },
  watch: {
    id (id) {
      if (id) {
        this.reload(id);
      } else {
        this.unit = new Unit();
      }
    }
  },
  methods: {
    reload (id) {
      api.get(`/units/${id}`)
      .then(res => this.unit = new Unit(res.data.data))
      .catch(err => console.error(err));
    },
    send () {
      if (this.id) {
        api.patch(`/units/${this.id}`, { data: this.unit.data })
        .then(() => this.reload(this.id))
        .then(() => this.$emit('update'))
        .catch(err => console.error(err));
      } else {
        api.post(`/units`, { data: this.unit.data })
        .then(res => new Unit(res.data.data))
        .then(unit => this.$router.push(`/unit/${unit.id}`))
        .then(() => this.$emit('update'))
        .catch(err => console.error(err));
      }
    },
    remove () {
      api.delete(`/units/${this.id}`)
      .then(() => this.$router.push(`/unit`))
      .then(() => this.$emit('update'))
      .catch(err => console.error(err));
    }
  }
}
</script>

<style lang="stylus">
</style>
