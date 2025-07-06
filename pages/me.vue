<template>
  <article class=" prose max-w-none">
    <span class="font-semibold ml-1 text-sm">Let's talk</span>
    <h1>About <span class="text-primary">Me</span></h1>
    <span class="text-lg">
      Hi, I'm Falsis. But that's not my name for sure. My name is Kerem. Nice to
      meet you. Currently, I'm
      <span
        class="text-primary tooltip tooltip-bottom"
        data-tip="Dec 30, 2007"
        >{{
          ((Date.now() - 1199002814000) / 1000 / 60 / 60 / 24 / 365.25).toFixed(
            2
          )
        }}</span
      >
      years old. So this means now I'm
      {{
        Date.now() <= 1749900506000
          ? "an 11th grade high-school student"
          : Date.now() <= 1781436506000
          ? "a 12th grade high school student"
          : "a university student"
      }}.
    </span>
    <div class="pr-4">
      <p class="text-md">
        I guess I should start with programming first. I am a Front-End
        developer who develops full-stack projects because sometimes uses
        back-end, but is far from being a full-stack developer due to
        shortcomings in back-end. I am developing web
        <NuxtLink to="/projects" class="text-primary">projects</NuxtLink>
        but I want to focus on Android application development in the future.
        That's why learning Kotlin is one of my goals but I am postponing it for
        later because I will have difficulty using Android Studio on my current
        computer. In short, I am currently a Front-End Web Developer. You can
        see the technologies i use
        <NuxtLink to="#techstack" class="text-primary">here</NuxtLink>.
        <br />
        <br />
        Academically I'm interested in Maths, Physics, History, Geography and
        Linguistics. Other than that, I love playing Football, Basketbool,
        Billiard, Guitar and something else.
        <br />
        I really love listening music. I love music genres such as Rock, Hard
        Rock, Turkish Rock, Metal, Alternative Metal, Numetal etc.
        <br />
        I use Instagram, YouTube, WhatsApp, Discord,
        <NuxtLink
          class="text-primary"
          to="https://open.spotify.com/user/315l5ib3a4fd2obidm76lipspxji?si=a4b244df74e44569"
          >Spotify</NuxtLink
        >, Pinterest, Reddit, Netflix,
        <NuxtLink class="text-primary" to="https://github.com/falsisdev"
          >GitHub</NuxtLink
        >, Wikipedia and Tiktok as social media platforms.
        <br />
        My main language is Turkish, my English level is about B2, I started
        learning German but it's still A1 or A2 for me. I can understand and
        speak A1 level Japanese too, besides I can read Kana (Hiragana &
        Katakana). Because of the language family of my main language. I can
        understand over 90% Oghuz languages such as Gagauz, Azerbaijani or
        Turkmen. I'm a Turkish muslim. So I'm familiar with Arabic alphabet too
        but I can't read it fr.
        <br />
        When it comes to talk about the series and movies, I like things that
        surprise me when I watch them, like Plot Twists. In this sense, I can
        watch any series or anime set in any period. I even watch sitcoms, even
        though they interest me less than others. I watched a lot of anime. You
        can check it out on
        <NuxtLink
          class="text-primary"
          to="https://myanimelist.net/profile/falsis"
          >MyAnimeList</NuxtLink
        >
        or
        <NuxtLink class="text-primary" to="https://anilist.co/user/Falsis/"
          >AniList</NuxtLink
        >. I also enjoy reading manga and occasionally light nove. You can check
        them out on those accounts too.
        <br />
        Regarding sports, I like to play and watch football. I am a Fenerbahçe
        fan. I don't support a team in the city I live in. I don't support any
        team in Europe other than Real Madrid and maybe Benfica. My card, which
        I need to enter the match as a spectator, is stamped Göztepe and I also
        have an expired Altınordu card. But I don't support them fr. I also like
        to play and watch basketball. I like the Euroleague more than the NBA
        because I prefer tactical games to show. I'm a huge Magnus Carlsen fan
        and I like to play chess. And Finally, I play billiards sometimes.
      </p>
    </div>
  </article>
</template>
