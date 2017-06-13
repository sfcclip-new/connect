export default class Unit {

  static list(object) {
    return Array.isArray(object.data) ? object.data.map(d => new Unit(d)) : [];
  }

  constructor(data) {
    this.rawData = {};
    if (data && data.type && data.type == 'units') {
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

  get data() {
    const data = this.rawData;
    data.type = "units";
    data.id = String(this.id);
    return data;
  }

  get text() {
    return `${this.id} ${this.attributes.name}`;
  }

}
