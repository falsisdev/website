<template>
  <main class="mt-20">
    <div class="hero">
      <div class="hero-content flex-col lg:flex-row-reverse">
        <img
          id="avatar"
          class="max-w-sm rounded-lg ml-20 hover:animate-pulse w-[256px] h-[256px]"
        />
        <div class="max-w-lg">
          <h1
            class="-ml-4 text-4xl font-extrabold leading-none tracking-tight md:text-5xl lg:text-6xl"
          >
            👋🏻Hi,
            <span
              class="underline underline-offset-3 decoration-8 decoration-primary"
              >Falsis</span
            >
            here!
          </h1>
          <p class="py-4">
            Hello there! I'm Falsis and only
            <span
              class="text-primary tooltip tooltip-bottom"
              data-tip="Dec 30, 2007"
              >{{
                (
                  (Date.now() - 1199002814000) /
                  1000 /
                  60 /
                  60 /
                  24 /
                  365.25
                ).toFixed(2)
              }}</span
            >
            years old!
          </p>
          <SpotifyCard />
        </div>
      </div>
    </div>
  </main>
</template>
