<template>
    <div class="dui-card lg:dui-card-side w-auto h-auto bg-base-100 p-[10px] rounded-lg bg-base-200 m-5">
        <figure class="rounded-lg"><img class="w-48 h-72 rounded shadow-md" :src="props.cover"/></figure>
        <div class="dui-card-body">
            <h2 class="dui-card-title">{{ props.name }} ({{ props.puan }}/10)</h2>
            <p>
                <span v-for="item of props.genres" :key="item" class="dui-badge dui-badge-lg m-1">{{ item.name }}</span>
                <br/>
                {{ props.description.length <= 200 ? props.description : props.description.substring(0, 197) + '...'}}
            </p>
            <div class="dui-card-actions justify-end">
            <button @click="myanimelist()" class="dui-btn dui-btn my-2">MyAnimeList</button>
            </div>
        </div>
    </div>
</template>
<script setup>
function myanimelist() {
    window.open(`https://myanimelist.net/${props.url}`, '_blank')
}
const props = defineProps({
    name: String,
    cover: String,
    description: String,
    url: String,
    genres: Array,
    puan: Number,
})
</script>