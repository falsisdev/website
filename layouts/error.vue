<template>
<div style="text-align: center;align-items: center;">
  <div>
      <title>{{ title }}</title>
      <br><br>
    <h1 v-if="error.statusCode === 404">Uh! Oh! We couldn't find the page you are looking for.</h1>
    <h1 v-else>Oh! Uh! An Error Occured</h1>
    <NuxtLink to="/">Ana Sayfa</NuxtLink>
  </div>
  </div>
</template>

<script>
  export default {
    props: ['error'],
    layout: 'error',
    data() {
        return {
          title: `Error - FalsisDev`
        }
      },
      head() {
        return {
        title: this.title,
        link: [{ rel: 'icon', type: 'image/x-icon', href: 'https://nuxtjs.org/design-kit/colored-logo.svg' }]
      }
      }
  }
</script>
