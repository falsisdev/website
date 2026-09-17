<script>
  let { projects = [] } = $props();

  /**
   * @type {{ title: string, description: string, tags: string[], url?: string }[]}
   */
  const defaultProjects = [
    {
      title: "Mangile",
      description: "Modern açık kaynak manga, manhwa ve light novel okuma platformu. Türkçe taramalar, gelişmiş okuyucu UX'i ve Sanity CMS entegrasyonu.",
      tags: ["Nuxt 4", "Vue", "TypeScript", "Tailwind CSS"],
      url: "https://github.com/falsisdev/mangile"
    },
    {
      title: "Anthology",
      description: "Stremio ve Nuvio için Türkçe film, dizi, anime ve canlı TV kataloğu. 42 aktif plugin, 35 video scraper.",
      tags: ["JavaScript", "QuickJS", "Stremio"],
      url: "https://github.com/falsisdev/anthology"
    },
    {
      title: "Portfolio Website",
      description: "Astro + Svelte + Tailwind CSS ile oluşturulmuş kişisel portfolio sitesi. GitHub Pages üzerinde statik hosting.",
      tags: ["Astro", "Svelte", "Tailwind CSS"],
      url: "https://github.com/falsisdev/website"
    }
  ];

  const items = projects.length > 0 ? projects : defaultProjects;
</script>

<section id="projects" class="px-6 py-24">
  <div class="mx-auto max-w-4xl">
    <h2 class="mb-2 text-sm font-medium tracking-widest text-brand-light uppercase">Projeler</h2>
    <p class="mb-12 text-3xl font-bold text-white">Üzerinde çalıştığım şeyler</p>

    <div class="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
      {#each items as project}
        <a
          href={project.url || "#"}
          target={project.url && project.url !== "#" ? "_blank" : undefined}
          rel={project.url && project.url !== "#" ? "noopener noreferrer" : undefined}
          class="group rounded-xl border border-border bg-surface-alt p-6 transition-all hover:border-brand/40 hover:-translate-y-1 hover:shadow-lg hover:shadow-brand/5"
        >
          <h3 class="mb-2 text-lg font-semibold text-white group-hover:text-brand-light transition-colors">
            {project.title}
          </h3>
          <p class="mb-4 text-sm leading-relaxed text-zinc-400">{project.description}</p>
          <div class="flex flex-wrap gap-2">
            {#each project.tags as tag}
              <span class="rounded-full bg-zinc-800 px-2.5 py-0.5 text-xs font-medium text-zinc-300">
                {tag}
              </span>
            {/each}
          </div>
        </a>
      {/each}
    </div>
  </div>
</section>
