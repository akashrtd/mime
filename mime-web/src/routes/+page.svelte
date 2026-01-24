<script>
  import { onMount } from "svelte";
  // Import logo dynamically so Vite handles the path/hashing
  import logoPath from "$lib/assets/mime-logo.png";
</script>

<svelte:head>
  <title>MIME | Integrated Modern Executor</title>
  <meta
    name="description"
    content="MIME: The headless browser automation engine for AI agents."
  />
  <!-- Force favicon via import to ensure correct path -->
  <link rel="icon" href={logoPath} />

  <!-- Open Graph / Link Preview -->
  <meta property="og:type" content="website" />
  <meta property="og:title" content="MIME | Integrated Modern Executor" />
  <meta
    property="og:description"
    content="The headless browser automation engine for AI agents. Built with Go and Rod, native to the Model Context Protocol."
  />
  <meta property="og:image" content={logoPath} />

  <!-- Twitter -->
  <meta name="twitter:card" content="summary" />
  <meta name="twitter:title" content="MIME | Integrated Modern Executor" />
  <meta
    name="twitter:description"
    content="The headless browser automation engine for AI agents."
  />
  <meta name="twitter:image" content={logoPath} />
</svelte:head>

<div
  class="min-h-screen bg-white font-mono text-[#0040DD] selection:bg-[#0040DD] selection:text-white"
