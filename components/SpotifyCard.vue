<template>
  <main id="spotifycheck" class="hidden">
    <div id="alertDiv" role="alert" class="alert alert-soft alert-success flex flex-row py-4 px-3">
      <span id="albumname" class="tooltip" data-tip="[Album Name]">
        <img
          id="albumart"
          src="https://i.pinimg.com/originals/ac/11/aa/ac11aa2add3b0193c8769e0a17d13535.jpg"
          class="w-12 h-12 rounded-lg -my-2"
        />
      </span>
      <span class="flex flex-col -ml-[0.2rem]">
        <b id="title" class="text-md -my-1">Yükleniyor...</b>
        <span id="artist" class="text-xs">Yükleniyor...</span>
      </span>
      <spam class="grow" />
      <span class="tooltip mt-1" data-tip="Listening to Spotify">
        <Icon name="mdi:spotify" class="text-lg mr-2" />
      </span>
    </div>
  </main>
</template>
