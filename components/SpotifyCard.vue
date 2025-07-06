<template>
  <main id="spotifycheck" class="hidden">
    <div id="alertDiv" role="alert" class="alert alert-soft alert-success flex flex-row py-4 px-3">
      <span id="albumname" class="tooltip" data-tip="[Album Name]">
        <img
          id="albumart"
          src="https://i.pinimg.com/originals/ac/11/aa/ac11aa2add3b0193c8769e0a17d13535.jpg"
          class="w-15 h-15 rounded-lg -my-1"
        />
      </span>
      <span class="flex flex-col -ml-[0.2rem]">
        <b id="title" class="text-md -my-1">Yükleniyor...</b>
        <span id="artist" class="text-xs">Yükleniyor...</span>
        <!-- Şarkı ilerleme çubuğu -->
        <div v-if="isPlaying" class="flex items-center -mb-1">
          <span class="text-[10px] text-right mr-1">{{ formatTime(current) }}</span>
          <progress
            class="progress h-1 w-24"
            :value="progress"
            :max="duration"
          ></progress>
          <span class="text-[10px] w-8 text-left ml-1">{{ formatTime(duration) }}</span>
        </div>
      </span>
      <spam class="grow" />
      <span class="tooltip mt-1" data-tip="Listening to Spotify">
        <Icon name="mdi:spotify" class="text-lg mr-3" />
      </span>
    </div>
  </main>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from "vue";

const start = ref(0);
const end = ref(0);
const isPlaying = ref(false);
const current = ref(0);
const duration = ref(0);
const progress = ref(0);
let timer = null;

function formatTime(ms) {
  if (!ms) return "0:00";
  const totalSeconds = Math.floor(ms / 1000);
  const min = Math.floor(totalSeconds / 60);
  const sec = totalSeconds % 60;
  return `${min}:${sec.toString().padStart(2, "0")}`;
}

function updateProgress() {
  if (!isPlaying.value) {
    current.value = 0;
    progress.value = 0;
    duration.value = 0;
    return;
  }
  const now = Date.now();
  current.value = Math.max(0, Math.min(now - start.value, end.value - start.value));
  duration.value = end.value - start.value;
  progress.value = current.value;
}

function handleSpotifyProgress(e) {
  start.value = e.detail.start;
  end.value = e.detail.end;
  isPlaying.value = e.detail.isPlaying;
  updateProgress();
  if (timer) clearInterval(timer);
  if (isPlaying.value) {
    timer = setInterval(updateProgress, 1000);
  }
}

onMounted(() => {
  window.addEventListener("spotify-progress", handleSpotifyProgress);
});

onUnmounted(() => {
  window.removeEventListener("spotify-progress", handleSpotifyProgress);
  if (timer) clearInterval(timer);
});
</script>
