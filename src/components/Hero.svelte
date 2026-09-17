<script>
  let { name = "Falsis", role = "Developer & Student", bio = "ECE undergrad @ Yıldız Technical University. Modern web teknolojileri ve Go ile projeler geliştiren bir developer." } = $props();

  let visible = $state(false);

  $effect(() => {
    visible = true;
  });
</script>

<section class="relative flex min-h-screen flex-col items-center justify-center px-6 text-center">
  <!-- Gradient orb background -->
  <div class="pointer-events-none absolute inset-0 overflow-hidden">
    <div class="absolute left-1/2 top-1/3 h-96 w-96 -translate-x-1/2 -translate-y-1/2 rounded-full bg-brand/20 blur-[120px]"></div>
    <div class="absolute right-1/4 top-2/3 h-64 w-64 rounded-full bg-purple-500/10 blur-[100px]"></div>
  </div>

  <div class="relative z-10 max-w-2xl transition-all duration-1000 {visible ? 'translate-y-0 opacity-100' : 'translate-y-8 opacity-0'}">
    <!-- Avatar placeholder -->
    <div class="mx-auto mb-8 flex h-24 w-24 items-center justify-center rounded-full border-2 border-brand/50 bg-surface-alt text-3xl font-bold text-brand">
      F
    </div>

    <p class="mb-2 text-sm font-medium tracking-widest text-brand-light uppercase">{role}</p>
    <h1 class="mb-4 text-5xl font-bold tracking-tight text-white sm:text-6xl">{name}</h1>
    <p class="mx-auto max-w-lg text-lg leading-relaxed text-zinc-400">{bio}</p>

    <!-- Social links -->
    <div class="mt-8 flex items-center justify-center gap-4">
      <a
        href="https://github.com/falsisdev"
        target="_blank"
        rel="noopener noreferrer"
        class="rounded-lg border border-border bg-surface-alt px-5 py-2.5 text-sm font-medium text-zinc-300 transition-all hover:border-brand/50 hover:text-white"
      >
        GitHub
      </a>
      <a
        href="#projects"
        class="rounded-lg bg-brand px-5 py-2.5 text-sm font-medium text-white transition-all hover:bg-brand-light"
      >
        Projelere Göz At
      </a>
    </div>
  </div>

  <!-- Scroll indicator -->
  <div class="absolute bottom-8 animate-bounce text-zinc-500">
    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
    </svg>
  </div>
</section>
