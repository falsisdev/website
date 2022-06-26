export default {
  components: true,
  generate: {
  dir: 'dist',
    subFolders: false
   },
  server: {
    host: 'localhost',
    port: 3000
  },
  tailwindcss: {
    viewer: true
  },
  loading: {
   color: "#57F287",
   failedColor: "#57F287",
   height: "3px",
   throttle: 0
  },
    script: [
        {
          src: 'https://cdnjs.cloudflare.com/ajax/libs/jquery/3.1.1/jquery.min.js'
        }
      ],
      link: [
        {
          rel: 'stylesheet',
          href: 'https://fonts.googleapis.com/css?family=Roboto&display=swap'
        }
      ],
    meta: [
        { 
          charset: 'utf-8' 
        },
        { 
          name: 'viewport', content: 'width=device-width, initial-scale=1' 
        },
        {
          author: 'FalsisDev',
          name: 'FalsisDev',
          color: "#57F287",
          content: 'FalsisDev | A Front-End Web Developer.'
        }
    ],
    modules: ['@nuxtjs/axios', '@nuxtjs/tailwindcss', '@nuxt/components'],
    target: 'static',
    css: ['~/assets/main.css'],
    devServerHandlers: [],
  }
