<template>
      <div class="m-10 text-white border-t border-zinc-600/20">
        <br/>
        <span v-if="isMobileOrTablet == false">
        <p class="float-left flex flex-row">
          © 2021 / 2022 / 2023 -
          <a
            class="duration-300 hover:text-primary font-medium flex flex-row mx-1"
            href="https://github.com/falsisdev/website"
            ><IconsGithub class="mx-2 w-5 pb-2"/> FalsisDev</a
          >
          - Some Rights
          <a
            class="duration-300 hover:text-primary font-medium ml-1"
            href="https://github.com/falsisdev/website/blob/third/LICENSE"
            >Reserved</a
          >.
        </p>
        <p class="text-white float-right">
          Designed with
          <span><font-awesome-icon icon="fa-solid fa-heart" class="mx-1 text-primary"/></span> by
          <a class="duration-300 hover:text-primary font-medium" href="/#top"
            >Falsis</a
          >
        </p>
        </span>
        <span v-else>
          <p class="float-left">
            <span class="flex flex-row">
          © 2021 / 2022 / 2023 -
          <a
            class="duration-300 hover:text-primary font-medium flex flex-row mx-1"
            href="https://github.com/falsisdev/website"
            ><IconsGithub class="mx-2 w-5 pb-2"/> FalsisDev</a
          >
          </span>
          Some Rights
          <a
            class="duration-300 hover:text-primary font-medium ml-1"
            href="https://github.com/falsisdev/website/blob/third/LICENSE"
            >Reserved</a
          >.
        </p>
        <p class="text-white">
          Designed with
          <span><font-awesome-icon icon="fa-solid fa-heart" class="mx-1 text-primary"/></span> by
          <a class="duration-300 hover:text-primary font-medium" href="/#top"
            >Falsis</a
          >
        </p>
        </span>
      </div>
</template>
<script setup>
const { isMobileOrTablet } = useDevice();
</script>