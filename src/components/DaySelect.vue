<template>
  <v-select v-model="day" :items="days" />
</template>
<script lang="ts">
import Component from "vue-class-component";
import { Model, Prop, Vue, Watch } from "vue-property-decorator";

@Component
export default class DaySelect extends Vue {
  @Model("update", { type: String }) day!: string;
  @Prop({ type: String }) Month!: string;
  days: string[] = [];

  beforeUpdate() {
    console.log("fired:beforeUpdate");
    this.setDays(this.day);
  }
  @Watch("day")
  dayChanged(value: string, oldValue: string) {
    this.$emit("update", value);
    this.day = value;
    this.setDays(value);
  }

  public setDays(v: string) {
    if (this.days.length > 0) {
      const old = this.days[0];
      if (old.substring(0, 7) == v.substring(0, 7)) {
        return;
      }
    }

    const y = Number(v.substring(0, 4));
    const m = Number(v.substring(5, 2));
    const d = new Date(y, m, 0);
    let l: string[] = new Array(d.getDay());
    for (var i = 0; i < d.getDay(); i++) {
      l[i] = `${y}-${m}-${i + 1}`;
    }
    this.days = l;
  }
}
</script>
