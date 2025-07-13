<script setup>
const theData = ref(null);
var query = `
query ($type: MediaType!, $userId: Int!, $sort: [MediaListSort]) {
  MediaListCollection(type: $type, userId: $userId, sort: $sort) {
    lists {
      name
      entries {
        id
        media {
          id
          title {
            romaji
          }
          idMal
          genres
          description
          coverImage {
            extraLarge
          }
          status
          meanScore
          format
        }
        startedAt {
          year
          month
          day
        }
        completedAt {
          year
          month
          day
        }
        score
      }
      status
    }
  }
}
`;

var variables = {
  userId: 5868307,
  type: "ANIME",
  sort: "SCORE_DESC",
};

var url = "https://graphql.anilist.co",
  options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
    body: JSON.stringify({
      query: query,
      variables: variables,
    }),
  };

fetch(url, options)
  .then(handleResponse)
  .then((data) => (theData.value = data))
  .catch(handleError);

function handleResponse(response) {
  return response.json().then(function (json) {
    return response.ok ? json : Promise.reject(json);
  });
}

function handleError(error) {
  alert("Error, check console");
  console.error(error);
}

import { ref, onMounted } from "vue";

const contextMenu = ref({
  visible: false,
  x: 0,
  y: 0,
  mediaId: null,
});

function openContextMenu(e, mediaId) {
  contextMenu.value.visible = true;
  contextMenu.value.x = e.clientX;
  contextMenu.value.y = e.clientY;
  contextMenu.value.mediaId = mediaId;
  document.addEventListener("click", closeContextMenu);
}

function closeContextMenu() {
  contextMenu.value.visible = false;
  document.removeEventListener("click", closeContextMenu);
}

function goToAnilistPage(mediaId) {
  window.open(`https://anilist.co/anime/${mediaId}`, "_blank");
  closeContextMenu();
}

function formatDate(y, m, d) {
  if (!y || !m || !d) return "";
  const date = new Date(y, m - 1, d);
  return date.toLocaleDateString("en-US", { year: "numeric", month: "short", day: "numeric" });
}
</script>
<template>
  <main>
    <article class="prose max-w-none mb-10">
      <span class="font-semibold text-sm ml-1">Check out my</span>
      <h1>Ani<span class="text-primary">List</span></h1>
      <p>
        This page visualizes my AniList data. It shows the animes i've watched
        so far. But this data doesn't have 100% accuracy about my data because
        of the missed series in the Anilist data. Right click for more info. Also you can visit my
        <a href="https://myanimelist.net/profile/falsis"
          ><Icon name="simple-icons:myanimelist" class="w-5 h-5 -mb-1 mx-1"
        /></a>
        profile.
        <br/>
        The order is based on score and it starts with Completed TV and then it continues with Completed Movies, Completed OVA, Completed Special, Dropped, Planning, On Hold and Watching.
        <br/>
        The start and end dates are indicated on the badges. If there are no start and end dates, I don't know about them. This means that the date I watched the series was quite old and I have no info of it.
      </p>
    </article>
    <article class="prose max-w-none mb-10">
      <h1>Animes</h1>
    </article>
    <div v-if="theData">
      <div
        v-for="list of theData.data.MediaListCollection.lists"
        :key="list"
        class="flex flex-row flex-wrap"
      >
        <div
          v-for="item of list.entries"
          :key="item"
          class="card w-72 mx-2 basis-1/4 mb-2 border border-base-200 rounded-2xl"
          @contextmenu.prevent="openContextMenu($event, item.media.id)"
        >
          <figure class="w-full h-64">
            <img :src="item.media.coverImage.extraLarge" alt="Cover" />
          </figure>
          <div class="card-body w-72 h-72">
            <span class="flex flex-col relative">
              <span class="whitespace-nowrap overflow-hidden">
                <span
                  :class="`inline-block ${
                    item.media.title['romaji'].length >= 20
                      ? 'animate-marquee'
                      : ''
                  }`"
                >
                  <h2 class="card-title">{{ item.media.title.romaji }}</h2>
                </span>
              </span>
              <span v-if="item.media.format" class="text-xs opacity-50">
                {{ item.media.format }}
              </span>
            </span>
            <span class="overflow-y-auto max-w-64 text-xs">
             <span class="flex flex-row no-wrap overflow-x-auto mb-2">
              <div
                v-for="genre of item.media.genres"
                :key="genre"
                class="badge badge-outline mr-1 badge-xs"
              >
                {{ genre }}
              </div>
            </span>
              <span v-html="item.media.description" />
            </span>
            <div class="card-actions flex flex-col flex-col-reverse items-end">
              <div class="flex flex-row w-full gap-1">
                <div class="badge badge-soft badge-warning badge-sm tooltip" data-tip="Score">
                  <Icon name="material-symbols:star" class="w-5 h-5 mr-1" />
                  {{ item.score == 0 ? "N/A" : item.score }}
                </div>
                <div class="badge badge-soft badge-sm tooltip" data-tip="Status">
                  <Icon name="material-symbols:movie-info" class="w-5 h-5 mr-1" />
                  {{
                    list.status
                      .replace("COMPLETED", "Completed")
                      .replace("PLANNING", "Planning")
                      .replace("DROPPED", "Dropped")
                      .replace("CURRENT", "Watching")
                      .replace("PAUSED", "On Hold")
                  }}
                </div>
              </div>
              <div class="flex flex-row w-full gap-1 mt-1">
                <div
                  v-if="item.startedAt?.year"
                  class="badge badge-soft badge-success badge-sm tooltip"
                  data-tip="Start Date"
                >
                  <Icon name="mdi:calendar-start" class="w-4 h-4 mr-1" />
                  {{ formatDate(item.startedAt.year, item.startedAt.month, item.startedAt.day) }}
                </div>
                <div
                  v-if="item.completedAt?.year"
                  class="badge badge-soft badge-error badge-sm tooltip"
                  data-tip="End Date"
                >
                  <Icon name="mdi:calendar-end" class="w-4 h-4 mr-1" />
                  {{ formatDate(item.completedAt.year, item.completedAt.month, item.completedAt.day) }}
                </div>
              </div>
              <span class="grow" />
              <!--<button class="btn btn-ghost btn-md">
            <Icon
              name="material-symbols:arrow-outward"
              class="w-5 h-5 mt-[1px]"
            />
          </button>-->
            </div>
          </div>
        </div>
      </div>
      <!-- Context Menu -->
      <div
        v-if="contextMenu.visible"
        :style="{ top: contextMenu.y + 'px', left: contextMenu.x + 'px' }"
        class="fixed z-50 bg-base-200 border border-base-300 rounded shadow-lg"
        @click="goToAnilistPage(contextMenu.mediaId)"
        @contextmenu.prevent
        style="min-width: 140px; cursor: pointer;"
      >
        <div class="px-4 py-2 hover:bg-base-300"><Icon name="simple-icons:anilist" class="text-sm mr-1" /> Details</div>
      </div>
    </div>
    <div v-else>
      <Icon name="mingcute:loading-line" class="animate-spin w-full h-32" />
    </div>
  </main>
</template>
<style>
@keyframes marquee {
  0% {
    transform: translateX(50%);
  }
  100% {
    transform: translateX(-100%);
  }
}

.animate-marquee {
  animation: marquee 10s linear infinite;
  animation-delay: 0.5s;
}
</style>
