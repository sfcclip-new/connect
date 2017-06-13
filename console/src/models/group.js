import Unit from './unit'

export default class Group {

  static list(object) {
    return Array.isArray(object.data) ? object.data.map(d => new Group(d)) : [];
  }

  constructor(data) {
    this.rawData = {};
    if (data && data.type && data.type == 'groups') {
      this.rawData = data;
    }
  }

  get id () {
    return this.rawData.id || 0;
  }

  get attributes () {
    this.rawData.attributes = this.rawData.attributes || {};
    return this.rawData.attributes;
  }

  get relationships () {
    this.rawData.relationships = this.rawData.relationships || {};
    return this.rawData.relationships || {};
  }

  get data() {
    const data = this.rawData;
    data.type = "groups";
    data.id = String(this.id);
    console.log(data);
    return data;
  }

  get units() {
    if (!this.relationships.units) return;
    return this.relationships.units.data.map(d => new Unit(d))
  }

  addUnit(unit) {
    this.relationships.units.data.push(unit.data);
  }

  removeUnit(unit) {
    this.relationships.units.data = this.relationships.units.data.filter(data => {
      console.log(data);
      console.log(unit);
      return (new Unit(data)).id != unit.id;
    });
  }

}
