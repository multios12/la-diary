<template>
  <v-app>
    <v-app-bar app dark>
      <v-app-bar-title>diary</v-app-bar-title>
    </v-app-bar>

    <v-main>
      <v-container>
        <v-card>
          <v-card-title><month-field v-model="month" /></v-card-title>
          <v-card-text>
            <v-row>
              <v-col><day-select v-model="day" /></v-col>
              <v-col>
                <v-text-field label="outline"></v-text-field>
              </v-col>
              <v-col>
                <v-btn>send</v-btn>
              </v-col>
            </v-row>
            <v-list></v-list>
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

@Component({ components: { MonthField, DaySelect } })
export default class App extends Vue {
  private month = new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
    .toISOString()
    .substr(0, 7);
  private day = new Date();
}
</script>
