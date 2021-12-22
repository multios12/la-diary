<template>
  <v-select v-model="value" :items="days" />
</template>
<script lang="ts">
import Component from "vue-class-component";
import { Model, Prop, Vue } from "vue-property-decorator";

@Component
export default class DaySelect extends Vue {
  @Model("update", { type: String }) day!: string;
  @Prop({ type: String }) Month!: string;
  days: string[] = [];

  get value() {
    return this.day;
  }
  set value(v: string) {
    this.$emit("update", v);
    this.day = v;
  }

  sestDays(v: string) {
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
