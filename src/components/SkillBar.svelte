<script>
  let {
    skills = [
      { name: "Go", level: 85 },
      { name: "TypeScript / JavaScript", level: 90 },
      { name: "Vue / Nuxt", level: 85 },
      { name: "Svelte / Astro", level: 65 },
      { name: "Tailwind CSS", level: 90 },
    ]
  } = $props();

  let visible = $state(false);

  $effect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          visible = true;
          observer.disconnect();
        }
      },
      { threshold: 0.3 }
    );

    const el = document.getElementById("skills-section");
    if (el) observer.observe(el);

    return () => observer.disconnect();
  });
</script>

<section id="skills-section" class="px-6 py-24">
  <div class="mx-auto max-w-2xl">
    <h2 class="mb-2 text-sm font-medium tracking-widest text-brand-light uppercase">Yetenekler</h2>
    <p class="mb-12 text-3xl font-bold text-white">Teknoloji stack'im</p>

    <div class="space-y-6">
      {#each skills as skill, i}
        <div>
          <div class="mb-2 flex items-center justify-between">
            <span class="text-sm font-medium text-zinc-300">{skill.name}</span>
            <span class="text-sm text-zinc-500">{skill.level}%</span>
          </div>
          <div class="h-2 overflow-hidden rounded-full bg-zinc-800">
            <div
              class="h-full rounded-full bg-gradient-to-r from-brand to-brand-light transition-all duration-1000 ease-out"
              style="width: {visible ? skill.level : 0}%; transition-delay: {i * 150}ms"
            ></div>
          </div>
        </div>
      {/each}
    </div>
  </div>
</section>
