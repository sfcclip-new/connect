<template>
  <v-container fluid>
    <v-data-table v-bind:headers="headers" :items="records" default-sort="id">
      <template slot="items" scope="props">
        <td>{{ props.item.id }}</td>
        <td>{{ props.item.created_at }}</td>
        <td>{{ props.item.attributes.unit_id }}</td>
        <td>{{ props.item.access_type }}</td>
        <td>{{ props.item.attributes.host }}</td>
        <td>{{ props.item.attributes.remote_addr }}</td>
        <td>{{ props.item.attributes.request_uri }}</td>
        <td>{{ props.item.attributes.user_agent }}</td>
        <td>{{ props.item.attributes.referer }}</td>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
import api from '../api'

export default {
  data () {
    return {
      records: [],
      headers: [
        { text: 'ID', value: 'id' },
        { text: 'Date and Time', value: 'created_at' },
        { text: 'Unit ID', value: 'attributes.unit_id', sortable: false },
        { text: 'Access Type', value: 'access_type', sortable: false },
        { text: 'Host', value: 'attributes.host', sortable: false },
        { text: 'Remote Address', value: 'attributes.remote_addr', sortable: false },
        { text: 'Request URI', value: 'attributes.request_uri', sortable: false },
        { text: 'User Agent', value: 'attributes.user_agent', sortable: false },
        { text: 'Referer', value: 'attributes.referer', sortable: false },
      ],
    };
  },
  created () {
    this.reload();
  },
  methods: {
    reload () {
      api.get(`/records`)
      .then(res => {
        const records = res.data.data;
        records.map(r => {
          r.created_at = r.attributes.created_at
          switch (r.attributes.access_type) {
            case 0: r.access_type = "Image"; break;
            case 1: r.access_type = "Open"; break;
          }
        })
        this.records = records
      })
      .catch(err => console.error(err));
    }
  }
}
</script>

<style lang="stylus">
</style>
