# falsisdev/website

Personal portfolio website built with modern web technologies — statically generated, zero-config deployment.

## Tech Stack

| Technology | Purpose |
|---|---|
| [Astro](https://astro.build) | Framework — routing, static site generation, zero-JS by default |
| [Svelte 5](https://svelte.dev) | Interactive UI components with runes (`$props`, `$state`, `$effect`) |
| [Tailwind CSS v4](https://tailwindcss.com) | Utility-first styling via Vite plugin |

## Architecture

The site follows Astro's [Islands Architecture](https://docs.astro.build/en/concepts/islands/). Pages are rendered as static HTML at build time. Interactive sections are hydrated selectively using `client:*` directives — only the JavaScript that's needed gets shipped to the browser.

```
src/
├── components/        # Svelte interactive islands
│   ├── Hero.svelte          # Animated hero section
│   ├── ProjectCard.svelte   # Project cards grid
│   └── SkillBar.svelte      # Scroll-triggered skill bars
├── layouts/
│   └── Layout.astro         # Base HTML layout (dark theme)
├── pages/
│   └── index.astro          # Homepage — assembles all components
└── styles/
    └── global.css           # Tailwind v4 theme configuration
```

For a detailed architectural overview, see [ARCHITECTURE.md](./ARCHITECTURE.md).

## Getting Started

**Prerequisites:** Node.js 18+

```bash
# Install dependencies
npm install

# Start development server
npm run dev
```

The dev server runs at `http://localhost:4321/website`.

## Build & Deploy

```bash
# Generate static output
npm run build

# Preview the production build locally
npm run preview
```

Static files are output to `dist/` and deployed to [GitHub Pages](https://falsisdev.github.io/website).

## License

MIT
