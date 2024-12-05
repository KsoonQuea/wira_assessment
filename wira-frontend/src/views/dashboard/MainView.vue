<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from 'axios';

import DataTable from 'datatables.net-vue3';
// import DataTablesCore from 'datatables.net-dt';
import DataTablesCore from 'datatables.net-bs5';
import 'datatables.net-select-dt';
import 'datatables.net-responsive-dt';

// import { DataTable } from "vue3-table-lite";
// import "vue3-table-lite/dist/vue3-table-lite.es.css";

const rankings  = ref([]);
const loading   = ref(true);
const error     = ref(null);

DataTable.use(DataTablesCore);

// const columns = [
//   { data: 'character_name', title: 'Character Name' },
//   { data: 'username',       title: 'Username' },
//   { data: 'email',          title: 'Email' },
//   { data: 'reward_score',   title: 'Score' },
//   { data: 'class_id',       title: 'Class' },
// ];
//
// const options = {
//   responsive: true,
//   select: true,
// };

const fetchRankings = async () => {
  try {
    const response = await axios.get('http://localhost:8080/dashboard/ranking_table', {
      headers: {
        'Content-Type': 'application/json',
      },
    });
    rankings.value = response.data;
    loading.value = false;
  } catch (err) {
    console.error('Error fetching rankings:', err);
    error.value = 'Failed to load rankings. Please try again later.';
    loading.value = false;
  }
};

onMounted(() => {
  fetchRankings();
});


</script>

<template>
  <div class="container mx-auto w-screen p-4 justify-center">
    <h1 class="text-2xl font-bold mb-4">Ranking Table</h1>
    <div v-if="loading" class="text-center">
      <p class="text-gray-600">Loading...</p>
    </div>
    <div v-else-if="error" class="text-center">
      <p class="text-red-600">{{ error }}</p>
    </div>
    <div v-else>
      <DataTable class="display">
        <thead>
        <tr>
          <th>No. </th>
          <th>Character</th>
          <th>Username</th>
          <th>Email</th>
          <th>Score</th>
          <th>Class</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(item, index) in rankings" :key="index" :class="{'bg-gray-100': index % 2 === 0}">
          <td class="py-2 px-4 border-b text-center">{{ index + 1 }}</td>
          <td class="py-2 px-4 border-b text-center">{{ item.character_name }}</td>
          <td class="py-2 px-4 border-b text-center">{{ item.username }}</td>
          <td class="py-2 px-4 border-b text-center">{{ item.email }}</td>
          <td class="py-2 px-4 border-b text-center">{{ item.reward_score }}</td>
          <td class="py-2 px-4 border-b text-center">{{ item.class_id }}</td>
        </tr>
        </tbody>
      </DataTable>

<!--      <DataTable-->
<!--        :columns="columns"-->
<!--        :options="options"-->
<!--        ajax=""-->
<!--        class="display nowrap"-->
<!--      />-->
    </div>
  </div>
</template>

<style scoped>
@import 'datatables.net-dt';
</style>