>
  <!-- Navbar -->
  <nav
    class="fixed top-0 z-50 flex w-full items-center justify-between border-b border-[#0040DD]/10 bg-white/90 px-6 py-6 backdrop-blur-md supports-[backdrop-filter]:bg-white/60 lg:px-12"
  >
    <div class="group flex cursor-pointer items-center gap-4">
      <img
        src={logoPath}
        alt="MIME Logo"
        class="h-10 w-10 transition-transform duration-500 group-hover:rotate-90"
      />
      <span class="text-2xl font-bold tracking-tighter">MIME</span>
    </div>
    <div class="hidden gap-8 text-sm font-bold tracking-tight md:flex">
      <a href="#features" class="transition-opacity hover:opacity-60"
        >CAPABILITIES</a
      >
      <a href="#integration" class="transition-opacity hover:opacity-60"
        >INTEGRATION</a
      >
      <a
        href="https://github.com/akashrtd/mime"
        target="_blank"
        class="border border-[#0040DD] px-5 py-2.5 text-xs transition-all hover:bg-[#0040DD] hover:text-white"
      >
        GITHUB
      </a>
    </div>
  </nav>

  <!-- Hero Section -->
  <header
    class="mx-auto grid max-w-[1600px] grid-cols-1 items-center gap-16 px-6 pt-40 pb-20 lg:grid-cols-2 lg:px-12"
  >
    <div class="relative z-10 space-y-10">
      <h1
        class="text-6xl leading-[0.85] font-black tracking-tighter select-none lg:text-[7rem]"
      >
        BROWSER<br />
        AUTOMATION<br />
        <span
          class="bg-gradient-to-r from-[#0040DD] to-cyan-500 bg-clip-text text-transparent"
          >FOR AGENTS.</span
        >
      </h1>
      <p
        class="max-w-lg text-xl leading-relaxed font-light text-gray-500 lg:text-2xl"
      >
        The <strong class="font-bold text-[#0040DD]"
          >Integrated Modern Executor</strong
        >. A headless, MCP-native browser engine designed for high-performance
        AI workflows.
      </p>
      <div class="flex flex-wrap gap-5 pt-4">
        <a
          href="#integration"
          class="bg-[#0040DD] px-10 py-5 text-lg font-bold text-white shadow-xl shadow-[#0040DD]/20 transition-all hover:-translate-y-1 hover:bg-[#0030AA]"
        >
          INSTALL ENGINE
        </a>
        <a
          href="https://github.com/akashrtd/mime"
          target="_blank"
          class="border-2 border-gray-200 bg-white px-10 py-5 text-lg font-bold text-gray-600 transition-all hover:border-black hover:text-black"
        >
          VIEW SOURCE
        </a>
      </div>
    </div>

    <!-- Interactive Terminal Visualization -->
    <!-- Video Visualization -->
    <div class="group perspective-1000 relative lg:translate-x-10">
      <!-- Glitchy background effect -->
      <div
        class="absolute inset-0 bg-[#0040DD] opacity-20 blur-3xl transition-opacity duration-1000 group-hover:opacity-30"
      ></div>

      <div
        class="relative z-10 skew-y-1 overflow-hidden rounded-xl border border-gray-800 bg-[#0A0A0A] shadow-2xl transition-all duration-700 ease-out group-hover:skew-y-0"
      >
        <video
          src="/video_tui.mp4"
          autoplay
          muted
          loop
          playsinline
          class="w-full h-auto object-cover opacity-90"
        ></video>

        <!-- Overlay Gradient -->
        <div
          class="absolute inset-0 bg-gradient-to-t from-[#0A0A0A] via-transparent to-transparent opacity-50"
        ></div>
      </div>
    </div>
  </header>

  <!-- Marquee / Divider -->
  <div class="mb-24 w-full overflow-hidden bg-[#0040DD] py-3 text-white">
    <div
      class="animate-marquee flex gap-12 text-sm font-bold tracking-widest whitespace-nowrap uppercase opacity-80"
    >
      {#each Array(10) as _}
        <span>Headless Architecture</span>
        <span>•</span>
        <span>MCP Native</span>
        <span>•</span>
        <span>Zero Config</span>
        <span>•</span>
      {/each}
    </div>
  </div>

  <!-- Feature Grid -->
  <section
    id="features"
    class="border-t border-gray-100 bg-gray-50/50 px-6 py-24 lg:px-12"
  >
    <div class="mx-auto max-w-[1600px]">
      <div
        class="mb-20 flex flex-col items-end justify-between gap-8 md:flex-row"
      >
        <h2 class="max-w-2xl text-5xl font-bold tracking-tighter lg:text-7xl">
          BUILT FOR<br />THE MACHINE.
        </h2>
        <div
          class="max-w-xs text-right text-sm font-bold tracking-widest text-gray-500 uppercase"
        >
          Performance Specs<br />
          v1.0.0 Stable
        </div>
      </div>

      <div class="grid grid-cols-1 gap-8 md:grid-cols-3">
        <div
          class="group border border-t-4 border-gray-100 bg-white p-10 shadow-sm transition-all duration-300 hover:border-[#0040DD] hover:shadow-2xl"
        >
          <div
            class="mb-8 flex h-16 w-16 items-center justify-center bg-blue-50 text-[#0040DD] transition-colors group-hover:bg-[#0040DD] group-hover:text-white"
          >
            <svg
              class="h-8 w-8"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              ><path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="1.5"
                d="M19.428 15.428a2 2 0 00-1.022-.547l-2.384-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z"
              ></path></svg
            >
          </div>
          <h3 class="mb-4 text-2xl font-bold">MCP Native</h3>
          <p class="leading-relaxed text-gray-500">
            Implements the Model Context Protocol out of the box. Plug MIME into
            Claude, Cursor, or Ensemble instantly.
          </p>
        </div>

        <div
          class="group border border-t-4 border-gray-100 bg-white p-10 shadow-sm transition-all duration-300 hover:border-[#0040DD] hover:shadow-2xl"
        >
          <div
            class="mb-8 flex h-16 w-16 items-center justify-center bg-blue-50 text-[#0040DD] transition-colors group-hover:bg-[#0040DD] group-hover:text-white"
          >
            <svg
              class="h-8 w-8"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              ><path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="1.5"
                d="M13 10V3L4 14h7v7l9-11h-7z"
              ></path></svg
            >
          </div>
          <h3 class="mb-4 text-2xl font-bold">Stealth Mode</h3>
          <p class="leading-relaxed text-gray-500">
            Advanced fingerprint masking and stealth navigation primitives
            ensure your agents never get blocked.
          </p>
        </div>

        <div
          class="group border border-t-4 border-gray-100 bg-white p-10 shadow-sm transition-all duration-300 hover:border-[#0040DD] hover:shadow-2xl"
        >
          <div
            class="mb-8 flex h-16 w-16 items-center justify-center bg-blue-50 text-[#0040DD] transition-colors group-hover:bg-[#0040DD] group-hover:text-white"
          >
            <svg
              class="h-8 w-8"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              ><path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="1.5"
                d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4m0 5c0 2.21-3.582 4-8 4s-8-1.79-8-4"
              ></path></svg
            >
          </div>
          <h3 class="mb-4 text-2xl font-bold">State Preservation</h3>
          <p class="leading-relaxed text-gray-500">
            Persistent sessions allow agents to maintain login states and
            cookies across multiple workflow steps.
          </p>
        </div>
      </div>
    </div>
  </section>

  <!-- Install -->
  <section id="integration" class="bg-[#0040DD] px-6 py-32 text-white lg:px-12">
    <div class="mx-auto max-w-4xl text-center">
      <h2 class="mb-10 text-4xl font-bold tracking-tight lg:text-5xl">
        DEPLOY THE AGENT
      </h2>

      <div
        class="mx-auto inline-flex flex-col items-stretch overflow-hidden rounded-lg bg-white p-1 text-[#0040DD] shadow-2xl md:flex-row"
      >
        <div
          class="flex flex-1 items-center gap-3 px-8 py-5 text-left font-mono"
        >
          <span class="opacity-50 select-none">$</span>
          <span class="font-bold"
            >go install github.com/akashrtd/mime/cmd/mime@latest</span
          >
        </div>
        <button
          class="bg-[#0040DD] px-8 py-4 font-bold tracking-wide text-white uppercase transition-colors hover:bg-black"
        >
          Copy Command
        </button>
      </div>

      <p class="mt-12 font-mono text-sm opacity-60">
        Requires Go 1.21+ • <a
          href="https://github.com/akashrtd/mime"
          class="underline hover:text-white">Read Documentation</a
        >
      </p>
    </div>
  </section>

  <footer
    class="border-t border-gray-100 bg-gray-50 py-16 text-center text-sm text-gray-400"
  >
    <img
      src={logoPath}
      alt="MIME"
      class="mx-auto mb-6 h-8 w-8 opacity-40 grayscale transition-all duration-500 hover:grayscale-0"
    />
    <p class="font-light tracking-wide">
      © 2026 MIME Project. Built for Agents.
    </p>
  </footer>
</div>

<style>
  /* Custom utility for marquee animation */
  @keyframes marquee {
    0% {
      transform: translateX(0);
    }
    100% {
      transform: translateX(-50%);
    }
  }
  .animate-marquee {
    animation: marquee 30s linear infinite;
  }
</style>
