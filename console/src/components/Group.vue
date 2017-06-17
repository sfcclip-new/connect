<template>
  <v-container fluid>
    <v-text-field v-if="group.id" name="id" label="グループID" v-model="group.id" readonly />
    <v-text-field name="name" label="グループ名" v-model="group.attributes.name" />

    <v-card>
      <v-subheader v-if="group.id">含まれるユニット</v-subheader>
      <v-list v-if="group.id" two-line>
        <template v-for="unit in relatedUnits" :keys="unit.id">
          <v-list-item>
            <v-divider />
            <v-list-tile @click.native.stop="removeUnit(unit)">
              <v-list-tile-content>
                <v-list-tile-title v-html="unit.id" />
                <v-list-tile-sub-title v-html="unit.attributes.name"></v-list-tile-sub-title>
              </v-list-tile-content>
            </v-list-tile>
          </v-list-item>
        </template>
      </v-list>
    </v-card>

    <v-container fluid>
      <v-layout row wrap>
        <v-flex xs11>
          <v-select v-if="group.id" :items="unrelatedUnits" v-model="selectedUnit" label="追加" dark single-line auto></v-select>
        </v-flex>
        <v-flex xs1>
          <v-btn v-if="group.id" @click.native.stop="addSelectedUnit" outline secondary light>追加</v-btn>
        </v-flex>
      </v-layout>
    </v-container>

    <v-btn outline primary @click.native.stop="send">登録</v-btn>
    <v-btn v-if="group.id" outline error @click.native.stop="remove">削除</v-btn>
  </v-container>
</template>

<script>
import api from '../api'
import { Group, Unit } from '../models'

export default {
  props: ["id", "units"],
  data () {
    return {
      group: new Group(),
      selectedUnit: null
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
        this.group = new Group();
      }
    },
  },
  computed: {
    relatedUnits () {
      return this.group.units
      .map(u => u.id)
      .map(id => this.units.filter(u => u.id == id)[0]);
    },
    unrelatedUnits () {
      const unrelatedUnits = this.units
      .filter(unit => !this.group.units.map(u => u.id).includes(unit.id));
      return unrelatedUnits;
    }
  },
  methods: {
    addSelectedUnit() {
      if (!this.selectedUnit) return;
      this.group.addUnit(this.selectedUnit);
      this.selectedUnit = null;
    },
    removeUnit (unit) {
      this.group.removeUnit(unit);
    },
    reload (id) {
      api.get(`/groups/${id}`)
      .then(res => this.group = new Group(res.data.data))
      .catch(err => console.error(err));
    },
    send () {
      if (this.id) {
        api.patch(`/groups/${this.id}`, { data: this.group.data })
        .then(() => this.reload(this.id))
        .then(() => this.$emit('update'))
        .catch(err => console.error(err));
      } else {
        api.post(`/groups`, { data: this.group.data })
        .then(res => new Group(res.data.data))
        .then(group => this.$router.push(`/group/${group.id}`))
        .then(() => this.$emit('update'))
        .catch(err => console.error(err));
      }
    },
    remove () {
      api.delete(`/groups/${this.id}`)
      .then(() => this.$router.push(`/group`))
      .then(() => this.$emit('update'))
      .catch(err => console.error(err));
    }
  }
}
</script>

<style lang="stylus">
</style>
