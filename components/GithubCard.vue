<template>
    <div class="max-w-sm p-6 border rounded-lg bg-base-200 border-gray-700 m-2 w-full">
        <a class="flex flex-row">
            <h5 class="mb-2 text-md font-bold tracking-tight text-white flex flex-row"><IconsRepo class="mr-2 mt-1"/> {{ props.reponame }} </h5>
            <span class="grow"></span>
            <IconsGithub class="w-5 h-5"/>
        </a>
        <p class="mb-3 font-normal text-gray-700 dark:text-gray-400 text-sm mt-2 h-12"> {{ props.repodesc }} </p>
        <span class="flex flex-col-reverse">
        <span class="flex flex-row">
        <p class="font-normal text-gray-400 text-xs flex flex-row mt-1.5">
            <span class="flex items-center"><font-awesome-icon icon="fa-solid fa-star" class="mt-[2px] mr-[1px] text-xs" /></span> <span class="ml-1 mr-2">{{ props.star }}</span>
            <span class="flex items-center mx-1"><span :class="`flex w-2 h-2 mt-[2px] mr-[4px] text-xs ${props.lang == 'CSS' ? 'bg-violet-900' : props.lang == 'Vue' ? 'bg-emerald-400' : props.lang == 'JavaScript' ? 'bg-yellow-200' : props.lang == 'Go' ? 'bg-sky-400' : props.lang == 'Dart' ? 'bg-teal-200' : props.lang == 'HTML' ? 'bg-orange-500' : ''} rounded-full mr-1 flex-shrink-0`"></span>{{ props.lang }}</span>
        </p>
        <span class="grow"></span>
        <a :href="props.url"><font-awesome-icon icon="fa-solid fa-arrow-up-right-from-square"/></a>
        </span>
        </span>
    </div>
</template>
<script setup>
const props = defineProps({
    reponame: String,
    repodesc: String,
    lang: String,
    star: Number,
    url: String
})
</script>