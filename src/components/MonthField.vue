<template>
  <v-menu
    v-model="menu"
    :close-on-content-click="false"
    :nudge-right="40"
    transition="scale-transition"
    offset-y
    min-width="auto"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-text-field
        v-model="value"
        prepend-icon="mdi-calendar-month"
        readonly
        v-bind="attrs"
        v-on="on"
      ></v-text-field>
    </template>
    <v-date-picker
      v-model="value"
      type="month"
      @input="menu = false"
    ></v-date-picker>
  </v-menu>
</template>
<script lang="ts">
import Component from "vue-class-component";
import { Model, Vue } from "vue-property-decorator";

@Component
export default class MonthField extends Vue {
  @Model("update", { type: String }) readonly month!: string;
  private menu = false;

  get value() {
    return this.month;
  }
  set value(v: string) {
    this.$emit("update", v);
  }
}
</script>
