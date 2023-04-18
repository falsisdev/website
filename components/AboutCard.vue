<template>
    <div class="flex flex-row">
        <div class="relative mx-10 flex items-center space-x-4">
            <span>
                <img id="avatar" class="w-12 h-12 rounded-full border border-gray-800" src="-" alt="Yükleniyor...">
                <span id="status1" class="bottom-0 left-8 absolute w-4 h-4 rounded-full animate-ping"></span>
                <span id="status2" class="bottom-0 left-8 absolute w-4 h-4 border border-black border-opacity-50 rounded-full"></span>
            </span>
            <div class="font-medium dark:text-white">
                <div><span id="uname">Yükleniyor...</span><span id="discrim" class="text-sm text-gray-500 dark:text-gray-400">#0000</span></div>
                <div id="spotify" class="text-sm text-gray-500 dark:text-gray-400"></div>
            </div>
        </div>
        <div class="grow"></div>
        <div id="spotifywarning" class="flex p-4 mr-10 text-sm text-green-800 border border-green-300 rounded-lg bg-green-50 dark:bg-gray-800 dark:text-green-400 dark:border-green-800" role="alert">
            <IconsSpotify class="w-[18px] h-[18px] mx-1"/>
            <div>
                Listening <span id="song" class="font-medium"></span> from <span id="artist" class="font-medium"></span>.
            </div>
        </div>
    </div>
    <article class="prose max-w-none mx-10 my-5">
        <h1 class="mb-4 text-5xl font-extrabold leading-none tracking-tight text-white"><span class="text-primary">Hello,</span> I'm Falsis.</h1>
        <span class="mb-6 text-lg font-normal text-gray-400">I'm from <a class="no-underline text-primary hover:opacity-80" href="https://en.wikipedia.org/wiki/%C4%B0zmir">İzmir</a>. I love <a class="no-underline text-gray-400 hover:opacity-80" href="https://en.wikipedia.org/wiki/Front-end_web_development">Front-End Web Development</a> and <a class="no-underline text-primary hover:opacity-80" href="https://en.wikipedia.org/wiki/User_interface_design">designing user interfaces</a>. Now, I'm a High-School student and I love to improve myself.</span>
    </article>
        <span></span>
</template>
<style scoped>
* {
    font-family: 'Arial', sans-serif;
}
.title {
    font-family: 'Arial black', sans-serif;
}
</style>
