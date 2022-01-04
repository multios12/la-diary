<template>
  <v-app>
    <v-app-bar app dark> diary </v-app-bar>

    <v-main>
      <v-container>
        <v-card>
          <v-card-title><month-field v-model="month" /></v-card-title>
          <v-card-text>
            <v-row>
              <v-col><v-text-field v-model="day" label="date" /></v-col>
              <v-col><v-text-field v-model="memo" label="outline" /></v-col>
              <v-col><v-btn @click="send">send</v-btn></v-col>
            </v-row>
          </v-card-text>
        </v-card>
        <v-card>
          <v-card-text>
            <v-list>
              <v-list-item v-for="item in items" :key="item.day">
                <v-list-item-content>
                  <v-list-item-title v-text="item.day" />
                  <v-list-item-content v-text="item.memo" />
                </v-list-item-content>
              </v-list-item>
            </v-list>
          </v-card-text>
        </v-card>
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import MonthField from "./components/MonthField.vue";
import DaySelect from "./components/DaySelect.vue";
import axios from "axios";
import { Watch } from "vue-property-decorator";

@Component({
  components: { MonthField, DaySelect },
  filters: {
    dateFormat(value: string) {
      return value.substring(5);
    },
  },
})
export default class App extends Vue {
  private month = new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
    .toISOString()
    .substr(0, 7);
  private day = "";
  private memo = "";
  private items = [];

  mounted() {
    const d = new Date();
    this.day = d.getFullYear() + "-";
    this.day += ("00" + (d.getMonth() + 1)).slice(-2) + "-";
    this.day += ("00" + (d.getDate() + 1)).slice(-2);
    this.showLines();
  }

  async send() {
    let url = this.day.replaceAll("-", "/");
    url = `./api/${url}`;

    await axios.post(url, { Memo: this.memo });
    this.showLines();
  }

  @Watch("month")
  showLines() {
    let url = this.month.replaceAll("-", "/");
    url = `./api/${url}`;

    axios.get(url).then((r) => {
      this.items = r.data;
    });
  }
}
</script>
