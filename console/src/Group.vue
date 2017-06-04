<template>
  <v-container fluid>
    <v-text-field name="input-1" label="グループID" id="testing" :value="group ? group.ID : id" disabled></v-text-field>
    <v-card>
      <v-card-text>
        <v-subheader>含まれるユニット</v-subheader>
        <v-checkbox dark v-for="unit in units" :key="unit.ID" :label="unit.ID" />
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
import axios from 'axios';

export default {
  data () {
    return {
      id: "",
      group: undefined,
      units: []
    };
  },
  created () {
    this.id = this.$route.params.id;

    axios.get('/unit')
    .then(res => {
      console.log(res);
      this.units = res.data || [];
    })
    .catch(err => console.error(err));
  },
  beforeRouteUpdate (to, from, next) {
    this.id = to.params.id;
    next();
  },
  watch: {
    id: id => {
      axios.get(`/group/${id}`)
      .then(res => {
        console.log(res);
        this.group = res.data;
      })
      .catch(err => console.error(err));
    },
  }
}
</script>

<style lang="stylus">
</style>
