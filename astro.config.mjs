import { defineConfig } from 'astro/config';
import svelte from '@astrojs/svelte';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
  site: 'https://falsisdev.github.io',
  base: '/website',
  integrations: [svelte()],
  vite: {
    plugins: [tailwindcss()],
  },
});
