<template>
    <div class="dui-card dui-card-side bg-base-100 m-3">
        <figure class="rounded-lg"><img class="w-72 h-64" :src="props.cover"/></figure>
        <div class="dui-card-body">
            <h2 class="dui-card-title">{{ props.name }} ({{ props.year }})</h2>
            <p>
                <span v-for="item of props.genres" :key="item" class="dui-badge dui-badge-lg m-1">{{ item.display_name }}</span>
                <br/>
                {{ props.description.length <= 200 ? props.description : props.description.substring(0, 197) + '...'}}
            </p>
            <div class="dui-card-actions justify-end">
            <button @click="animecix()" class="dui-btn dui-btn my-2">Animecix</button>
            </div>
        </div>
    </div>
</template>
<script setup>
function animecix() {
    window.open(`https://animecix.net/titles/${props.id}`, '_blank')
}
const props = defineProps({
    name: String,
    cover: String,
    description: String,
    year: Number,
    genres: Array,
    id: Number,
})
</script>